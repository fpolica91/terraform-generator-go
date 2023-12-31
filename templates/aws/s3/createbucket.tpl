{% for bucket in buckets %}

{% set bucket_name = "" %}
  
{% if bucket.Configuration['bucket'] %}
  {% set bucket_name = bucket.Configuration['bucket'] %}
  {% else %}
  {% set bucket_name = "bucket_"|add:forloop.Counter %}
{% endif %}

  resource "aws_s3_bucket" "{{ bucket_name }}" {
    {%- for key, value in bucket.Configuration -%}
      {%- if value -%}
        {% if key == "tags" %}
          tags = {
              {%- for tag in value -%}
              {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
              {%- endfor -%}
          }
        {% else %}
          {{ key }} = "{{ value }}"
        {% endif %}
      {%- endif -%}
    {%- endfor -%}
  }

    resource "aws_s3_bucket_ownership_controls" "ownership_controls_{{ forloop.Counter }}" {
        bucket = aws_s3_bucket.{{bucket_name}}.id
        rule {
            object_ownership = "BucketOwnerPreferred"
        }
    }

    {% if bucket.PublicAccessBlock %}
        resource "aws_s3_bucket_public_access_block" "public_access_block_{{ forloop.Counter }}" {
            bucket = aws_s3_bucket.{{bucket_name}}.id
            {%- for key, value in bucket.PublicAccessBlock -%}
                {% if key != "acl" %}
                    {{ key }} = {{ value | lower }}
                {% endif %}
            {%- endfor -%}
        }
    {% endif %}



    {% if bucket.PublicAccessBlock.acl %}
        resource "aws_s3_bucket_acl" "acl_bucket_{{ forloop.Counter }}" {
            depends_on = [aws_s3_bucket_ownership_controls.ownership_controls_{{ forloop.Counter }}, aws_s3_bucket_public_access_block.public_access_block_{{ forloop.Counter }}]
            bucket = aws_s3_bucket.{{bucket_name}}.id
            acl    = "{{ bucket.PublicAccessBlock.acl }}"
        }
    {% endif %}

       {% if bucket.WebsiteConfiguration.host_website %}
          resource "aws_s3_bucket_website_configuration" "website_config_{{ forloop.Counter }}" {
            bucket = aws_s3_bucket.{{bucket_name}}.id
            index_document  {
              suffix = "{{ bucket.WebsiteConfiguration.index_document }}"
            }
            {% if bucket.WebsiteConfiguration.error_document %}
              error_document  {
                key = "{{ bucket.WebsiteConfiguration.error_document }}"
              }
            {% endif %}
          {% if bucket.WebsiteConfiguration.routing_rules %}
              {%- for rule in bucket.WebsiteConfiguration.routing_rules -%}
                routing_rule {
                  {% if rule.redirect %}
                    redirect {
                      {% if rule.redirect.host_name %}
                        host_name = "{{ rule.redirect.host_name }}"
                      {%- endif -%}
                      {% if rule.redirect.http_redirect_code %}
                        http_redirect_code = {{ rule.redirect.http_redirect_code | integer}}
                      {%- endif -%}
                      {% if rule.redirect.protocol %}
                        protocol = "{{ rule.redirect.protocol }}"
                      {%- endif -%}
                      {% if rule.redirect.replace_key_prefix_with %}
                        replace_key_prefix_with = "{{ rule.redirect.replace_key_prefix_with }}"
                      {%- endif -%}
                      {% if rule.redirect.replace_key_with %}
                        replace_key_with = "{{ rule.redirect.replace_key_with }}"
                      {%-endif %}
                    }
                  {% endif %}
                  {% if rule.condition %}
                    condition {
                      {% if rule.condition.http_error_code_returned_equals %}
                        http_error_code_returned_equals = {{ rule.condition.http_error_code_returned_equals | integer}}
                      {%- endif -%}
                      {% if rule.condition.key_prefix_equals %}
                        key_prefix_equals = "{{ rule.condition.key_prefix_equals }}"
                      {% endif %}
                    }
                  {% endif %}
                }
              {%- endfor -%}
          {% endif %}
          }
       {%endif%}

  
    {% if bucket.CorsConfiguration %}
      resource "aws_s3_bucket_cors_configuration" "bucket_{{ loop.index }}_cors" {
        bucket = aws_s3_bucket.bucket_{{ bucket_name }}.id
        {% for rule in bucket.CorsConfiguration.cors_rules_list %}
          cors_rule {
            {% for key, value in rule %}
              {% if key != 'max_age_seconds' %}
              {{ key }} = [
                {%- for item in value -%}
                  "{{ item }}"{% if not forloop.Last %}, {% endif %}
                {%- endfor -%}
              ]
              {% else %}
              {{ key }} = {{ value  | integer}}
              {% endif %}
            {% endfor %}
          }
        {% endfor %}
      }
    {% endif %}
{%- endfor -%}

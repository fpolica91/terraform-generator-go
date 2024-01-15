{% for bucket in storage_objects %}
  {% if bucket.Configuration %}
    {% if bucket.Configuration.name %}
      resource "aws_s3_bucket" "{{ bucket.Configuration.name }}" {
        {%- for key, value in bucket.Configuration -%}
          {%- if value -%}
            {% if key == "name" %}
              bucket = "{{ value }}"

            {% elif key == "tags" %}
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
    {% endif %}
  {%- endif -%}

    resource "aws_s3_bucket_ownership_controls" "ownership_controls_{{ forloop.Counter }}" {
        bucket = aws_s3_bucket.{{bucket.Configuration.name}}.id
        rule {
            object_ownership = "BucketOwnerPreferred"
        }
    }

    {% if bucket.PublicAccessBlock %}
        resource "aws_s3_bucket_public_access_block" "public_access_block_{{ forloop.Counter }}" {
              bucket = aws_s3_bucket.{{bucket.Configuration.name}}.id
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
            bucket = aws_s3_bucket.{{bucket.Configuration.name}}.id
            acl    = "{{ bucket.PublicAccessBlock.acl }}"
        }
    {%- endif -%}

       {% if bucket.WebsiteConfiguration.host_website %}
          resource "aws_s3_bucket_website_configuration" "website_config_{{ forloop.Counter }}" {
            bucket = aws_s3_bucket.{{bucket.Configuration.name}}.id
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

  
    {% if bucket.CorsConfiguration.cors_rules_list %}
      resource "aws_s3_bucket_cors_configuration" "bucket_{{ loop.index }}_cors" {
        bucket = aws_s3_bucket.bucket_{{ bucket.Configuration.name }}.id
        {% for rule in bucket.CorsConfiguration.cors_rules_list %}
          cors_rule {
              {% for key, value in rule %}
                {%- if value  -%}
                  {% if key == 'max_age_seconds' %}
                   {{ key }} = {{ value  | integer}}
                  {% elif key == 'allowed_methods' or key == 'allowed_headers' %}
                    {{ key }} = [
                      {%- for item in value -%}
                        "{{ item }}"
                      {%- endfor -%}
                    ]
                  {% elif key == 'allowed_origins' %}
                    {{ key }} = [
                      {%- for item in value -%}
                        "{{ item.origin }}"
                      {%- endfor -%}
                    ]
                  {%- else -%}
                    {{ key }} = "{{ value }}"
                  {%- endif -%}
                {% endif %}
              {% endfor %}
          }
        {% endfor %}
      }
    {% endif %}
{%- endfor -%}
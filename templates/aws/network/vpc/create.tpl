{% for vpc in network_units %}

  resource "aws_vpc" "{{vpc.Configuration.name}}" {
    {%- for key, value in vpc.Configuration -%}
      {% if value and key != "name"%}
        {% if key == "tags" %}
          tags = {
              {%- for tag in value -%}
               {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
              {%- endfor -%}
          }
        {% elif key == "enable_dns_support" or key == "enable_dns_hostnames" %}
          {{ key }} = {{ value  | lower }}
        {% else %}
          {{ key }} = "{{ value }}"
        {%- endif -%}
      {% endif %}
    {% endfor %}
  } 
{%- endfor -%}
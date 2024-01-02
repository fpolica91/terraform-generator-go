{% for vpc in vpcs %}
  {%set vpc_name = ""%}
  {% if vpc.Configuration['name'] %}
    {% set vpc_name = vpc.Configuration['name'] %}
    {% else %}
    {% set vpc_name = "vpc_"|add:forloop.Counter %}
  {% endif %}

  resource "aws_vpc" "{{vpc_name}}" {
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

  {% for subnet in vpc.SubnetConfiguration %}
    {% set subnet_name = "" %}
      {% if subnet.name %}
      {% set subnet_name = subnet.name %}
      {% else %}
      {% set subnet_name = "subnet_"|add:forloop.Counter %}
    {% endif %}

    resource "aws_subnet" "{{ subnet_name }}" {
      vpc_id     = aws_vpc.{{vpc_name}}.id
      {%- for key, value in subnet -%}
        {% if value and key != "name" %}
          {% if key == "tags" %}
            tags = {
                {%- for tag in value -%}
                 {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
                {%- endfor -%}
            }
          {% else %}
            {{ key }} = "{{ value }}"
          {% endif %}
        {% endif %}
      {%-endfor-%}
    }
  {% endfor %}


  {% for gateway in vpc.GatewayConfiguration %}
    {% set gateway_name = "" %}
      {% if gateway.name %}
      {% set gateway_name = gateway.name %}
      {% else %}
      {% set gateway_name = "gateway_"|add:forloop.Counter %}
    {% endif %}


    {% if gateway.type == "default" %}
       resource "aws_internet_gateway" "{{ gateway_name }}" {
          vpc_id = aws_vpc.{{vpc_name}}.id
          {%- for key, value in gateway -%}
            {% if value and key != "name" and key != "type" %}
              {% if key == "tags" %}
                tags = {
                    {%- for tag in value -%}
                     {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
                    {%- endfor -%}
                }
              {% else %}
                {{ key }} = "{{ value }}"
              {% endif %}
            {% endif %}
          {%-endfor-%}
       }
       {% elif gateway.type == 'nat' and gateway.public_gateway_name %}

        resource "aws_nat_gateway" "{{ gateway_name }}" {
          subnet_id     = aws_subnet.{{gateway.subnet_name}}.id
          depends_on = [aws_internet_gateway.{{gateway.public_gateway_name}}]
          {%- for key, value in gateway -%}
            {% if key == "tags" %}
              tags = {
                    {%- for tag in value -%}
                     {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
                    {%- endfor -%}
                }
            {% endif %}
          {% endfor %}
        }
    {% endif %}
  {% endfor %}

{%- endfor -%}
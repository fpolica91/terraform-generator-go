{% for gateway in network_units %}
  {% if gateway.type ==  "nat_gateway" %}
  resource "aws_nat_gateway" "{{ gateway.Configuration.name }}" {
      subnet_id     = aws_subnet.{{gateway.Configuration.subnet_name}}.id
      depends_on = [aws_internet_gateway.{{gateway.Configuration.public_gateway_name}}]
      {%- for key, value in gateway.Configuration -%}
        {% if key == "tags" and key != "name" and key != "subnet_name" and key != "public_gateway_name" %}
          tags = {
                {%- for tag in value -%}
                  {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
                {%- endfor -%}
            }
        {% endif %}
      {% endfor %}
    }
  {%- else -%}
   resource "aws_internet_gateway" "{{ gateway.Configuration.name }}" {
      vpc_id = aws_vpc.{{gateway.Configuration.vpc_name}}.id
        {%- for key, value in gateway.Configuration -%}
          {% if value and key != "name" and key != "type" and key != "vpc_name"  and key != "name"%}
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
        {% endfor %}
       }
  {%- endif -%}
{%- endfor -%}
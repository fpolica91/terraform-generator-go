{% for subnet in network_units %}
  resource "aws_subnet" "{{ subnet.Configuration.name }}" {
      {% if subnet.Configuration.vpc_name %}
        vpc_id     = aws_vpc.{{subnet.Configuration.vpc_name}}.id
      {% else %}
        vpc_id     = aws_vpc.{{"your_vpc_name"}}.id
      {% endif %}
      {%- for key, value in subnet.Configuration -%}
        {% if value and key != "name" and key != "vpc_name" %}
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
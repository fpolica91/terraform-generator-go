{% for ec2 in compute %}
  resource "aws_instance" "{{ec2.Configuration.name}}" {
   {%- for key, value in ec2.Configuration -%}
    {%- if value -%}
      {% if key == "tags" %}
        tags = {
            {%- for tag in value -%}
            {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
            {%- endfor -%}
        }
        {% elif key == "associate_public_ip_address" %}
          associate_public_ip_address = {{ value | lower }}
      {% else %}
        {{ key }} = "{{ value }}"
      {% endif %}
    {%- endif -%}
  {%- endfor -%}
}

{% endfor %}
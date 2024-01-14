{% for ec2 in compute %}
  {%set ec2_name = ""%}
  {% if ec2.Configuration.name %}
    {% set ec2_name = ec2.Configuration.name %}
    {% else %}
    {% set ec2_name = "ec2_"|add:forloop.Counter %}
  {% endif %}

  resource "aws_instance" "{{ec2_name}}" {
   {%- for key, value in ec2.Configuration -%}
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

{% endfor %}
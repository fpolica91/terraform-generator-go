{% for lambda in compute %}
{% set lambda_name = "" %}
  
{% if lambda.Configuration.name %}
  {% set lambda_name = lambda.Configuration.name %}
  {% else %}
  {% set lambda_name = "lambda_"|add:forloop.Counter %}
{% endif %}

resource "aws_lambda_function" "{{ lambda_name }}" {
  {%- for key, value in lambda.Configuration -%}
    {%- if value -%}
      {% if key == "tags" %}
        tags = {
            {%- for tag in value -%}
            {{ tag.key }} = "{{ tag.value }}"{% if not forloop.Last %},{% endif %}
            {%- endfor -%}
        }
        {% elif key == "publish" %}
          {{key}} = {{ value | lower }}
      {% else %}
        {{ key }} = "{{ value }}"
      {% endif %}
    {%- endif -%}
  {%- endfor -%}
}
{% endfor %}
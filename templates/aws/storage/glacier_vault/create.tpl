{% for volume in storage_objects %}
{% if volume.Configuration.sns_configuration %}
resource "aws_sns_topic" "{{ volume.Configuration.sns_configuration.name }}" {
  name = "{{ volume.Configuration.sns_configuration.name }}"
}
{% endif %}
resource "aws_glacier_vault" "{{ volume.Configuration.name }}" {
  {% if volume.name %}
    name               = "{{ volume.Configuration.name }}"
  {%- endif -%}

  {% if volume.Configuration.access_policy %}
    access_policy      =   "
      <<POLICY
        {{ volume.Configuration.access_policy}}
      POLICY
    "
  {%- endif -%}


  {% if volume.Configuration.sns_configuration %}
    notification_configuration {
      sns_topic = "${aws_sns_topic.{{ volume.Configuration.sns_configuration.name }}.arn}"
      events = [
        {% for event in volume.Configuration.sns_configuration.events %}
          "{{ event }}"
        {% endfor %}
      ]
    }
  {%- endif -%}

  {% if volume.Configuration.tags %}
    tags = {
      {% for tag in volume.Configuration.tags %}
        {{ tag.key }} = "{{ tag.value }}"
      {% endfor %}
    }
  {%- endif -%}

}
{% endfor %}
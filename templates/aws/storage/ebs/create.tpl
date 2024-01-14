
{% for volume in storage_objects %}

resource "aws_ebs_volume" "{{ volume.Configuration.name }}" {
  {% if volume.name %}
    name               = "{{ volume.Configuration.name }}"
  {%- endif -%}


  {% if volume.Configuration.availability_zone %}
    availability_zone = "{{ volume.Configuration.availability_zone }}"
    {% if volume.Configuration.availability_zone == "all" %}
      availability_zone = "${data.aws_availability_zones.available.names}"
    {%- endif -%}
    {%-else-%}
      availability_zone = "us-east-1"
  {% endif %}

  {%- if volume.Configuration.size -%}
    size               = {{ volume.Configuration.size | integer }}
  {%- endif -%}

  {% if volume.Configuration.type %}
    type               = "{{ volume.Configuration.type }}"
  {%- endif -%}

  {% if volume.Configuration.encrypted %}
    encrypted          = "{{ volume.Configuration.encrypted }}"
  {%- endif -%}

  {% if volume.Configuration.kms_key_id %}
    kms_key_id         = "{{ volume.Configuration.kms_key_id }}"
  {%- endif -%}

  {% if volume.Configuration.snapshot_id %}
    snapshot_id        = "{{ volume.Configuration.snapshot_id }}"
  {%- endif -%}

  {% if volume.Configuration.tags %}
    tags = {
      {% for tag in volume.Configuration.tags %}
        {{ tag.key }} = "{{ tag.value }}"
      {% endfor %}
    }
  {% endif %}
}
  

{% if volume.Configuration.name %}
  {% for attachment in volume.Configuration.attachments %}
    resource "aws_volume_attachment" "{{ attachment.Configuration.name }}" {
      device_name = "{{ attachment.Configuration.device_name }}"
      volume_id   = "${aws_ebs_volume.{{ volume.Configuration.name }}.id}"
      instance_id = "{{ attachment.Configuration.instance_id }}"
    }
  {% endfor %}
{% endif %}

{% if volume.Configuration.name %}
  {% for snapshot in volume.Configuration.snapshots %}
    resource "aws_ebs_snapshot" "{{ snapshot.name }}" {
      volume_id = "${aws_ebs_volume.{{ volume.Configuration.name }}.id}"
      tags {
        Name = "{{ snapshot.name }}"
      }
    }
  {% endfor %}
{% endif %}

  


{% endfor %}
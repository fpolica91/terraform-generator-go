{% for item in network_units %}
  resource "aws_route53_zone" "{{ item.Configuration.name }}" {
    name = "{{item.Configuration.domain}}"  # Replace with your domain name
  }
  {% for record in item.Configuration.route53_records %}
      resource "aws_route53_record" "{{ item.Configuration.name }}" {
      zone_id = aws_route53_zone.{{ item.Configuration.name }}.zone_id
      name    = "{{record.record_domain_name}}"  # Replace with your desired record name
      type    =   "{{record.type}}"               # Type of the DNS record
      ttl     =  {{record.ttl}}              # Time to live
      records = ["{{record.value}}"]      # Replace with the IP address you want to point to
    }
  {% endfor %}
{% endfor %}
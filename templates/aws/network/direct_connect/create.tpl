{% for item in network_units %}
resource "aws_dx_connection" "{{ item.Configuration.name }}" {
  name            = "{{item.Configuration.connection_name}}"  # Replace with your connection name
  bandwidth       = "{{item.Configuration.bandwidth}}"            # Specify the connection bandwidth
  location        = "{{item.Configuration.location}}"           # Replace with the AWS Direct Connect location
  # The location value must be a valid AWS Direct Connect location. 
  # You can get a list of locations using `aws directconnect describe-locations` AWS CLI command.
}
{%endfor%}
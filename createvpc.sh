curl -d '{
  "vpcs": [
    {
      "configuration": 
        {
          "name": "my-vpc2",
          "cidr_block": "10.0.0.0/16",
          "enable_dns_hostnames": true,
          "enable_dns_support": true,
          "instance_tenancy": "default",
          "tags": [
            {"key": "Environment", "value": "Dev"},
            {"key": "Project", "value": "ProjectX"}
          ]
        },
      "gateway_configuration": [
        {
          "name": "gateway1",
          "type": "default",
          "tags": [
            {"key": "Name", "value": "gateway1"}
          ]
        },
        {
          "name": "gateway2",
          "type": "nat",
          "subnet_name": "subnet1",
          "public_gateway_name": "gateway1",
          "tags": [
            {"key": "Name", "value": "gateway2"}
          ]
        }
      ],
       "subnet_configuration": [
        {
          "name": "subnet1",
          "cidr_block": "10.0.1.0/24",
          "availability_zone": "us-west-2a",
          "tags": [
            {"key": "Name", "value": "subnet1"}
          ]
        }
      ]
    }
  ]
}' \
-H "Content-Type: application/json" \
-X POST http://localhost:8080/vpcs

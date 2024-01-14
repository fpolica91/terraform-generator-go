package pkgs

type NetworkUnit struct {
	Configuration        map[string]interface{}   `json:"configuration"`
	GatewayConfiguration []map[string]interface{} `json:"gateway_configuration"`
	SubnetConfiguration  []map[string]interface{} `json:"subnet_configuration"`
}

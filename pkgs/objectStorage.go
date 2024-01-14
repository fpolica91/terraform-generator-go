package pkgs

type ObjectStorage struct {
	Id                   string                 `json:"id"`
	Configuration        map[string]interface{} `json:"configuration"`
	PublicAccessBlock    map[string]interface{} `json:"public_access_block"`
	WebsiteConfiguration map[string]interface{} `json:"website_configuration"`
	CorsConfiguration    map[string]interface{} `json:"cors_configuration"`
}

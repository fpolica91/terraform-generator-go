package pkgs

type Bucket struct {
	Configuration        map[string]interface{} `json:"configuration"`
	PublicAccessBlock    map[string]interface{} `json:"public_access_block"`
	WebsiteConfiguration map[string]interface{} `json:"website_configuration"`
	CorsConfiguration    map[string]interface{} `json:"cors_configuration"`
}

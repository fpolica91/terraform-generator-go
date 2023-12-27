


terraform {
  required_providers {
    {{provider.Provider}} = {
      source  = "{{provider.ProviderSource}}"
      version = "{{provider.ProviderVersion}}"
    }
  }
}

provider "{{provider.Provider}}" {
  access_key = var.access_key
  secret_key = var.secret_key
  region     = var.region
}



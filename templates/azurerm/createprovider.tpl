
terraform {
  required_providers {
    {{provider.Provider}} = {
      source  = "{{provider.ProviderSource}}"
      version = "{{provider.ProviderVersion}}"
    }
  }
}

provider "azurerm" {
  features {}
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

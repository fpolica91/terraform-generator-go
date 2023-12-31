variable "subscription_id" {
  type        = string
  description = "The Subscription ID used for the AzureRM Provider."
}

variable "client_id" {
  type       = string
  description = "The Client ID used for the AzureRM Provider."
}

variable "client_secret" {
  type        = string  
  description = "The Client Secret used for the AzureRM Provider."
}

variable "tenant_id" {
  type       = string
  description = "The Tenant ID used for the AzureRM Provider."
}

variable "topic_name" {
  description = "The name of the bucket. This must be globally unique."
  type        = string
}

variable "stack_id" {
  description = "The ID of the Nitric stack"
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group in which to create the topic."
  type        = string
}

variable "location" {
  description = "The location/region where the topic should be created."
  type        = string
}

variable "subscribers" {
  description = "A list of subscribers to the topic."
  type = map(object({
    url         = string
    event_token = string
    service_principal_client_id = string
    service_principal_tenant_id = string
  }))
}

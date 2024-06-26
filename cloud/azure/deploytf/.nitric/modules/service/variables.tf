variable "name" {
  description = "The name of the service"
  type        = string
}
variable "container_app_environment_id" {
  description = "The ID of the container app environment"
  type        = string
}

variable "resource_group_name" {
  description = "The name of the resource group"
  type        = string
}

variable "application_client_id" {
  description = "The client ID of the application for which to create this services service principal"
  type        = string
}

variable "registry_server" {
  description = "The server of the container registry"
  type        = string
}

variable "registry_username" {
  description = "The username of the container registry"
  type        = string
}

variable "registry_password" {
  description = "The password of the container registry"
  type        = string
}

variable "tenant_id" {
  description = "The tenant ID of the application for which to create this services service principal"
  type        = string
}

variable "client_secret" {
  description = "The client secret of the application for which to create this services service principal"
  type        = string
}



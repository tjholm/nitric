variable "name" {
  description = "The name of the API Gateway"
  type = string
}

variable "resource_group_name" {
  description = "The name of the resource group"
  type = string
}

variable "stack_id" {
  description = "The ID of the stack"
  type = string
}

variable "spec" {
  description = "Open API spec"
  type = string
}

variable "location" {
  description = "The location of the API Gateway"
  type = string
}
variable "publisher_name" {
  description = "The name of the publisher"
  type = string
}

variable "publisher_email" {
  description = "The email of the publisher"
  type = string
}
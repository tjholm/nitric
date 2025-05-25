variable "nitric" {
    type = object({
        name = string
        image_id = string
    })
}

variable "timeout" {
    type = number
    default = 300
}

variable "memory" {
    type = number
    default = 1024
}

variable "ephemeral_storage" {
    type = number
    default = 1024
}

variable "environment" {
    type = map(string)
    default = {}
}

variable "subnet_ids" {
    type = list(string)
    default = []
}

variable "security_group_ids" {
    type = list(string)
    default = []
}

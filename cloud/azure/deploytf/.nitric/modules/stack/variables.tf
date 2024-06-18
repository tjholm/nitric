
variable stack_name {
    type = string
    description = "The name of the nitric stack"
}

variable region {
    type = string
    description = "The Azure region to deploy resources to"
}

variable deploy_storage {
    type = bool
    description = "Whether to deploy an Azure Storage Account for this stack"
}

variable deploy_key_vault {
    type = bool
    description = "Whether to deploy an Azure Key Vault for this stack"
}
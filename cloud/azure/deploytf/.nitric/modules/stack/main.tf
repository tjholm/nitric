resource "random_string" "id" {
  length  = 8
  special = false
  upper   = false
}

# Create an azure resource group
resource "azurerm_resource_group" "rg" {
  name     = "${var.stack_name}-${random_string.id.result}"
  location = var.region
}

# Deploy an Azure Storage Account
resource "azurerm_storage_account" "sa" {
  count = var.deploy_storage ? 1 : 0

  name                     = "${var.stack_name}${random_string.id.result}"
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

# Deploy an Azure Key Vault
resource "azurerm_key_vault" "kv" {
  count = var.deploy_key_vault ? 1 : 0

  name                = "${var.stack_name}${random_string.id.result}"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  tenant_id           = data.azurerm_client_config.current.tenant_id
  soft_delete_retention_days = 0
  sku_name            = "standard"
}

# Deploy a container apps registry
resource "azurerm_container_registry" "acr" {
  name                = "${var.stack_name}${random_string.id.result}"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  # TODO: Add SKU configurability
  sku                 = "Basic"
}

# Deploy a new operational insights workspace
resource "azurerm_log_analytics_workspace" "law" {
  name                = "${var.stack_name}${random_string.id.result}"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  sku                 = "PerGB2018"
}

# Deploy an Azure container apps environment
resource "azurerm_container_app_environment" "environment" {
  name                       = "${var.stack_name}-${random_string.id.result}-environment"
  location                   = azurerm_resource_group.rg.location
  resource_group_name        = azurerm_resource_group.rg.name
  log_analytics_workspace_id = azurerm_log_analytics_workspace.law.id
}

# Get the azure AD client config
data "azurerm_client_config" "current" {}

# Deploy an AzureAD Application to enable secure webhook delivery
resource "azuread_application" "webhook" {
  display_name = "${var.stack_name}-${random_string.id.result}-webhook"
  owners = [ data.azurerm_client_config.current.object_id ]
  app_role {
    allowed_member_types = ["Application"]
    description          = "Enables webhook subscriptions to authenticate using this application"
    display_name         = "AzureEventGridSecureWebhookSubscriber"
    id                   = "4962773b-9cdb-44cf-a8bf-237846a00ab7"
    value                = "4962773b-9cdb-44cf-a8bf-237846a00ab7"
  }
}
# Create a new azure api management service
resource "azurerm_api_management" "service" {
  name                = var.name
  location            = var.location
  resource_group_name = var.resource_group_name
  publisher_name      = var.publisher_name
  publisher_email     = var.publisher_email

  sku_name = "Consumption"
}

# Create a new api management api
resource "azurerm_api_management_api" "api" {
  name                = var.name
  resource_group_name = var.resource_group_name
  api_management_name = azurerm_api_management.service.name
  revision            = "1"
  display_name        = var.name
  path                = "/"
  protocols           = ["https"]
  import {
    content_format = "openapi+json"
    content_value  = var.spec
  }
}

# For each reachable service create a new api management operation
resource "azurerm_api_management_api_operation_policy" "operations" {
  api_name            = azurerm_api_management_api.api.name
  api_management_name = azurerm_api_management.service.name
  resource_group_name = var.resource_group_name
  operation_id        = each.value.operation_id

  xml_content = templatefile("${path.module}/policies.xml.tfpl", {
    endpoint = each.value.endpoint
    resource_id = each.value.resource_id
    client_id = each.value.client_id
    jwts = each.value.jwts
  })
}

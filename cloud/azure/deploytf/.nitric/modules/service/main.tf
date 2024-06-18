# Create a random password for the event token for services
resource "random_password" "event_token" {
  length           = 32
  special          = false
  override_special = ["/", "+", "=", "@", "?", "&", "%", "#", "!", "-", "_"]
}

# Create a new service principal for the Nitric service
resource "azurerm_service_principal" "service_principal" {
  application_id = var.application_id
}

# create a new app role assignment
resource "azuread_app_role_assignment" "example" {
  app_role_id         = azuread_service_principal.msgraph.app_role_ids["User.Read.All"]
  principal_object_id = azuread_service_principal.example.object_id
  resource_object_id  = azuread_service_principal.msgraph.object_id
}
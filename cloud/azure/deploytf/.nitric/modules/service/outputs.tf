output "host_url" {
  value = azurerm_container_app.container_app.host_url
}

output "container_app" {
    value = azurerm_container_app.container_app
}

output "service_principal_client_id" {
  value = azuread_service_principal.sp.client_id
}

output "service_principal_tenant_id" {
  value = azuread_service_principal.sp.application_tenant_id
}

output "event_token" {
  value = azurerm_container_app.random_password.event_token.result
}

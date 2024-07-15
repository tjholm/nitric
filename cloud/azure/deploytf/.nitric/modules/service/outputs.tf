output "host_url" {
  value = azurerm_container_app.container_app.host_url
}

output "container_app" {
    value = azurerm_container_app.container_app
}

output "service_principal_id" {
  value = azuread_service_principal.sp
}

output "event_token" {
  value = azurerm_container_app.random_password.event_token.result
}
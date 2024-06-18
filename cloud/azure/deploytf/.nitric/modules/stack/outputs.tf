output "stack_id" {
  description = "The randomized Id of the nitric stack"
  value       =  random_string.id.result
}

output "resource_group_name" {
  description = "The name of the Azure resource group"
  value       = azurerm_resource_group.rg.name
}

output "storage_account_name" {
  description = "The name of the Azure storage account"
  value       = azurerm_storage_account.sa[0].name
}

output "key_vault_name" {
  description = "The name of the Azure key vault"
  value       = azurerm_key_vault.kv[0].name
}
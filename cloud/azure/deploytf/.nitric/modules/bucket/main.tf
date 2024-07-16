# deploy an azure storage container
resource "azurerm_storage_container" "container" {
  name                 = var.bucket_name
  storage_account_name = var.storage_account_name
}

# for each subscriber create an eventgrid subscription to the above container
resource "azurerm_eventgrid_event_subscription" "bucket_subscriptions" {
  for_each = var.bucket_subcribers

  name  = "${var.bucket_name}-${each.key}"
  scope = azurerm_storage_container.container.id

  webhook_endpoint {
    url = "${each.value.host_url}/${each.value.event_token}/${var.bucket_name}"
    max_events_per_batch = 1
    active_directory_app_id_or_uri = each.value.sp_client_id
    active_directory_tenant_id = each.value.sp_tenant_id
  }

  retry_policy {
    max_delivery_attempts = 30
    event_time_to_live    = 5
  }

  // TODO: Only apply correct event types
  included_event_types = ["Microsoft.Storage.BlobCreated", "Microsoft.Storage.BlobDeleted"]

  subject_filter {
    subject_begins_with = each.value.prefix_filter
  }
} 

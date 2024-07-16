# deploy an azure event grid topic
resource "azurerm_eventgrid_topic" "topic" {
  name                = var.topic_name
  resource_group_name = var.resource_group_name
  location            = var.location
  tags                = {
    "x-nitric-${var.stack_id}-name" = var.topic_name,
    "x-nitric-${var.stack_id}-type" = "topic",
  }
}

# for each subscriber create a subscription to the above topic
resource "azurerm_eventgrid_event_subscription" "subscription" {
  for_each = var.subscribers
  name                = "${var.topic_name}-${var.subscribers[count.index].name}"
  scope               = azurerm_eventgrid_topic.topic.id
  event_delivery_schema = "EventGridSchema"
  webhook_endpoint {
    url = "${each.value.url}/${each.value.event_token}/${var.topic_name}"
    max_events_per_batch = 1
    active_directory_app_id_or_uri = each.value.service_principal_client_id
    active_directory_tenant_id = each.value.service_principal_tenant_id
  }
}
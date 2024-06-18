
# Create a custom azure role definition for reading blobs
resource "azurerm_role_definition" "kvstore_read" {
  name        = "${var.stack_id}-KeyValueStoreRead"
  description = "Nitric KeyValue store read access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    data_actions = [
      "Microsoft.Storage/storageAccounts/tableServices/tables/entities/read"
    ]
    actions     = []
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "queue_enqueue" {
  name        = "${var.stack_id}-QueueEnqueue"
  description = "Nitric Queue send access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = [
      "Microsoft.Storage/storageAccounts/queueServices/queues/read"
    ]
    data_actions = [
      "Microsoft.Storage/storageAccounts/queueServices/queues/messages/write"
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "queue_dequeue" {
  name        = "${var.stack_id}-QueueDequeue"
  description = "Nitric Queue receive access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = [
      "Microsoft.Storage/storageAccounts/queueServices/queues/read"
    ]
    data_actions = [
      "Microsoft.Storage/storageAccounts/queueServices/queues/messages/read",
      "Microsoft.Storage/storageAccounts/queueServices/queues/messages/delete"
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "kvstore_write" {
  name        = "${var.stack_id}-KvStoreWrite"
  description = "Nitric KeyValue write access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.Storage/storageAccounts/tableServices/tables/entities/write",
      // Delete is required for upserting
      "Microsoft.Storage/storageAccounts/tableServices/tables/entities/delete",

    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "kvstore_delete" {
  name        = "${var.stack_id}-KvStoreDelete"
  description = "Nitric KeyValue write access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.Storage/storageAccounts/tableServices/tables/entities/delete",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "bucket_file_read" {
  name        = "${var.stack_id}-BucketFileRead"
  description = "Nitric Bucket read access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/read"
    ]
    data_actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "bucket_file_put" {
  name        = "${var.stack_id}-BucketFilePut"
  description = "Nitric Bucket put access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/write",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "bucket_file_list" {
  name        = "${var.stack_id}-BucketFileList"
  description = "Nitric Bucket put access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/read",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "bucket_file_delete" {
  name        = "${var.stack_id}-BucketFileDelete"
  description = "Nitric Bucket delete access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.Storage/storageAccounts/blobServices/containers/blobs/delete",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "topic_publish" {
  name        = "${var.stack_id}-TopicPublish"
  description = "Nitric Topic publish access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = [
        "Microsoft.EventGrid/topics/read",
		"Microsoft.EventGrid/topics/*/write",
    ]
    data_actions = [
      "Microsoft.EventGrid/events/send/action",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}


resource "azurerm_role_definition" "secret_access" {
  name        = "${var.stack_id}-SecretAccess"
  description = "Nitric Secret access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = []
    data_actions = [
      "Microsoft.KeyVault/vaults/secrets/getSecret/action",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

resource "azurerm_role_definition" "secret_put" {
  name        = "${var.stack_id}-SecretPut"
  description = "Nitric Secret put access"
  scope       = "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"

  permissions {
    actions = [
        "Microsoft.KeyVault/vaults/secrets/write"
    ]
    data_actions = [
      "Microsoft.KeyVault/vaults/secrets/setSecret/action",
    ]
    not_actions = []
  }

  assignable_scopes = [
    "/subscriptions/${var.subscription_id}/resourceGroups/${var.resource_group_name}"
  ]
}

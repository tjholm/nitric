terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.2"
    }
  }
}
provider "docker" {
  registry_auth {
    address  = var.registry_server
    username = var.registry_user
    password = var.registry_pass
  }
}

locals {
  # The name of the ECR repository
  repository_url = "${var.registry_server}/${var.name}"
}

# Tag the provided docker image with the ECR repository u     rl
resource "docker_tag" "tag" {
  source_image = var.image
  target_image = local.repository_url
}

# Push the tagged image to the ECR repository
resource "docker_registry_image" "push" {
  name = local.repository_url
  triggers = {
    source_image_id = docker_tag.tag.source_image_id
  }
}

# Create a random password for the event token for services
resource "random_password" "event_token" {
  length           = 32
  special          = false
  override_special = ["/", "+", "=", "@", "?", "&", "%", "#", "!", "-", "_"]
}

data "azuread_client_config" "current" {}

resource "azuread_service_principal" "sp" {
  client_id                    = var.application_client_id
  app_role_assignment_required = false
  owners                       = [data.azuread_client_config.current.object_id]
}

# create a new app role assignment
resource "azuread_app_role_assignment" "role_assignment" {
  app_role_id         = "4962773b-9cdb-44cf-a8bf-237846a00ab7"
  principal_object_id = azuread_service_principal.current.object_id
  resource_object_id  = azuread_service_principal.sp.id
}

resource "azuread_service_principal_password" "sp_password" {
  service_principal_id = azuread_service_principal.sp.id
}

// Create a new role assignement for each of the RoleDefinitions
// Leave thease out for now

# Create a new container app

resource "azurerm_container_app" "container_app" {
  name                         = var.name
  container_app_environment_id = var.container_app_environment_id
  resource_group_name          = var.resource_group_name
  revision_mode                = "Single"

  ingress {
    external_enabled = true
    target_port = 9001

    traffic_weight {
      percentage = 100
    }
  }

  registry {
    server   = var.registry_server
    username = var.registry_username
    password_secret_name = "pwd"
  }

  dapr {
    app_id = var.name
    app_port = 9001
    app_protocol = "http"
  }

  secret {
    name = "pwd"
    value = var.registry_password
  }

  secret {
    name = "client_id"
    value = azuread_service_principal.sp.client_id
  }

  secret {
    name = "tenant-id"
    value = azuread_service_principal.sp.application_tenant_id
  }

  secret {
    name = "client_secret"
    value = azuread_service_principal_password.sp_password.value
  }

  template {
    container {
      name   = var.name
      image  = "mcr.microsoft.com/azuredocs/containerapps-helloworld:latest"
      # TODO: Configure memory and cpu
      cpu    = 0.25
      memory = "0.5Gi"
    }
  }
}

# 
resource "azapi_resource_action" "my_app_auth" {
  depends_on = [azurerm_container_app.my_app]

  type        = "Microsoft.App/containerApps/authConfigs@2024-03-01"
  resource_id = "${azurerm_container_app.container_app.id}/authConfigs/current"
  method      = "PUT"

  body = jsonencode({
    location = azurerm_container_app.container_app.location
    properties = {
      globalValidation = {
        unauthenticatedClientAction = "Return401"
      }
      identityProviders = {
        azureActiveDirectory = {
          enabled = true
          registration = {
            clientId                = var.client_id
            clientSecretSettingName = "client-secret"
            openIdIssuer            = "https://sts.windows.net/${var.tenant_id}/v2.0"
          }
          validation = {
            allowedAudiences = [var.managed_user_client_id]
          }
        }
      }
      platform = {
        enabled = true
      }
    }
  })
}
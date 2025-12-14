resource "azurerm_container_app" "hangar_app" {
  name                         = "hangar-api-${var.environment_id}"
  container_app_environment_id = azurerm_container_app_environment.hangar_app_env.id
  resource_group_name          = azurerm_resource_group.hangar_rg.name
  revision_mode                = "Single"

  template {
    min_replicas = 1
    max_replicas = 3

    container {
      name   = "hangarapp-${var.environment_id}"
      image  = "mcr.microsoft.com/k8se/quickstart:latest"
      cpu    = 0.25
      memory = "0.5Gi"
    }
  }

  ingress {
    allow_insecure_connections = false
    external_enabled           = true
    target_port                = 8080

    traffic_weight {
      percentage      = 100
      label           = "primary"
      latest_revision = true
    }
  }

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

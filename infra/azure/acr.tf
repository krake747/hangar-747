# Azure Container Registry
# A managed Docker registry service for storing and managing container images.
# Required to store Docker images built in CI/CD pipelines for deployment.
# https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/container_registry

resource "azurerm_container_registry" "hangar_acr" {
  name                          = "hangarregistry${var.environment_id}"
  resource_group_name           = azurerm_resource_group.hangar_rg.name
  location                      = azurerm_resource_group.hangar_rg.location
  sku                           = "Basic"
  admin_enabled                 = true
  public_network_access_enabled = true

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

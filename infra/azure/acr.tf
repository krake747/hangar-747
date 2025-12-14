resource "azurerm_container_registry" "krake_acr" {
  name                          = "krakeregistry${var.environment_id}"
  resource_group_name           = azurerm_resource_group.krake_rg.name
  location                      = azurerm_resource_group.krake_rg.location
  sku                           = "Basic"
  admin_enabled                 = true
  public_network_access_enabled = true

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

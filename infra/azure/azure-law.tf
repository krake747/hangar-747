resource "azurerm_log_analytics_workspace" "hangar_law" {
  name                = "hangarlaw${var.environment_id}"
  location            = azurerm_resource_group.hangar_rg.location
  resource_group_name = azurerm_resource_group.hangar_rg.name
  sku                 = "PerGB2018"
  retention_in_days   = 30

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

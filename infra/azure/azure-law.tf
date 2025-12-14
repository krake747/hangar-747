resource "azurerm_log_analytics_workspace" "krake_law" {
  name                = "krakelaw${var.environment_id}"
  location            = azurerm_resource_group.krake_rg.location
  resource_group_name = azurerm_resource_group.krake_rg.name
  sku                 = "PerGB2018"
  retention_in_days   = 30

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

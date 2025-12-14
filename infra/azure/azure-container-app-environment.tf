resource "azurerm_container_app_environment" "krake_app_env" {
  name                       = "krakeappenv${var.environment_id}"
  location                   = azurerm_resource_group.krake_rg.location
  resource_group_name        = azurerm_resource_group.krake_rg.name
  log_analytics_workspace_id = azurerm_log_analytics_workspace.krake_law.id

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

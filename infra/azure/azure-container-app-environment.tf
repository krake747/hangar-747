resource "azurerm_container_app_environment" "hangar_app_env" {
  name                       = "hangarappenv${var.environment_id}"
  location                   = azurerm_resource_group.hangar_rg.location
  resource_group_name        = azurerm_resource_group.hangar_rg.name
  log_analytics_workspace_id = azurerm_log_analytics_workspace.hangar_law.id

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

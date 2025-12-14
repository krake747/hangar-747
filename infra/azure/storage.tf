resource "azurerm_storage_account" "terraform_state" {
  name                     = "krakeiacfiles"
  resource_group_name      = azurerm_resource_group.reference_rg.name
  location                 = var.region_location
  account_tier             = "Standard"
  account_replication_type = "LRS"
  account_kind             = "StorageV2"

  blob_properties {
    versioning_enabled = true
    delete_retention_policy {
      days = 7
    }
  }

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

resource "azurerm_storage_container" "terraform_state" {
  name                  = "terraform"
  storage_account_id    = azurerm_storage_account.terraform_state.id
  container_access_type = "private"
}

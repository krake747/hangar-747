resource "azurerm_resource_group" "krake_rg" {
  name     = "krake-rg"
  location = var.region_location

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

resource "azurerm_resource_group" "reference_rg" {
  name     = "reference-rg"
  location = var.region_location

  tags = {
    environment = var.environment_id
    src         = var.src_key
  }
}

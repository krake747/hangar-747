# Azure Resource Groups
# Resource groups are fundamental containers in Azure that hold related resources.
# We use them to organize resources by environment and manage access/policies collectively.
# https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/resource_group

resource "azurerm_resource_group" "hangar_rg" {
  name     = "hangar-rg"
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

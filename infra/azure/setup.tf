terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.55.0"
    }
  }

  # Initial setup required: Comment out this backend block, run 'tofu apply' to create storage,
  # then uncomment and run 'tofu init' to migrate state to Azure Blob Storage.
  backend "azurerm" {
    resource_group_name  = "reference-rg"
    storage_account_name = "krakeiacfiles"
    container_name       = "terraform"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
}

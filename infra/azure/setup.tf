# Terraform Backend and Provider Configuration
# Configures the backend for state storage in Azure Blob Storage and sets up the Azure provider.
# Ensures state is stored remotely for collaboration and CI/CD.
# https://www.terraform.io/language/settings/backends/azurerm
# https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs

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
    storage_account_name = "hangariacfiles"
    container_name       = "terraform"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}

  subscription_id = var.subscription_id
}

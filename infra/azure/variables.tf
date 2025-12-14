# Variable Definitions
# Defines input variables for the Terraform configuration.
# Allows customization of resource names and settings per environment.

variable "environment_id" {
  type        = string
  description = "Environment identifier"
  default     = "dev"
}

variable "src_key" {
  type        = string
  description = "Infrastructure source"
  default     = "terraform"
}

variable "subscription_id" {
  type        = string
  description = "Azure Subscription ID"
}

variable "region_location" {
  type        = string
  description = "Azure region for the resources"
  default     = "West Europe"

  validation {
    condition     = contains(["West Europe", "North Europe", "France Central"], var.region_location)
    error_message = "Invalid Azure region."
  }
}

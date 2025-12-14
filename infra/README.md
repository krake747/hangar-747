# Hangar 747 Infra

This repository contains infrastructure as code configurations for deploying cloud resources.

## Azure Infrastructure

The `azure/` folder contains Terraform (OpenTofu) configurations for deploying Azure services, including
container apps, storage, and more. For details on how to deploy Azure services, see [azure/README.md](azure/README.md).

## CI/CD Pipeline

The release-and-deploy workflow automates the deployment of the managed API resource. It runs in the CI/CD pipeline to deploy
the following resource: [Web Api](https://krake-api-dev.mangoplant-23fd7721.westeurope.azurecontainerapps.io/)

### Deploy Infrastructure to Azure Workflow

This workflow automates the provisioning and management of Azure cloud resources using OpenTofu.
It deploys infrastructure components such as container apps, storage accounts, and networking resources
defined in the `infra/azure/` directory. Triggered by changes to infrastructure code on the main branch,
it ensures consistent and repeatable deployments.
This helps me build foundational knowledge in DevOps practices by demonstrating automated infrastructure management,
eliminating local environment inconsistencies, and letting CI/CD pipelines handle all deployments.

### Release and Deploy Workflow

This workflow manages the end-to-end release process for the .NET API application, from versioning to
production deployment. It handles semantic versioning, builds and tests the application,
packages it into a Docker image, and deploys it to Azure Container Apps. Manually triggered via workflow dispatch,
it supports patch, minor, and major version bumps. This reinforces my DevOps foundations through automated
release management, preventing local changes from affecting production, and ensuring CI/CD pipelines
manage the entire deployment lifecycle.

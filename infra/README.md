# Hangar 747 Infra

This repository contains infrastructure as code configurations for deploying cloud resources.

## Azure Infrastructure

The `azure/` folder contains Terraform (OpenTofu) configurations for deploying Azure services, including
container apps, storage, and more. For details on how to deploy Azure services, see [azure/README.md](azure/README.md).

## CI/CD Pipeline

The template workflow automates the personal IAC for a private project. It runs in the CI/CD pipeline to deploy
the following resource: [Web Api](https://krake-api-dev.mangoplant-23fd7721.westeurope.azurecontainerapps.io/)

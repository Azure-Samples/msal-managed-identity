# Managed Identity with Microsoft Authentication Library (MSAL)

This repository contains samples that show how to use [Azure Managed Identity](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview) with Microsoft [Authentication Library](https://learn.microsoft.com/entra/msal).

## Current samples

| Sample | Platform | Build status | Description |
|:-------|:---------|:-------------|:------------|
| [`ms-activedirectory-managedidentity`](src/dotnet/README.md) | .NET | | This sample showcases how to acquire a secret from an Azure Key Vault using the Microsoft identity platform. It shows you how to use the [managed identity for app service](https://learn.microsoft.com/azure/app-service/overview-managed-identity) and acquire a token for an Azure Key Vault resource. | 
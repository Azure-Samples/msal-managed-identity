# Managed Identity with Microsoft Authentication Library (MSAL)

This repository contains samples that show how to use [Azure Managed Identity](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview) with [Microsoft Authentication Library](https://learn.microsoft.com/entra/msal).

## Current samples

| Sample | Platform | Build status | Description |
|:-------|:---------|:-------------|:------------|
| [`ms-activedirectory-managedidentity`](src/dotnet/README.md) | .NET | [![.NET Build](https://github.com/Azure-Samples/msal-managed-identity/actions/workflows/ms-activedirectory-managedidentity-build.yml/badge.svg)](https://github.com/Azure-Samples/msal-managed-identity/actions/workflows/ms-activedirectory-managedidentity-build.yml) | This sample showcases how to acquire a secret from an Azure Key Vault using the Microsoft identity platform. It shows you how to use the [managed identity for app service](https://learn.microsoft.com/azure/app-service/overview-managed-identity) and acquire a token for an Azure Key Vault resource. |

## Additional resources

* [Announcing Microsoft Authentication Library for .NET 4.54.0, with General Availability of Managed Identity APIs](https://devblogs.microsoft.com/identity/msal-net-managed-identity-ga/)
* [Managed identity with MSAL.NET](https://learn.microsoft.com/entra/msal/dotnet/advanced/managed-identity)

## Authors

* [@gladjohn](https://github.com/@gladjohn)
* [@neha-bhargava](https://github.com/@neha-bhargava)

## Get support

If you found a bug or want to suggest a new (for example, a new feature, use case, or sample), please [submit an issue](/issues).

If you have questions, comments, or need help with code, we're here to help - join us on Stack Overflow at the [`azure-ad-msal`](https://stackoverflow.com/questions/tagged/azure-ad-msal) tag.

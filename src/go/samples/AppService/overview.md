---
services: active-directory
platforms: go
author: andrewohart
level: 400
client: Azure functions
service: Azure Key Vault
endpoint: Microsoft identity platform
page_type: sample
languages:
  - go  
products:
  - azure
  - azure-active-directory  
  - go
  - managed identity
  - key vault
description: "This sample showcases how to use the App Service Source and develop an Azure function that gets a secret from a key Vault using Managed Identities."
---
# App Service

This sample showcases the [App Service source](https://learn.microsoft.com/en-us/azure/app-service/) and how to acquire a secret from an [Azure Key Vault](https://vault.azure.net) using Azure Managed Identity  
We will cover [Azure Functions](https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows) and how they can be used to return a token from the key vault

## Topology

![Overview](../../images/call-kv.png)

### Scenario

You want to access an Azure Key Vault secret from a function deployed in Azure. And, you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An IDE of your choice, for example [Visual Studio Code](https://code.visualstudio.com/download)
- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.

### Step 1: Create & Publish your function

[Deploy your function](https://learn.microsoft.com/en-gb/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cmacos) using an IDE of your choice, for example Visual Studio Code

### Step 2:  Modify the Key Vault URI and Secret name values in the code

Following are the changes you need to make:

- In the [handler.go](handler.go) file under the getSecretFromAzureVault method modify the following values,

```go
  keyVaultUri := "your-key-vault-uri"
  secretName := "your-secret-name"
```

- Change these to match your key vault uri and secret name. These can be found in the following locations:

1. Key Vault URI - In your Azure home page, go to your key vault, on the Overview page our key vault URI can be found under **'Essentials'**
1. Secret Name - On the Key Vault Overview page, go to the Panel on the left and expand the **'Objects'** dropdown  
Click into **'Secrets'**  
Click into the secret you want to use  
Click ont the version you would like to use  
Copy the part after the key vault URI and use that as your secret name  

## After you deploy the sample to Azure

There are few important settings you need to change for this sample to work:

### Enable managed identity on the function

- After you publish the function to Azure, go to your resource in the [Azure Portal](https://portal.azure.com/)
- Select the `Identity` blade of the function
- [Enable the System Assigned managed identity](https://learn.microsoft.com/azure/azure-functions/functions-identity-access-azure-sql-with-managed-identity#enable-system-assigned-managed-identity-on-azure-function) of the resource.

### Assign Azure roles using the Azure portal

Azure role-based access control (Azure RBAC) is the authorization system you use to manage access to Azure resources. To grant access, you assign roles to users, groups, service principals, or managed identities at a particular scope. This [article](https://learn.microsoft.com/azure/role-based-access-control/role-assignments-portal) describes how to assign roles using the Azure portal.

You will need to authorize the managed identity resource to access the vault.

![RBAC](../../images/rbac.png)

## Launch the function

To launch the function you can use the following:

1. {your host}/api/AcquireTokenMsi - to acquire a token for system assigned managed identity
2. {your host}/api/AcquireTokenMsi?userAssignedClientId=<client id of the user assigned managed identity> - to acquire a token for a user assigned managed identity.
3. {your host}/api/AcquireTokenMsi?userAssignedResourceId=<resource id of the user assigned managed identity> - to acquire a token for a user assigned managed identity.

> **Note**
> Did the sample not work for you as expected? Did you encounter issues trying this sample? Then please reach out to us using the [GitHub Issues](https://github.com/Azure-Samples/msal-managed-identity/issues) page.

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
description: "This sample showcases how to use the Cloud Shell Source"
---
# Cloud Shell

This sample shows you how you can build and run MSAL GO using [Cloud Shell](https://learn.microsoft.com/en-us/azure/cloud-shell/overview) as your Managed Identity Source.  
This source can be tested entirely through the Azure Portals built in terminal

## Scenario

You want to access an Azure Key Vault secret from a Cloud Shell source, and you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.

### Step 1:  Access Cloud Shell

- Access your [Azure Portal](https://portal.azure.com/#home)
- At the top of the screen you can see the terminal logo, beside the search button. Click it
- This will open the CloudShell terminal

### Step 2:  Setup Environment for Cloud Shell

- Cloud Shell is similar to PowerShell, so similar commands should work
- For example you can call

```Shell
Get-ChildItem Env: 
```

- This will show environment variables. You can observe **MSI_ENDPOINT** is contained here which is used for Cloud Shell.

## Running the sample

- Run the following to get the MSAL GO sdk

```Shell
git clone https://github.com/Azure-Samples/msal-managed-identity.git
```

- Then run this to access the sample app directory

```Shell
cd ./microsoft-authentication-library-for-go/apps/tests/devapps/managedidentity
```

- Run the following to execute the sample app

```Shell
go run .
```

- You should see a token expiry date printed if successful

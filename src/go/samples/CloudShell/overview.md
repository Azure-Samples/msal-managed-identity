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

- You may need to do an update before installing go, so run the following commands to do that, and then install Golang

```bash
sudo apt-get update
sudo apt-get install golang
```

- If asked are you sure you want to install, say yes
- Create a new GO Project using the following

```bash
go mod init <ProjectName>
go get github.com/AzureAD/microsoft-authentication-library-for-go
touch main.go
```

- Modify the ***main.go*** file by using the following commands if using terminal, otherwise just open the file and modify as needed

```bash
vi main.go
```

- Press 'I' and 'Enter' to enter ***Insert*** mode
- You can then use the arrow keys to navigate and modify the code as needed, such as changing the **<ProjectName>** etc

``` go
package <ProjectName>
 
import (
    mi "github.com/AzureAD/microsoft-authentication-library-for-go/apps/managedidentity"
)
 
func main() {
    client, err := mi.New(mi.SystemAssigned())
    if err != nil {
        log.Fatal(err)
    }
    result, err := client.AcquireToken(context.TODO(), "https://management.azure.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("token expire at : ", result.ExpiresOn)
}
 
```

- When you are finished, press ***Escape*** and then type ***:wq*** and press ***Enter*** to save the file
- Run the sample app

```bash
go run .
```

- You should see a token expiry date printed if successful

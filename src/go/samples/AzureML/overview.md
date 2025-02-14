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
description: "This sample showcases how to use the Azure Machine Learning Source"
---
# Azure Machine Learning

This sample shows you how you can build and run MSAL GO using [Azure Machine Learning](https://learn.microsoft.com/en-us/azure/machine-learning/overview-what-is-azure-machine-learning) as your Managed Identity Source.  
We will talk through creating some Azure resources that we will use for the Azure ML source using [MSAL (Microsoft Authentication Library) for GO](https://github.com/AzureAD/microsoft-authentication-library-for-go) to obtain an access token.  

## Scenario

You want to access an Azure Key Vault secret from a Azure Machine Learning source, and you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An IDE of your choice, for example [Visual Studio Code](https://code.visualstudio.com/download)
- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.

### Step 1:  Clone or download this repository

From your terminal of choice:

```bash
git clone https://github.com/Azure-Samples/msal-managed-identity.git
```

or download and extract the repository `.ZIP` file.

The GO sample is located in the [`/src/go/devapps/managedidentity`](https://github.com/AzureAD/microsoft-authentication-library-for-go/blob/a0bb7862bd71c187c09214b1efa20016410d0824/apps/tests/devapps/managedidentity/managedidentity_sample.go) folder.

### Step 2:  Setup Environment for Azure ML

Open your terminal of choice
You will want to install Azure CLI, see instructions [here](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)
If you want to login to your existing Azure account you can do so with the following command:

```bash
az login
```

From here you can run the following

```bash
az ml workspace create --name whatevernameforyourworkspace -g yourresourcegroupname
```

Where '*whatevernameforyourworkspace*' is a name for your workspace, and '*yourresourcegroupname*' is a resource group you have created already on Azure
When creating your Resource Group on Azure, you can also go ahead and create a Managed Identity resource if you wish
Follow the next few steps to setup your Azure ML environment:

- Go to your [Azure homepage](https://portal.azure.com/#home)
- Under your created resources, find the resource group you created earlier
- Click into the Resource Group
- Find the workspace you created earlier and click into it
- Under the essentials section you will see '**Studio Web URL**', open this URL
- Under '**Manage**' on the left hand side, click '**Compute**'
- Create the cheapest Virtual Machine you can
- From here you can select System or User Assigned Managed Identity, for now lets use System Assigned
- You may or may not need to start the VM from the Compute page

## Running the sample

Now that you have set up the Machine Learning environment, we can begin using it

- While still in your **Studio Web** instance, on the left hand side click **Notebooks**, which is under the **Authoring** section
- Click on the terminal button
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

- Currently in MSAL GO, we default to using System Assigned. If you want to try it for User Assigned, see steps below

### Modifying Sample to run User Assigned Managed Identity

- Modify the ***managedidentity_sample.go*** file

```go
func runIMDSUserAssigned() {
    miUserAssigned, err := mi.New(mi.UserAssignedClientID("YOUR_CLIENT_ID"))
    if err != nil {
        log.Fatal(err)
    }
    result, err := miUserAssigned.AcquireToken(context.TODO(), "https://management.azure.com")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("token expires at: ", result.ExpiresOn)
}
```

- For **YOUR_CLIENT_ID**, enter the client ID associated with the Managed Identity resource you have previously created on Azure. It is found under the *Essentials* section of the Managed Identity resource
- Once you have modified this, we need to update the Azure ML VM
- Go back to **Compute** and click on the VM name we created previously
- On the right hand side of the **Compute** section we will see a Managed Identity section
- Click on the small edit icon
- You can swap between system and user assigned here
- When doing User Assigned, you can chose the Managed Identity resource we created previously

> **Note**
> Did the sample not work for you as expected? Did you encounter issues trying this sample? Then please reach out to us using the [GitHub Issues](https://github.com/Azure-Samples/msal-managed-identity/issues) page.
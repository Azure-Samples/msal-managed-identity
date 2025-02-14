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
description: "This sample showcases how to use the IMDS Source"
---
# IMDS (Instance Metadata Service)

This sample shows you how you can build and run MSAL GO using [IMDS](https://learn.microsoft.com/en-us/azure/virtual-machines/instance-metadata-service) as your Managed Identity Source.  
We will talk through creating some Azure resources that we will use for the IMDS source using [MSAL (Microsoft Authentication Library) for GO](https://github.com/AzureAD/microsoft-authentication-library-for-go) to obtain an access token.  
IMDS acts as the default source for Managed Identity, and we default to this if no other Source is to be used

## Scenario

You want to access an Azure Key Vault secret from a IMDS source, and you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An IDE of your choice, for example [Visual Studio Code](https://code.visualstudio.com/download)
- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.

### Setup Virtual Environment

To test locally we will require a virtual machine, which we can set up in [Azure Portal](https://portal.azure.com/#home)
The next few steps will go over each step of the process

#### Setup Resource Group

We need to set up a resource group that our virtual machine will use
The following steps are very detailed as it is what worked during testing, but you may need to alter on your end

1. On the Azure home screen linked above, you should see **'Resource Groups'**, click this
2. Click **'Create'**
3. Select your Subscription, and below that enter a Resource Group name, such as ***'go-lang-rg'***
4. Select a region, for example **'West Europe'**
5. At the bottom of the screen click **'Review + Create'**
6. At the bottom of the screen click **'Create'**
7. Your resource group is created, you can now go back to the Azure homepage

#### Setup Virtual Machine

Next, we need to set up a Virtual Machine for IMDS to run on. Make sure you are on the Azure Homepage

- Click **'Create a Resource'**
- Under **'Virtual Machine'** click **'Create'**
- Select the same Subscription you did for the Resource Group creation
- Under Resource Group you should see the one you created prior, in this example, ***'go-lang-rg'***. Select it
- Create a name for your Virtual Machine, such as ***'go-lang-machine'***
- Select the region, such as West Europe
- Most of the options can be kept the same

```bash
- Availability Options = Availability Zone
- Zone Options = Self Selected Zone
- Security Type = Trusted Launch Virtual Machines
- Image = Ubuntu Server 24.04 LTS -x64 Gen2
- VM Architecture = x64
- Run with Azure Spot Discount = Off
- Size = Standard_D2s_v3 - 2 vcpus, 8 GiB memory
- Enable Hibernation = Off
- Authentication Type = SSH Public Key
- SSH Public Key Source = Generate New Key Pair
- SSH Key Type = RSA SSH Format
- Public Inbound Ports = Allow Selected Ports
- Select Inbound Ports = SSH(22)
```

- For username, enter whatever you want, for example ***'go-lang-machine'***
- For Key Pair Name, set it to whatever you want, i.e ***'go-lang-machine-key'***
- At the bottom of the screen click on **'Review + Create'**
- Click **'Create'** at the bottom of the next page
- You will see a popup to Generate a New Key Pair
- Click **'Download private key and create resource'**
- Once downloaded it should redirect to a page that shows the Virtual Machine deployment. When it completes you can go back to the Home Screen
- To do the following steps you need to ensure you Virtual Machine is running. On the homepage you can see your virtual machine i.e ***'go-lang-machine'***, click on it
- In the **'Overview'** of the virtual machine, you can see if it is started or not. If it is not started click on **'Start'**
- Go back to the homepage

### Setup Local Machine

The next step involves using SSH and setting up the repo on the Virtual Machine
In this example it is done using Mac, if using Windows just make the required adjustments

- You should have the private key downloaded using the name you set previously, i.e ***'go-lang-machine-key'***. In this example the key is saved in the Downloads folder
- Open up Terminal and run the following command, using the key name name you setup prior

```bash
chmod 700 go-lang-machine-key.pem
```

- On your Azure home page, click into your Virtual Machine that you created prior, i.e ***'go-lang-machine'***
- In the left hand panel click expand the **'Connect'** section, and click on **'Connect'**, ensure your virtual machine is started
- Select the **'Native SSH'** option
- In the panel that opens to the right, in section 3 you can copy the SSH command i.e `ssh -i ~/Downloads/KEY-NAME.pem VIRTUAL-MACHINE-NAME@PUBLIC-IP-ADDRESS`
- In Terminal on your local machine, run the command you copied
- It will ask you if you are sure you want to continue connecting, type **'yes'**, this will add your public IP to the list of known hosts
- We should now be connected to the VM via SSH, if not, run the command again
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

- You should see any changes be committed to the SSH instance of the library, and receive some error along the lines of **"Identity not found"**
- The next steps will talk through running System Assigned and User Assigned

### Run System Assigned Test

- From the Azure homepage click on your virtual machine
- In the left hand menu click on Security, and then Identity
- You should see two tabs, System Assigned and User Assigned. Click on System Assigned if not already accepted
- Click Status so it is **'On'**
- Click Save and then Yes
- Wait for it to finish
- Run the sample app from managed identity directory we entered earlier

```bash
go run .
```

- You should see **'token expire at :  some expiry date'**, where **'some expiry date'** is an expiry that is not all 0's, i.e
`2024-09-26 22:05:11.532734044 +0000 UTC m=+86400.490900710`

### Run User Assigned Test

- From Azure homepage click on **'Create a resource'**
- Search for **'Managed Identity'**
- You should see **'User Assigned Managed Identity'**, under it click **'Create'**
- Under **'Create'** click on **'User Assigned Managed Identity'**
- Select your subscription
- Select the resource group you created earlier
- Select your region, i.e West Europe
- Put in a name i.e ***'go-lang-mi'***
- Click **'Review + Create'** at the bottom of the page
- Click on **'Create'** at the bottom
- When it is deployed go back to the Azure homepage
- Click on the virtual machine you created earlier
- In the left hand panel click on **'Security'**, in the expanded menu click on **'Identity'**
- At the top select the **'User Assigned'** tab
- Click on **'Add'**
- When it has deployed click into the managed identity in the User Assigned tab
- Copy the client ID
- In your local instance of **'microsoft-authentication-library-for-go'**, open `managedidentity_sample.go`
- Change the following:

```go
miSystemAssigned, err := mi.New(mi.SystemAssigned())
```

to be

```go
miUserAssigned, err := mi.New(mi.UserAssignedClientID(CLIENT_ID_YOU_COPIED))
```

- Update anything that was previously `miSystemAssigned`, to be `miUserAssigned`
- Run the sample app from managed identity directory we entered earlier

```bash
go run .
```

- You should see **'token expire at :  some expiry date'**, where **'some expiry date'** is an expiry that is not all 0's, i.e
`2024-09-26 22:05:11.532734044 +0000 UTC m=+86400.490900710`

## Useful command for local testing

This command first synchronizes the local microsoft-authentication-library-for-go directory (including code changes), with the corresponding directory on a remote virtual machine using rsync. After the synchronization, it connects to the remote machine via SSH and runs the go application in the correct directory.
This is useful when the developer is not working on the server machine itself

```bash
rsync -avz -e "ssh -i PATH_TO_YOUR_PEM_FILE.pem" PATH_TO_THE_GO_LIB/microsoft-authentication-library-for-go/VIRTUAL-MACHINE-NAME@PUBLIC-IP-ADDRESS:/home/VIRTUAL-MACHINE-NAME/PATH_TO_GO_LIB/microsoft-authentication-library-for-go && ssh -i PATH_TO_YOUR_PEM_FILE.pem VIRTUAL-MACHINE-NAME@PUBLIC-IP-ADDRESS 'cd microsoft-authentication-library-for-go/apps/tests/devapps/managedidentity && go run managedidentity_sample.go'
```

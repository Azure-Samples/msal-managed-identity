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
description: "This sample showcases how to use the Azure Arc Source"
---
# Azure Arc

This sample shows you how you can build and run MSAL GO using [Azure Arc](https://learn.microsoft.com/en-us/azure/azure-arc/servers/overview) as your Managed Identity Source.  
The documentation will cover testing on Windows and Linux as Azure Arc has some special logic for determining the working environment

## Scenario

You want to access an Azure Key Vault secret from an Azure Arc source, and you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An IDE of your choice, for example [Visual Studio Code](https://code.visualstudio.com/download)
- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.
- Windows or Linux VM

## Setting Up Azure Arc on Ubuntu (Linux)

This guide provides a step-by-step approach to set up Azure Arc on Ubuntu, including troubleshooting tips for common issues. You can use your Linux distribution of choice and modify steps as required

### Prerequisites For Ubuntu

1. **Development Environment**: Download Ubuntu or use Hyper-V to create an Ubuntu VM.
2. **Administrative Access**: Ensure you have administrative rights on your terminal.

### Step 1: Download and Set Up Ubuntu

1. **Download Ubuntu** or create a new VM using Hyper-V. For enhanced session mode to enable copy-paste and fix keyboard issues, refer to this [guide](https://www.nakivo.com/blog/install-ubuntu-20-04-on-hyper-v-with-enhanced-session/).
1. **Development Environment**: Download Windows or use Hyper-V to create an Windows VM.

- [Microsoft Hyper-V Documentation](https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/)

### Steps to setup if you are using Hyper-V for Linux

#### Step 1.1: Enable Hyper-V

1. **Open Control Panel**.
2. Navigate to **Programs** > **Turn Windows features on or off**.
3. Check the box for **Hyper-V** and click **OK**.
4. Restart your computer if prompted.

#### Step 1.2: Open Hyper-V Manager

1. Press `Windows + X` and select **Hyper-V Manager**.
2. If it’s not listed, you can search for it in the Start menu.

#### Step 1.3: Create a New Virtual Machine

1. In Hyper-V Manager, select your computer name from the left pane.
2. In the right pane, click on **New** > **Virtual Machine**.
3. Click **Next** on the Wizard.

##### Step 1.3.1: Specify Name and Location

- Enter a name for your VM.
- Optionally, specify a different location to store the VM files.
- Click **Next**.

##### Step 1.3.2: Specify Generation

- Choose between **Generation 1** and **Generation 2** based on your needs (Gen 2 is recommended for newer OS).
- Click **Next**.

##### Step 1.3.3: Assign Memory

- Specify the amount of memory (RAM) for your VM.
- You can also enable **Dynamic Memory** if desired.
- Click **Next**.

##### Step 1.3.4: Configure Networking

- Select a virtual switch to connect your VM to the network. If you don't have one, you'll need to create it.
- Click **Next**.

##### Step 1.3.5: Connect Virtual Hard Disk

- Choose to create a new virtual hard disk, use an existing one, or attach a virtual disk later.
- Specify the name and size of the new disk.
- Click **Next**.

##### Step 1.3.6: Installation Options

- Choose how you want to install the operating system:
  - Install an operating system from a bootable CD/DVD-ROM.
  - Install an operating system from a bootable ISO file.
  - Install an operating system over the network.
- Follow the prompts to select your installation method and click **Next**.

#### Step 1.4: Review and Finish

- Review your settings and click **Finish** to create the VM.

#### Step 1.5: Start the Virtual Machine

1. Right-click on the VM you just created in Hyper-V Manager.
2. Click **Start**.
3. Right-click again and select **Connect** to open the VM console.
4. Follow the on-screen instructions to install your operating system.

5. **Administrative Access**: Ensure you have administrative rights on your Powershell.

6. **Set Up Ubuntu**: Follow the installation prompts until the OS is ready to use.

### Step 2: Troubleshooting Keyboard Issues if Any

If your keyboard doesn't function correctly:

- Use the **On-Screen Keyboard** as a temporary workaround.

### Step 3: Access Azure Portal

1. Navigate to the Azure portal: [Azure Portal](https://portal.azure.com/#view/Microsoft_Azure_ArcCenterUX/ArcCenterInfrastructure.ReactView)
2. Search Azure Arc in the search bar and open it.
3. Click on **Add Resource** in Overview
4. In Machine select **Add a machine** in Add or Create menu.
5. Follow next step.

### Step 4: Set Up Script for Windows or Linux

1. Choose the appropriate setup/server script for your environment (Windows or Linux).
2. Add Resource this from the resource and region
3. Select the operating system **Windows or Linux**
4. **Download the Script**: It’s recommended to download the script instead of copying and pasting it to avoid issues.

### Step 5: Run the Script

1. Open an **Admin Terminal**.
2. Execute the downloaded script.

### Step 6: Installation and running

- Follow this guide for additional setup: [Azure Arc Jumpstart](https://azurearcjumpstart.com/azure_arc_jumpstart/azure_arc_servers/azure/azure_arm_template_linux/).
- **If you encounter an error** about not being able to create Azure on the VM, refer to these resources:

- [Azure Arc on Virtual Machines](https://learn.microsoft.com/en-gb/azure/azure-arc/servers/plan-evaluate-on-azure-virtual-machine)
- Install the Azure CLI:

```bash
curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
```

- Verify installation:

```bash
az --version
```

- **Clone Microsoft Authentication Library**:

```bash
git clone https://github.com/AzureAD/microsoft-authentication-library-for-go.git
cd microsoft-authentication-library-for-go
git switch YOUR_BRANCH_NAME
```

- **Install Go**:

```bash
sudo apt-get update
sudo apt-get install golang
```

- Navigate to the managed identity sample:

```bash
cd apps/tests/devapps/managedidentity
go run managedidentity_sample.go
```

- **Set Environment Variables**:

- Edit your `.bashrc`:

```bash
nano ~/.bashrc
```

- Add the following lines:

```bash
IMDS_IMDS_ENDPOINT=http://localhost:40342
IDENTITY_ENDPOINT=http://localhost:40342/metadata/identity/oauth2/token
```

- Source the changes:

```bash
source ~/.bashrc
```

- **Add Sudo Access for Environment Variables**:

```bash
sudo visudo
```

- Add the line:

```bash
Defaults env_keep += "IMDS_ENDPOINT IDENTITY_ENDPOINT"
```

- **Run the Managed Identity Sample Again**:

```bash
sudo -E go run managedidentity_sample.go
```

- If issues persist, please reset the VM after running the script to verify if it resolves the setup problems.

## Step 7: HIMDS(Hybrid IMDS) support for ARC

1. There is a possibility that the env variables might not be set, in this case we fallback on HIMDS apth
2. For Linux the file is at path : `/opt/azcmagent/bin/himds`
3. This file will be present when you run the azure script.

### Notes on Environment Variables

- **Linux Environment Variables**: Ensure that processes and daemons see the environment variables; refer to [this documentation](https://eng.ms/docs/cloud-ai-platform/azure-core/azure-management-and-platforms/control-plane-bburns/hybrid-resource-provider/azure-arc-for-servers/specs/extension_authoring).

This guide should assist you in setting up Azure Arc on your Ubuntu environment effectively. If you encounter any issues, consult the linked resources for additional support.

## Setting Up Azure Arc on Windows

### Step 1:Prerequisites For Windows

1. **Development Environment**: Download Windows or use Hyper-V to create an Windows VM.

- [Microsoft Hyper-V Documentation](https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/)

#### Step 1.1: Enable Hyper-V

1. **Open Control Panel**.
2. Navigate to **Programs** > **Turn Windows features on or off**.
3. Check the box for **Hyper-V** and click **OK**.
4. Restart your computer if prompted.

#### Step 1.2: Open Hyper-V Manager

1. Press `Windows + X` and select **Hyper-V Manager**.
2. If it’s not listed, you can search for it in the Start menu.

#### Step 1.3: Create a New Virtual Machine

1. In Hyper-V Manager, select your computer name from the left pane.
2. In the right pane, click on **New** > **Virtual Machine**.
3. Click **Next** on the Wizard.

##### Step 1.3.1: Specify Name and Location

- Enter a name for your VM.
- Optionally, specify a different location to store the VM files.
- Click **Next**.

##### Step 1.3.2: Specify Generation

- Choose between **Generation 1** and **Generation 2** based on your needs (Gen 2 is recommended for newer OS).
- Click **Next**.

##### Step 1.3.3: Assign Memory

- Specify the amount of memory (RAM) for your VM.
- You can also enable **Dynamic Memory** if desired.
- Click **Next**.

##### Step 1.3.4: Configure Networking

- Select a virtual switch to connect your VM to the network. If you don't have one, you'll need to create it.
- Click **Next**.

##### Step 1.3.5: Connect Virtual Hard Disk

- Choose to create a new virtual hard disk, use an existing one, or attach a virtual disk later.
- Specify the name and size of the new disk.
- Click **Next**.

##### Step 1.3.6: Installation Options

- Choose how you want to install the operating system:
  - Install an operating system from a bootable CD/DVD-ROM.
  - Install an operating system from a bootable ISO file.
  - Install an operating system over the network.
- Follow the prompts to select your installation method and click **Next**.

#### Step 1.4: Review and Finish

- Review your settings and click **Finish** to create the VM.

#### Step 1.5: Start the Virtual Machine

1. Right-click on the VM you just created in Hyper-V Manager.
2. Click **Start**.
3. Right-click again and select **Connect** to open the VM console.
4. Follow the on-screen instructions to install your operating system.

5. **Administrative Access**: Ensure you have administrative rights on your Powershell.

### Step 2: Access Azure Portal

1. Navigate to the Azure portal: [Azure Portal](https://portal.azure.com/#view/Microsoft_Azure_ArcCenterUX/ArcCenterInfrastructure.ReactView).

### Step 3: Set Up Script for Windows

1. Choose the appropriate setup script for your environment (Windows or Linux).
2. **Download the Script**: It’s recommended to download the script instead of copying and pasting it to avoid issues.

### Step 4: Run the Script

1. Open an **Admin Powershell**.
2. Execute the downloaded script.

### Step 5: Installation and running

- **Install Go Lang**: [Go Lang](https://go.dev/doc/install)

- **Clone Microsoft Authentication Library**

```bash
git clone https://github.com/AzureAD/microsoft-authentication-library-for-go.git
cd microsoft-authentication-library-for-go
git switch YOUR_BRANCH_NAME
```

- **Navigate to the managed identity sample**:

```bash
cd apps/tests/devapps/managedidentity
go run managedidentity_sample.go
```

### Step 6 HIMDS(Hybrid IMDS) support for ARC

1. There is a possibility that the env variables might not be set, in this case we fallback on HIMDS apth
2. For Windows the file is at path : `\AzureConnectedMachineAgent\himds.exe`
3. This file will be present when you run the azure script.

#### Notes on Environment Variables

- **Windows Environment Variables**: To manage environment variables, you can use PowerShell:

```powershell
Remove-Item -Path Env:\IDENTITY_ENDPOINT
Remove-Item -Path Env:\IMDS_ENDPOINT

[System.Environment]::SetEnvironmentVariable("IDENTITY_ENDPOINT", "http://localhost:40342/metadata/identity/oauth2/token", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("IMDS_ENDPOINT", "http://localhost:40342", [System.EnvironmentVariableTarget]::User)
```

This guide should assist you in setting up Azure Arc on your Ubuntu environment effectively. If you encounter any issues, consult the linked resources for additional support.
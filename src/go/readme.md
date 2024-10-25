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
description: "This sample showcases how to develop an Azure function that gets a secret from a key Vault using Managed Identities."
---
# GO - Acquire a secret from an Azure Key Vault using Azure Managed Identity

## About this sample

### Overview

This sample showcases how to acquire a secret from an Azure Key Vault using Azure Managed Identity. It shows you how to use the [Azure Functions](https://learn.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows) and acquire a token for an Azure Key Vault resource.

The sample shows how to use [MSAL (Microsoft Authentication Library) for GO](https://github.com/AzureAD/microsoft-authentication-library-for-go) to obtain an access token for [Azure Key Vault](https://vault.azure.net). Specifically, the sample shows how to retrieve the secret value from a vault.

Finally, the sample also demonstrates how to use the different [types of managed identities](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview#managed-identity-types) to get an access token.

For more information about how the protocols work in this scenario and other scenarios, see [Authentication Scenarios for Azure AD](http://go.microsoft.com/fwlink/?LinkId=394414).

For more information about Managed Identity, please visit the [Managed Identities for Azure Resources homepage](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview).

## Topology

<img alt="Overview" src="./images/call-kv.png" />

### Scenario

You want to access an Azure Key Vault secret from a function deployed in Azure. And, you don't want to worry about managing secrets or app credentials.

## How To Run This Sample

To run this sample, you'll need:

- An IDE of your choice, for example [Visual Studio Code](https://code.visualstudio.com/download)
- An Internet connection
- An Azure account to create, deploy, and manage applications. If you do not have an Azure Account, follow the [instructions](https://azure.microsoft.com/free/) to get a free account.

### Step 1:  Clone or download this repository

From your shell or command line:

```Shell
git clone https://github.com/Azure-Samples/msal-managed-identity.git
```

or download and extract the repository `.ZIP` file.

The GO sample is located in the [`/src/go/devapps/managedidentity`](https://github.com/AzureAD/microsoft-authentication-library-for-go/blob/c5febcbae287a26a0cfedd45f4edeaf3c41ad7dc/apps/tests/devapps/managedidentity/managedidentity_sample.go) folder.

### Step 2:  Modify the Key Vault URI and Secret name values in the code

Following are the changes you need to make:

- In the [`handler.go`](https://github.com/Azure-Samples/msal-managed-identity/blob/main/src/go/sample/AcquireTokenMSI.go) file under the getSecretFromAzureVault method modify the following values, 

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

## Publish your function

[Deploy your function](https://learn.microsoft.com/en-gb/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cmacos) using an IDE of your choice, for example Visual Studio Code

## After you deploy the sample to Azure

There are few important settings you need to change for this sample to work:

### Enable managed identity on the function

- After you publish the function to Azure, go to your resource in the [Azure Portal](https://portal.azure.com/)
- Select the `Identity` blade of the function
- [Enable the System Assigned managed identity](https://learn.microsoft.com/azure/azure-functions/functions-identity-access-azure-sql-with-managed-identity#enable-system-assigned-managed-identity-on-azure-function) of the resource.

### Assign Azure roles using the Azure portal

Azure role-based access control (Azure RBAC) is the authorization system you use to manage access to Azure resources. To grant access, you assign roles to users, groups, service principals, or managed identities at a particular scope. This [article](https://learn.microsoft.com/azure/role-based-access-control/role-assignments-portal) describes how to assign roles using the Azure portal.

You will need to authorize the managed identity resource to access the vault.

<img alt="RBAC" src="./images/rbac.png" />

## Launch the function

To launch the function you can use the following:

1. {your host}/api/AcquireTokenMsi - to acquire a token for system assigned managed identity
2. {your host}/api/AcquireTokenMsi?userAssignedClientId=<client id of the user assigned managed identity> - to acquire a token for a user assigned managed identity.
3. {your host}/api/AcquireTokenMsi?userAssignedResourceId=<resource id of the user assigned managed identity> - to acquire a token for a user assigned managed identity.

> **Note**
> Did the sample not work for you as expected? Did you encounter issues trying this sample? Then please reach out to us using the [GitHub Issues](https://github.com/Azure-Samples/msal-managed-identity/issues) page.

## About the code

Here there's a quick guide to the most interesting authentication-related bits of the sample.

### Acquiring the managed identity token

MSAL GO supports acquiring tokens through the managed identity capability when used with applications running inside Azure infrastructure. You can read more about MSAL GO support for managed identities in the [official documentation](https://learn.microsoft.com/entra/msal/go/advanced/managed-identity).

### Using the access tokens in the app

The `AcquireToken` function in the `managedidentity.go` class demonstrates how to take advantage of the Managed Identity Client for calling Microsoft Key Vault without having to worry about secrets or certificates.

Here is the relevant code:

```go
// Get the access token using MSAL, or an error if there was one 
accessToken, err := miClient.AcquireToken(context.Background(), "https://vault.azure.net")
if err != nil {
    log.Fatalf("failed to acquire token: %v", err)
    return
}

// Create an HttpClient and set the Authorization Header
req, err := http.NewRequest("GET", url, nil)
if err != nil {
    log.Fatalf("Error creating request: %v", err)
}

req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))

// Create the URL and send the request, then read and save the response body, handling errors along the way
url := fmt.Sprintf("%ssecrets/%s?api-version=7.2", keyVaultUri, secretName)
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    log.Fatalf("Error sending request: %v", err)
}
defer resp.Body.Close()

// Combine all received buffer streams into one buffer, and then into a string
var parsedData map[string]interface{}
if err := json.Unmarshal(body, &parsedData); err != nil {
    log.Fatalf("Error parsing JSON: %v", err)
}

// Print the response body which contains the secret value
println(fmt.Sprintf("The secret, %s, has a value of: %s", secretName, string(body)))
```

- First, you need to acquire an access token.
- Once you have the access token, create an HTTP client to call the Key Vault APIs.
- Add the access token as an authorization header to the HTTP client. Set the Authorization header to Bearer followed by the access token.
- Once you have set up the authorization header, you can make calls to Key Vault APIs.  

The terminal will then show you the secret response that contains the value of the secret

## Common Errors

Following are the most common errors you would see if any step was missed during setup:

### An attempt was made to access a socket in a way forbidden by its access permissions. (169.254.169.254:80)

This error indicates that the managed identity endpoint is not reachable. Please refer to the Azure Web App documentation on [how to turn on Managed Identity](https://learn.microsoft.com/azure/azure-app-configuration/howto-integrate-azure-managed-service-identity?pivots=framework-dotnet&tabs=core5x#add-a-managed-identity).

> Causes: Managed identity is not enabled for the Azure Resource.

### Access Denied errors

```json
{
  "error": {
    "code": "Forbidden",
    "message": "The user, group or application 'appid=xyz;oid=xyz;iss=https://sts.windows.net/xyz/' does not have secrets get permission on key vault '<key vault name>;location=xyz'. For help resolving this issue, please see https://go.microsoft.com/fwlink/?linkid=2125287",
    "innererror": {
      "code": "AccessDenied"
    }
  }
}
```

This error indicates that the managed identity service principal was not granted access to the key vault. Please refer to ["Assign a Key Vault access policy"](https://learn.microsoft.com/azure/key-vault/general/assign-access-policy?tabs=azure-portal) for more information.

> Causes: Managed identity resource was not granted access to the Key Vault

## Community Help and Support

Use [Stack Overflow](http://stackoverflow.com/questions/tagged/azure-ad-msal) to get support from the community. Ask your questions on Stack Overflow first and browse existing issues to see if someone has asked your question before.

Make sure that your questions or comments are tagged with `azure-ad-msal`, `go`, and `microsoft-graph`.

If you find a bug in the sample, please raise the issue on [GitHub Issues](/issues).

## Contributing

If you'd like to contribute to this sample, see [our contribution guidelines](https://github.com/Azure-Samples/msal-managed-identity/blob/main/CONTRIBUTING.md).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

## More information

For more information, refer to the [MSAL GO documentation](https://learn.microsoft.com/en-us/entra/msal/go/).

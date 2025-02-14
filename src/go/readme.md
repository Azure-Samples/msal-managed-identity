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
description: "This file is to give a general overview of Managed Identity for MSAL GO and the various sources and samples that can be used"
---
# GO - Samples Overview

This document is to give samples and step by step instructions on how to run the various Managed Identity Source samples

## Overview

MSAL GO supports acquiring tokens through the managed identity capability when used with applications running inside Azure infrastructure.
Samples can all be tested by downloading the [MSAL (Microsoft Authentication Library) for GO](https://github.com/AzureAD/microsoft-authentication-library-for-go)

The Sources with samples are as follows:

- [IMDS](samples/IMDS/overview.md)
- [Azure Arc](samples/AzureArc/overview.md)
- [Azure Machine Learning](samples/AzureML/overview.md)
- [App Service](samples/AppService/overview.md)
- [Cloud Shell](samples/CloudShell/overview.md)
- Service Fabric

 You can read more about MSAL GO support for managed identities in the [official documentation](https://learn.microsoft.com/entra/msal/go/advanced/managed-identity).

## Creating a GO Project for your samples

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

### System and User Assigned Managed Identities

By default, all samples use **System Assigned** Managed Identity.  
**User Assigned** can be assigned in the form of ***Client ID***, ***Object ID*** and ***Resource ID***
You can read more about [types of managed identities here](https://learn.microsoft.com/azure/active-directory/managed-identities-azure-resources/overview#managed-identity-types)

#### Changing to User Assigned Managed Identity

In the case you want to use **User Assigned**, make sure that a Managed Identity Resource has been set up correctly on Azure.  
Then you will need to modify the code:

- Modify the ***main.go*** file
- You can change ***UserAssignedClientID*** to be ***UserAssignedObjectID*** or ***UserAssignedResourceID*** as needed

```go
func main() {
    client, err := mi.New(mi.UserAssignedClientID("YOUR_CLIENT_ID"))
    if err != nil {
        log.Fatal(err)
    }
    result, err := client.AcquireToken(context.TODO(), "https://management.azure.com")
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

### About the code

Then depending on the Source, the Request will be created.
Here there's a quick guide to the most interesting authentication-related bits of the sample.

The `New` function is used to create a new Managed Identity client. You can pass in a System or User Assigned managed identity type here.
From here MSAL attempt to get the token for the request.

```go
// Creates an instance of Managed Identity Client using System Assigned managed identity
miSystemAssigned, err := mi.New(mi.SystemAssigned())
if err != nil {
    log.Fatal(err)
}
```

MSAL will then save to the cache and return the Auth Result.
You can then call the `AcquireToken` function to get a token or throw an error

```go
// Get the access token using MSAL, or an error if there was one, you can see this in the sample app
accessToken, err := miClient.AcquireToken(context.Background(), "https://vault.azure.net")
if err != nil {
    log.Fatalf("failed to acquire token: %v", err)
    return
}
```

The `AcquireToken` function in the `managedidentity.go` class demonstrates how to take advantage of the Managed Identity Client for calling Microsoft Key Vault without having to worry about secrets or certificates.

```go
// In AcquireToken, when claims are empty, we get token from the cache, otherwise acquire a new one
if o.claims == "" {
    storageTokenResponse, err := cacheManager.Read(ctx, c.authParams)
    if err != nil {
        return base.AuthResult{}, err
    }
    ar, err := base.AuthResultFromStorage(storageTokenResponse)
    if err == nil {
        ar.AccessToken, err = c.authParams.AuthnScheme.FormatAccessToken(ar.AccessToken)
        return ar, err
    }
}
```

Then depending on the Source, the Request will be created
Creation of the Request is different per Source, but usually consists of getting the URL from environment variables and adding in the correct Headers

From here we attempt to get the token for the request

```go
tokenResponse, err := c.getTokenForRequest(req, resource)
if err != nil {
    return base.AuthResult{}, err
}
```

The below code shows how a user can print the ***ExpiresOn*** from the returned AuthResult

```go
accessToken, err := miClient.AcquireToken(context.Background(), "https://vault.azure.net")
if err != nil {
    log.Fatalf("failed to acquire token: %v", err)
    return
}
fmt.Println("token expire at : ", accessToken.ExpiresOn)
```

### Common Errors

The following are the most common errors you would see if any step was missed during setup:

#### An attempt was made to access a socket in a way forbidden by its access permissions. (169.254.169.254:80)

This error indicates that the managed identity endpoint is not reachable. Please refer to the Azure Web App documentation on [how to turn on Managed Identity](https://learn.microsoft.com/azure/azure-app-configuration/howto-integrate-azure-managed-service-identity?pivots=framework-dotnet&tabs=core5x#add-a-managed-identity).

> Causes: Managed identity is not enabled for the Azure Resource.

#### Access Denied errors

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

### Community Help and Support

Use [Stack Overflow](http://stackoverflow.com/questions/tagged/azure-ad-msal) to get support from the community. Ask your questions on Stack Overflow first and browse existing issues to see if someone has asked your question before.

Make sure that your questions or comments are tagged with `azure-ad-msal`, `go`, and `microsoft-graph`.

If you find a bug in the sample, please raise the issue on [GitHub Issues](/issues).

### Contributing

If you'd like to contribute to this sample, see [our contribution guidelines](https://github.com/Azure-Samples/msal-managed-identity/blob/main/CONTRIBUTING.md).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

### More information

For more information, refer to the [MSAL GO documentation](https://learn.microsoft.com/en-us/entra/msal/go/).

### Issues

Did the sample not work for you as expected? Did you encounter issues trying this sample? Then please reach out to us using the [GitHub Issues](https://github.com/Azure-Samples/msal-managed-identity/issues) page

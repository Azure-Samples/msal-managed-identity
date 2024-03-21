package com.microsoft.msi.sample;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpResponse;
import java.util.*;

import com.microsoft.aad.msal4j.IAuthenticationResult;
import com.microsoft.aad.msal4j.ManagedIdentityApplication;
import com.microsoft.aad.msal4j.ManagedIdentityId;
import com.microsoft.aad.msal4j.ManagedIdentityParameters;
import com.microsoft.azure.functions.annotation.*;
import com.microsoft.azure.functions.*;

/**
 * Azure Function to acquire a token for managed identity using MSAL Java.
 */
public class AcquireTokenMsi {
    /**
     * This function listens at endpoint "/api/AcquireTokenMsi". Two ways to invoke it using "curl" command in bash:
     * 1. curl -d "HTTP Body" {your host}/api/AcquireTokenMsi
     * 2. curl {your host}/api/AcquireTokenMsi?userAssignedClientId={client id of the user assigned managed identity}
     * 3. curl {your host}/api/AcquireTokenMsi?userAssignedResourceId={resource id of the user assigned managed identity}
     */
    @FunctionName("AcquireTokenMsi")
    public HttpResponseMessage run(
            @HttpTrigger(name = "req", methods = {HttpMethod.GET, HttpMethod.POST}, authLevel = AuthorizationLevel.FUNCTION) HttpRequestMessage<Optional<String>> request,
            final ExecutionContext context) {
        context.getLogger().info("Initializing Acquire Token request.");

        String resource = "https://vault.azure.net";
        String kvUri = "https://<your-key-vault-name>.vault.azure.net/";
        String secretName = "<secret name>";

        // Parse query parameter
        Map<String, String> queryParameters = request.getQueryParameters();
        String userAssignedClientId = queryParameters.get("userAssignedClientId");
        String userAssignedResourceId = queryParameters.get("userAssignedResourceId");

        StringBuilder response = new StringBuilder("Acquire a token using managed identity to use the key vault.\n");

        try {
            ManagedIdentityApplication msiApp = buildManagedIdentityApplication(userAssignedClientId, userAssignedResourceId, response);
            IAuthenticationResult result = msiApp.acquireTokenForManagedIdentity(ManagedIdentityParameters.builder(resource).build()).get();

            context.getLogger().info("Access token acquired.");

            HttpClient httpClient = HttpClient.newHttpClient();

            java.net.http.HttpRequest httpRequest = java.net.http.HttpRequest.newBuilder()
                    .GET()
                    .uri(new URI(kvUri + "/secrets/" + secretName + "?api-version=7.2"))
                    .header("Authorization", "Bearer " + result.accessToken())
                    .build();

            HttpResponse<String> httpResponse = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofString());

            context.getLogger().info("Secret fetch success.");

            response.append(httpResponse.body());
        } catch (Exception exception) {
            context.getLogger().severe("Error occured. " + exception.getCause());
            response.append(exception.getMessage());
        }

        return request.createResponseBuilder(HttpStatus.OK).body(response).build();
    }

    private ManagedIdentityApplication buildManagedIdentityApplication(String userAssignedClientId, String userAssignedResourceId, StringBuilder response) {
        // By default acquire a token for a system assigned managed identity.
        ManagedIdentityId managedIdentityId = ManagedIdentityId.systemAssigned();

        // Acquire a token for user assigned managed identity using client id of the managed identity.
        if (userAssignedClientId != null) {
            response.append("Using user assigned client id.\n");
            managedIdentityId = ManagedIdentityId.userAssignedClientId(userAssignedClientId);
        } else if (userAssignedResourceId != null) {
            response.append("Using user assigned resource id.\n");
            managedIdentityId = ManagedIdentityId.userAssignedResourceId(userAssignedResourceId);
        } else {
            response.append("Using system assigned managed identity\n");
        }

        return ManagedIdentityApplication
                .builder(managedIdentityId)
                .logPii(true)
                .build();
    }
}

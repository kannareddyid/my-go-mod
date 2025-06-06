package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// GetAccessToken retrieves an Azure AD access token using the client credentials grant flow.
func GetAccessToken(clientID, clientSecret, tenantID, scope string) (string, error) {
	// Create a client secret credential
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create credential: %v", err)
	}

	// Request the access token
	token, err := cred.GetToken(context.Background(), azidentity.TokenRequestOptions{
		Scopes: []string{scope},
	})
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	return token.Token, nil
}

func main() {
	// Example usage
	clientID := "your-client-id"
	clientSecret := "your-client-secret"
	tenantID := "your-tenant-id"
	scope := "https://graph.microsoft.com/.default" // Example scope for Microsoft Graph

	token, err := GetAccessToken(clientID, clientSecret, tenantID, scope)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Access Token: %s\n", token)
}
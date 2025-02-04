//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/UpdatePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		"<private-endpoint-connection-name>",
		armsearch.PrivateEndpointConnection{
			Properties: &armsearch.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armsearch.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armsearch.PrivateLinkServiceConnectionStatusRejected.ToPtr(),
				},
			},
		},
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientUpdateResult)
}

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/GetPrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		"<private-endpoint-connection-name>",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientGetResult)
}

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/DeletePrivateEndpointConnection.json
func ExamplePrivateEndpointConnectionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	res, err := client.Delete(ctx,
		"<resource-group-name>",
		"<search-service-name>",
		"<private-endpoint-connection-name>",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionsClientDeleteResult)
}

// x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/ListPrivateEndpointConnectionsByService.json
func ExamplePrivateEndpointConnectionsClient_ListByService() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsearch.NewPrivateEndpointConnectionsClient("<subscription-id>", cred, nil)
	pager := client.ListByService("<resource-group-name>",
		"<search-service-name>",
		&armsearch.SearchManagementRequestOptions{ClientRequestID: nil})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

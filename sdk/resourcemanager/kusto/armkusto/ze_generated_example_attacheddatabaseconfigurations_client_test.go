//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationCheckNameAvailability.json
func ExampleAttachedDatabaseConfigurationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.CheckNameAvailability(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.AttachedDatabaseConfigurationsCheckNameRequest{
			Name: to.StringPtr("<name>"),
			Type: to.StringPtr("<type>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDatabaseConfigurationsClientCheckNameAvailabilityResult)
}

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsListByCluster.json
func ExampleAttachedDatabaseConfigurationsClient_ListByCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.ListByCluster(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDatabaseConfigurationsClientListByClusterResult)
}

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsGet.json
func ExampleAttachedDatabaseConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<attached-database-configuration-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDatabaseConfigurationsClientGetResult)
}

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
func ExampleAttachedDatabaseConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<attached-database-configuration-name>",
		armkusto.AttachedDatabaseConfiguration{
			Location: to.StringPtr("<location>"),
			Properties: &armkusto.AttachedDatabaseConfigurationProperties{
				ClusterResourceID:                 to.StringPtr("<cluster-resource-id>"),
				DatabaseName:                      to.StringPtr("<database-name>"),
				DefaultPrincipalsModificationKind: armkusto.DefaultPrincipalsModificationKind("Union").ToPtr(),
				TableLevelSharingProperties: &armkusto.TableLevelSharingProperties{
					ExternalTablesToExclude: []*string{
						to.StringPtr("ExternalTable2")},
					ExternalTablesToInclude: []*string{
						to.StringPtr("ExternalTable1")},
					MaterializedViewsToExclude: []*string{
						to.StringPtr("MaterializedViewTable2")},
					MaterializedViewsToInclude: []*string{
						to.StringPtr("MaterializedViewTable1")},
					TablesToExclude: []*string{
						to.StringPtr("Table2")},
					TablesToInclude: []*string{
						to.StringPtr("Table1")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AttachedDatabaseConfigurationsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsDelete.json
func ExampleAttachedDatabaseConfigurationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewAttachedDatabaseConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		"<attached-database-configuration-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

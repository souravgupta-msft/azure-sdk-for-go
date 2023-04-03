//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/getActivityInAModule.json
func ExampleActivityClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActivityClient().Get(ctx, "rg", "myAutomationAccount33", "OmsCompositeResources", "Add-AzureRmAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Activity = armautomation.Activity{
	// 	Name: to.Ptr("Add-AzureRmAccount"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Add-AzureRmAccount"),
	// 	Properties: &armautomation.ActivityProperties{
	// 		Description: to.Ptr("The Add-AzureRmAcccount cmdlet adds an authenticated Azure account to use for Azure Resource Manager cmdlet requests.\n\nYou can use this authenticated account only with Azure Resource Manager cmdlets. To add an authenticated account for use with Service Management cmdlets, use the Add-AzureAccount or the Import-AzurePublishSettingsFile cmdlet."),
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
	// 		Definition: to.Ptr(""),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
	// 		OutputTypes: []*armautomation.ActivityOutputType{
	// 			{
	// 				Name: to.Ptr("Microsoft.Azure.Commands.Profile.Models.PSAzureProfile"),
	// 				Type: to.Ptr("Microsoft.Azure.Commands.Profile.Models.PSAzureProfile"),
	// 		}},
	// 		ParameterSets: []*armautomation.ActivityParameterSet{
	// 			{
	// 				Name: to.Ptr("SubscriptionId"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("ServicePrincipal"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ServicePrincipal"),
	// 						Type: to.Ptr("System.Management.Automation.SwitchParameter"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("SubscriptionName"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("SubscriptionName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(true),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("User"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("Credential"),
	// 						Type: to.Ptr("System.Management.Automation.PSCredential"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("AccessToken"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("AccessToken"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("AccountId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 			},
	// 			{
	// 				Name: to.Ptr("ServicePrincipalCertificate"),
	// 				Parameters: []*armautomation.ActivityParameter{
	// 					{
	// 						Name: to.Ptr("ApplicationId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("CertificateThumbprint"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("Environment"),
	// 						Type: to.Ptr("Microsoft.Azure.Common.Authentication.Models.AzureEnvironment"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("EnvironmentName"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(false),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("ServicePrincipal"),
	// 						Type: to.Ptr("System.Management.Automation.SwitchParameter"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("TenantId"),
	// 						Type: to.Ptr("System.String"),
	// 						Description: to.Ptr("Specify the feature description"),
	// 						IsDynamic: to.Ptr(false),
	// 						IsMandatory: to.Ptr(true),
	// 						Position: to.Ptr[int64](-2147483648),
	// 						ValidationSet: []*armautomation.ActivityParameterValidationSet{
	// 						},
	// 						ValueFromPipeline: to.Ptr(false),
	// 						ValueFromPipelineByPropertyName: to.Ptr(false),
	// 						ValueFromRemainingArguments: to.Ptr(false),
	// 				}},
	// 		}},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listActivitiesByModule.json
func ExampleActivityClient_NewListByModulePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActivityClient().NewListByModulePager("rg", "myAutomationAccount33", "OmsCompositeResources", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ActivityListResult = armautomation.ActivityListResult{
		// 	Value: []*armautomation.Activity{
		// 		{
		// 			Name: to.Ptr("Add-AzureRmAccount"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Add-AzureRmAccount"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Add-AzureRmAcccount cmdlet adds an authenticated Azure account to use for Azure Resource Manager cmdlet requests.\n\nYou can use this authenticated account only with Azure Resource Manager cmdlets. To add an authenticated account for use with Service Management cmdlets, use the Add-AzureAccount or the Import-AzurePublishSettingsFile cmdlet."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Add-AzureRmEnvironment"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Add-AzureRmEnvironment"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Add-AzureRmEnvironment cmdlet adds endpoints and metadata to enable Azure Resource Manager cmdlets to connect with a new instance of Azure Resource Manager. The built-in environments AzureCloud and AzureChinaCloud target existing public instances of Azure Resource Manager."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Disable-AzureRmDataCollection"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Disable-AzureRmDataCollection"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr(""),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Enable-AzureRmDataCollection"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Enable-AzureRmDataCollection"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr(""),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Get-AzureRmContext"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Get-AzureRmContext"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Get-AzureRmContext cmdlet gets the current metadata used to authenticate Azure Resource Manager requests.\n\nThis cmdlet gets the Active Directory account, Active Directory tenant, Azure subscription, and the targeted Azure environment. Azure Resource Manager cmdlets use these settings by default when making Azure Resource Manager requests."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.713+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Get-AzureRmEnvironment"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Get-AzureRmEnvironment"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Get-AzureRmEnvironment cmdlet gets endpoints and metadata for an instance of Azure services."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Get-AzureRmSubscription"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Get-AzureRmSubscription"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Get-AzureRmSubscription cmdlet gets the subscription ID, subscription name, and home tenant for subscriptions that the current account can access."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Get-AzureRmTenant"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Get-AzureRmTenant"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Get-AzureRmTenant cmdlet gets tenants authorized for the current user."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Remove-AzureRmEnvironment"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Remove-AzureRmEnvironment"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Remove-AzureRmEnvironment cmdlet removes endpoints and metadata information for connecting to a given Azure instance."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Save-AzureRmProfile"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Save-AzureRmProfile"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Save-AzureRmProfile cmdlet saves the current authentication information for use in other PowerShell sessions."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Select-AzureRmProfile"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Select-AzureRmProfile"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Select-AzureRmProfile cmdlet loads authentication information from a file to set the Azure environment and context. Cmdlets that you run in the current session use this information to authenticate requests to Azure Resource Manager."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Set-AzureRmContext"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Set-AzureRmContext"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Set-AzureRmContext cmdlet sets authentication information for cmdlets that you run in the current session. The context includes tenant, subscription, and environment information."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Set-AzureRmEnvironment"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/modules/AzureRM.Profile/activities/Set-AzureRmEnvironment"),
		// 			Properties: &armautomation.ActivityProperties{
		// 				Description: to.Ptr("The Set-AzureRMEnvironment cmdlet sets endpoints and metadata for connecting to an instance of Azure."),
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 				Definition: to.Ptr(""),
		// 				LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-17T17:05:21.697+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

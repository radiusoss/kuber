package sql

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// DatabaseOperationsClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type DatabaseOperationsClient struct {
	ManagementClient
}

// NewDatabaseOperationsClient creates an instance of the DatabaseOperationsClient client.
func NewDatabaseOperationsClient(subscriptionID string) DatabaseOperationsClient {
	return NewDatabaseOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDatabaseOperationsClientWithBaseURI creates an instance of the DatabaseOperationsClient client.
func NewDatabaseOperationsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseOperationsClient {
	return DatabaseOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Cancel cancels the asynchronous operation on the database.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database. operationID is the operation identifier.
func (client DatabaseOperationsClient) Cancel(resourceGroupName string, serverName string, databaseName string, operationID uuid.UUID) (result autorest.Response, err error) {
	req, err := client.CancelPreparer(resourceGroupName, serverName, databaseName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", resp, "Failure sending request")
		return
	}

	result, err = client.CancelResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "Cancel", resp, "Failure responding to request")
	}

	return
}

// CancelPreparer prepares the Cancel request.
func (client DatabaseOperationsClient) CancelPreparer(resourceGroupName string, serverName string, databaseName string, operationID uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/operations/{operationId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CancelSender sends the Cancel request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseOperationsClient) CancelSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CancelResponder handles the response to the Cancel request. The method always
// closes the http.Response Body.
func (client DatabaseOperationsClient) CancelResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ListByDatabase gets a list of operations performed on the database.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from the
// Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of the
// database.
func (client DatabaseOperationsClient) ListByDatabase(resourceGroupName string, serverName string, databaseName string) (result DatabaseOperationListResult, err error) {
	req, err := client.ListByDatabasePreparer(resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure responding to request")
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client DatabaseOperationsClient) ListByDatabasePreparer(resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/operations", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DatabaseOperationsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client DatabaseOperationsClient) ListByDatabaseResponder(resp *http.Response) (result DatabaseOperationListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabaseNextResults retrieves the next set of results, if any.
func (client DatabaseOperationsClient) ListByDatabaseNextResults(lastResults DatabaseOperationListResult) (result DatabaseOperationListResult, err error) {
	req, err := lastResults.DatabaseOperationListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure sending next results request")
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.DatabaseOperationsClient", "ListByDatabase", resp, "Failure responding to next results request")
	}

	return
}

// ListByDatabaseComplete gets all elements from the list without paging.
func (client DatabaseOperationsClient) ListByDatabaseComplete(resourceGroupName string, serverName string, databaseName string, cancel <-chan struct{}) (<-chan DatabaseOperation, <-chan error) {
	resultChan := make(chan DatabaseOperation)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByDatabase(resourceGroupName, serverName, databaseName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByDatabaseNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

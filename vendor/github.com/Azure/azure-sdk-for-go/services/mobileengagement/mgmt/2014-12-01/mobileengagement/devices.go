package mobileengagement

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// DevicesClient is the microsoft Azure Mobile Engagement REST APIs.
type DevicesClient struct {
	BaseClient
}

// NewDevicesClient creates an instance of the DevicesClient client.
func NewDevicesClient(subscriptionID string) DevicesClient {
	return NewDevicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDevicesClientWithBaseURI creates an instance of the DevicesClient client.
func NewDevicesClientWithBaseURI(baseURI string, subscriptionID string) DevicesClient {
	return DevicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetByDeviceID get the information associated to a device running an application.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name. deviceID is device identifier.
func (client DevicesClient) GetByDeviceID(ctx context.Context, resourceGroupName string, appCollection string, appName string, deviceID string) (result Device, err error) {
	req, err := client.GetByDeviceIDPreparer(ctx, resourceGroupName, appCollection, appName, deviceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByDeviceID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByDeviceIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByDeviceID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByDeviceIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByDeviceID", resp, "Failure responding to request")
	}

	return
}

// GetByDeviceIDPreparer prepares the GetByDeviceID request.
func (client DevicesClient) GetByDeviceIDPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, deviceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"deviceId":          autorest.Encode("path", deviceID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/{deviceId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByDeviceIDSender sends the GetByDeviceID request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesClient) GetByDeviceIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByDeviceIDResponder handles the response to the GetByDeviceID request. The method always
// closes the http.Response Body.
func (client DevicesClient) GetByDeviceIDResponder(resp *http.Response) (result Device, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByUserID get the information associated to a device running an application using the user identifier.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name. userID is user identifier.
func (client DevicesClient) GetByUserID(ctx context.Context, resourceGroupName string, appCollection string, appName string, userID string) (result Device, err error) {
	req, err := client.GetByUserIDPreparer(ctx, resourceGroupName, appCollection, appName, userID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByUserID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByUserIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByUserID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByUserIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "GetByUserID", resp, "Failure responding to request")
	}

	return
}

// GetByUserIDPreparer prepares the GetByUserID request.
func (client DevicesClient) GetByUserIDPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, userID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"userId":            autorest.Encode("path", userID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/users/{userId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByUserIDSender sends the GetByUserID request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesClient) GetByUserIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetByUserIDResponder handles the response to the GetByUserID request. The method always
// closes the http.Response Body.
func (client DevicesClient) GetByUserIDResponder(resp *http.Response) (result Device, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List query the information associated to the devices running an application.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name. top is number of devices to return with each call. Defaults to 100 and cannot return more. Passing a
// greater value is ignored. The response contains a `nextLink` property describing the URI path to get the next page
// of results if not all results could be returned at once. selectParameter is by default all `meta` and `appInfo`
// properties are returned, this property is used to restrict the output to the desired properties. It also excludes
// all devices from the output that have none of the selected properties. In other terms, only devices having at least
// one of the selected property being set is part of the results. Examples: - `$select=appInfo` : select all devices
// having at least 1 appInfo, return them all and don’t return any meta property. - `$select=meta` : return only meta
// properties in the output. - `$select=appInfo,meta/firstSeen,meta/lastSeen` : return all `appInfo`, plus meta object
// containing only firstSeen and lastSeen properties. The format is thus a comma separated list of properties to
// select. Use `appInfo` to select all appInfo properties, `meta` to select all meta properties. Use `appInfo/{key}`
// and `meta/{key}` to select specific appInfo and meta properties. filter is filter can be used to reduce the number
// of results. Filter is a boolean expression that can look like the following examples: * `$filter=deviceId gt
// 'abcdef0123456789abcdef0123456789'` * `$filter=lastModified le 1447284263690L` * `$filter=(deviceId ge
// 'abcdef0123456789abcdef0123456789') and (deviceId lt 'bacdef0123456789abcdef0123456789') and (lastModified gt
// 1447284263690L)` The first example is used automatically for paging when returning the `nextLink` property. The
// filter expression is a combination of checks on some properties that can be compared to their value. The available
// operators are: * `gt`  : greater than * `ge`  : greater than or equals * `lt`  : less than * `le`  : less than or
// equals * `and` : to add multiple checks (all checks must pass), optional parentheses can be used. The properties
// that can be used in the expression are the following: * `deviceId {operator} '{deviceIdValue}'` : a lexicographical
// comparison is made on the deviceId value, use single quotes for the value. * `lastModified {operator} {number}L` :
// returns only meta properties or appInfo properties whose last value modification timestamp compared to the specified
// value is matching (value is milliseconds since January 1st, 1970 UTC). Please note the `L` character after the
// number of milliseconds, its required when the number of milliseconds exceeds `2^31 - 1` (which is always the case
// for recent timestamps). Using `lastModified` excludes all devices from the output that have no property matching the
// timestamp criteria, like `$select`. Please note that the internal value of `lastModified` timestamp for a given
// property is never part of the results.
func (client DevicesClient) List(ctx context.Context, resourceGroupName string, appCollection string, appName string, top *int32, selectParameter string, filter string) (result DevicesQueryResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, appCollection, appName, top, selectParameter, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dqr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "List", resp, "Failure sending request")
		return
	}

	result.dqr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DevicesClient) ListPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, top *int32, selectParameter string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = autorest.Encode("query", selectParameter)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DevicesClient) ListResponder(resp *http.Response) (result DevicesQueryResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DevicesClient) listNextResults(lastResults DevicesQueryResult) (result DevicesQueryResult, err error) {
	req, err := lastResults.devicesQueryResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DevicesClient) ListComplete(ctx context.Context, resourceGroupName string, appCollection string, appName string, top *int32, selectParameter string, filter string) (result DevicesQueryResultIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, appCollection, appName, top, selectParameter, filter)
	return
}

// TagByDeviceID update the tags registered for a set of devices running an application. Updates are performed
// asynchronously, meaning that a few seconds are needed before the modifications appear in the results of the Get
// device command.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name.
func (client DevicesClient) TagByDeviceID(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters DeviceTagsParameters) (result DeviceTagsResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Tags", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "mobileengagement.DevicesClient", "TagByDeviceID")
	}

	req, err := client.TagByDeviceIDPreparer(ctx, resourceGroupName, appCollection, appName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByDeviceID", nil, "Failure preparing request")
		return
	}

	resp, err := client.TagByDeviceIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByDeviceID", resp, "Failure sending request")
		return
	}

	result, err = client.TagByDeviceIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByDeviceID", resp, "Failure responding to request")
	}

	return
}

// TagByDeviceIDPreparer prepares the TagByDeviceID request.
func (client DevicesClient) TagByDeviceIDPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters DeviceTagsParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/devices/tag", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TagByDeviceIDSender sends the TagByDeviceID request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesClient) TagByDeviceIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// TagByDeviceIDResponder handles the response to the TagByDeviceID request. The method always
// closes the http.Response Body.
func (client DevicesClient) TagByDeviceIDResponder(resp *http.Response) (result DeviceTagsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// TagByUserID update the tags registered for a set of users running an application. Updates are performed
// asynchronously, meaning that a few seconds are needed before the modifications appear in the results of the Get
// device command.
//
// resourceGroupName is the name of the resource group. appCollection is application collection. appName is application
// resource name.
func (client DevicesClient) TagByUserID(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters DeviceTagsParameters) (result DeviceTagsResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Tags", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "mobileengagement.DevicesClient", "TagByUserID")
	}

	req, err := client.TagByUserIDPreparer(ctx, resourceGroupName, appCollection, appName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByUserID", nil, "Failure preparing request")
		return
	}

	resp, err := client.TagByUserIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByUserID", resp, "Failure sending request")
		return
	}

	result, err = client.TagByUserIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mobileengagement.DevicesClient", "TagByUserID", resp, "Failure responding to request")
	}

	return
}

// TagByUserIDPreparer prepares the TagByUserID request.
func (client DevicesClient) TagByUserIDPreparer(ctx context.Context, resourceGroupName string, appCollection string, appName string, parameters DeviceTagsParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"appCollection":     autorest.Encode("path", appCollection),
		"appName":           autorest.Encode("path", appName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileEngagement/appcollections/{appCollection}/apps/{appName}/users/tag", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// TagByUserIDSender sends the TagByUserID request. The method will close the
// http.Response Body if it receives an error.
func (client DevicesClient) TagByUserIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// TagByUserIDResponder handles the response to the TagByUserID request. The method always
// closes the http.Response Body.
func (client DevicesClient) TagByUserIDResponder(resp *http.Response) (result DeviceTagsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

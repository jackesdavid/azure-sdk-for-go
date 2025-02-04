//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlhsc

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ConfigurationsClientListByServerGroupPager provides operations for iterating over paged responses.
type ConfigurationsClientListByServerGroupPager struct {
	client    *ConfigurationsClient
	current   ConfigurationsClientListByServerGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConfigurationsClientListByServerGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConfigurationsClientListByServerGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConfigurationsClientListByServerGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerGroupConfigurationListResult.NextLink == nil || len(*p.current.ServerGroupConfigurationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByServerGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConfigurationsClientListByServerGroupResponse page.
func (p *ConfigurationsClientListByServerGroupPager) PageResponse() ConfigurationsClientListByServerGroupResponse {
	return p.current
}

// ConfigurationsClientListByServerPager provides operations for iterating over paged responses.
type ConfigurationsClientListByServerPager struct {
	client    *ConfigurationsClient
	current   ConfigurationsClientListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ConfigurationsClientListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ConfigurationsClientListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ConfigurationsClientListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerConfigurationListResult.NextLink == nil || len(*p.current.ServerConfigurationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ConfigurationsClientListByServerResponse page.
func (p *ConfigurationsClientListByServerPager) PageResponse() ConfigurationsClientListByServerResponse {
	return p.current
}

// ServerGroupsClientListByResourceGroupPager provides operations for iterating over paged responses.
type ServerGroupsClientListByResourceGroupPager struct {
	client    *ServerGroupsClient
	current   ServerGroupsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServerGroupsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServerGroupsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServerGroupsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerGroupListResult.NextLink == nil || len(*p.current.ServerGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServerGroupsClientListByResourceGroupResponse page.
func (p *ServerGroupsClientListByResourceGroupPager) PageResponse() ServerGroupsClientListByResourceGroupResponse {
	return p.current
}

// ServerGroupsClientListPager provides operations for iterating over paged responses.
type ServerGroupsClientListPager struct {
	client    *ServerGroupsClient
	current   ServerGroupsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServerGroupsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServerGroupsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServerGroupsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerGroupListResult.NextLink == nil || len(*p.current.ServerGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServerGroupsClientListResponse page.
func (p *ServerGroupsClientListPager) PageResponse() ServerGroupsClientListResponse {
	return p.current
}

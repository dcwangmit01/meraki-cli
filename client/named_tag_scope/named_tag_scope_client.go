// Code generated by go-swagger; DO NOT EDIT.

package named_tag_scope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new named tag scope API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for named tag scope API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNetworkSmTargetGroup(params *CreateNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkSmTargetGroupCreated, error)

	DeleteNetworkSmTargetGroup(params *DeleteNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkSmTargetGroupNoContent, error)

	GetNetworkSmTargetGroup(params *GetNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSmTargetGroupOK, error)

	GetNetworkSmTargetGroups(params *GetNetworkSmTargetGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSmTargetGroupsOK, error)

	UpdateNetworkSmTargetGroup(params *UpdateNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSmTargetGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNetworkSmTargetGroup creates network sm target group

  Add a target group
*/
func (a *Client) CreateNetworkSmTargetGroup(params *CreateNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNetworkSmTargetGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkSmTargetGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNetworkSmTargetGroup",
		Method:             "POST",
		PathPattern:        "/networks/{networkId}/sm/targetGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNetworkSmTargetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNetworkSmTargetGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNetworkSmTargetGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkSmTargetGroup deletes network sm target group

  Delete a target group from a network
*/
func (a *Client) DeleteNetworkSmTargetGroup(params *DeleteNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNetworkSmTargetGroupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkSmTargetGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNetworkSmTargetGroup",
		Method:             "DELETE",
		PathPattern:        "/networks/{networkId}/sm/targetGroups/{targetGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkSmTargetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNetworkSmTargetGroupNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNetworkSmTargetGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkSmTargetGroup gets network sm target group

  Return a target group
*/
func (a *Client) GetNetworkSmTargetGroup(params *GetNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSmTargetGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkSmTargetGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkSmTargetGroup",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/sm/targetGroups/{targetGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkSmTargetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkSmTargetGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkSmTargetGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkSmTargetGroups gets network sm target groups

  List the target groups in this network
*/
func (a *Client) GetNetworkSmTargetGroups(params *GetNetworkSmTargetGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNetworkSmTargetGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkSmTargetGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetworkSmTargetGroups",
		Method:             "GET",
		PathPattern:        "/networks/{networkId}/sm/targetGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkSmTargetGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetworkSmTargetGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkSmTargetGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateNetworkSmTargetGroup updates network sm target group

  Update a target group
*/
func (a *Client) UpdateNetworkSmTargetGroup(params *UpdateNetworkSmTargetGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateNetworkSmTargetGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkSmTargetGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNetworkSmTargetGroup",
		Method:             "PUT",
		PathPattern:        "/networks/{networkId}/sm/targetGroups/{targetGroupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateNetworkSmTargetGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateNetworkSmTargetGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateNetworkSmTargetGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

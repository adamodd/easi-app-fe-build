// Code generated by go-swagger; DO NOT EDIT.

package enumeration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new enumeration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for enumeration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	EnumerationFindList(params *EnumerationFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnumerationFindListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  EnumerationFindList retrieves a list of enumerations based on a list of enumeration names passed in

  Retrieve a list of enumerations based on a list of enumeration names passed in.
*/
func (a *Client) EnumerationFindList(params *EnumerationFindListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnumerationFindListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnumerationFindListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enumerationFindList",
		Method:             "GET",
		PathPattern:        "/enumeration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnumerationFindListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnumerationFindListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enumerationFindList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
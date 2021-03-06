// Code generated by go-swagger; DO NOT EDIT.

package room_directory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new room directory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for room directory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteRoomAlias removes a mapping of room alias to room ID

Remove a mapping of room alias to room ID.

Servers may choose to implement additional access control checks here, for instance that room aliases can only be deleted by their creator or a server administrator.
*/
func (a *Client) DeleteRoomAlias(params *DeleteRoomAliasParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRoomAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRoomAliasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRoomAlias",
		Method:             "DELETE",
		PathPattern:        "/_matrix/client/unstable/directory/room/{roomAlias}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoomAliasReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRoomAliasOK), nil

}

/*
GetRoomIDByAlias gets the room ID corresponding to this room alias

Requests that the server resolve a room alias to a room ID.

The server will use the federation API to resolve the alias if the
domain part of the alias does not correspond to the server's own
domain.
*/
func (a *Client) GetRoomIDByAlias(params *GetRoomIDByAliasParams) (*GetRoomIDByAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoomIDByAliasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoomIdByAlias",
		Method:             "GET",
		PathPattern:        "/_matrix/client/unstable/directory/room/{roomAlias}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRoomIDByAliasReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoomIDByAliasOK), nil

}

/*
SetRoomAlias creates a new mapping from room alias to room ID
*/
func (a *Client) SetRoomAlias(params *SetRoomAliasParams, authInfo runtime.ClientAuthInfoWriter) (*SetRoomAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRoomAliasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setRoomAlias",
		Method:             "PUT",
		PathPattern:        "/_matrix/client/unstable/directory/room/{roomAlias}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetRoomAliasReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetRoomAliasOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the yggdrasil client
type API interface {
	/*
	   GetControlMessageForDevice Get control message for device API*/
	GetControlMessageForDevice(ctx context.Context, params *GetControlMessageForDeviceParams) (*GetControlMessageForDeviceOK, error)
	/*
	   GetDataMessageForDevice Get data message for device API*/
	GetDataMessageForDevice(ctx context.Context, params *GetDataMessageForDeviceParams) (*GetDataMessageForDeviceOK, error)
	/*
	   PostControlMessageForDevice Post control message for device API*/
	PostControlMessageForDevice(ctx context.Context, params *PostControlMessageForDeviceParams) (*PostControlMessageForDeviceOK, error)
	/*
	   PostDataMessageForDevice Post data message for device API*/
	PostDataMessageForDevice(ctx context.Context, params *PostDataMessageForDeviceParams) (*PostDataMessageForDeviceOK, *PostDataMessageForDeviceAlreadyReported, error)
}

// New creates a new yggdrasil API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for yggdrasil API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetControlMessageForDevice Get control message for device API
*/
func (a *Client) GetControlMessageForDevice(ctx context.Context, params *GetControlMessageForDeviceParams) (*GetControlMessageForDeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetControlMessageForDevice",
		Method:             "GET",
		PathPattern:        "/control/{device_id}/in",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetControlMessageForDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetControlMessageForDeviceOK), nil

}

/*
GetDataMessageForDevice Get data message for device API
*/
func (a *Client) GetDataMessageForDevice(ctx context.Context, params *GetDataMessageForDeviceParams) (*GetDataMessageForDeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDataMessageForDevice",
		Method:             "GET",
		PathPattern:        "/data/{device_id}/in",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataMessageForDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDataMessageForDeviceOK), nil

}

/*
PostControlMessageForDevice Post control message for device API
*/
func (a *Client) PostControlMessageForDevice(ctx context.Context, params *PostControlMessageForDeviceParams) (*PostControlMessageForDeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostControlMessageForDevice",
		Method:             "POST",
		PathPattern:        "/control/{device_id}/out",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostControlMessageForDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostControlMessageForDeviceOK), nil

}

/*
PostDataMessageForDevice Post data message for device API
*/
func (a *Client) PostDataMessageForDevice(ctx context.Context, params *PostDataMessageForDeviceParams) (*PostDataMessageForDeviceOK, *PostDataMessageForDeviceAlreadyReported, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostDataMessageForDevice",
		Method:             "POST",
		PathPattern:        "/data/{device_id}/out",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDataMessageForDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostDataMessageForDeviceOK:
		return value, nil, nil
	case *PostDataMessageForDeviceAlreadyReported:
		return nil, value, nil
	}
	return nil, nil, nil

}

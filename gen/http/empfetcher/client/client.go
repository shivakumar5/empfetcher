// Code generated by goa v3.2.0, DO NOT EDIT.
//
// empfetcher client HTTP transport
//
// Command:
// $ goa gen github.com/flexera/empfetcher/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the empfetcher service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// List Doer is the HTTP client used to make requests to the list endpoint.
	ListDoer goahttp.Doer

	// Show Doer is the HTTP client used to make requests to the show endpoint.
	ShowDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// Restore Doer is the HTTP client used to make requests to the restore
	// endpoint.
	RestoreDoer goahttp.Doer

	// Viewdeleted Doer is the HTTP client used to make requests to the viewdeleted
	// endpoint.
	ViewdeletedDoer goahttp.Doer

	// Search Doer is the HTTP client used to make requests to the search endpoint.
	SearchDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the empfetcher service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		UpdateDoer:          doer,
		ListDoer:            doer,
		ShowDoer:            doer,
		DeleteDoer:          doer,
		RestoreDoer:         doer,
		ViewdeletedDoer:     doer,
		SearchDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the empfetcher service
// add server.
func (c *Client) Add() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddRequest(c.encoder)
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the empfetcher
// service update server.
func (c *Client) Update() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateRequest(c.encoder)
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "update", err)
		}
		return decodeResponse(resp)
	}
}

// List returns an endpoint that makes HTTP requests to the empfetcher service
// list server.
func (c *Client) List() goa.Endpoint {
	var (
		decodeResponse = DecodeListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "list", err)
		}
		return decodeResponse(resp)
	}
}

// Show returns an endpoint that makes HTTP requests to the empfetcher service
// show server.
func (c *Client) Show() goa.Endpoint {
	var (
		decodeResponse = DecodeShowResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildShowRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ShowDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "show", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the empfetcher
// service delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "delete", err)
		}
		return decodeResponse(resp)
	}
}

// Restore returns an endpoint that makes HTTP requests to the empfetcher
// service restore server.
func (c *Client) Restore() goa.Endpoint {
	var (
		decodeResponse = DecodeRestoreResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRestoreRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RestoreDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "restore", err)
		}
		return decodeResponse(resp)
	}
}

// Viewdeleted returns an endpoint that makes HTTP requests to the empfetcher
// service viewdeleted server.
func (c *Client) Viewdeleted() goa.Endpoint {
	var (
		decodeResponse = DecodeViewdeletedResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildViewdeletedRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ViewdeletedDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "viewdeleted", err)
		}
		return decodeResponse(resp)
	}
}

// Search returns an endpoint that makes HTTP requests to the empfetcher
// service search server.
func (c *Client) Search() goa.Endpoint {
	var (
		decodeResponse = DecodeSearchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSearchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SearchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("empfetcher", "search", err)
		}
		return decodeResponse(resp)
	}
}

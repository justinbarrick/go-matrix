// Code generated by go-swagger; DO NOT EDIT.

package room_membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetJoinedRoomsParams creates a new GetJoinedRoomsParams object
// with the default values initialized.
func NewGetJoinedRoomsParams() *GetJoinedRoomsParams {

	return &GetJoinedRoomsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJoinedRoomsParamsWithTimeout creates a new GetJoinedRoomsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJoinedRoomsParamsWithTimeout(timeout time.Duration) *GetJoinedRoomsParams {

	return &GetJoinedRoomsParams{

		timeout: timeout,
	}
}

// NewGetJoinedRoomsParamsWithContext creates a new GetJoinedRoomsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJoinedRoomsParamsWithContext(ctx context.Context) *GetJoinedRoomsParams {

	return &GetJoinedRoomsParams{

		Context: ctx,
	}
}

// NewGetJoinedRoomsParamsWithHTTPClient creates a new GetJoinedRoomsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJoinedRoomsParamsWithHTTPClient(client *http.Client) *GetJoinedRoomsParams {

	return &GetJoinedRoomsParams{
		HTTPClient: client,
	}
}

/*GetJoinedRoomsParams contains all the parameters to send to the API endpoint
for the get joined rooms operation typically these are written to a http.Request
*/
type GetJoinedRoomsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get joined rooms params
func (o *GetJoinedRoomsParams) WithTimeout(timeout time.Duration) *GetJoinedRoomsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get joined rooms params
func (o *GetJoinedRoomsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get joined rooms params
func (o *GetJoinedRoomsParams) WithContext(ctx context.Context) *GetJoinedRoomsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get joined rooms params
func (o *GetJoinedRoomsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get joined rooms params
func (o *GetJoinedRoomsParams) WithHTTPClient(client *http.Client) *GetJoinedRoomsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get joined rooms params
func (o *GetJoinedRoomsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetJoinedRoomsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

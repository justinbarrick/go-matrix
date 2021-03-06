// Code generated by go-swagger; DO NOT EDIT.

package device_management

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

	models "github.com/justinbarrick/go-matrix/pkg/models"
)

// NewUpdateDeviceParams creates a new UpdateDeviceParams object
// with the default values initialized.
func NewUpdateDeviceParams() *UpdateDeviceParams {
	var ()
	return &UpdateDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceParamsWithTimeout creates a new UpdateDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceParamsWithTimeout(timeout time.Duration) *UpdateDeviceParams {
	var ()
	return &UpdateDeviceParams{

		timeout: timeout,
	}
}

// NewUpdateDeviceParamsWithContext creates a new UpdateDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceParamsWithContext(ctx context.Context) *UpdateDeviceParams {
	var ()
	return &UpdateDeviceParams{

		Context: ctx,
	}
}

// NewUpdateDeviceParamsWithHTTPClient creates a new UpdateDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceParamsWithHTTPClient(client *http.Client) *UpdateDeviceParams {
	var ()
	return &UpdateDeviceParams{
		HTTPClient: client,
	}
}

/*UpdateDeviceParams contains all the parameters to send to the API endpoint
for the update device operation typically these are written to a http.Request
*/
type UpdateDeviceParams struct {

	/*Body
	  New information for the device.

	*/
	Body *models.UpdateDeviceParamsBody
	/*DeviceID
	  The device to update.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device params
func (o *UpdateDeviceParams) WithTimeout(timeout time.Duration) *UpdateDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device params
func (o *UpdateDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device params
func (o *UpdateDeviceParams) WithContext(ctx context.Context) *UpdateDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device params
func (o *UpdateDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device params
func (o *UpdateDeviceParams) WithHTTPClient(client *http.Client) *UpdateDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device params
func (o *UpdateDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device params
func (o *UpdateDeviceParams) WithBody(body *models.UpdateDeviceParamsBody) *UpdateDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device params
func (o *UpdateDeviceParams) SetBody(body *models.UpdateDeviceParamsBody) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the update device params
func (o *UpdateDeviceParams) WithDeviceID(deviceID string) *UpdateDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device params
func (o *UpdateDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetNetworkSmCertsParams creates a new GetNetworkSmCertsParams object
// with the default values initialized.
func NewGetNetworkSmCertsParams() *GetNetworkSmCertsParams {
	var ()
	return &GetNetworkSmCertsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkSmCertsParamsWithTimeout creates a new GetNetworkSmCertsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkSmCertsParamsWithTimeout(timeout time.Duration) *GetNetworkSmCertsParams {
	var ()
	return &GetNetworkSmCertsParams{

		timeout: timeout,
	}
}

// NewGetNetworkSmCertsParamsWithContext creates a new GetNetworkSmCertsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkSmCertsParamsWithContext(ctx context.Context) *GetNetworkSmCertsParams {
	var ()
	return &GetNetworkSmCertsParams{

		Context: ctx,
	}
}

// NewGetNetworkSmCertsParamsWithHTTPClient creates a new GetNetworkSmCertsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkSmCertsParamsWithHTTPClient(client *http.Client) *GetNetworkSmCertsParams {
	var ()
	return &GetNetworkSmCertsParams{
		HTTPClient: client,
	}
}

/*GetNetworkSmCertsParams contains all the parameters to send to the API endpoint
for the get network sm certs operation typically these are written to a http.Request
*/
type GetNetworkSmCertsParams struct {

	/*DeviceID*/
	DeviceID string
	/*NetworkID*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network sm certs params
func (o *GetNetworkSmCertsParams) WithTimeout(timeout time.Duration) *GetNetworkSmCertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network sm certs params
func (o *GetNetworkSmCertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network sm certs params
func (o *GetNetworkSmCertsParams) WithContext(ctx context.Context) *GetNetworkSmCertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network sm certs params
func (o *GetNetworkSmCertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network sm certs params
func (o *GetNetworkSmCertsParams) WithHTTPClient(client *http.Client) *GetNetworkSmCertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network sm certs params
func (o *GetNetworkSmCertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get network sm certs params
func (o *GetNetworkSmCertsParams) WithDeviceID(deviceID string) *GetNetworkSmCertsParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get network sm certs params
func (o *GetNetworkSmCertsParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNetworkID adds the networkID to the get network sm certs params
func (o *GetNetworkSmCertsParams) WithNetworkID(networkID string) *GetNetworkSmCertsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network sm certs params
func (o *GetNetworkSmCertsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkSmCertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

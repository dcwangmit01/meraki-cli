// Code generated by go-swagger; DO NOT EDIT.

package switch_settings

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

	"github.com/cisco-sso/meraki-cli/models"
)

// NewUpdateNetworkSwitchSettingsMulticastParams creates a new UpdateNetworkSwitchSettingsMulticastParams object
// with the default values initialized.
func NewUpdateNetworkSwitchSettingsMulticastParams() *UpdateNetworkSwitchSettingsMulticastParams {
	var ()
	return &UpdateNetworkSwitchSettingsMulticastParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSwitchSettingsMulticastParamsWithTimeout creates a new UpdateNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateNetworkSwitchSettingsMulticastParamsWithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsMulticastParams {
	var ()
	return &UpdateNetworkSwitchSettingsMulticastParams{

		timeout: timeout,
	}
}

// NewUpdateNetworkSwitchSettingsMulticastParamsWithContext creates a new UpdateNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateNetworkSwitchSettingsMulticastParamsWithContext(ctx context.Context) *UpdateNetworkSwitchSettingsMulticastParams {
	var ()
	return &UpdateNetworkSwitchSettingsMulticastParams{

		Context: ctx,
	}
}

// NewUpdateNetworkSwitchSettingsMulticastParamsWithHTTPClient creates a new UpdateNetworkSwitchSettingsMulticastParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateNetworkSwitchSettingsMulticastParamsWithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsMulticastParams {
	var ()
	return &UpdateNetworkSwitchSettingsMulticastParams{
		HTTPClient: client,
	}
}

/*UpdateNetworkSwitchSettingsMulticastParams contains all the parameters to send to the API endpoint
for the update network switch settings multicast operation typically these are written to a http.Request
*/
type UpdateNetworkSwitchSettingsMulticastParams struct {

	/*NetworkID*/
	NetworkID string
	/*UpdateNetworkSwitchSettingsMulticast*/
	UpdateNetworkSwitchSettingsMulticast *models.UpdateNetworkSwitchSettingsMulticast

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) WithTimeout(timeout time.Duration) *UpdateNetworkSwitchSettingsMulticastParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) WithContext(ctx context.Context) *UpdateNetworkSwitchSettingsMulticastParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) WithHTTPClient(client *http.Client) *UpdateNetworkSwitchSettingsMulticastParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) WithNetworkID(networkID string) *UpdateNetworkSwitchSettingsMulticastParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSwitchSettingsMulticast adds the updateNetworkSwitchSettingsMulticast to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) WithUpdateNetworkSwitchSettingsMulticast(updateNetworkSwitchSettingsMulticast *models.UpdateNetworkSwitchSettingsMulticast) *UpdateNetworkSwitchSettingsMulticastParams {
	o.SetUpdateNetworkSwitchSettingsMulticast(updateNetworkSwitchSettingsMulticast)
	return o
}

// SetUpdateNetworkSwitchSettingsMulticast adds the updateNetworkSwitchSettingsMulticast to the update network switch settings multicast params
func (o *UpdateNetworkSwitchSettingsMulticastParams) SetUpdateNetworkSwitchSettingsMulticast(updateNetworkSwitchSettingsMulticast *models.UpdateNetworkSwitchSettingsMulticast) {
	o.UpdateNetworkSwitchSettingsMulticast = updateNetworkSwitchSettingsMulticast
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSwitchSettingsMulticastParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}

	if o.UpdateNetworkSwitchSettingsMulticast != nil {
		if err := r.SetBodyParam(o.UpdateNetworkSwitchSettingsMulticast); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package radio_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkDeviceWirelessRadioSettingsReader is a Reader for the UpdateNetworkDeviceWirelessRadioSettings structure.
type UpdateNetworkDeviceWirelessRadioSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkDeviceWirelessRadioSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkDeviceWirelessRadioSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkDeviceWirelessRadioSettingsOK creates a UpdateNetworkDeviceWirelessRadioSettingsOK with default headers values
func NewUpdateNetworkDeviceWirelessRadioSettingsOK() *UpdateNetworkDeviceWirelessRadioSettingsOK {
	return &UpdateNetworkDeviceWirelessRadioSettingsOK{}
}

/*UpdateNetworkDeviceWirelessRadioSettingsOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkDeviceWirelessRadioSettingsOK struct {
	Payload interface{}
}

func (o *UpdateNetworkDeviceWirelessRadioSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/devices/{serial}/wireless/radioSettings][%d] updateNetworkDeviceWirelessRadioSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkDeviceWirelessRadioSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkDeviceWirelessRadioSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
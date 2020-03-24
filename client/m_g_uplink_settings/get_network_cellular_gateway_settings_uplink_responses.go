// Code generated by go-swagger; DO NOT EDIT.

package m_g_uplink_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkCellularGatewaySettingsUplinkReader is a Reader for the GetNetworkCellularGatewaySettingsUplink structure.
type GetNetworkCellularGatewaySettingsUplinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkCellularGatewaySettingsUplinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkCellularGatewaySettingsUplinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkCellularGatewaySettingsUplinkOK creates a GetNetworkCellularGatewaySettingsUplinkOK with default headers values
func NewGetNetworkCellularGatewaySettingsUplinkOK() *GetNetworkCellularGatewaySettingsUplinkOK {
	return &GetNetworkCellularGatewaySettingsUplinkOK{}
}

/*GetNetworkCellularGatewaySettingsUplinkOK handles this case with default header values.

Successful operation
*/
type GetNetworkCellularGatewaySettingsUplinkOK struct {
	Payload interface{}
}

func (o *GetNetworkCellularGatewaySettingsUplinkOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/cellularGateway/settings/uplink][%d] getNetworkCellularGatewaySettingsUplinkOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCellularGatewaySettingsUplinkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkCellularGatewaySettingsUplinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
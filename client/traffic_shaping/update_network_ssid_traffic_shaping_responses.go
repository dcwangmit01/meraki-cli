// Code generated by go-swagger; DO NOT EDIT.

package traffic_shaping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkSsidTrafficShapingReader is a Reader for the UpdateNetworkSsidTrafficShaping structure.
type UpdateNetworkSsidTrafficShapingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSsidTrafficShapingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSsidTrafficShapingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkSsidTrafficShapingOK creates a UpdateNetworkSsidTrafficShapingOK with default headers values
func NewUpdateNetworkSsidTrafficShapingOK() *UpdateNetworkSsidTrafficShapingOK {
	return &UpdateNetworkSsidTrafficShapingOK{}
}

/*UpdateNetworkSsidTrafficShapingOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkSsidTrafficShapingOK struct {
	Payload interface{}
}

func (o *UpdateNetworkSsidTrafficShapingOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/ssids/{number}/trafficShaping][%d] updateNetworkSsidTrafficShapingOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSsidTrafficShapingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSsidTrafficShapingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package switch_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchStacksReader is a Reader for the GetNetworkSwitchStacks structure.
type GetNetworkSwitchStacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchStacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchStacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSwitchStacksOK creates a GetNetworkSwitchStacksOK with default headers values
func NewGetNetworkSwitchStacksOK() *GetNetworkSwitchStacksOK {
	return &GetNetworkSwitchStacksOK{}
}

/*GetNetworkSwitchStacksOK handles this case with default header values.

Successful operation
*/
type GetNetworkSwitchStacksOK struct {
	Payload interface{}
}

func (o *GetNetworkSwitchStacksOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switchStacks][%d] getNetworkSwitchStacksOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStacksOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchStacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

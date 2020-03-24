// Code generated by go-swagger; DO NOT EDIT.

package meraki_auth_users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkMerakiAuthUsersReader is a Reader for the GetNetworkMerakiAuthUsers structure.
type GetNetworkMerakiAuthUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkMerakiAuthUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkMerakiAuthUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkMerakiAuthUsersOK creates a GetNetworkMerakiAuthUsersOK with default headers values
func NewGetNetworkMerakiAuthUsersOK() *GetNetworkMerakiAuthUsersOK {
	return &GetNetworkMerakiAuthUsersOK{}
}

/*GetNetworkMerakiAuthUsersOK handles this case with default header values.

Successful operation
*/
type GetNetworkMerakiAuthUsersOK struct {
	Payload interface{}
}

func (o *GetNetworkMerakiAuthUsersOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/merakiAuthUsers][%d] getNetworkMerakiAuthUsersOK  %+v", 200, o.Payload)
}

func (o *GetNetworkMerakiAuthUsersOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkMerakiAuthUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package http_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateNetworkHTTPServerReader is a Reader for the UpdateNetworkHTTPServer structure.
type UpdateNetworkHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkHTTPServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkHTTPServerOK creates a UpdateNetworkHTTPServerOK with default headers values
func NewUpdateNetworkHTTPServerOK() *UpdateNetworkHTTPServerOK {
	return &UpdateNetworkHTTPServerOK{}
}

/*UpdateNetworkHTTPServerOK handles this case with default header values.

Successful operation
*/
type UpdateNetworkHTTPServerOK struct {
	Payload interface{}
}

func (o *UpdateNetworkHTTPServerOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/httpServers/{id}][%d] updateNetworkHttpServerOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkHTTPServerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkHTTPServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package s_s_i_ds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSsidReader is a Reader for the GetNetworkSsid structure.
type GetNetworkSsidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSsidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSsidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkSsidOK creates a GetNetworkSsidOK with default headers values
func NewGetNetworkSsidOK() *GetNetworkSsidOK {
	return &GetNetworkSsidOK{}
}

/*GetNetworkSsidOK handles this case with default header values.

Successful operation
*/
type GetNetworkSsidOK struct {
	Payload interface{}
}

func (o *GetNetworkSsidOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/ssids/{number}][%d] getNetworkSsidOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSsidOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSsidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

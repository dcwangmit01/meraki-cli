// Code generated by go-swagger; DO NOT EDIT.

package s_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateNetworkSmAppPolarisReader is a Reader for the CreateNetworkSmAppPolaris structure.
type CreateNetworkSmAppPolarisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSmAppPolarisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateNetworkSmAppPolarisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNetworkSmAppPolarisOK creates a CreateNetworkSmAppPolarisOK with default headers values
func NewCreateNetworkSmAppPolarisOK() *CreateNetworkSmAppPolarisOK {
	return &CreateNetworkSmAppPolarisOK{}
}

/*CreateNetworkSmAppPolarisOK handles this case with default header values.

Successful operation
*/
type CreateNetworkSmAppPolarisOK struct {
	Payload interface{}
}

func (o *CreateNetworkSmAppPolarisOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/app/polaris][%d] createNetworkSmAppPolarisOK  %+v", 200, o.Payload)
}

func (o *CreateNetworkSmAppPolarisOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSmAppPolarisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

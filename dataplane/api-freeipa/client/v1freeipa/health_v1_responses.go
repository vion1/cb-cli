// Code generated by go-swagger; DO NOT EDIT.

package v1freeipa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// HealthV1Reader is a Reader for the HealthV1 structure.
type HealthV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHealthV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthV1OK creates a HealthV1OK with default headers values
func NewHealthV1OK() *HealthV1OK {
	return &HealthV1OK{}
}

/*HealthV1OK handles this case with default header values.

successful operation
*/
type HealthV1OK struct {
	Payload *model.HealthDetailsFreeIpaV1Response
}

func (o *HealthV1OK) Error() string {
	return fmt.Sprintf("[GET /v1/freeipa/health][%d] healthV1OK  %+v", 200, o.Payload)
}

func (o *HealthV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.HealthDetailsFreeIpaV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

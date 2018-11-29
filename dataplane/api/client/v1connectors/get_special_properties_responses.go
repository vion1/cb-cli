// Code generated by go-swagger; DO NOT EDIT.

package v1connectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetSpecialPropertiesReader is a Reader for the GetSpecialProperties structure.
type GetSpecialPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpecialPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSpecialPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpecialPropertiesOK creates a GetSpecialPropertiesOK with default headers values
func NewGetSpecialPropertiesOK() *GetSpecialPropertiesOK {
	return &GetSpecialPropertiesOK{}
}

/*GetSpecialPropertiesOK handles this case with default header values.

successful operation
*/
type GetSpecialPropertiesOK struct {
	Payload *model.SpecialParametersJSON
}

func (o *GetSpecialPropertiesOK) Error() string {
	return fmt.Sprintf("[POST /v1/connectors/custom][%d] getSpecialPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetSpecialPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SpecialParametersJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

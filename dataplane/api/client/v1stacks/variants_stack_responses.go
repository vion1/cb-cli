// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// VariantsStackReader is a Reader for the VariantsStack structure.
type VariantsStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VariantsStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVariantsStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVariantsStackOK creates a VariantsStackOK with default headers values
func NewVariantsStackOK() *VariantsStackOK {
	return &VariantsStackOK{}
}

/*VariantsStackOK handles this case with default header values.

successful operation
*/
type VariantsStackOK struct {
	Payload *model.PlatformVariantsJSON
}

func (o *VariantsStackOK) Error() string {
	return fmt.Sprintf("[GET /v1/stacks/platformVariants][%d] variantsStackOK  %+v", 200, o.Payload)
}

func (o *VariantsStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformVariantsJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

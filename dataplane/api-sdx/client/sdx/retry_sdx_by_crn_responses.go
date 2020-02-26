// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// RetrySdxByCrnReader is a Reader for the RetrySdxByCrn structure.
type RetrySdxByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrySdxByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrySdxByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrySdxByCrnOK creates a RetrySdxByCrnOK with default headers values
func NewRetrySdxByCrnOK() *RetrySdxByCrnOK {
	return &RetrySdxByCrnOK{}
}

/*RetrySdxByCrnOK handles this case with default header values.

successful operation
*/
type RetrySdxByCrnOK struct {
	Payload *model.FlowIdentifier
}

func (o *RetrySdxByCrnOK) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/retry][%d] retrySdxByCrnOK  %+v", 200, o.Payload)
}

func (o *RetrySdxByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

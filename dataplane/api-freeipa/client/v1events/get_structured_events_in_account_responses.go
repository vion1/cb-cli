// Code generated by go-swagger; DO NOT EDIT.

package v1events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-freeipa/model"
)

// GetStructuredEventsInAccountReader is a Reader for the GetStructuredEventsInAccount structure.
type GetStructuredEventsInAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStructuredEventsInAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStructuredEventsInAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStructuredEventsInAccountOK creates a GetStructuredEventsInAccountOK with default headers values
func NewGetStructuredEventsInAccountOK() *GetStructuredEventsInAccountOK {
	return &GetStructuredEventsInAccountOK{}
}

/*GetStructuredEventsInAccountOK handles this case with default header values.

successful operation
*/
type GetStructuredEventsInAccountOK struct {
	Payload *model.CDPStructuredEventContainer
}

func (o *GetStructuredEventsInAccountOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/{name}/structured][%d] getStructuredEventsInAccountOK  %+v", 200, o.Payload)
}

func (o *GetStructuredEventsInAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CDPStructuredEventContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

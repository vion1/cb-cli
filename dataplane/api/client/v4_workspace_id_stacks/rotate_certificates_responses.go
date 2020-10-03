// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RotateCertificatesReader is a Reader for the RotateCertificates structure.
type RotateCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRotateCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRotateCertificatesOK creates a RotateCertificatesOK with default headers values
func NewRotateCertificatesOK() *RotateCertificatesOK {
	return &RotateCertificatesOK{}
}

/*RotateCertificatesOK handles this case with default header values.

successful operation
*/
type RotateCertificatesOK struct {
	Payload *model.CertificatesRotationV4Response
}

func (o *RotateCertificatesOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/stacks/internal/{name}/rotate_certificates][%d] rotateCertificatesOK  %+v", 200, o.Payload)
}

func (o *RotateCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.CertificatesRotationV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

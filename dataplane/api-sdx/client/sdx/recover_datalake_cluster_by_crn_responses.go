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

// RecoverDatalakeClusterByCrnReader is a Reader for the RecoverDatalakeClusterByCrn structure.
type RecoverDatalakeClusterByCrnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecoverDatalakeClusterByCrnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecoverDatalakeClusterByCrnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRecoverDatalakeClusterByCrnOK creates a RecoverDatalakeClusterByCrnOK with default headers values
func NewRecoverDatalakeClusterByCrnOK() *RecoverDatalakeClusterByCrnOK {
	return &RecoverDatalakeClusterByCrnOK{}
}

/*RecoverDatalakeClusterByCrnOK handles this case with default header values.

successful operation
*/
type RecoverDatalakeClusterByCrnOK struct {
	Payload *model.SdxRecoveryResponse
}

func (o *RecoverDatalakeClusterByCrnOK) Error() string {
	return fmt.Sprintf("[POST /sdx/crn/{crn}/recover_failed_upgrade][%d] recoverDatalakeClusterByCrnOK  %+v", 200, o.Payload)
}

func (o *RecoverDatalakeClusterByCrnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SdxRecoveryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

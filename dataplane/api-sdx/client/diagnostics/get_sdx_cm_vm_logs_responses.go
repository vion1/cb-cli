// Code generated by go-swagger; DO NOT EDIT.

package diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-sdx/model"
)

// GetSdxCmVMLogsReader is a Reader for the GetSdxCmVMLogs structure.
type GetSdxCmVMLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSdxCmVMLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSdxCmVMLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSdxCmVMLogsOK creates a GetSdxCmVMLogsOK with default headers values
func NewGetSdxCmVMLogsOK() *GetSdxCmVMLogsOK {
	return &GetSdxCmVMLogsOK{}
}

/*GetSdxCmVMLogsOK handles this case with default header values.

successful operation
*/
type GetSdxCmVMLogsOK struct {
	Payload *model.VMLogsResponse
}

func (o *GetSdxCmVMLogsOK) Error() string {
	return fmt.Sprintf("[GET /diagnostics/logs][%d] getSdxCmVmLogsOK  %+v", 200, o.Payload)
}

func (o *GetSdxCmVMLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.VMLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

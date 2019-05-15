// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// GetAccessConfigsForWorkspaceReader is a Reader for the GetAccessConfigsForWorkspace structure.
type GetAccessConfigsForWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessConfigsForWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccessConfigsForWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccessConfigsForWorkspaceOK creates a GetAccessConfigsForWorkspaceOK with default headers values
func NewGetAccessConfigsForWorkspaceOK() *GetAccessConfigsForWorkspaceOK {
	return &GetAccessConfigsForWorkspaceOK{}
}

/*GetAccessConfigsForWorkspaceOK handles this case with default header values.

successful operation
*/
type GetAccessConfigsForWorkspaceOK struct {
	Payload *model.PlatformAccessConfigsV1Response
}

func (o *GetAccessConfigsForWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v1/platform_resources/access_configs][%d] getAccessConfigsForWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetAccessConfigsForWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PlatformAccessConfigsV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

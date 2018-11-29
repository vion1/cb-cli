// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DeleteSmartSenseSubscriptionInWorkspaceReader is a Reader for the DeleteSmartSenseSubscriptionInWorkspace structure.
type DeleteSmartSenseSubscriptionInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSmartSenseSubscriptionInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteSmartSenseSubscriptionInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSmartSenseSubscriptionInWorkspaceOK creates a DeleteSmartSenseSubscriptionInWorkspaceOK with default headers values
func NewDeleteSmartSenseSubscriptionInWorkspaceOK() *DeleteSmartSenseSubscriptionInWorkspaceOK {
	return &DeleteSmartSenseSubscriptionInWorkspaceOK{}
}

/*DeleteSmartSenseSubscriptionInWorkspaceOK handles this case with default header values.

successful operation
*/
type DeleteSmartSenseSubscriptionInWorkspaceOK struct {
	Payload *model.SmartSenseSubscriptionJSON
}

func (o *DeleteSmartSenseSubscriptionInWorkspaceOK) Error() string {
	return fmt.Sprintf("[DELETE /v3/{workspaceId}/smartsensesubscriptions/{name}][%d] deleteSmartSenseSubscriptionInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *DeleteSmartSenseSubscriptionInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

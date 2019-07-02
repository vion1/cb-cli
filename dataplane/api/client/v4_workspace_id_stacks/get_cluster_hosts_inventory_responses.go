// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetClusterHostsInventoryReader is a Reader for the GetClusterHostsInventory structure.
type GetClusterHostsInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterHostsInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterHostsInventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterHostsInventoryOK creates a GetClusterHostsInventoryOK with default headers values
func NewGetClusterHostsInventoryOK() *GetClusterHostsInventoryOK {
	return &GetClusterHostsInventoryOK{}
}

/*GetClusterHostsInventoryOK handles this case with default header values.

successful operation
*/
type GetClusterHostsInventoryOK struct {
	Payload string
}

func (o *GetClusterHostsInventoryOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/stacks/{name}/inventory][%d] getClusterHostsInventoryOK  %+v", 200, o.Payload)
}

func (o *GetClusterHostsInventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

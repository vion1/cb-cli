// Code generated by go-swagger; DO NOT EDIT.

package v1proxyconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPrivatesProxyConfigReader is a Reader for the GetPrivatesProxyConfig structure.
type GetPrivatesProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesProxyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesProxyConfigOK creates a GetPrivatesProxyConfigOK with default headers values
func NewGetPrivatesProxyConfigOK() *GetPrivatesProxyConfigOK {
	return &GetPrivatesProxyConfigOK{}
}

/*GetPrivatesProxyConfigOK handles this case with default header values.

successful operation
*/
type GetPrivatesProxyConfigOK struct {
	Payload []*model.ProxyConfigResponse
}

func (o *GetPrivatesProxyConfigOK) Error() string {
	return fmt.Sprintf("[GET /v1/proxyconfigs/user][%d] getPrivatesProxyConfigOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesProxyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

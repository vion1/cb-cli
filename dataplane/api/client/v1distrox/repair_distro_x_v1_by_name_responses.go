// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// RepairDistroXV1ByNameReader is a Reader for the RepairDistroXV1ByName structure.
type RepairDistroXV1ByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairDistroXV1ByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRepairDistroXV1ByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRepairDistroXV1ByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRepairDistroXV1ByNameOK creates a RepairDistroXV1ByNameOK with default headers values
func NewRepairDistroXV1ByNameOK() *RepairDistroXV1ByNameOK {
	return &RepairDistroXV1ByNameOK{}
}

/*RepairDistroXV1ByNameOK handles this case with default header values.

successful operation
*/
type RepairDistroXV1ByNameOK struct {
	Payload *model.FlowIdentifier
}

func (o *RepairDistroXV1ByNameOK) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/name/{name}/manual_repair][%d] repairDistroXV1ByNameOK  %+v", 200, o.Payload)
}

func (o *RepairDistroXV1ByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.FlowIdentifier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRepairDistroXV1ByNameDefault creates a RepairDistroXV1ByNameDefault with default headers values
func NewRepairDistroXV1ByNameDefault(code int) *RepairDistroXV1ByNameDefault {
	return &RepairDistroXV1ByNameDefault{
		_statusCode: code,
	}
}

/*RepairDistroXV1ByNameDefault handles this case with default header values.

unsuccessful operation
*/
type RepairDistroXV1ByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the repair distro x v1 by name default response
func (o *RepairDistroXV1ByNameDefault) Code() int {
	return o._statusCode
}

func (o *RepairDistroXV1ByNameDefault) Error() string {
	return fmt.Sprintf("[POST /v1/distrox/name/{name}/manual_repair][%d] repairDistroXV1ByName default ", o._statusCode)
}

func (o *RepairDistroXV1ByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

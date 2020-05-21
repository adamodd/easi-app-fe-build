// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/gen/models"
)

// IntakegovernancePOST4Reader is a Reader for the IntakegovernancePOST4 structure.
type IntakegovernancePOST4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntakegovernancePOST4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntakegovernancePOST4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIntakegovernancePOST4BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewIntakegovernancePOST4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewIntakegovernancePOST4NotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntakegovernancePOST4OK creates a IntakegovernancePOST4OK with default headers values
func NewIntakegovernancePOST4OK() *IntakegovernancePOST4OK {
	return &IntakegovernancePOST4OK{}
}

/*IntakegovernancePOST4OK handles this case with default header values.

OK
*/
type IntakegovernancePOST4OK struct {
	Payload *models.IntakegovernancePOSTResponse
}

func (o *IntakegovernancePOST4OK) Error() string {
	return fmt.Sprintf("[POST /intake/governance][%d] intakegovernancePOST4OK  %+v", 200, o.Payload)
}

func (o *IntakegovernancePOST4OK) GetPayload() *models.IntakegovernancePOSTResponse {
	return o.Payload
}

func (o *IntakegovernancePOST4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntakegovernancePOSTResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntakegovernancePOST4BadRequest creates a IntakegovernancePOST4BadRequest with default headers values
func NewIntakegovernancePOST4BadRequest() *IntakegovernancePOST4BadRequest {
	return &IntakegovernancePOST4BadRequest{}
}

/*IntakegovernancePOST4BadRequest handles this case with default header values.

Bad Request
*/
type IntakegovernancePOST4BadRequest struct {
}

func (o *IntakegovernancePOST4BadRequest) Error() string {
	return fmt.Sprintf("[POST /intake/governance][%d] intakegovernancePOST4BadRequest ", 400)
}

func (o *IntakegovernancePOST4BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIntakegovernancePOST4Unauthorized creates a IntakegovernancePOST4Unauthorized with default headers values
func NewIntakegovernancePOST4Unauthorized() *IntakegovernancePOST4Unauthorized {
	return &IntakegovernancePOST4Unauthorized{}
}

/*IntakegovernancePOST4Unauthorized handles this case with default header values.

Access Denied
*/
type IntakegovernancePOST4Unauthorized struct {
}

func (o *IntakegovernancePOST4Unauthorized) Error() string {
	return fmt.Sprintf("[POST /intake/governance][%d] intakegovernancePOST4Unauthorized ", 401)
}

func (o *IntakegovernancePOST4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIntakegovernancePOST4NotImplemented creates a IntakegovernancePOST4NotImplemented with default headers values
func NewIntakegovernancePOST4NotImplemented() *IntakegovernancePOST4NotImplemented {
	return &IntakegovernancePOST4NotImplemented{}
}

/*IntakegovernancePOST4NotImplemented handles this case with default header values.

Internal Server Error
*/
type IntakegovernancePOST4NotImplemented struct {
}

func (o *IntakegovernancePOST4NotImplemented) Error() string {
	return fmt.Sprintf("[POST /intake/governance][%d] intakegovernancePOST4NotImplemented ", 501)
}

func (o *IntakegovernancePOST4NotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

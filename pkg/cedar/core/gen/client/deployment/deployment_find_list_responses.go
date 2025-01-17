// Code generated by go-swagger; DO NOT EDIT.

package deployment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
)

// DeploymentFindListReader is a Reader for the DeploymentFindList structure.
type DeploymentFindListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeploymentFindListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeploymentFindListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeploymentFindListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeploymentFindListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeploymentFindListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeploymentFindListOK creates a DeploymentFindListOK with default headers values
func NewDeploymentFindListOK() *DeploymentFindListOK {
	return &DeploymentFindListOK{}
}

/* DeploymentFindListOK describes a response with status code 200, with default header values.

OK
*/
type DeploymentFindListOK struct {
	Payload *models.DeploymentFindResponse
}

func (o *DeploymentFindListOK) Error() string {
	return fmt.Sprintf("[GET /deployment][%d] deploymentFindListOK  %+v", 200, o.Payload)
}
func (o *DeploymentFindListOK) GetPayload() *models.DeploymentFindResponse {
	return o.Payload
}

func (o *DeploymentFindListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentFindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentFindListBadRequest creates a DeploymentFindListBadRequest with default headers values
func NewDeploymentFindListBadRequest() *DeploymentFindListBadRequest {
	return &DeploymentFindListBadRequest{}
}

/* DeploymentFindListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeploymentFindListBadRequest struct {
	Payload *models.Response
}

func (o *DeploymentFindListBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployment][%d] deploymentFindListBadRequest  %+v", 400, o.Payload)
}
func (o *DeploymentFindListBadRequest) GetPayload() *models.Response {
	return o.Payload
}

func (o *DeploymentFindListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentFindListUnauthorized creates a DeploymentFindListUnauthorized with default headers values
func NewDeploymentFindListUnauthorized() *DeploymentFindListUnauthorized {
	return &DeploymentFindListUnauthorized{}
}

/* DeploymentFindListUnauthorized describes a response with status code 401, with default header values.

Access Denied
*/
type DeploymentFindListUnauthorized struct {
	Payload *models.Response
}

func (o *DeploymentFindListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment][%d] deploymentFindListUnauthorized  %+v", 401, o.Payload)
}
func (o *DeploymentFindListUnauthorized) GetPayload() *models.Response {
	return o.Payload
}

func (o *DeploymentFindListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeploymentFindListInternalServerError creates a DeploymentFindListInternalServerError with default headers values
func NewDeploymentFindListInternalServerError() *DeploymentFindListInternalServerError {
	return &DeploymentFindListInternalServerError{}
}

/* DeploymentFindListInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeploymentFindListInternalServerError struct {
	Payload *models.Response
}

func (o *DeploymentFindListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployment][%d] deploymentFindListInternalServerError  %+v", 500, o.Payload)
}
func (o *DeploymentFindListInternalServerError) GetPayload() *models.Response {
	return o.Payload
}

func (o *DeploymentFindListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

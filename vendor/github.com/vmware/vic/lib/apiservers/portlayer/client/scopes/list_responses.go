package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// ListReader is a Reader for the List structure.
type ListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOK creates a ListOK with default headers values
func NewListOK() *ListOK {
	return &ListOK{}
}

/*ListOK handles this case with default header values.

OK
*/
type ListOK struct {
	Payload []*models.ScopeConfig
}

func (o *ListOK) Error() string {
	return fmt.Sprintf("[GET /scopes/{idName}][%d] listOK  %+v", 200, o.Payload)
}

func (o *ListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNotFound creates a ListNotFound with default headers values
func NewListNotFound() *ListNotFound {
	return &ListNotFound{}
}

/*ListNotFound handles this case with default header values.

Not found
*/
type ListNotFound struct {
	Payload *models.Error
}

func (o *ListNotFound) Error() string {
	return fmt.Sprintf("[GET /scopes/{idName}][%d] listNotFound  %+v", 404, o.Payload)
}

func (o *ListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDefault creates a ListDefault with default headers values
func NewListDefault(code int) *ListDefault {
	return &ListDefault{
		_statusCode: code,
	}
}

/*ListDefault handles this case with default header values.

error
*/
type ListDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list default response
func (o *ListDefault) Code() int {
	return o._statusCode
}

func (o *ListDefault) Error() string {
	return fmt.Sprintf("[GET /scopes/{idName}][%d] List default  %+v", o._statusCode, o.Payload)
}

func (o *ListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

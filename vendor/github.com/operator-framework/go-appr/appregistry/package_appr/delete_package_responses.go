// Code generated by go-swagger; DO NOT EDIT.

package package_appr

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/operator-framework/go-appr/models"
)

// DeletePackageReader is a Reader for the DeletePackage structure.
type DeletePackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeletePackageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeletePackageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePackageOK creates a DeletePackageOK with default headers values
func NewDeletePackageOK() *DeletePackageOK {
	return &DeletePackageOK{}
}

/*DeletePackageOK handles this case with default header values.

successful operation
*/
type DeletePackageOK struct {
	Payload *models.Package
}

func (o *DeletePackageOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] deletePackageOK  %+v", 200, o.Payload)
}

func (o *DeletePackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackageUnauthorized creates a DeletePackageUnauthorized with default headers values
func NewDeletePackageUnauthorized() *DeletePackageUnauthorized {
	return &DeletePackageUnauthorized{}
}

/*DeletePackageUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type DeletePackageUnauthorized struct {
	Payload *models.Error
}

func (o *DeletePackageUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] deletePackageUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePackageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePackageNotFound creates a DeletePackageNotFound with default headers values
func NewDeletePackageNotFound() *DeletePackageNotFound {
	return &DeletePackageNotFound{}
}

/*DeletePackageNotFound handles this case with default header values.

Package not found
*/
type DeletePackageNotFound struct {
	Payload *models.Error
}

func (o *DeletePackageNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/packages/{namespace}/{package}/{release}/{media_type}][%d] deletePackageNotFound  %+v", 404, o.Payload)
}

func (o *DeletePackageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

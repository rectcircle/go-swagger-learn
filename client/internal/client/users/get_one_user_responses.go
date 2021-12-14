// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rectcircle/go-swagger-learn/client/internal/models"
)

// GetOneUserReader is a Reader for the GetOneUser structure.
type GetOneUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOneUserOK creates a GetOneUserOK with default headers values
func NewGetOneUserOK() *GetOneUserOK {
	return &GetOneUserOK{}
}

/* GetOneUserOK describes a response with status code 200, with default header values.

GetOneUserOK get one user o k
*/
type GetOneUserOK struct {
	Payload *GetOneUserOKBody
}

func (o *GetOneUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}][%d] getOneUserOK  %+v", 200, o.Payload)
}
func (o *GetOneUserOK) GetPayload() *GetOneUserOKBody {
	return o.Payload
}

func (o *GetOneUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOneUserOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOneUserOKBody get one user o k body
swagger:model GetOneUserOKBody
*/
type GetOneUserOKBody struct {

	// 错误码
	Code int64 `json:"code,omitempty"`

	// 错误信息
	Message string `json:"message,omitempty"`

	// data
	Data *models.User `json:"data,omitempty"`
}

// Validate validates this get one user o k body
func (o *GetOneUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOneUserOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOneUserOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOneUserOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get one user o k body based on the context it is used
func (o *GetOneUserOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOneUserOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOneUserOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOneUserOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOneUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOneUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetOneUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
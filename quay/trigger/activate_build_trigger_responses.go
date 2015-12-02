package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type ActivateBuildTriggerReader struct {
	formats strfmt.Registry
}

func (o *ActivateBuildTriggerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ActivateBuildTriggerOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ActivateBuildTriggerBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("activateBuildTriggerBadRequest", &result, response.Code())

	case 401:
		var result ActivateBuildTriggerUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("activateBuildTriggerUnauthorized", &result, response.Code())

	case 403:
		var result ActivateBuildTriggerForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("activateBuildTriggerForbidden", &result, response.Code())

	case 404:
		var result ActivateBuildTriggerNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("activateBuildTriggerNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

/*ActivateBuildTriggerOK

Successful invocation
*/
type ActivateBuildTriggerOK struct {
}

func (o *ActivateBuildTriggerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ActivateBuildTriggerBadRequest

Bad Request
*/
type ActivateBuildTriggerBadRequest struct {
	Payload *models.GeneralError
}

func (o *ActivateBuildTriggerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*ActivateBuildTriggerUnauthorized

Session required
*/
type ActivateBuildTriggerUnauthorized struct {
}

func (o *ActivateBuildTriggerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ActivateBuildTriggerForbidden

Unauthorized access
*/
type ActivateBuildTriggerForbidden struct {
}

func (o *ActivateBuildTriggerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ActivateBuildTriggerNotFound

Not found
*/
type ActivateBuildTriggerNotFound struct {
}

func (o *ActivateBuildTriggerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
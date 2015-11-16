package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

type DeleteBuildTriggerReader struct {
	formats strfmt.Registry
}

func (o *DeleteBuildTriggerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		var result DeleteBuildTriggerNoContent
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result DeleteBuildTriggerBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteBuildTriggerBadRequest", &result, response.Code())

	case 401:
		var result DeleteBuildTriggerUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteBuildTriggerUnauthorized", &result, response.Code())

	case 403:
		var result DeleteBuildTriggerForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteBuildTriggerForbidden", &result, response.Code())

	case 404:
		var result DeleteBuildTriggerNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("deleteBuildTriggerNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Deleted
*/
type DeleteBuildTriggerNoContent struct {
}

func (o *DeleteBuildTriggerNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type DeleteBuildTriggerBadRequest struct {
	Payload *models.GeneralError
}

func (o *DeleteBuildTriggerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*
Session required
*/
type DeleteBuildTriggerUnauthorized struct {
}

func (o *DeleteBuildTriggerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type DeleteBuildTriggerForbidden struct {
}

func (o *DeleteBuildTriggerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type DeleteBuildTriggerNotFound struct {
}

func (o *DeleteBuildTriggerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListUserLogsParams creates a new ListUserLogsParams object
// with the default values initialized.
func NewListUserLogsParams() *ListUserLogsParams {
	var ()
	return &ListUserLogsParams{}
}

/*ListUserLogsParams contains all the parameters to send to the API endpoint
for the list user logs operation typically these are written to a http.Request
*/
type ListUserLogsParams struct {

	/*Endtime
	  Latest time to which to get logs. (%m/%d/%Y %Z)

	*/
	Endtime *string
	/*NextPage
	  The page token for the next page

	*/
	NextPage *string
	/*Performer
	  Username for which to filter logs.

	*/
	Performer *string
	/*Starttime
	  Earliest time from which to get logs. (%m/%d/%Y %Z)

	*/
	Starttime *string
}

// WithEndtime adds the endtime to the list user logs params
func (o *ListUserLogsParams) WithEndtime(Endtime *string) *ListUserLogsParams {
	o.Endtime = Endtime
	return o
}

// WithNextPage adds the nextPage to the list user logs params
func (o *ListUserLogsParams) WithNextPage(NextPage *string) *ListUserLogsParams {
	o.NextPage = NextPage
	return o
}

// WithPerformer adds the performer to the list user logs params
func (o *ListUserLogsParams) WithPerformer(Performer *string) *ListUserLogsParams {
	o.Performer = Performer
	return o
}

// WithStarttime adds the starttime to the list user logs params
func (o *ListUserLogsParams) WithStarttime(Starttime *string) *ListUserLogsParams {
	o.Starttime = Starttime
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime string
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

	}

	if o.NextPage != nil {

		// query param next_page
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("next_page", qNextPage); err != nil {
				return err
			}
		}

	}

	if o.Performer != nil {

		// query param performer
		var qrPerformer string
		if o.Performer != nil {
			qrPerformer = *o.Performer
		}
		qPerformer := qrPerformer
		if qPerformer != "" {
			if err := r.SetQueryParam("performer", qPerformer); err != nil {
				return err
			}
		}

	}

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime string
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
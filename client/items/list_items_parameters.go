// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListItemsParams creates a new ListItemsParams object
// with the default values initialized.
func NewListItemsParams() *ListItemsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &ListItemsParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListItemsParamsWithTimeout creates a new ListItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListItemsParamsWithTimeout(timeout time.Duration) *ListItemsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &ListItemsParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewListItemsParamsWithContext creates a new ListItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListItemsParamsWithContext(ctx context.Context) *ListItemsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &ListItemsParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewListItemsParamsWithHTTPClient creates a new ListItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListItemsParamsWithHTTPClient(client *http.Client) *ListItemsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &ListItemsParams{
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*ListItemsParams contains all the parameters to send to the API endpoint
for the list items operation typically these are written to a http.Request
*/
type ListItemsParams struct {

	/*PageSize
	  Amount of items to return in a single page

	*/
	PageSize *int32
	/*SinceID
	  The last id that was seen.

	*/
	SinceID *int64
	/*Status
	  the status to filter by

	*/
	Status []string
	/*Tags
	  the tags to filter by

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list items params
func (o *ListItemsParams) WithTimeout(timeout time.Duration) *ListItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list items params
func (o *ListItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list items params
func (o *ListItemsParams) WithContext(ctx context.Context) *ListItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list items params
func (o *ListItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list items params
func (o *ListItemsParams) WithHTTPClient(client *http.Client) *ListItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list items params
func (o *ListItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the list items params
func (o *ListItemsParams) WithPageSize(pageSize *int32) *ListItemsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list items params
func (o *ListItemsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSinceID adds the sinceID to the list items params
func (o *ListItemsParams) WithSinceID(sinceID *int64) *ListItemsParams {
	o.SetSinceID(sinceID)
	return o
}

// SetSinceID adds the sinceId to the list items params
func (o *ListItemsParams) SetSinceID(sinceID *int64) {
	o.SinceID = sinceID
}

// WithStatus adds the status to the list items params
func (o *ListItemsParams) WithStatus(status []string) *ListItemsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list items params
func (o *ListItemsParams) SetStatus(status []string) {
	o.Status = status
}

// WithTags adds the tags to the list items params
func (o *ListItemsParams) WithTags(tags []string) *ListItemsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the list items params
func (o *ListItemsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *ListItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.SinceID != nil {

		// query param sinceId
		var qrSinceID int64
		if o.SinceID != nil {
			qrSinceID = *o.SinceID
		}
		qSinceID := swag.FormatInt64(qrSinceID)
		if qSinceID != "" {
			if err := r.SetQueryParam("sinceId", qSinceID); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "pipes")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

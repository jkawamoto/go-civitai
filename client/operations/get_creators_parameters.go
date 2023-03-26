// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCreatorsParams creates a new GetCreatorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCreatorsParams() *GetCreatorsParams {
	return &GetCreatorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreatorsParamsWithTimeout creates a new GetCreatorsParams object
// with the ability to set a timeout on a request.
func NewGetCreatorsParamsWithTimeout(timeout time.Duration) *GetCreatorsParams {
	return &GetCreatorsParams{
		timeout: timeout,
	}
}

// NewGetCreatorsParamsWithContext creates a new GetCreatorsParams object
// with the ability to set a context for a request.
func NewGetCreatorsParamsWithContext(ctx context.Context) *GetCreatorsParams {
	return &GetCreatorsParams{
		Context: ctx,
	}
}

// NewGetCreatorsParamsWithHTTPClient creates a new GetCreatorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCreatorsParamsWithHTTPClient(client *http.Client) *GetCreatorsParams {
	return &GetCreatorsParams{
		HTTPClient: client,
	}
}

/*
GetCreatorsParams contains all the parameters to send to the API endpoint

	for the get creators operation.

	Typically these are written to a http.Request.
*/
type GetCreatorsParams struct {

	/* Limit.

	   The number of results to be returned per page. This can be a number between 0 and 200. By default, each page will return 20 results. If set to 0, it'll return all the creators.

	   Format: int64
	   Default: 20
	*/
	Limit *int64

	/* Page.

	   The page from which to start fetching creators.

	   Format: int64
	*/
	Page *int64

	/* Query.

	   Search query to filter creators by username.
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get creators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCreatorsParams) WithDefaults() *GetCreatorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get creators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCreatorsParams) SetDefaults() {
	var (
		limitDefault = int64(20)
	)

	val := GetCreatorsParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get creators params
func (o *GetCreatorsParams) WithTimeout(timeout time.Duration) *GetCreatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get creators params
func (o *GetCreatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get creators params
func (o *GetCreatorsParams) WithContext(ctx context.Context) *GetCreatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get creators params
func (o *GetCreatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get creators params
func (o *GetCreatorsParams) WithHTTPClient(client *http.Client) *GetCreatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get creators params
func (o *GetCreatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get creators params
func (o *GetCreatorsParams) WithLimit(limit *int64) *GetCreatorsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get creators params
func (o *GetCreatorsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get creators params
func (o *GetCreatorsParams) WithPage(page *int64) *GetCreatorsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get creators params
func (o *GetCreatorsParams) SetPage(page *int64) {
	o.Page = page
}

// WithQuery adds the query to the get creators params
func (o *GetCreatorsParams) WithQuery(query *string) *GetCreatorsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get creators params
func (o *GetCreatorsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
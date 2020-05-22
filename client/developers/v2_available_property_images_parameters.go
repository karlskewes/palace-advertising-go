// Code generated by go-swagger; DO NOT EDIT.

package developers

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
)

// NewV2AvailablePropertyImagesParams creates a new V2AvailablePropertyImagesParams object
// with the default values initialized.
func NewV2AvailablePropertyImagesParams() *V2AvailablePropertyImagesParams {
	var ()
	return &V2AvailablePropertyImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2AvailablePropertyImagesParamsWithTimeout creates a new V2AvailablePropertyImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2AvailablePropertyImagesParamsWithTimeout(timeout time.Duration) *V2AvailablePropertyImagesParams {
	var ()
	return &V2AvailablePropertyImagesParams{

		timeout: timeout,
	}
}

// NewV2AvailablePropertyImagesParamsWithContext creates a new V2AvailablePropertyImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2AvailablePropertyImagesParamsWithContext(ctx context.Context) *V2AvailablePropertyImagesParams {
	var ()
	return &V2AvailablePropertyImagesParams{

		Context: ctx,
	}
}

// NewV2AvailablePropertyImagesParamsWithHTTPClient creates a new V2AvailablePropertyImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2AvailablePropertyImagesParamsWithHTTPClient(client *http.Client) *V2AvailablePropertyImagesParams {
	var ()
	return &V2AvailablePropertyImagesParams{
		HTTPClient: client,
	}
}

/*V2AvailablePropertyImagesParams contains all the parameters to send to the API endpoint
for the v2 available property images operation typically these are written to a http.Request
*/
type V2AvailablePropertyImagesParams struct {

	/*PropertyCode
	  Internal 'Property Code'

	*/
	PropertyCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) WithTimeout(timeout time.Duration) *V2AvailablePropertyImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) WithContext(ctx context.Context) *V2AvailablePropertyImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) WithHTTPClient(client *http.Client) *V2AvailablePropertyImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPropertyCode adds the propertyCode to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) WithPropertyCode(propertyCode string) *V2AvailablePropertyImagesParams {
	o.SetPropertyCode(propertyCode)
	return o
}

// SetPropertyCode adds the propertyCode to the v2 available property images params
func (o *V2AvailablePropertyImagesParams) SetPropertyCode(propertyCode string) {
	o.PropertyCode = propertyCode
}

// WriteToRequest writes these params to a swagger request
func (o *V2AvailablePropertyImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param PropertyCode
	if err := r.SetPathParam("PropertyCode", o.PropertyCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

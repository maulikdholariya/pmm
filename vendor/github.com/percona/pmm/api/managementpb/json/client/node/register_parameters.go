// Code generated by go-swagger; DO NOT EDIT.

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRegisterParams creates a new RegisterParams object
// with the default values initialized.
func NewRegisterParams() *RegisterParams {
	var ()
	return &RegisterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterParamsWithTimeout creates a new RegisterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterParamsWithTimeout(timeout time.Duration) *RegisterParams {
	var ()
	return &RegisterParams{

		timeout: timeout,
	}
}

// NewRegisterParamsWithContext creates a new RegisterParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterParamsWithContext(ctx context.Context) *RegisterParams {
	var ()
	return &RegisterParams{

		Context: ctx,
	}
}

// NewRegisterParamsWithHTTPClient creates a new RegisterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterParamsWithHTTPClient(client *http.Client) *RegisterParams {
	var ()
	return &RegisterParams{
		HTTPClient: client,
	}
}

/*RegisterParams contains all the parameters to send to the API endpoint
for the register operation typically these are written to a http.Request
*/
type RegisterParams struct {

	/*Body*/
	Body RegisterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register params
func (o *RegisterParams) WithTimeout(timeout time.Duration) *RegisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register params
func (o *RegisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register params
func (o *RegisterParams) WithContext(ctx context.Context) *RegisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register params
func (o *RegisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register params
func (o *RegisterParams) WithHTTPClient(client *http.Client) *RegisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register params
func (o *RegisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register params
func (o *RegisterParams) WithBody(body RegisterBody) *RegisterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register params
func (o *RegisterParams) SetBody(body RegisterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

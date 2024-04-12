// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package bgp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetBgpRoutesParams creates a new GetBgpRoutesParams object
//
// There are no default values defined in the spec.
func NewGetBgpRoutesParams() GetBgpRoutesParams {

	return GetBgpRoutesParams{}
}

// GetBgpRoutesParams contains all the bound params for the get bgp routes operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetBgpRoutes
type GetBgpRoutesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Address Family Indicator (AFI) of a BGP route
	  Required: true
	  In: query
	*/
	Afi string
	/*IP address specifying a BGP neighbor.
	Has to be specified only when table type is adj-rib-in or adj-rib-out.

	  In: query
	*/
	Neighbor *string
	/*Autonomous System Number (ASN) identifying a BGP virtual router instance.
	If not specified, all virtual router instances are selected.

	  In: query
	*/
	RouterAsn *int64
	/*Subsequent Address Family Indicator (SAFI) of a BGP route
	  Required: true
	  In: query
	*/
	Safi string
	/*BGP Routing Information Base (RIB) table type
	  Required: true
	  In: query
	*/
	TableType string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBgpRoutesParams() beforehand.
func (o *GetBgpRoutesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAfi, qhkAfi, _ := qs.GetOK("afi")
	if err := o.bindAfi(qAfi, qhkAfi, route.Formats); err != nil {
		res = append(res, err)
	}

	qNeighbor, qhkNeighbor, _ := qs.GetOK("neighbor")
	if err := o.bindNeighbor(qNeighbor, qhkNeighbor, route.Formats); err != nil {
		res = append(res, err)
	}

	qRouterAsn, qhkRouterAsn, _ := qs.GetOK("router_asn")
	if err := o.bindRouterAsn(qRouterAsn, qhkRouterAsn, route.Formats); err != nil {
		res = append(res, err)
	}

	qSafi, qhkSafi, _ := qs.GetOK("safi")
	if err := o.bindSafi(qSafi, qhkSafi, route.Formats); err != nil {
		res = append(res, err)
	}

	qTableType, qhkTableType, _ := qs.GetOK("table_type")
	if err := o.bindTableType(qTableType, qhkTableType, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAfi binds and validates parameter Afi from query.
func (o *GetBgpRoutesParams) bindAfi(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("afi", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("afi", "query", raw); err != nil {
		return err
	}
	o.Afi = raw

	return nil
}

// bindNeighbor binds and validates parameter Neighbor from query.
func (o *GetBgpRoutesParams) bindNeighbor(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Neighbor = &raw

	return nil
}

// bindRouterAsn binds and validates parameter RouterAsn from query.
func (o *GetBgpRoutesParams) bindRouterAsn(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("router_asn", "query", "int64", raw)
	}
	o.RouterAsn = &value

	return nil
}

// bindSafi binds and validates parameter Safi from query.
func (o *GetBgpRoutesParams) bindSafi(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("safi", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("safi", "query", raw); err != nil {
		return err
	}
	o.Safi = raw

	return nil
}

// bindTableType binds and validates parameter TableType from query.
func (o *GetBgpRoutesParams) bindTableType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("table_type", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("table_type", "query", raw); err != nil {
		return err
	}
	o.TableType = raw

	if err := o.validateTableType(formats); err != nil {
		return err
	}

	return nil
}

// validateTableType carries on validations for parameter TableType
func (o *GetBgpRoutesParams) validateTableType(formats strfmt.Registry) error {

	if err := validate.EnumCase("table_type", "query", o.TableType, []interface{}{"loc-rib", "adj-rib-in", "adj-rib-out"}, true); err != nil {
		return err
	}

	return nil
}
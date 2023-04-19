// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package responses

import (
    "github.com/pb33f/libopenapi-validator/errors"
    "github.com/pb33f/libopenapi/datamodel/high/v3"
    "net/http"
)

// ResponseBodyValidator is an interface that defines the methods for validating response bodies for Operations.
//
//	ValidateResponseBody method accepts an *http.Request and returns true if validation passed,
//	                     false if validation failed and a slice of ValidationError pointers.
type ResponseBodyValidator interface {
    ValidateResponseBody(request *http.Request, response *http.Response) (bool, []*errors.ValidationError)
    SetPathItem(path *v3.PathItem, pathValue string)
}

// SetPathItem will set the pathItem for the ResponseBodyValidator, all validations will be performed
// against this pathItem otherwise if not set, each validation will perform a lookup for the
// pathItem based on the *http.Request
func (v *responseBodyValidator) SetPathItem(path *v3.PathItem, pathValue string) {
    v.pathItem = path
    v.pathValue = pathValue
}

// NewResponseBodyValidator will create a new ResponseBodyValidator from an OpenAPI 3+ document
func NewResponseBodyValidator(document *v3.Document) ResponseBodyValidator {
    return &responseBodyValidator{document: document}
}

type responseBodyValidator struct {
    document  *v3.Document
    pathItem  *v3.PathItem
    pathValue string
    errors    []*errors.ValidationError
}
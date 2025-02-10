// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
	"net/http"
)

type PatchConnectionRequest struct {
	ConnectionID           string                        `pathParam:"style=simple,explode=false,name=connectionId"`
	ConnectionPatchRequest shared.ConnectionPatchRequest `request:"mediaType=application/json"`
}

func (o *PatchConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *PatchConnectionRequest) GetConnectionPatchRequest() shared.ConnectionPatchRequest {
	if o == nil {
		return shared.ConnectionPatchRequest{}
	}
	return o.ConnectionPatchRequest
}

type PatchConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse        *http.Response
	ConnectionResponse *shared.ConnectionResponse
}

func (o *PatchConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PatchConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PatchConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PatchConnectionResponse) GetConnectionResponse() *shared.ConnectionResponse {
	if o == nil {
		return nil
	}
	return o.ConnectionResponse
}

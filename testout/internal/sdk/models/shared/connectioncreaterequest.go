// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ConnectionCreateRequest struct {
	// Optional name of the connection
	Name          *string `json:"name,omitempty"`
	SourceID      string  `json:"sourceId"`
	DestinationID string  `json:"destinationId"`
}

func (o *ConnectionCreateRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionCreateRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *ConnectionCreateRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

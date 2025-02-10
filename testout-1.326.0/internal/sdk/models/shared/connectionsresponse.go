// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionsResponse struct {
	Previous *string              `json:"previous,omitempty"`
	Next     *string              `json:"next,omitempty"`
	Data     []ConnectionResponse `json:"data"`
}

func (o *ConnectionsResponse) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ConnectionsResponse) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ConnectionsResponse) GetData() []ConnectionResponse {
	if o == nil {
		return []ConnectionResponse{}
	}
	return o.Data
}

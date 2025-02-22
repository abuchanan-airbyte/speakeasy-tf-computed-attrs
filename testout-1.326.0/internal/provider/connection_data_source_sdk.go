// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/speakeasy/terraform-provider-terraform/internal/provider/types"
	"github.com/speakeasy/terraform-provider-terraform/internal/sdk/models/shared"
)

func (r *ConnectionDataSourceModel) RefreshFromSharedConnectionResponse(resp *shared.ConnectionResponse) {
	if resp != nil {
		if resp.Configurations == nil {
			r.Configurations = nil
		} else {
			r.Configurations = &tfTypes.StreamConfigurations{}
			r.Configurations.Streams = []tfTypes.StreamConfiguration{}
			if len(r.Configurations.Streams) > len(resp.Configurations.Streams) {
				r.Configurations.Streams = r.Configurations.Streams[:len(resp.Configurations.Streams)]
			}
			for streamsCount, streamsItem := range resp.Configurations.Streams {
				var streams1 tfTypes.StreamConfiguration
				streams1.CursorField = []types.String{}
				for _, v := range streamsItem.CursorField {
					streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
				}
				streams1.Name = types.StringValue(streamsItem.Name)
				streams1.PrimaryKey = nil
				for _, primaryKeyItem := range streamsItem.PrimaryKey {
					var primaryKey1 []types.String
					primaryKey1 = []types.String{}
					for _, v := range primaryKeyItem {
						primaryKey1 = append(primaryKey1, types.StringValue(v))
					}
					streams1.PrimaryKey = append(streams1.PrimaryKey, primaryKey1)
				}
				if streamsCount+1 > len(r.Configurations.Streams) {
					r.Configurations.Streams = append(r.Configurations.Streams, streams1)
				} else {
					r.Configurations.Streams[streamsCount].CursorField = streams1.CursorField
					r.Configurations.Streams[streamsCount].Name = streams1.Name
					r.Configurations.Streams[streamsCount].PrimaryKey = streams1.PrimaryKey
				}
			}
		}
		r.ConnectionID = types.StringValue(resp.ConnectionID)
		r.CreatedAt = types.Int64Value(resp.CreatedAt)
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.WorkspaceID = types.StringPointerValue(resp.WorkspaceID)
	}
}

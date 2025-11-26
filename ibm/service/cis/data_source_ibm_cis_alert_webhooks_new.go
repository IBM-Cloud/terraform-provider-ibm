package cis

import (
	"context"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSource = &CISWebhooksDataSource{}

type CISWebhooksDataSource struct {
	client interface{}
}

func NewCISWebhooksDataSource(client interface{}) datasource.DataSource {
	return &CISWebhooksDataSource{
		client: client,
	}
}

func (d *CISWebhooksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cis_alert_webhooks_new"
}

func (d *CISWebhooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "CIS alert webhooks data source",

		Attributes: map[string]schema.Attribute{
			"cis_id": schema.StringAttribute{
				MarkdownDescription: "CIS instance crn",
				Required:            true,
			},
			"cis_webhooks": schema.ListAttribute{
				MarkdownDescription: "Collection of Webhook details",
				Computed:            true,
				ElementType: types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"webhook_id": types.StringType,
						"name":       types.StringType,
						"url":        types.StringType,
						"type":       types.StringType,
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Data source identifier",
			},
		},
	}
}

func (d *CISWebhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CISWebhooksDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	sess, err := d.client.(conns.ClientSession).CisWebhookSession()
	if err != nil {
		resp.Diagnostics.AddError("Error getting CIS webhook session", err.Error())
		return
	}

	crn := data.CisID.ValueString()
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewListWebhooksOptions()

	result, resp1, err := sess.ListWebhooks(opt)
	if err != nil || result == nil {
		resp.Diagnostics.AddError("Error listing webhooks", flex.FmtErrorf("[ERROR] Error Listing all Webhooks: %s %s", err, resp1).Error())
		return
	}

	webhooks := make([]attr.Value, 0)

	for _, instance := range result.Result {
		webhook, _ := types.ObjectValue(
			map[string]attr.Type{
				"webhook_id": types.StringType,
				"name":       types.StringType,
				"url":        types.StringType,
				"type":       types.StringType,
			},
			map[string]attr.Value{
				"webhook_id": basetypes.NewStringValue(*instance.ID),
				"name":       basetypes.NewStringValue(*instance.Name),
				"url":        basetypes.NewStringValue(*instance.URL),
				"type":       basetypes.NewStringValue(*instance.Type),
			},
		)
		webhooks = append(webhooks, webhook)
	}

	data.CisWebhooks, _ = types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"webhook_id": types.StringType,
				"name":       types.StringType,
				"url":        types.StringType,
				"type":       types.StringType,
			},
		},
		webhooks,
	)

	data.Id = basetypes.NewStringValue(time.Now().UTC().String())
	data.CisID = basetypes.NewStringValue(crn)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type CISWebhooksDataSourceModel struct {
	CisID       types.String `tfsdk:"cis_id"`
	CisWebhooks types.List   `tfsdk:"cis_webhooks"`
	Id          types.String `tfsdk:"id"`
}

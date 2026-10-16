// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/action/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ action.Action              = &crPlanUpgradeAction{}
	_ action.ActionWithConfigure = &crPlanUpgradeAction{}
)

// NewCrPlanUpgradeAction is a one-way,
// account-level operation with nothing to destroy, so it is modelled as an action.
func NewCrPlanUpgradeAction() action.Action {
	return &crPlanUpgradeAction{}
}

type crPlanUpgradeAction struct {
	client *containerregistryv1.ContainerRegistryV1
}

type crPlanUpgradeModel struct {
	Plan     types.String `tfsdk:"plan"`
	Endpoint types.String `tfsdk:"endpoint"`
}

func (a *crPlanUpgradeAction) Metadata(ctx context.Context, req action.MetadataRequest, resp *action.MetadataResponse) {
	resp.TypeName = "ibm_cr_plan_upgrade"
}

func (a *crPlanUpgradeAction) Schema(ctx context.Context, req action.SchemaRequest, resp *action.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Upgrades the IBM Cloud Container Registry plan of the account for one registry region. The action reads the current plan first and does nothing if the account already has the requested plan.",
		Attributes: map[string]schema.Attribute{
			"plan": schema.StringAttribute{
				Optional:    true,
				Description: "The target plan. Only 'Standard' is accepted by the API today. Default: Standard.",
			},
			"endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "Regional Container Registry API endpoint, for example 'https://us.icr.io' or 'https://private.de.icr.io'. Defaults to the endpoint the provider is configured with.",
			},
		},
	}
}

func (a *crPlanUpgradeAction) Configure(ctx context.Context, req action.ConfigureRequest, resp *action.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, err := crClientFromProviderData(req.ProviderData)
	if err != nil {
		resp.Diagnostics.AddError("Unable to Create Container Registry Client", err.Error())
		return
	}
	a.client = client
}

func (a *crPlanUpgradeAction) Invoke(ctx context.Context, req action.InvokeRequest, resp *action.InvokeResponse) {
	var config crPlanUpgradeModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client, err := crClientForEndpoint(a.client, config.Endpoint)
	if err != nil {
		resp.Diagnostics.AddError("Invalid Container Registry Endpoint", err.Error())
		return
	}

	plan := "Standard"
	if !config.Plan.IsNull() && config.Plan.ValueString() != "" {
		plan = config.Plan.ValueString()
	}

	current, _, err := client.GetPlansWithContext(ctx, client.NewGetPlansOptions())
	if err != nil {
		resp.Diagnostics.AddError("Unable to Read Container Registry Plan", err.Error())
		return
	}
	if current != nil && current.Plan != nil && strings.EqualFold(*current.Plan, plan) {
		resp.SendProgress(action.InvokeProgressEvent{
			Message: fmt.Sprintf("Container Registry plan is already '%s', nothing to do", *current.Plan),
		})
		return
	}

	if _, err := client.UpdatePlansWithContext(ctx, client.NewUpdatePlansOptions().SetPlan(plan)); err != nil {
		resp.Diagnostics.AddError("Unable to Upgrade Container Registry Plan", err.Error())
		return
	}

	resp.SendProgress(action.InvokeProgressEvent{
		Message: fmt.Sprintf("Container Registry plan upgraded to '%s'", plan),
	})
}

// crClientFromProviderData pulls the Container Registry client out of the shared
// provider session. It is shared by every framework construct in this package.
func crClientFromProviderData(providerData any) (*containerregistryv1.ContainerRegistryV1, error) {
	session, ok := providerData.(conns.ClientSession)
	if !ok {
		return nil, fmt.Errorf("expected conns.ClientSession, got: %T. Please report this issue to the provider developers", providerData)
	}
	return session.ContainerRegistryV1()
}

// crClientForEndpoint returns a copy of the client pointed at the given regional
// endpoint, so one provider block can manage several registry regions.
func crClientForEndpoint(client *containerregistryv1.ContainerRegistryV1, endpoint types.String) (*containerregistryv1.ContainerRegistryV1, error) {
	if client == nil {
		return nil, fmt.Errorf("the Container Registry client is not configured")
	}
	if endpoint.IsNull() || endpoint.IsUnknown() || endpoint.ValueString() == "" {
		return client, nil
	}
	url := endpoint.ValueString()
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + strings.TrimPrefix(url, "//")
	}
	clone := client.Clone()
	if err := clone.SetServiceURL(url); err != nil {
		return nil, err
	}
	return clone, nil
}

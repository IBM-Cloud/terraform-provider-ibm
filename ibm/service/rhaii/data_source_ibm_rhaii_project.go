// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &rhaiiProjectDataSource{}
	_ datasource.DataSourceWithConfigure = &rhaiiProjectDataSource{}
)

// NewRhaiiProjectDataSource returns the ibm_rhaii_project data source, which
// reads an existing Red Hat AI Inference project by name or by ID.
func NewRhaiiProjectDataSource() datasource.DataSource {
	return &rhaiiProjectDataSource{}
}

type rhaiiProjectDataSource struct {
	meta conns.ClientSession
}

type rhaiiProjectDataSourceModel struct {
	Name            types.String `tfsdk:"name"`
	Identifier      types.String `tfsdk:"identifier"`
	ResourceGroupID types.String `tfsdk:"resource_group_id"`
	Location        types.String `tfsdk:"location"`
	PlanName        types.String `tfsdk:"plan_name"`
	Tags            types.Set    `tfsdk:"tags"`
	AccessTags      types.Set    `tfsdk:"access_tags"`
	rhaiiProjectComputedModel
}

func (d *rhaiiProjectDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rhaii_project"
}

func (d *rhaiiProjectDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The name of the project.",
		},
		"identifier": schema.StringAttribute{
			Optional:    true,
			Description: "The ID (GUID) or CRN of the project.",
			Validators: []validator.String{
				stringvalidator.ExactlyOneOf(path.MatchRoot("name"), path.MatchRoot("identifier")),
				stringvalidator.ConflictsWith(path.MatchRoot("resource_group_id"), path.MatchRoot("location")),
			},
		},
		"resource_group_id": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The ID of the resource group to search in.",
		},
		"location": schema.StringAttribute{
			Optional:    true,
			Computed:    true,
			Description: "The region of the project.",
		},
		"plan_name": schema.StringAttribute{
			Computed:    true,
			Description: "The name of the pricing plan of the project.",
		},
		"tags": schema.SetAttribute{
			ElementType: types.StringType,
			Computed:    true,
			Description: "The user tags of the project.",
		},
		"access_tags": schema.SetAttribute{
			ElementType: types.StringType,
			Computed:    true,
			Description: "The access management tags of the project.",
		},
	}
	for k, v := range rhaiiProjectDataSourceComputedAttributes() {
		attributes[k] = v
	}
	resp.Schema = schema.Schema{
		Description: "Reads a Red Hat AI Inference project.",
		Attributes:  attributes,
	}
}

func (d *rhaiiProjectDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	meta, ok := req.ProviderData.(conns.ClientSession)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected conns.ClientSession, got %T.", req.ProviderData))
		return
	}
	d.meta = meta
}

func (d *rhaiiProjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var m rhaiiProjectDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}

	rcClient, err := d.meta.ResourceControllerV2API()
	if err != nil {
		resp.Diagnostics.AddError("Error initializing the Resource Controller client", err.Error())
		return
	}

	var instance *rc.ResourceInstance
	if !m.Identifier.IsNull() {
		id := m.Identifier.ValueString()
		instance, _, err = rcClient.GetResourceInstanceWithContext(ctx, &rc.GetResourceInstanceOptions{ID: &id})
		if err != nil {
			resp.Diagnostics.AddError("Error reading the Red Hat AI Inference project", fmt.Sprintf("GetResourceInstanceWithContext failed: %s", err))
			return
		}
		if !isRhaiiInstance(instance) {
			resp.Diagnostics.AddError("Not a Red Hat AI Inference project",
				fmt.Sprintf("Resource instance %s is not a Red Hat AI Inference project (service is %q, expected %q).", id, strDeref(instance.ResourceID), rhaiiServiceName))
			return
		}
	} else {
		instance, err = rhaiiFindProjectByName(ctx, rcClient, m.Name.ValueString(), m.ResourceGroupID.ValueString(), m.Location.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Error finding the Red Hat AI Inference project", err.Error())
			return
		}
	}

	values, diags := flattenRhaiiProject(ctx, d.meta, instance, &m.rhaiiProjectComputedModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	m.Name = types.StringValue(values.Name)
	m.Location = types.StringValue(values.Location)
	m.PlanName = types.StringValue(values.PlanName)
	m.ResourceGroupID = types.StringValue(values.ResourceGroupID)

	tags, accessTags, diags := rhaiiReadTags(ctx, d.meta, strDeref(instance.CRN))
	resp.Diagnostics.Append(diags...)
	m.Tags = tags
	m.AccessTags = accessTags

	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}

// rhaiiFindProjectByName lists the RHAII instances with a name and returns the
// only match. Removed and pending reclamation instances are ignored.
func rhaiiFindProjectByName(ctx context.Context, rcClient *rc.ResourceControllerV2, name, resourceGroupID, location string) (*rc.ResourceInstance, error) {
	listOptions := &rc.ListResourceInstancesOptions{
		Name:       &name,
		ResourceID: ptr(rhaiiServiceName),
	}
	if resourceGroupID != "" {
		listOptions.ResourceGroupID = &resourceGroupID
	}

	var matches []rc.ResourceInstance
	for {
		list, _, err := rcClient.ListResourceInstancesWithContext(ctx, listOptions)
		if err != nil {
			return nil, fmt.Errorf("ListResourceInstancesWithContext failed: %s", err)
		}
		for i := range list.Resources {
			instance := list.Resources[i]
			if isRhaiiInstanceGone(&instance) {
				continue
			}
			if location != "" && strDeref(instance.RegionID) != location {
				continue
			}
			matches = append(matches, instance)
		}
		start, err := rhaiiNextStart(list.NextURL)
		if err != nil {
			return nil, fmt.Errorf("error parsing next_url of the resource instance list: %s", err)
		}
		if start == "" {
			break
		}
		listOptions.Start = &start
	}

	switch len(matches) {
	case 0:
		return nil, fmt.Errorf("no Red Hat AI Inference project found with name %q. Check the name, resource_group_id and location", name)
	case 1:
		return &matches[0], nil
	default:
		return nil, fmt.Errorf("found %d Red Hat AI Inference projects with name %q. Set resource_group_id or location to narrow the search, or use identifier", len(matches), name)
	}
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dsschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	rsschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rhaiiProjectComputedModel holds the computed attributes of a project that the
// resource and the data source have in common. It is embedded in both models.
type rhaiiProjectComputedModel struct {
	ID               types.String `tfsdk:"id"`
	ProjectID        types.String `tfsdk:"project_id"`
	Endpoint         types.String `tfsdk:"endpoint"`
	PrivateEndpoint  types.String `tfsdk:"private_endpoint"`
	GUID             types.String `tfsdk:"guid"`
	CRN              types.String `tfsdk:"crn"`
	Service          types.String `tfsdk:"service"`
	State            types.String `tfsdk:"state"`
	DashboardURL     types.String `tfsdk:"dashboard_url"`
	AccountID        types.String `tfsdk:"account_id"`
	ResourceGroupCRN types.String `tfsdk:"resource_group_crn"`
	PlanID           types.String `tfsdk:"plan_id"`
	TargetCRN        types.String `tfsdk:"target_crn"`
	CreatedAt        types.String `tfsdk:"created_at"`
	CreatedBy        types.String `tfsdk:"created_by"`
	UpdatedAt        types.String `tfsdk:"updated_at"`
	UpdatedBy        types.String `tfsdk:"updated_by"`
	Locked           types.Bool   `tfsdk:"locked"`
	LastOperation    types.List   `tfsdk:"last_operation"`
}

// rhaiiLastOperationAttrTypes is the object type of a last_operation element.
var rhaiiLastOperationAttrTypes = map[string]attr.Type{
	"type":        types.StringType,
	"state":       types.StringType,
	"async":       types.BoolType,
	"description": types.StringType,
}

// rhaiiComputedAttr describes one computed string or bool attribute once, so
// that the resource and the data source schemas cannot drift apart.
type rhaiiComputedAttr struct {
	name        string
	description string
	isBool      bool
	// stable attributes never change after create; the resource keeps their
	// state value in the plan instead of showing "known after apply".
	stable bool
}

var rhaiiProjectComputedAttrs = []rhaiiComputedAttr{
	{name: "id", description: "The unique identifier of the project. This is the CRN of the resource instance.", stable: true},
	{name: "project_id", description: "The ID of the project. Use this value as `project_id` in the Red Hat AI Inference API.", stable: true},
	{name: "endpoint", description: "The public base URL of the Red Hat AI Inference API for this project.", stable: true},
	{name: "private_endpoint", description: "The private base URL of the Red Hat AI Inference API for this project. It can be reached only from the IBM Cloud private network.", stable: true},
	{name: "guid", description: "The GUID of the resource instance. Same value as `project_id`.", stable: true},
	{name: "crn", description: "The CRN of the project.", stable: true},
	{name: "service", description: "The service name of the project. Always `instructlab`.", stable: true},
	{name: "state", description: "The state of the project, for example `active`."},
	{name: "dashboard_url", description: "The relative URL of the project in the IBM Cloud console.", stable: true},
	{name: "account_id", description: "The ID of the account that owns the project.", stable: true},
	{name: "resource_group_crn", description: "The CRN of the resource group of the project.", stable: true},
	{name: "plan_id", description: "The global catalog ID of the pricing plan of the project.", stable: true},
	{name: "target_crn", description: "The deployment CRN of the project in the global catalog.", stable: true},
	{name: "created_at", description: "The date when the project was created.", stable: true},
	{name: "created_by", description: "The subject who created the project.", stable: true},
	{name: "updated_at", description: "The date when the project was last updated."},
	{name: "updated_by", description: "The subject who last updated the project."},
	{name: "locked", description: "Whether the project is locked.", isBool: true},
}

const rhaiiLastOperationDescription = "The last operation on the project."

var rhaiiLastOperationDescriptions = map[string]string{
	"type":        "The type of the operation, for example `create`.",
	"state":       "The state of the operation, for example `succeeded`.",
	"async":       "Whether the operation is asynchronous.",
	"description": "The description of the operation.",
}

// rhaiiProjectResourceComputedAttributes returns the computed attributes for the resource schema.
func rhaiiProjectResourceComputedAttributes() map[string]rsschema.Attribute {
	out := map[string]rsschema.Attribute{}
	for _, a := range rhaiiProjectComputedAttrs {
		if a.isBool {
			out[a.name] = rsschema.BoolAttribute{Computed: true, Description: a.description}
			continue
		}
		s := rsschema.StringAttribute{Computed: true, Description: a.description}
		if a.stable {
			s.PlanModifiers = []planmodifier.String{stringplanmodifier.UseStateForUnknown()}
		}
		out[a.name] = s
	}
	nested := map[string]rsschema.Attribute{}
	for k, d := range rhaiiLastOperationDescriptions {
		if k == "async" {
			nested[k] = rsschema.BoolAttribute{Computed: true, Description: d}
		} else {
			nested[k] = rsschema.StringAttribute{Computed: true, Description: d}
		}
	}
	out["last_operation"] = rsschema.ListNestedAttribute{
		Computed:     true,
		Description:  rhaiiLastOperationDescription,
		NestedObject: rsschema.NestedAttributeObject{Attributes: nested},
	}
	return out
}

// rhaiiProjectDataSourceComputedAttributes returns the computed attributes for the data source schema.
func rhaiiProjectDataSourceComputedAttributes() map[string]dsschema.Attribute {
	out := map[string]dsschema.Attribute{}
	for _, a := range rhaiiProjectComputedAttrs {
		if a.isBool {
			out[a.name] = dsschema.BoolAttribute{Computed: true, Description: a.description}
		} else {
			out[a.name] = dsschema.StringAttribute{Computed: true, Description: a.description}
		}
	}
	nested := map[string]dsschema.Attribute{}
	for k, d := range rhaiiLastOperationDescriptions {
		if k == "async" {
			nested[k] = dsschema.BoolAttribute{Computed: true, Description: d}
		} else {
			nested[k] = dsschema.StringAttribute{Computed: true, Description: d}
		}
	}
	out["last_operation"] = dsschema.ListNestedAttribute{
		Computed:     true,
		Description:  rhaiiLastOperationDescription,
		NestedObject: dsschema.NestedAttributeObject{Attributes: nested},
	}
	return out
}

// rhaiiInstanceValues holds the values of a resource instance that map to the
// user facing arguments of a project.
type rhaiiInstanceValues struct {
	Name            string
	Location        string
	PlanName        string
	ResourceGroupID string
}

// flattenRhaiiProject fills the computed model from a resource instance and
// returns the argument values. It reads the plan name from the global catalog.
func flattenRhaiiProject(ctx context.Context, meta conns.ClientSession, instance *rc.ResourceInstance, m *rhaiiProjectComputedModel) (rhaiiInstanceValues, diag.Diagnostics) {
	var diags diag.Diagnostics
	values := rhaiiInstanceValues{
		Name:            strDeref(instance.Name),
		Location:        strDeref(instance.RegionID),
		ResourceGroupID: strDeref(instance.ResourceGroupID),
	}
	projectID := strDeref(instance.GUID)

	m.ID = types.StringValue(strDeref(instance.ID))
	m.ProjectID = types.StringValue(projectID)
	m.GUID = types.StringValue(projectID)
	m.CRN = types.StringValue(strDeref(instance.CRN))
	m.Service = types.StringValue(rhaiiServiceName)
	m.State = types.StringValue(strDeref(instance.State))
	m.DashboardURL = types.StringValue(strDeref(instance.DashboardURL))
	m.AccountID = types.StringValue(strDeref(instance.AccountID))
	m.ResourceGroupCRN = types.StringValue(strDeref(instance.ResourceGroupCRN))
	m.PlanID = types.StringValue(strDeref(instance.ResourcePlanID))
	m.TargetCRN = types.StringValue(strDeref(instance.TargetCRN))
	m.CreatedBy = types.StringValue(strDeref(instance.CreatedBy))
	m.UpdatedBy = types.StringValue(strDeref(instance.UpdatedBy))
	m.Locked = types.BoolValue(instance.Locked != nil && *instance.Locked)

	m.Endpoint = types.StringValue("")
	m.PrivateEndpoint = types.StringValue("")
	if projectID != "" && values.Location != "" {
		m.Endpoint = types.StringValue(rhaiiProjectEndpoint(values.Location, projectID))
		m.PrivateEndpoint = types.StringValue(rhaiiProjectPrivateEndpoint(values.Location, projectID))
	}
	m.CreatedAt = types.StringValue("")
	if instance.CreatedAt != nil {
		m.CreatedAt = types.StringValue(instance.CreatedAt.String())
	}
	m.UpdatedAt = types.StringValue("")
	if instance.UpdatedAt != nil {
		m.UpdatedAt = types.StringValue(instance.UpdatedAt.String())
	}

	elemType := types.ObjectType{AttrTypes: rhaiiLastOperationAttrTypes}
	ops := []attr.Value{}
	if op := instance.LastOperation; op != nil {
		obj, d := types.ObjectValue(rhaiiLastOperationAttrTypes, map[string]attr.Value{
			"type":        types.StringValue(strDeref(op.Type)),
			"state":       types.StringValue(strDeref(op.State)),
			"async":       types.BoolValue(op.Async != nil && *op.Async),
			"description": types.StringValue(strDeref(op.Description)),
		})
		diags.Append(d...)
		ops = append(ops, obj)
	}
	list, d := types.ListValue(elemType, ops)
	diags.Append(d...)
	m.LastOperation = list

	if instance.ResourcePlanID != nil {
		gcClient, err := meta.GlobalCatalogV1API()
		if err != nil {
			diags.AddError("Error initializing the Global Catalog client", err.Error())
			return values, diags
		}
		planName, err := rhaiiPlanName(ctx, gcClient, *instance.ResourcePlanID)
		if err != nil {
			diags.AddError("Error reading the plan of the Red Hat AI Inference project", err.Error())
			return values, diags
		}
		values.PlanName = planName
	}
	return values, diags
}

// rhaiiReadTags reads the user tags and the access tags of a CRN for the state.
func rhaiiReadTags(ctx context.Context, meta conns.ClientSession, crn string) (types.Set, types.Set, diag.Diagnostics) {
	tags, diags := flex.FrameworkReadTagsUsingCRN(ctx, meta, crn, "", rhaiiUserTagType)
	accessTags, d := flex.FrameworkReadTagsUsingCRN(ctx, meta, crn, "", rhaiiAccessTagType)
	diags.Append(d...)
	return tags, accessTags, diags
}

package vpc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &MigrationNewResource{}

func NewMigrationNewResource() resource.Resource {
	return &MigrationNewResource{}
}

type MigrationNewResource struct{}

type MigrationNewResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	SampleAttribute    types.String `tfsdk:"sample_attribute"`
	NonSampleAttribute types.String `tfsdk:"non_sample_attribute"`
}

func (r *MigrationNewResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_new_resource_new"
}

func (r *MigrationNewResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"sample_attribute": schema.StringAttribute{
				Required:    true,
				Description: "This attribute forces a new resource when changed",
				// This makes the attribute force a new resource
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"non_sample_attribute": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *MigrationNewResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MigrationNewResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idFromAPI := "my-id"
	sample := plan.SampleAttribute.ValueString()

	plan.Id = types.StringValue(idFromAPI + sample)
	plan.NonSampleAttribute = types.StringValue(idFromAPI + "non-sample")

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *MigrationNewResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MigrationNewResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idFromAPI := state.Id.ValueString()
	state.NonSampleAttribute = types.StringValue(idFromAPI + "non-sample")

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *MigrationNewResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update does nothing as sample_attribute forces a new resource
	var state MigrationNewResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *MigrationNewResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Delete sets the id as empty, which effectively removes the resource from state
	var state MigrationNewResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Setting Id to an empty string effectively removes the resource from state
	state.Id = types.StringValue("")

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

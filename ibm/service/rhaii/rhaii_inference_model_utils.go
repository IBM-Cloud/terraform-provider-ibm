// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"regexp"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/rhaii/rhaiiv1"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rhaiiModelIDRegexp is the pattern of the model path parameter in the API spec.
var rhaiiModelIDRegexp = regexp.MustCompile(`^[a-z0-9-]+$`)

var (
	rhaiiPricingAttrTypes = map[string]attr.Type{
		"input_measure":  types.StringType,
		"output_measure": types.StringType,
	}
	rhaiiValidationAttrTypes = map[string]attr.Type{
		"status":       types.StringType,
		"error":        types.StringType,
		"completed_at": types.StringType,
	}
)

// rhaiiInferenceModelMetadataModel holds the attributes built from the IBM
// metadata of an inference model.
type rhaiiInferenceModelMetadataModel struct {
	Name        types.String `tfsdk:"name"`
	DisplayName types.String `tfsdk:"display_name"`
	State       types.String `tfsdk:"state"`
	Status      types.String `tfsdk:"status"`
	TotalParams types.Int64  `tfsdk:"total_params"`
	ConfigJSON  types.String `tfsdk:"config_json"`
	Pricing     types.List   `tfsdk:"pricing"`
	Validation  types.List   `tfsdk:"validation"`
}

// rhaiiInferenceModelMetadataAttributes returns the schema of the IBM metadata
// attributes. total_params is Int64, which is 64-bit on every platform in the
// framework, so 70553706496 does not overflow on 32-bit builds.
func rhaiiInferenceModelMetadataAttributes() map[string]schema.Attribute {
	str := func(d string) schema.Attribute { return schema.StringAttribute{Computed: true, Description: d} }
	return map[string]schema.Attribute{
		"name":         str("The name of the model."),
		"display_name": str("The display name of the model."),
		"state":        str("The state of the model, for example `model_loaded`."),
		"status":       str("The status of the model, for example `loaded`."),
		"total_params": schema.Int64Attribute{Computed: true, Description: "The total number of parameters of the model."},
		"config_json":  str("The model configuration (architecture details and hyperparameters) as a JSON string."),
		"pricing": schema.ListNestedAttribute{
			Computed:    true,
			Description: "The pricing measures of the model.",
			NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"input_measure":  str("The billing measure for input tokens."),
				"output_measure": str("The billing measure for output tokens."),
			}},
		},
		"validation": schema.ListNestedAttribute{
			Computed:    true,
			Description: "The validation information of a custom model.",
			NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"status":       str("The validation status."),
				"error":        str("The validation error message, if any."),
				"completed_at": str("The date when the validation completed."),
			}},
		},
	}
}

// flattenRhaiiInferenceModelMetadata converts the IBM metadata of a model.
func flattenRhaiiInferenceModelMetadata(m *rhaiiv1.InferenceModelMetadata) (rhaiiInferenceModelMetadataModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	pricingType := types.ObjectType{AttrTypes: rhaiiPricingAttrTypes}
	validationType := types.ObjectType{AttrTypes: rhaiiValidationAttrTypes}
	out := rhaiiInferenceModelMetadataModel{
		Name:        types.StringValue(""),
		DisplayName: types.StringValue(""),
		State:       types.StringValue(""),
		Status:      types.StringValue(""),
		TotalParams: types.Int64Null(),
		ConfigJSON:  types.StringValue(""),
		Pricing:     types.ListValueMust(pricingType, []attr.Value{}),
		Validation:  types.ListValueMust(validationType, []attr.Value{}),
	}
	if m == nil {
		return out, diags
	}

	out.Name = types.StringValue(strDeref(m.Name))
	out.DisplayName = types.StringValue(strDeref(m.DisplayName))
	out.State = types.StringValue(strDeref(m.State))
	out.Status = types.StringValue(strDeref(m.Status))
	if m.TotalParams != nil {
		out.TotalParams = types.Int64Value(*m.TotalParams)
	}
	configJSON, err := rhaiiToJSON(m.Config)
	if err != nil {
		diags.AddError("Error converting the model configuration to JSON", err.Error())
	}
	out.ConfigJSON = types.StringValue(configJSON)

	if p := m.Pricing; p != nil {
		obj, d := types.ObjectValue(rhaiiPricingAttrTypes, map[string]attr.Value{
			"input_measure":  types.StringValue(strDeref(p.InputMeasure)),
			"output_measure": types.StringValue(strDeref(p.OutputMeasure)),
		})
		diags.Append(d...)
		out.Pricing = types.ListValueMust(pricingType, []attr.Value{obj})
	}
	if v := m.Validation; v != nil {
		obj, d := types.ObjectValue(rhaiiValidationAttrTypes, map[string]attr.Value{
			"status":       types.StringValue(strDeref(v.Status)),
			"error":        types.StringValue(strDeref(v.Error)),
			"completed_at": types.StringValue(strDeref(v.CompletedAt)),
		})
		diags.Append(d...)
		out.Validation = types.ListValueMust(validationType, []attr.Value{obj})
	}
	return out, diags
}

// rhaiiInferenceModelMetadataAttrTypes returns the object attribute types of
// the metadata attributes, for list elements of the models data source.
func rhaiiInferenceModelMetadataAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":         types.StringType,
		"display_name": types.StringType,
		"state":        types.StringType,
		"status":       types.StringType,
		"total_params": types.Int64Type,
		"config_json":  types.StringType,
		"pricing":      types.ListType{ElemType: types.ObjectType{AttrTypes: rhaiiPricingAttrTypes}},
		"validation":   types.ListType{ElemType: types.ObjectType{AttrTypes: rhaiiValidationAttrTypes}},
	}
}

// metadataAttrValues returns the metadata as attribute values, for an object.
func (m rhaiiInferenceModelMetadataModel) metadataAttrValues() map[string]attr.Value {
	return map[string]attr.Value{
		"name":         m.Name,
		"display_name": m.DisplayName,
		"state":        m.State,
		"status":       m.Status,
		"total_params": m.TotalParams,
		"config_json":  m.ConfigJSON,
		"pricing":      m.Pricing,
		"validation":   m.Validation,
	}
}

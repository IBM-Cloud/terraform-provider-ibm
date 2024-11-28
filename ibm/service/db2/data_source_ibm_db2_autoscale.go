package db2

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func DataSourceIBMDB2Autoscale() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMDB2AutoscaleRead,

		Schema: map[string]*schema.Schema{
			"auto_scaling_allow_plan_limit": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates the maximum number of scaling actions that are allowed within a specified time period",
			},

			"auto_scaling_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if automatic scaling is enabled or no",
			},

			"auto_scaling_max_storage": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The maximum number of storage in GB",
			},

			"auto_scaling_over_time_period": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Defines the time period over which auto-scaling adjustments are monitored and applied",
			},

			"auto_scaling_pause_limit": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the duration to pause auto-scaling actions after a scaling event has occurred",
			},

			"auto_scaling_threshold": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the minimum threshold for the autoscaling group",
			},

			"storage_unit": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the storage unit for the autoscaling group",
			},

			"storage_utilization_percentage": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Specifies the percentage for the autoscaling group",
			},

			"support_auto_scaling": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the autoscaling group is enabled or no",
			},
		},
	}
}

func DataSourceIBMDB2AutoscaleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2SaasV1Client, err := meta.(conns.ClientSession).DB2SaasV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_autoscale", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAutoScaleOptions := &db2saasv1.GetDb2SaasAutoscaleOptions{}

	getAutoScaleOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	result, response, err := db2SaasV1Client.GetDb2SaasAutoscaleWithContext(context, getAutoScaleOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAutoScaleWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_db2_autoscale", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_allow_plan_limit", result.AutoScalingAllowPlanLimit); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_allow_plan_limit: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_enabled", result.AutoScalingEnabled); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_enabled: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_max_storage", result.AutoScalingMaxStorage); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_max_storage: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_over_time_period", result.AutoScalingOverTimePeriod); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_over_time_period: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_pause_limit", result.AutoScalingPauseLimit); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_pause_limit: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("auto_scaling_threshold", result.AutoScalingThreshold); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting auto_scaling_threshold: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("storage_unit", result.StorageUnit); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting storage_unit: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("storage_utilization_percentage", result.StorageUtilizationPercentage); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting storage_utilization_percentage: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("support_auto_scaling", result.SupportAutoScaling); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting support_auto_scaling: %s", err), "(Data) ibm_db2_autoscale", "read")
		return tfErr.GetDiag()
	}

	return nil

}

// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type dataSourceIBMDatabasePointInTimeRecoveryBackend interface {
	Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
}

func pickDataSourcePointInTimeRecoveryBackend(d *schema.ResourceData, meta interface{}) (dataSourceIBMDatabasePointInTimeRecoveryBackend, error) {
	deploymentID := d.Get("deployment_id").(string)
	if deploymentID == "" {
		return nil, fmt.Errorf("deployment_id is required to read point-in-time recovery data")
	}

	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, err
	}

	instance, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{ID: &deploymentID})
	if err != nil {
		return nil, fmt.Errorf("failed to get resource instance: %s", describeRequestFailure(err, response))
	}

	if instance.ResourcePlanID == nil {
		return nil, fmt.Errorf("resource instance %s has no ResourcePlanID", deploymentID)
	}

	if isGen2Plan(*instance.ResourcePlanID) {
		return newDataSourceIBMDatabasePointInTimeRecoveryGen2Backend(instance), nil
	}
	return newDataSourceIBMDatabasePointInTimeRecoveryClassicBackend(), nil
}

// describeRequestFailure shows the body when the SDK error is only the status text, and adds the
// Handbook trace that support cases must quote, because the SDK keeps it only in the decoded body.
func describeRequestFailure(err error, response *core.DetailedResponse) string {
	if response == nil || (response.StatusCode >= 200 && response.StatusCode < 300) {
		return err.Error()
	}

	statusText := http.StatusText(response.StatusCode)
	message := err.Error()
	if message == "" || message == statusText {
		message = errorResponseBody(response)
	} else if trace := handbookTrace(response); trace != "" {
		message = fmt.Sprintf("%s (trace: %s)", message, trace)
	}
	if message == "" {
		message = statusText
	}
	return fmt.Sprintf("HTTP %d: %s", response.StatusCode, message)
}

func errorResponseBody(response *core.DetailedResponse) string {
	if body := strings.TrimSpace(string(response.RawResult)); body != "" {
		return body
	}
	result, ok := response.GetResultAsMap()
	if !ok {
		return ""
	}
	var body strings.Builder
	encoder := json.NewEncoder(&body)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(result); err != nil {
		return ""
	}
	return strings.TrimSpace(body.String())
}

func handbookTrace(response *core.DetailedResponse) string {
	body, ok := response.GetResultAsMap()
	if !ok {
		return ""
	}
	trace, _ := body["trace"].(string)
	return trace
}

func DataSourceIBMDatabasePointInTimeRecovery() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMDatabasePointInTimeRecoveryRead,

		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Deployment ID.",
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_database_point_in_time_recovery",
					"deployment_id"),
			},
			"earliest_point_in_time_recovery_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The earliest time that a restore can target. Gen2: Empty when the instance cannot be restored.",
			},
			"latest_point_in_time_recovery_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gen2 only. The latest time that a restore can target when the data source was read. Empty whenever not_restorable_reason is set.",
			},
			"retention_days": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Gen2 only. The days of history the instance keeps for point-in-time recovery.",
			},
			"archiving_status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gen2 only. The state of the instance's change archiving: healthy, delayed or unknown.",
			},
			"unavailable_periods": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gen2 only. Periods inside the window that cannot be restored to.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start of the period, rounded down to the millisecond, so from itself may still be restorable; treat it as unavailable.",
						},
						"until": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The first time after the period that can be restored to again. Empty while the period is still open.",
						},
					},
				},
			},
			"not_restorable_reason": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gen2 only. Why the instance cannot be restored now: archive_not_started, no_eligible_backup, archive_delayed or archive_gap. Empty when it can be restored. The data source returns a warning whenever it is set.",
			},
		},
	}
}

func DataSourceIBMDatabasePointInTimeRecoveryValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "deployment_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			Required:                   true,
			CloudDataType:              "cloud-database",
			CloudDataRange:             []string{"resolved_to:id"}})

	iBMDatabasePointInTimeRecoveryValidator := validate.ResourceValidator{ResourceName: "ibm_database_point_in_time_recovery", Schema: validateSchema}
	return &iBMDatabasePointInTimeRecoveryValidator
}

func DataSourceIBMDatabasePointInTimeRecoveryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	b, err := pickDataSourcePointInTimeRecoveryBackend(d, meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return b.Read(context, d, meta)
}

type dataSourceIBMDatabasePointInTimeRecoveryClassicBackend struct{}

func newDataSourceIBMDatabasePointInTimeRecoveryClassicBackend() dataSourceIBMDatabasePointInTimeRecoveryBackend {
	return &dataSourceIBMDatabasePointInTimeRecoveryClassicBackend{}
}

func (c *dataSourceIBMDatabasePointInTimeRecoveryClassicBackend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPitrDataOptions := &clouddatabasesv5.GetPitrDataOptions{}

	getPitrDataOptions.SetID(d.Get("deployment_id").(string))

	pointInTimeRecoveryData, response, err := cloudDatabasesClient.GetPitrDataWithContext(context, getPitrDataOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPitrDataWithContext failed: %s", describeRequestFailure(err, response)), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(d.Get("deployment_id").(string))

	if pointInTimeRecoveryData.PointInTimeRecoveryData.EarliestPointInTimeRecoveryTime != nil {
		pitr := pointInTimeRecoveryData.PointInTimeRecoveryData.EarliestPointInTimeRecoveryTime
		if err = d.Set("earliest_point_in_time_recovery_time", pitr); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting earliest_point_in_time_recovery_time: %s", err), "(Data) ibm_database_point_in_time_recovery", "read")
			return tfErr.GetDiag()
		}
	}

	return nil
}

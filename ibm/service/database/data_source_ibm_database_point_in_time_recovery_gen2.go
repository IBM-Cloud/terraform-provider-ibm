// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

type pointInTimeRecoveryWindow struct {
	EarliestRestorableAt *string                                `json:"earliest_restorable_at"`
	LatestRestorableAt   *string                                `json:"latest_restorable_at"`
	RetentionDays        *int64                                 `json:"retention_days"`
	Archiving            *pointInTimeRecoveryArchiving          `json:"archiving"`
	UnavailablePeriods   []pointInTimeRecoveryUnavailablePeriod `json:"unavailable_periods"`
	NotRestorableReason  *string                                `json:"not_restorable_reason"`
}

type pointInTimeRecoveryArchiving struct {
	Status *string `json:"status"`
}

type pointInTimeRecoveryUnavailablePeriod struct {
	From  *string `json:"from"`
	Until *string `json:"until"`
}

// dataSourceIBMDatabasePointInTimeRecoveryGen2Backend keeps the instance read by
// pickDataSourcePointInTimeRecoveryBackend, whose extensions carry the window URL.
type dataSourceIBMDatabasePointInTimeRecoveryGen2Backend struct {
	instance *rc.ResourceInstance
}

func newDataSourceIBMDatabasePointInTimeRecoveryGen2Backend(instance *rc.ResourceInstance) dataSourceIBMDatabasePointInTimeRecoveryBackend {
	return &dataSourceIBMDatabasePointInTimeRecoveryGen2Backend{instance: instance}
}

func (g *dataSourceIBMDatabasePointInTimeRecoveryGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	deploymentID := d.Get("deployment_id").(string)

	windowURL, err := restorableWindowURL(deploymentID, g.instance.Extensions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// The Resource Controller client carries the provider's IAM authenticator.
	window, response, err := fetchPointInTimeRecoveryWindow(context, rsConClient.Service, windowURL)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GET point-in-time recovery window failed: %s", describeRequestFailure(err, response)), "(Data) ibm_database_point_in_time_recovery", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(deploymentID)

	for key, value := range flattenPointInTimeRecoveryWindow(window) {
		if err := d.Set(key, value); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", key, err), "(Data) ibm_database_point_in_time_recovery", "read")
			return tfErr.GetDiag()
		}
	}

	if reason := flex.StringValue(window.NotRestorableReason); reason != "" {
		return diag.Diagnostics{{
			Severity: diag.Warning,
			Summary:  "No point-in-time restore of the instance can start now",
			Detail:   notRestorableDetail(deploymentID, reason),
		}}
	}
	return nil
}

func notRestorableDetail(deploymentID, reason string) string {
	detail := fmt.Sprintf("%s cannot be restored now (not_restorable_reason %s), so latest_point_in_time_recovery_time is empty and a restore that uses it is refused.", deploymentID, reason)
	switch reason {
	case "archive_delayed":
		return detail + " Retry once archiving_status is healthy again."
	case "no_eligible_backup":
		return detail + " Retry after the instance's next backup completes, and check that the caller can read its backups."
	case "archive_not_started", "archive_gap":
		return detail + " Retry after the instance's next backup completes."
	}
	return detail
}

// restorableWindowURL refuses any URL but HTTPS to a host name under cloud.ibm.com, because the request carries the provider's IAM token.
func restorableWindowURL(deploymentID string, extensions map[string]interface{}) (string, error) {
	dataservices, _ := extensions["dataservices"].(map[string]interface{})
	status, _ := dataservices["point_in_time_recovery_status"].(map[string]interface{})
	windowURL, _ := status["restorable_window_url"].(string)
	if windowURL == "" {
		return "", fmt.Errorf("the Resource Controller extensions of %s publish no dataservices.point_in_time_recovery_status.restorable_window_url. "+
			"If point-in-time recovery is available for this service in its region, update the instance's parameters or plan and wait for the update to finish, which refreshes its extensions (a tag change does not), or ask IBM Cloud support to sync it", deploymentID)
	}

	parsed, err := url.Parse(windowURL)
	// The zone of an IPv6 literal can end in .cloud.ibm.com while its address points anywhere.
	if err != nil || parsed.Scheme != "https" || strings.ContainsAny(parsed.Hostname(), ":%") || !strings.HasSuffix(parsed.Hostname(), ".cloud.ibm.com") {
		return "", fmt.Errorf("the restorable window URL %q of %s is not an https URL under cloud.ibm.com", windowURL, deploymentID)
	}
	return windowURL, nil
}

func fetchPointInTimeRecoveryWindow(ctx context.Context, service *core.BaseService, windowURL string) (*pointInTimeRecoveryWindow, *core.DetailedResponse, error) {
	builder := core.NewRequestBuilder(core.GET).WithContext(ctx)
	if _, err := builder.ResolveRequestURL(windowURL, "", nil); err != nil {
		return nil, nil, err
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return nil, nil, err
	}

	var window pointInTimeRecoveryWindow
	response, err := service.Request(request, &window)
	if err != nil {
		return nil, response, err
	}
	return &window, response, nil
}

// flattenPointInTimeRecoveryWindow reports a null time as "", which Terraform state can hold for a computed string.
func flattenPointInTimeRecoveryWindow(window *pointInTimeRecoveryWindow) map[string]interface{} {
	archivingStatus := ""
	if window.Archiving != nil {
		archivingStatus = flex.StringValue(window.Archiving.Status)
	}

	retentionDays := 0
	if window.RetentionDays != nil {
		retentionDays = int(*window.RetentionDays)
	}

	unavailablePeriods := make([]map[string]interface{}, 0, len(window.UnavailablePeriods))
	for _, period := range window.UnavailablePeriods {
		unavailablePeriods = append(unavailablePeriods, map[string]interface{}{
			"from":  flex.StringValue(period.From),
			"until": flex.StringValue(period.Until),
		})
	}

	return map[string]interface{}{
		"earliest_point_in_time_recovery_time": flex.StringValue(window.EarliestRestorableAt),
		"latest_point_in_time_recovery_time":   flex.StringValue(window.LatestRestorableAt),
		"retention_days":                       retentionDays,
		"archiving_status":                     archivingStatus,
		"unavailable_periods":                  unavailablePeriods,
		"not_restorable_reason":                flex.StringValue(window.NotRestorableReason),
	}
}

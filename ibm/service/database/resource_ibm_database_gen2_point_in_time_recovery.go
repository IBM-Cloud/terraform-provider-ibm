// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	sdkretry "github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	pointInTimeRecoveryDeploymentIDKey  = "point_in_time_recovery_deployment_id"
	pointInTimeRecoveryTimeKey          = "point_in_time_recovery_time"
	pointInTimeRecoveryRetentionDaysKey = "backups.0.point_in_time_recovery.0.retention_days"
)

// A value that is unknown at plan time counts as set; its checks run again at apply.
type gen2PointInTimeRestore struct {
	sourceCRN        string
	pointInTime      string
	service          string
	location         string
	sourceSet        bool
	sourceKnown      bool
	pointInTimeSet   bool
	pointInTimeKnown bool
	backupIDSet      bool
}

func restoresToPointInTime(d *schema.ResourceData) bool {
	_, ok := d.GetOk(pointInTimeRecoveryDeploymentIDKey)
	return ok
}

func (g *resourceIBMDatabaseGen2Backend) buildGen2CreateParameters(d *schema.ResourceData, serviceName string, meta interface{}, catalogCRN string) (map[string]interface{}, error) {
	parameters, err := g.buildGen2Parameters(d, serviceName, meta, catalogCRN)
	if err != nil {
		return nil, err
	}

	// Kept out of buildGen2Parameters, which updates reuse: the restore arguments are immutable and retention has its own update.
	dataservices := parameters["dataservices"].(map[string]interface{})
	if err := addPointInTimeRestore(d, serviceName, dataservices); err != nil {
		return nil, err
	}
	addPointInTimeRecoveryRetention(d, dataservices)

	return parameters, nil
}

func addPointInTimeRestore(d *schema.ResourceData, serviceName string, dataservices map[string]interface{}) error {
	if !restoresToPointInTime(d) {
		return nil
	}

	_, backupIDSet := d.GetOk("backup_id")
	restore := gen2PointInTimeRestore{
		sourceCRN:        d.Get(pointInTimeRecoveryDeploymentIDKey).(string),
		pointInTime:      d.Get(pointInTimeRecoveryTimeKey).(string),
		service:          serviceName,
		location:         d.Get("location").(string),
		sourceSet:        true,
		sourceKnown:      true,
		pointInTimeSet:   true,
		pointInTimeKnown: true,
		backupIDSet:      backupIDSet,
	}
	if err := validateGen2PointInTimeRestore(restore); err != nil {
		return err
	}

	pointInTime, err := canonicalPointInTime(restore.pointInTime)
	if err != nil {
		return err
	}
	dataservices["source_dataservice_crn"] = restore.sourceCRN
	dataservices["point_in_time"] = pointInTime
	return nil
}

// canonicalPointInTime returns the UTC form the platform stores, so the value sent matches the value read back.
func canonicalPointInTime(value string) (string, error) {
	if strings.TrimSpace(value) == "" {
		return "", errors.New("point_in_time_recovery_time must be an RFC 3339 timestamp such as 2026-09-27T09:30:00Z: Gen2 databases have no restore to the latest time")
	}

	pointInTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", fmt.Errorf("point_in_time_recovery_time %q is not an RFC 3339 timestamp with an offset, for example 2026-09-27T09:30:00Z", value)
	}
	return pointInTime.UTC().Format(time.RFC3339Nano), nil
}

func validateGen2PointInTimeRestore(restore gen2PointInTimeRestore) error {
	if !restore.sourceSet && !restore.pointInTimeSet {
		return nil
	}
	if restore.sourceSet != restore.pointInTimeSet {
		return errors.New("point_in_time_recovery_deployment_id and point_in_time_recovery_time must be set together for Gen2 databases")
	}
	if restore.backupIDSet {
		return errors.New("backup_id cannot be combined with point_in_time_recovery_deployment_id: a Gen2 point-in-time restore selects its base backup itself")
	}
	if restore.pointInTimeKnown {
		if _, err := canonicalPointInTime(restore.pointInTime); err != nil {
			return err
		}
	}
	if restore.sourceKnown {
		return validateGen2PointInTimeRestoreSource(restore.sourceCRN, restore.service, restore.location)
	}
	return nil
}

// validateGen2PointInTimeRestoreSource checks the CRN segments
// crn:v1:<cname>:<ctype>:<service>:<location>:<scope>:<instance>:<resource-type>:<resource>.
func validateGen2PointInTimeRestoreSource(sourceCRN, service, location string) error {
	segments := strings.Split(sourceCRN, ":")
	if len(segments) != 10 || segments[0] != "crn" || segments[7] == "" || segments[8] != "" || segments[9] != "" {
		return fmt.Errorf("point_in_time_recovery_deployment_id %q must be the CRN of the source instance, such as crn:v1:bluemix:public:databases-for-postgresql:ca-mon:a/<account-id>:<instance-id>::; to restore a backup, use backup_id", sourceCRN)
	}
	if service != "" && segments[4] != service {
		return fmt.Errorf("point_in_time_recovery_deployment_id must be an instance of %s, not %s", service, segments[4])
	}
	if location != "" && segments[5] != location {
		return fmt.Errorf("the source is in %s but location is %s: Gen2 point-in-time recovery restores into the source's region only", segments[5], location)
	}
	return nil
}

// ValidatePointInTimeRecoveryDiff validates a point-in-time restore when the instance is planned for creation.
func (g *resourceIBMDatabaseGen2Backend) ValidatePointInTimeRecoveryDiff(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
	// The restore arguments apply only at create; ApplyOnce suppresses later edits.
	if diff.Id() != "" {
		return nil
	}

	restore := pointInTimeRestoreFromRawConfig(diff.GetRawConfig())
	restore.service = diff.Get("service").(string)
	restore.location = diff.Get("location").(string)
	return validateGen2PointInTimeRestore(restore)
}

// pointInTimeRestoreFromRawConfig reads the raw configuration because GetOk treats "" as unset.
func pointInTimeRestoreFromRawConfig(raw cty.Value) gen2PointInTimeRestore {
	var restore gen2PointInTimeRestore
	restore.sourceCRN, restore.sourceSet, restore.sourceKnown = rawConfigString(raw, pointInTimeRecoveryDeploymentIDKey)
	restore.pointInTime, restore.pointInTimeSet, restore.pointInTimeKnown = rawConfigString(raw, pointInTimeRecoveryTimeKey)

	backupID, backupIDSet, backupIDKnown := rawConfigString(raw, "backup_id")
	restore.backupIDSet = backupIDSet && (!backupIDKnown || backupID != "")
	return restore
}

func rawConfigString(raw cty.Value, attr string) (value string, set, known bool) {
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute(attr) {
		return "", false, false
	}

	attrValue := raw.GetAttr(attr)
	if attrValue.IsNull() {
		return "", false, false
	}
	if !attrValue.IsKnown() {
		return "", true, false
	}
	return attrValue.AsString(), true, true
}

func pointInTimeRecoveryRetention(retentionDays int) map[string]interface{} {
	return map[string]interface{}{
		"point_in_time_recovery": map[string]interface{}{
			"retention_days": retentionDays,
		},
	}
}

func addPointInTimeRecoveryRetention(d *schema.ResourceData, dataservices map[string]interface{}) {
	if retentionDays, ok := d.GetOk(pointInTimeRecoveryRetentionDaysKey); ok {
		dataservices["backups"] = pointInTimeRecoveryRetention(retentionDays.(int))
	}
}

// pointInTimeRecoveryRetentionUpdate carries only the retention, so no other setting is resent.
func pointInTimeRecoveryRetentionUpdate(d *schema.ResourceData) (map[string]interface{}, bool) {
	if !d.HasChange(pointInTimeRecoveryRetentionDaysKey) {
		return nil, false
	}
	retentionDays, ok := d.GetOk(pointInTimeRecoveryRetentionDaysKey)
	if !ok {
		return nil, false
	}

	return map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backups": pointInTimeRecoveryRetention(retentionDays.(int)),
		},
	}, true
}

func (g *resourceIBMDatabaseGen2Backend) applyPointInTimeRecoveryRetentionWithDiagnostics(ctx context.Context, d *schema.ResourceData, rsConClient *rc.ResourceControllerV2, instanceID string) diag.Diagnostics {
	parameters, ok := pointInTimeRecoveryRetentionUpdate(d)
	if !ok {
		return nil
	}

	if err := g.updateResourceInstanceParameters(rsConClient, instanceID, parameters); err != nil {
		return diagError("error updating point-in-time recovery retention: %s", err)
	}

	stateConf := gen2InstanceOperationStateChangeConf(rsConClient, instanceID, d.Timeout(schema.TimeoutUpdate))
	if _, err := stateConf.WaitForStateContext(ctx); err != nil {
		return diagError("error waiting for the point-in-time recovery retention update of (%s) to complete: %s", d.Id(), err)
	}
	return nil
}

// An asynchronous update leaves the instance active and refreshes its extensions only when it ends, so follow last_operation.
func gen2InstanceOperationStateChangeConf(rsConClient *rc.ResourceControllerV2, instanceID string, timeout time.Duration) *sdkretry.StateChangeConf {
	return &sdkretry.StateChangeConf{
		Pending: []string{rc.ResourceInstanceLastOperationStateInProgressConst, databaseInstanceInactiveStatus},
		Target:  []string{rc.ResourceInstanceLastOperationStateSucceededConst, databaseInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{ID: &instanceID})
			if err != nil {
				return nil, "", wrapAPIError("get resource instance", err, response)
			}

			lastOperation := instance.LastOperation
			if lastOperation != nil && lastOperation.Async != nil && *lastOperation.Async {
				state := flex.StringValue(lastOperation.State)
				if state == rc.ResourceInstanceLastOperationStateFailedConst {
					return instance, state, fmt.Errorf("the update of resource instance %s failed: %s", instanceID, flex.StringValue(lastOperation.Description))
				}
				return instance, state, nil
			}

			state := flex.StringValue(instance.State)
			if state == databaseInstanceFailStatus {
				return instance, state, fmt.Errorf("resource instance %s is in the %s state", instanceID, state)
			}
			return instance, state, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
}

// flattenPointInTimeRecoveryRetention returns nil when the instance never set a retention,
// because the platform does not store its default.
func flattenPointInTimeRecoveryRetention(extensions map[string]interface{}) []map[string]interface{} {
	dataservices, _ := extensions["dataservices"].(map[string]interface{})
	backups, _ := dataservices["backups"].(map[string]interface{})
	pointInTimeRecovery, _ := backups["point_in_time_recovery"].(map[string]interface{})
	retentionDays, ok := pointInTimeRecovery["retention_days"].(float64)
	if !ok {
		return nil
	}

	return []map[string]interface{}{
		{
			"point_in_time_recovery": []map[string]interface{}{
				{"retention_days": int(retentionDays)},
			},
		},
	}
}

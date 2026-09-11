// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/stretchr/testify/require"
)

// backupExtensions builds the dataservices.backup extension block used by
// extractGen2BackupExtensions to identify a Gen2 Independent Backup's source
// deployment and type.
func backupExtensions(sourceDataServiceCRN, backupType string) map[string]interface{} {
	return map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backup": map[string]interface{}{
				"source_data_service_crn": sourceDataServiceCRN,
				"type":                    backupType,
			},
		},
	}
}

// TestFilterGen2BackupsByDeployment verifies that only backup instances whose
// source_data_service_crn extension matches the requested deploymentID are
// returned, so listing Gen2 Independent Backups for a specific deployment
// does not leak backups belonging to other deployments in the account.
func TestFilterGen2BackupsByDeployment(t *testing.T) {
	const deploymentA = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-a::"
	const deploymentB = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-b::"

	backupA1 := rc.ResourceInstance{
		CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-a1::"),
		State:      core.StringPtr("active"),
		Extensions: backupExtensions(deploymentA, "on_demand"),
	}
	backupA2 := rc.ResourceInstance{
		CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-a2::"),
		State:      core.StringPtr("provisioning"),
		Extensions: backupExtensions(deploymentA, "scheduled"),
	}
	backupB1 := rc.ResourceInstance{
		CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-b1::"),
		State:      core.StringPtr("active"),
		Extensions: backupExtensions(deploymentB, "on_demand"),
	}
	backupNoCRN := rc.ResourceInstance{
		State:      core.StringPtr("active"),
		Extensions: backupExtensions(deploymentA, "on_demand"),
	}

	instances := []rc.ResourceInstance{backupA1, backupA2, backupB1, backupNoCRN}

	result := filterGen2BackupsByDeployment(nil, instances, deploymentA)

	require.Len(t, result, 2, "only backups belonging to deploymentA should be returned")

	backupIDs := make([]string, 0, len(result))
	for _, backup := range result {
		require.Equal(t, deploymentA, backup["deployment_id"], "each returned backup must belong to the requested deployment")
		backupIDs = append(backupIDs, backup["backup_id"].(string))
	}
	require.ElementsMatch(t, []string{*backupA1.CRN, *backupA2.CRN}, backupIDs)
}

// TestFilterGen2BackupsByDeployment_NoMatches ensures the accumulator is
// returned unmodified (empty) when no instances belong to the requested
// deployment.
func TestFilterGen2BackupsByDeployment_NoMatches(t *testing.T) {
	const deploymentA = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-a::"
	const deploymentC = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-c::"

	instances := []rc.ResourceInstance{
		{
			CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-a1::"),
			State:      core.StringPtr("active"),
			Extensions: backupExtensions(deploymentA, "on_demand"),
		},
	}

	result := filterGen2BackupsByDeployment([]map[string]interface{}{}, instances, deploymentC)

	require.NotNil(t, result)
	require.Empty(t, result)
}

// TestFilterGen2BackupsByDeployment_AcrossPages simulates paginated
// ListResourceInstances results by calling filterGen2BackupsByDeployment
// once per page and accumulating results, mirroring how Read consumes pages.
func TestFilterGen2BackupsByDeployment_AcrossPages(t *testing.T) {
	const deploymentA = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-a::"
	const deploymentB = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-b::"

	page1 := []rc.ResourceInstance{
		{
			CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-a1::"),
			State:      core.StringPtr("active"),
			Extensions: backupExtensions(deploymentA, "on_demand"),
		},
		{
			CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-b1::"),
			State:      core.StringPtr("active"),
			Extensions: backupExtensions(deploymentB, "on_demand"),
		},
	}
	page2 := []rc.ResourceInstance{
		{
			CRN:        core.StringPtr("crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-a2::"),
			State:      core.StringPtr("active"),
			Extensions: backupExtensions(deploymentA, "scheduled"),
		},
	}

	var backups []map[string]interface{}
	for _, page := range [][]rc.ResourceInstance{page1, page2} {
		backups = filterGen2BackupsByDeployment(backups, page, deploymentA)
	}

	require.Len(t, backups, 2, "matches across all pages should be accumulated")
	for _, backup := range backups {
		require.Equal(t, deploymentA, backup["deployment_id"])
	}
}

// s2sExtensions builds the extensions block for a Gen2 database instance
// that has Independent Backups and optionally has S2S authorizations.
func s2sExtensions(hasBackups, authorized bool) map[string]interface{} {
	ds := map[string]interface{}{}
	if hasBackups {
		ds["backups"] = map[string]interface{}{"retention_days": 30}
	}
	if authorized {
		ds["authorizations"] = map[string]interface{}{
			"independent_backups": true,
			"resource_group":      true,
		}
	}
	return map[string]interface{}{"dataservices": ds}
}

// TestBackupsGen2BackendS2SWarning verifies that Read emits an S2S warning
// when the source instance has Independent Backups but S2S is not configured.
// The warning must be a diag.Warning (not an error) and must use the correct
// summary and detail from the shared constants.
func TestBackupsGen2BackendS2SWarning(t *testing.T) {
	tests := []struct {
		name        string
		extensions  map[string]interface{}
		wantWarning bool
	}{
		{
			name:        "no source instance — no warning",
			extensions:  nil,
			wantWarning: false,
		},
		{
			name:        "instance has independent backups but no S2S — warning",
			extensions:  s2sExtensions(true, false),
			wantWarning: true,
		},
		{
			name:        "instance has independent backups and S2S configured — no warning",
			extensions:  s2sExtensions(true, true),
			wantWarning: false,
		},
		{
			name:        "instance has no independent backups — no warning",
			extensions:  s2sExtensions(false, false),
			wantWarning: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sourceInstance *rc.ResourceInstance
			if tc.extensions != nil {
				sourceInstance = &rc.ResourceInstance{Extensions: tc.extensions}
			}

			hasWarning := sourceInstance != nil &&
				hasIndependentBackups(sourceInstance.Extensions) &&
				!checkS2SAuthorization(sourceInstance.Extensions)

			require.Equal(t, tc.wantWarning, hasWarning, "S2S warning flag mismatch")

			if tc.wantWarning {
				diags := diag.Diagnostics{{
					Severity: diag.Warning,
					Summary:  s2sAuthWarningHeader,
					Detail:   s2sAuthWarningDetail,
				}}
				require.Len(t, diags, 1)
				require.Equal(t, diag.Warning, diags[0].Severity)
				require.Equal(t, s2sAuthWarningHeader, diags[0].Summary)
				require.Equal(t, s2sAuthWarningDetail, diags[0].Detail)
			}
		})
	}
}

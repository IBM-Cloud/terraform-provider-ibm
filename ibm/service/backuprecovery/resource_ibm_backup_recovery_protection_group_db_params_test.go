// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
)

func TestProtectionGroupDb2ParamsMapper(t *testing.T) {
	in := map[string]interface{}{
		"backup_job_arguments": []interface{}{map[string]interface{}{"key": "compress", "value": "true"}},
		"concurrency":          2,
		"exclude_object_ids":   []interface{}{7},
		"mounts":               1,
		"objects": []interface{}{
			map[string]interface{}{"id": 11, "name": "SAMPLE"},
			map[string]interface{}{"id": 0, "name": "OTHER"},
		},
		"source_id": 5,
	}
	got, err := backuprecovery.ResourceIbmBackupRecoveryProtectionGroupMapToDB2ProtectionGroupParams(in)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if *got.SourceID != 5 || *got.Concurrency != 2 || *got.Mounts != 1 {
		t.Fatalf("unexpected scalars: %+v", got)
	}
	if len(got.Objects) != 2 || *got.Objects[0].ID != 11 || got.Objects[1].ID != nil || *got.Objects[1].Name != "OTHER" {
		t.Fatalf("unexpected objects: %+v", got.Objects)
	}
	if len(got.BackupJobArguments) != 1 || *got.BackupJobArguments[0].Key != "compress" {
		t.Fatalf("unexpected backup job arguments: %+v", got.BackupJobArguments)
	}
	if len(got.ExcludeObjectIds) != 1 || got.ExcludeObjectIds[0] != 7 {
		t.Fatalf("unexpected exclude ids: %+v", got.ExcludeObjectIds)
	}
}

func TestProtectionGroupSapHanaParamsMapper(t *testing.T) {
	in := map[string]interface{}{
		"concurrency": 0,
		"delta":       "incremental",
		"objects":     []interface{}{map[string]interface{}{"id": 3, "name": ""}},
		"source_id":   9,
	}
	got, err := backuprecovery.ResourceIbmBackupRecoveryProtectionGroupMapToSapHanaProtectionGroupParams(in)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if *got.Delta != "incremental" || *got.SourceID != 9 || got.Concurrency != nil {
		t.Fatalf("unexpected result: %+v", got)
	}
	if len(got.Objects) != 1 || *got.Objects[0].ID != 3 || got.Objects[0].Name != nil {
		t.Fatalf("unexpected objects: %+v", got.Objects)
	}
}

func TestProtectionGroupDbParamsSchema(t *testing.T) {
	r := backuprecovery.ResourceIbmBackupRecoveryProtectionGroup()
	for _, k := range []string{"db2_params", "sap_hana_params"} {
		s, ok := r.Schema[k]
		if !ok || !s.Optional || s.MaxItems != 1 {
			t.Fatalf("%s missing or misconfigured: %+v", k, s)
		}
	}
	if err := r.InternalValidate(nil, true); err != nil {
		t.Fatalf("schema invalid: %s", err)
	}
}

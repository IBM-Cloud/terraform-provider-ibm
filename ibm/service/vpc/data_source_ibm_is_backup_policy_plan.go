// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsBackupPolicyPlan() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsBackupPolicyPlanRead,

		Schema: map[string]*schema.Schema{
			"backup_policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The backup policy identifier.",
			},
			"identifier": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The backup policy plan identifier.",
			},
			"name": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ExactlyOneOf: []string{"name", "identifier"},
				Description:  "The unique user-defined name for this backup policy plan.",
			},
			"active": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the plan is active.",
			},
			"attach_user_tags": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User tags to attach to each resource created by this plan.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"copy_user_tags": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to copy the source's user tags to the created resource.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the backup policy plan was created.",
			},
			"cron_spec": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cron specification for the backup schedule.",
			},
			"deletion_trigger": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_after": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of days to keep each backup after creation.",
						},
						"delete_over_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of recent backups to keep. If absent, there is no maximum.",
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this backup policy plan.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of this backup policy plan.",
			},
			"clone_policy": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_snapshots": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of recent snapshots (per source) that will keep clones.",
						},
						"zones": &schema.Schema{
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
							Description: "The zone this backup policy plan will create snapshot clones in.",
						},
					},
				},
			},
			"remote_region_policy": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Policies for creating remote copies of this backup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_over_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of recent remote copies to keep in this region.",
						},
						"encryption_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Services Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The globally unique name for this region.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of resource referenced.",
			},
		},
	}
}

func dataSourceIBMIsBackupPolicyPlanRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := vpcClient(meta)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("vpcClient creation failed: %s", err.Error()), "(Data) ibm_is_backup_policy_plan", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	var backupPolicyPlan *vpcv1.BackupPolicyPlan

	if v, ok := d.GetOk("identifier"); ok {
		id := v.(string)
		getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}

		getBackupPolicyPlanOptions.SetBackupPolicyID(d.Get("backup_policy_id").(string))
		getBackupPolicyPlanOptions.SetID(id)

		backupPolicyPlanInfo, response, err := sess.GetBackupPolicyPlanWithContext(context, getBackupPolicyPlanOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBackupPolicyPlanWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_is_backup_policy_plan", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		backupPolicyPlan = backupPolicyPlanInfo
	} else if v, ok := d.GetOk("name"); ok {

		name := v.(string)
		listBackupPolicyPlansOptions := &vpcv1.ListBackupPolicyPlansOptions{}

		listBackupPolicyPlansOptions.SetBackupPolicyID(d.Get("backup_policy_id").(string))

		backupPolicyPlanCollection, response, err := sess.ListBackupPolicyPlansWithContext(context, listBackupPolicyPlansOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBackupPolicyPlansWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_is_backup_policy_plan", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		for _, backupPolicyPlanInfo := range backupPolicyPlanCollection.Plans {
			if *backupPolicyPlanInfo.Name == name {
				backupPolicyPlan = &backupPolicyPlanInfo
				break
			}
		}
		if backupPolicyPlan == nil {
			err = fmt.Errorf("[ERROR] No backup policy plan found with name (%s)", name)
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBackupPolicyPlansWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_is_backup_policy_plan", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	d.SetId(*backupPolicyPlan.ID)

	if err = d.Set("active", backupPolicyPlan.Active); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting active: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-active").GetDiag()
	}
	if err = d.Set("attach_user_tags", backupPolicyPlan.AttachUserTags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting active: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-attach_user_tags").GetDiag()
	}
	if err = d.Set("copy_user_tags", backupPolicyPlan.CopyUserTags); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting copy_user_tags: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-copy_user_tags").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(backupPolicyPlan.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("cron_spec", backupPolicyPlan.CronSpec); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cron_spec: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-cron_spec").GetDiag()
	}

	if backupPolicyPlan.DeletionTrigger != nil {
		err = d.Set("deletion_trigger", dataSourceBackupPolicyPlanFlattenDeletionTrigger(*backupPolicyPlan.DeletionTrigger))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting deletion_trigger: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-deletion_trigger").GetDiag()
		}
	}
	if backupPolicyPlan.ClonePolicy != nil {
		backupPolicyPlanClonePolicyMap := []map[string]interface{}{}
		finalList := map[string]interface{}{}

		if backupPolicyPlan.ClonePolicy.MaxSnapshots != nil {
			finalList["max_snapshots"] = flex.IntValue(backupPolicyPlan.ClonePolicy.MaxSnapshots)
		}
		if backupPolicyPlan.ClonePolicy.Zones != nil && len(backupPolicyPlan.ClonePolicy.Zones) != 0 {
			zoneList := []string{}
			for i := 0; i < len(backupPolicyPlan.ClonePolicy.Zones); i++ {
				zoneList = append(zoneList, string(*(backupPolicyPlan.ClonePolicy.Zones[i].Name)))
			}
			finalList["zones"] = flex.NewStringSet(schema.HashString, zoneList)
		}
		backupPolicyPlanClonePolicyMap = append(backupPolicyPlanClonePolicyMap, finalList)
		d.Set("clone_policy", backupPolicyPlanClonePolicyMap)
	}
	if err = d.Set("href", backupPolicyPlan.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-href").GetDiag()
	}
	if err = d.Set("lifecycle_state", backupPolicyPlan.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("name", backupPolicyPlan.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-name").GetDiag()
	}
	remoteRegionPolicies := []map[string]interface{}{}
	if backupPolicyPlan.RemoteRegionPolicies != nil {
		for _, remoteCopyPolicy := range backupPolicyPlan.RemoteRegionPolicies {
			remoteRegionPoliciesMap, err := dataSourceIBMIsVPCBackupPolicyPlanRemoteCopyPolicyItemToMap(&remoteCopyPolicy)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_backup_policy_plan", "read", "remote_region_policies-to-map").GetDiag()
			}
			remoteRegionPolicies = append(remoteRegionPolicies, remoteRegionPoliciesMap)
		}
	}
	if err = d.Set("remote_region_policy", remoteRegionPolicies); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting remote_region_policies: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-remote_region_policies").GetDiag()
	}
	if err = d.Set("resource_type", backupPolicyPlan.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_backup_policy_plan", "read", "set-resource_type").GetDiag()
	}

	return nil
}

func dataSourceBackupPolicyPlanFlattenDeletionTrigger(result vpcv1.BackupPolicyPlanDeletionTrigger) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceBackupPolicyPlanDeletionTriggerToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceBackupPolicyPlanDeletionTriggerToMap(deletionTriggerItem vpcv1.BackupPolicyPlanDeletionTrigger) (deletionTriggerMap map[string]interface{}) {
	deletionTriggerMap = map[string]interface{}{}

	if deletionTriggerItem.DeleteAfter != nil {
		deletionTriggerMap["delete_after"] = deletionTriggerItem.DeleteAfter
	}
	if deletionTriggerItem.DeleteOverCount != nil {
		deletionTriggerMap["delete_over_count"] = deletionTriggerItem.DeleteOverCount
	}

	return deletionTriggerMap
}

func dataSourceIBMIsVPCBackupPolicyPlanRemoteCopyPolicyItemToMap(remoteCopyPolicyItem *vpcv1.BackupPolicyPlanRemoteRegionPolicy) (map[string]interface{}, error) {
	remoteCopyPolicyItemMap := make(map[string]interface{})
	if remoteCopyPolicyItem.DeleteOverCount != nil {
		remoteCopyPolicyItemMap["delete_over_count"] = *remoteCopyPolicyItem.DeleteOverCount
	}
	if remoteCopyPolicyItem.EncryptionKey != nil {
		remoteCopyPolicyItemMap["encryption_key"] = *remoteCopyPolicyItem.EncryptionKey.CRN
	}
	if remoteCopyPolicyItem.Region.Name != nil {
		remoteCopyPolicyItemMap["region"] = *remoteCopyPolicyItem.Region.Name
	}
	return remoteCopyPolicyItemMap, nil
}

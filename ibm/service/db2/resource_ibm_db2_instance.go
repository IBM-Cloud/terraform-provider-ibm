// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package db2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	RsInstanceSuccessStatus       = "active"
	RsInstanceProgressStatus      = "in progress"
	RsInstanceProvisioningStatus  = "provisioning"
	RsInstanceInactiveStatus      = "inactive"
	RsInstanceFailStatus          = "failed"
	RsInstanceRemovedStatus       = "removed"
	RsInstanceReclamation         = "pending_reclamation"
	RsInstanceUpdateSuccessStatus = "succeeded"
	PerformanceSubscription       = "PerformanceSubscription"
)

func ResourceIBMDb2Instance() *schema.Resource {
	riSchema := resourcecontroller.ResourceIBMResourceInstance().Schema

	riSchema["high_availability"] = &schema.Schema{
		Description: "If you require high availability, please choose this option",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["instance_type"] = &schema.Schema{
		Description: "Available machine type flavours (default selection will assume smallest configuration)",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["backup_location"] = &schema.Schema{
		Description: "Cross Regional backups can be stored across multiple regions in a zone. Regional backups are stored in only specific region.",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["disk_encryption_instance_crn"] = &schema.Schema{
		Description: "Cross Regional disk encryption crn",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["disk_encryption_key_crn"] = &schema.Schema{
		Description: "Cross Regional disk encryption crn",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["oracle_compatibility"] = &schema.Schema{
		Description: "Indicates whether is has compatibility for oracle or not",
		Optional:    true,
		Type:        schema.TypeString,
	}

	riSchema["subscription_id"] = &schema.Schema{
		Description: "For PerformanceSubscription plans a Subscription ID is required. It is not required for Performance plans.",
		Optional:    true,
		Type:        schema.TypeString,
	}

	return &schema.Resource{
		Create:   resourceIBMDb2InstanceCreate,
		Read:     resourcecontroller.ResourceIBMResourceInstanceRead,
		Update:   resourcecontroller.ResourceIBMResourceInstanceUpdate,
		Delete:   resourcecontroller.ResourceIBMResourceInstanceDelete,
		Exists:   resourcecontroller.ResourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),

		Schema: riSchema,
	}
}

func resourceIBMDb2InstanceCreate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}

	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := rc.CreateResourceInstanceOptions{
		Name: &name,
	}

	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving service offering: %s", err)
	}

	if metadata, ok := serviceOff[0].Metadata.(*models.ServiceResourceMetadata); ok {
		if !metadata.Service.RCProvisionable {
			return fmt.Errorf("%s cannot be provisioned by resource controller", serviceName)
		}
	} else {
		return fmt.Errorf("[ERROR] Cannot create instance of resource %s\nUse 'ibm_service_instance' if the resource is a Cloud Foundry service", serviceName)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving plan: %s", err)
	}
	rsInst.ResourcePlanID = &servicePlan

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return fmt.Errorf("[ERROR] Error retrieving deployment for plan %s : %s", plan, err)
	}
	if len(deployments) == 0 {
		return fmt.Errorf("[ERROR] No deployment found for service plan : %s", plan)
	}
	deployments, supportedLocations := resourcecontroller.FilterDeployments(deployments, location)

	if len(deployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		return fmt.Errorf("[ERROR] No deployment found for service plan %s at location %s.\nValid location(s) are: %q.\nUse 'ibm_service_instance' if the service is a Cloud Foundry service", plan, location, locationList)
	}

	rsInst.Target = &deployments[0].CatalogCRN

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rg := rsGrpID.(string)
		rsInst.ResourceGroup = &rg
	} else {
		defaultRg, err := flex.DefaultResourceGroup(meta)
		if err != nil {
			return err
		}
		rsInst.ResourceGroup = &defaultRg
	}

	params := map[string]interface{}{}

	if serviceEndpoints, ok := d.GetOk("service_endpoints"); ok {
		params["service-endpoints"] = serviceEndpoints.(string)
	}
	if highAvailability, ok := d.GetOk("high_availability"); ok {
		params["high_availability"] = highAvailability.(string)
	}
	if instanceType, ok := d.GetOk("instance_type"); ok {
		params["instance_type"] = instanceType.(string)
	}
	if backupLocation, ok := d.GetOk("backup_location"); ok {
		params["backup-locations"] = backupLocation.(string)
	}

	if diskEncryptionInstanceCrn, ok := d.GetOk("disk_encryption_instance_crn"); ok {
		params["disk_encryption_instance_crn"] = diskEncryptionInstanceCrn.(string)
	}

	if diskEncryptionKeyCrn, ok := d.GetOk("disk_encryption_key_crn"); ok {
		params["disk_encryption_key_crn"] = diskEncryptionKeyCrn.(string)
	}

	if oracleCompatibility, ok := d.GetOk("oracle_compatibility"); ok {
		params["oracle_compatibility"] = oracleCompatibility.(string)
	}

	if plan == PerformanceSubscription {
		if subscriptionId, ok := d.GetOk("subscription_id"); ok {
			params["subscription_id"] = subscriptionId.(string)
		} else {
			return fmt.Errorf("[ERROR] Missing required field 'subscription_id' while creating an instance for plan: %s", plan)
		}
	}

	if parameters, ok := d.GetOk("parameters"); ok {
		temp := parameters.(map[string]interface{})
		for k, v := range temp {
			if v == "true" || v == "false" {
				b, _ := strconv.ParseBool(v.(string))
				params[k] = b
			} else if strings.HasPrefix(v.(string), "[") && strings.HasSuffix(v.(string), "]") {
				//transform v.(string) to be []string
				arrayString := v.(string)
				result := []string{}
				trimLeft := strings.TrimLeft(arrayString, "[")
				trimRight := strings.TrimRight(trimLeft, "]")
				if len(trimRight) == 0 {
					params[k] = result
				} else {
					array := strings.Split(trimRight, ",")
					for _, a := range array {
						result = append(result, strings.Trim(a, "\""))
					}
					params[k] = result
				}
			} else {
				params[k] = v
			}
		}

	}

	if s, ok := d.GetOk("parameters_json"); ok {
		json.Unmarshal([]byte(s.(string)), &params)
	}

	rsInst.Parameters = params

	//Start to create resource instance
	instance, resp, err := rsConClient.CreateResourceInstance(&rsInst)
	if err != nil {
		log.Printf(
			"Error when creating resource instance: %s, Instance info  NAME->%s, LOCATION->%s, GROUP_ID->%s, PLAN_ID->%s",
			err, *rsInst.Name, *rsInst.Target, *rsInst.ResourceGroup, *rsInst.ResourcePlanID)
		return fmt.Errorf("[ERROR] Error when creating resource instance: %s with resp code: %s", err, resp)
	}

	d.SetId(*instance.ID)

	_, err = waitForResourceInstanceCreate(d, meta)
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for create resource instance (%s) to be succeeded: %s", d.Id(), err)
	}

	db2SaasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		return err
	}

	if allolistConfigRaw, ok := d.GetOk("allowlist_config"); ok {
		if allolistConfigRaw == nil || reflect.ValueOf(allolistConfigRaw).IsNil() {
			fmt.Println("No allowlisting config is provided, Skipping.")
		} else {
			allowelistConfig := allolistConfigRaw.([]interface{})[0].(map[string]interface{})

			ipAddress := make([]db2saasv1.IpAddress, 0, len(allowelistConfig["ip_addresses"].([]interface{})))

			for _, ip := range ipAddress {
				if err = validateIPAddress(ip); err != nil {
					return err
				}
			}

			input := &db2saasv1.PostDb2SaasAllowlistOptions{
				XDeploymentID: core.StringPtr(*instance.CRN),
				IpAddresses:   ipAddress,
			}

			result, response, err := db2SaasClient.PostDb2SaasAllowlist(input)
			if err != nil {
				log.Printf("Error while posting allowlist to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}
	}

	if autoscaleConfigRaw, ok := d.GetOk("autoscale_config"); ok {
		if autoscaleConfigRaw == nil || reflect.ValueOf(autoscaleConfigRaw).IsNil() {
			fmt.Println("No autoscaling config is provided, Skipping.")
		} else {
			autoscalingConfig := autoscaleConfigRaw.([]interface{})[0].(map[string]interface{})

			var (
				autoScalingThreshold      int
				autoScalingOverTimePeriod int
				autoScalingPauseLimit     int
			)

			if autoscalingConfig["auto_scaling_threshold"] != nil {
				autoScalingThreshold, err = strconv.Atoi(autoscalingConfig["auto_scaling_threshold"].(string))
				if err != nil {
					return err
				}
			}

			if autoscalingConfig["auto_scaling_over_time_period"] != nil {
				autoScalingOverTimePeriod, err = strconv.Atoi(autoscalingConfig["auto_scaling_over_time_period"].(string))
				if err != nil {
					return err
				}
			}

			if autoscalingConfig["auto_scaling_pause_limit"] != nil {
				autoScalingPauseLimit, err = strconv.Atoi(autoscalingConfig["auto_scaling_pause_limit"].(string))
				if err != nil {
					return err
				}
			}

			input := &db2saasv1.PutDb2SaasAutoscaleOptions{
				XDbProfile:                core.StringPtr(*instance.CRN),
				AutoScalingEnabled:        core.StringPtr("YES"),
				AutoScalingAllowPlanLimit: core.StringPtr("YES"),
				AutoScalingThreshold:      core.Int64Ptr(int64(autoScalingThreshold)),
				AutoScalingOverTimePeriod: core.Float64Ptr(float64(autoScalingOverTimePeriod)),
				AutoScalingPauseLimit:     core.Int64Ptr(int64(autoScalingPauseLimit)),
			}

			result, response, err := db2SaasClient.PutDb2SaasAutoscale(input)
			if err != nil {
				log.Printf("Error while updating autoscaling config to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}

	}

	if userDetailsRaw, ok := d.GetOk("db2_userdetails"); ok {
		if userDetailsRaw == nil || reflect.ValueOf(userDetailsRaw).IsNil() {
			fmt.Println("No user details configs provided; skipping")
		} else {
			userDetais := userDetailsRaw.([]interface{})[0].(map[string]interface{})
			fmt.Println(userDetailsRaw)

			input := &db2saasv1.PostDb2SaasUserOptions{
				XDeploymentID: core.StringPtr(*instance.CRN),
				ID:            core.StringPtr(userDetais["id"].(string)),
				Iam:           core.BoolPtr(userDetais["iam"].(bool)),
				Ibmid:         core.StringPtr(userDetais["ibmid"].(string)),
				Name:          core.StringPtr(userDetais["name"].(string)),
				Password:      core.StringPtr(userDetais["password"].(string)),
				Role:          core.StringPtr(userDetais["role"].(string)),
				Email:         core.StringPtr(userDetais["email"].(string)),
				Locked:        core.StringPtr(userDetais["locked"].(string)),
				Authentication: &db2saasv1.CreateUserAuthentication{
					Method:   core.StringPtr("internal"),
					PolicyID: core.StringPtr("Default"),
				},
			}

			result, response, err := db2SaasClient.PostDb2SaasUser(input)
			if err != nil {
				log.Printf("Error while posting users to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}
	}

	if backupConfigRaw, ok := d.GetOk("db2_backup"); ok {
		if backupConfigRaw == nil || reflect.ValueOf(backupConfigRaw).IsNil() {
			fmt.Println("No backup configs provided; skipping.")
		} else {
			backupConfig := backupConfigRaw.([]interface{})[0].(map[string]interface{})
			fmt.Println(backupConfig)

			input := &db2saasv1.PostDb2SaasBackupOptions{
				XDbProfile: core.StringPtr(*instance.CRN),
			}

			result, response, err := db2SaasClient.PostDb2SaasBackup(input)
			if err != nil {
				log.Printf("Error while posting backup to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}

		}
	}

	if customSettingRaw, ok := d.GetOk("dbm_configuration"); ok {
		if customSettingRaw == nil || reflect.ValueOf(customSettingRaw).IsNil() {
			fmt.Println("No custom setting configs provided; skipping.")
		} else {
			customerSettingConfig := customSettingRaw.([]interface{})[0].(map[string]interface{})
			fmt.Println(customerSettingConfig)

			var (
				registryConfig, dbConfig, dbmConfig map[string]interface{}
				registry                            *db2saasv1.CreateCustomSettingsRegistry
				db                                  *db2saasv1.CreateCustomSettingsDb
				dbm                                 *db2saasv1.CreateCustomSettingsDbm
			)

			registryConfigRaw := customerSettingConfig["registry"]
			if registryConfigRaw == nil || reflect.ValueOf(registryConfigRaw).IsNil() {
				fmt.Println("No custom setting registry configs provided; skipping.")
			} else {
				registryConfig = customerSettingConfig["registry"].(map[string]interface{})

				registry = &db2saasv1.CreateCustomSettingsRegistry{
					DB2BIDI:                      core.StringPtr(registryConfig["DB2BIDI"].(string)),
					DB2COMPOPT:                   core.StringPtr(registryConfig["DB2COMPOPT"].(string)),
					DB2LOCKTORB:                  core.StringPtr(registryConfig["DB2LOCK_TO_RB"].(string)),
					DB2STMM:                      core.StringPtr(registryConfig["DB2STMM"].(string)),
					DB2ALTERNATEAUTHZBEHAVIOUR:   core.StringPtr(registryConfig["DB2_ALTERNATE_AUTHZ_BEHAVIOUR"].(string)),
					DB2ANTIJOIN:                  core.StringPtr(registryConfig["DB2_ANTIJOIN"].(string)),
					DB2ATSENABLE:                 core.StringPtr(registryConfig["DB2_ATS_ENABLE"].(string)),
					DB2DEFERREDPREPARESEMANTICS:  core.StringPtr(registryConfig["DB2_DEFERRED_PREPARE_SEMANTICS"].(string)),
					DB2EVALUNCOMMITTED:           core.StringPtr(registryConfig["DB2_EVALUNCOMMITTED"].(string)),
					DB2EXTENDEDOPTIMIZATION:      core.StringPtr(registryConfig["DB2_EXTENDED_OPTIMIZATION"].(string)),
					DB2INDEXPCTFREEDEFAULT:       core.StringPtr(registryConfig["DB2_INDEX_PCTFREE_DEFAULT"].(string)),
					DB2INLISTTONLJN:              core.StringPtr(registryConfig["DB2_INLIST_TO_NLJN"].(string)),
					DB2MINIMIZELISTPREFETCH:      core.StringPtr(registryConfig["DB2_MINIMIZE_LISTPREFETCH"].(string)),
					DB2OBJECTTABLEENTRIES:        core.StringPtr(registryConfig["DB2_OBJECT_TABLE_ENTRIES"].(string)),
					DB2OPTPROFILE:                core.StringPtr(registryConfig["DB2_OPTPROFILE"].(string)),
					DB2OPTSTATSLOG:               core.StringPtr(registryConfig["DB2_OPTSTATS_LOG"].(string)),
					DB2OPTMAXTEMPSIZE:            core.StringPtr(registryConfig["DB2_OPT_MAX_TEMP_SIZE"].(string)),
					DB2PARALLELIO:                core.StringPtr(registryConfig["DB2_PARALLEL_IO"].(string)),
					DB2REDUCEDOPTIMIZATION:       core.StringPtr(registryConfig["DB2_REDUCED_OPTIMIZATION"].(string)),
					DB2SELECTIVITY:               core.StringPtr(registryConfig["DB2_SELECTIVITY"].(string)),
					DB2SKIPDELETED:               core.StringPtr(registryConfig["DB2_SKIPDELETED"].(string)),
					DB2SKIPINSERTED:              core.StringPtr(registryConfig["DB2_SKIPINSERTED"].(string)),
					DB2SYNCRELEASELOCKATTRIBUTES: core.StringPtr(registryConfig["DB2_SYNC_RELEASE_LOCK_ATTRIBUTES"].(string)),
					DB2TRUNCATEREUSESTORAGE:      core.StringPtr(registryConfig["DB2_TRUNCATE_REUSESTORAGE"].(string)),
					DB2USEALTERNATEPAGECLEANING:  core.StringPtr(registryConfig["DB2_USE_ALTERNATE_PAGE_CLEANING"].(string)),
					DB2VIEWREOPTVALUES:           core.StringPtr(registryConfig["DB2_VIEW_REOPT_VALUES"].(string)),
					DB2WLMSETTINGS:               core.StringPtr(registryConfig["DB2_WLM_SETTINGS"].(string)),
					DB2WORKLOAD:                  core.StringPtr(registryConfig["DB2_WORKLOAD"].(string)),
				}
			}

			dbConfigRaw := customerSettingConfig["db"]
			if dbConfigRaw == nil || reflect.ValueOf(dbConfigRaw).IsNil() {
				fmt.Println("No custom setting db configs provided; skipping.")
			} else {
				dbConfig = customerSettingConfig["db"].(map[string]interface{})

				db = &db2saasv1.CreateCustomSettingsDb{
					ACTSORTMEMLIMIT:    core.StringPtr(dbConfig["ACT_SORTMEM_LIMIT"].(string)),
					ALTCOLLATE:         core.StringPtr(dbConfig["ALT_COLLATE"].(string)),
					APPGROUPMEMSZ:      core.StringPtr(dbConfig["APPGROUP_MEM_SZ"].(string)),
					APPLHEAPSZ:         core.StringPtr(dbConfig["APPLHEAPSZ"].(string)),
					APPLMEMORY:         core.StringPtr(dbConfig["APPL_MEMORY"].(string)),
					APPCTLHEAPSZ:       core.StringPtr(dbConfig["APP_CTL_HEAP_SZ"].(string)),
					ARCHRETRYDELAY:     core.StringPtr(dbConfig["ARCHRETRYDELAY"].(string)),
					AUTHNCACHEDURATION: core.StringPtr(dbConfig["AUTHN_CACHE_DURATION"].(string)),
					AUTORESTART:        core.StringPtr(dbConfig["AUTORESTART"].(string)),
					AUTOCGSTATS:        core.StringPtr(dbConfig["AUTO_CG_STATS"].(string)),
					AUTOMAINT:          core.StringPtr(dbConfig["AUTO_MAINT"].(string)),
					AUTOREORG:          core.StringPtr(dbConfig["AUTO_REORG"].(string)),
					AUTOREVAL:          core.StringPtr(dbConfig["AUTO_REVAL"].(string)),
					AUTORUNSTATS:       core.StringPtr(dbConfig["AUTO_RUNSTATS"].(string)),
					AUTOSAMPLING:       core.StringPtr(dbConfig["AUTO_SAMPLING"].(string)),
					AUTOSTATSVIEWS:     core.StringPtr(dbConfig["AUTO_STATS_VIEWS"].(string)),
					AUTOSTMTSTATS:      core.StringPtr(dbConfig["AUTO_STMT_STATS"].(string)),
					AUTOTBLMAINT:       core.StringPtr(dbConfig["AUTO_TBL_MAINT"].(string)),
					AVGAPPLS:           core.StringPtr(dbConfig["AVG_APPLS"].(string)),
					CATALOGCACHESZ:     core.StringPtr(dbConfig["CATALOGCACHE_SZ"].(string)),
					CHNGPGSTHRESH:      core.StringPtr(dbConfig["CHNGPGS_THRESH"].(string)),
					CURCOMMIT:          core.StringPtr(dbConfig["CUR_COMMIT"].(string)),
					DATABASEMEMORY:     core.StringPtr(dbConfig["DATABASE_MEMORY"].(string)),
					DBHEAP:             core.StringPtr(dbConfig["DBHEAP"].(string)),
					DBCOLLNAME:         core.StringPtr(dbConfig["DB_COLLNAME"].(string)),
					DBMEMTHRESH:        core.StringPtr(dbConfig["DB_MEM_THRESH"].(string)),
					DDLCOMPRESSIONDEF:  core.StringPtr(dbConfig["DDL_COMPRESSION_DEF"].(string)),
					DDLCONSTRAINTDEF:   core.StringPtr(dbConfig["DDL_CONSTRAINT_DEF"].(string)),
					DECFLTROUNDING:     core.StringPtr(dbConfig["DECFLT_ROUNDING"].(string)),
					DECARITHMETIC:      core.StringPtr(dbConfig["DEC_ARITHMETIC"].(string)),
					DECTOCHARFMT:       core.StringPtr(dbConfig["DEC_TO_CHAR_FMT"].(string)),
					DFTDEGREE:          core.StringPtr(dbConfig["DFT_DEGREE"].(string)),
					DFTEXTENTSZ:        core.StringPtr(dbConfig["DFT_EXTENT_SZ"].(string)),
					DFTLOADRECSES:      core.StringPtr(dbConfig["DFT_LOADREC_SES"].(string)),
					DFTMTTBTYPES:       core.StringPtr(dbConfig["DFT_MTTB_TYPES"].(string)),
					DFTPREFETCHSZ:      core.StringPtr(dbConfig["DFT_PREFETCH_SZ"].(string)),
					DFTQUERYOPT:        core.StringPtr(dbConfig["DFT_QUERYOPT"].(string)),
					DFTREFRESHAGE:      core.StringPtr(dbConfig["DFT_REFRESH_AGE"].(string)),
					DFTSCHEMASDCC:      core.StringPtr(dbConfig["DFT_SCHEMAS_DCC"].(string)),
					DFTSQLMATHWARN:     core.StringPtr(dbConfig["DFT_SQLMATHWARN"].(string)),
					DFTTABLEORG:        core.StringPtr(dbConfig["DFT_TABLE_ORG"].(string)),
					DLCHKTIME:          core.StringPtr(dbConfig["DLCHKTIME"].(string)),
					ENABLEXMLCHAR:      core.StringPtr(dbConfig["ENABLE_XMLCHAR"].(string)),
					EXTENDEDROWSZ:      core.StringPtr(dbConfig["EXTENDED_ROW_SZ"].(string)),
					GROUPHEAPRATIO:     core.StringPtr(dbConfig["GROUPHEAP_RATIO"].(string)),
					INDEXREC:           core.StringPtr(dbConfig["INDEXREC"].(string)),
					LARGEAGGREGATION:   core.StringPtr(dbConfig["LARGE_AGGREGATION"].(string)),
					LOCKLIST:           core.StringPtr(dbConfig["LOCKLIST"].(string)),
					LOCKTIMEOUT:        core.StringPtr(dbConfig["LOCKTIMEOUT"].(string)),
					LOGINDEXBUILD:      core.StringPtr(dbConfig["LOGINDEXBUILD"].(string)),
					LOGAPPLINFO:        core.StringPtr(dbConfig["LOG_APPL_INFO"].(string)),
					LOGDDLSTMTS:        core.StringPtr(dbConfig["LOG_DDL_STMTS"].(string)),
					LOGDISKCAP:         core.StringPtr(dbConfig["LOG_DISK_CAP"].(string)),
					MAXAPPLS:           core.StringPtr(dbConfig["MAXAPPLS"].(string)),
					MAXFILOP:           core.StringPtr(dbConfig["MAXFILOP"].(string)),
					MAXLOCKS:           core.StringPtr(dbConfig["MAXLOCKS"].(string)),
					MINDECDIV3:         core.StringPtr(dbConfig["MIN_DEC_DIV_3"].(string)),
					MONACTMETRICS:      core.StringPtr(dbConfig["MON_ACT_METRICS"].(string)),
					MONDEADLOCK:        core.StringPtr(dbConfig["MON_DEADLOCK"].(string)),
					MONLCKMSGLVL:       core.StringPtr(dbConfig["MON_LCK_MSG_LVL"].(string)),
					MONLOCKTIMEOUT:     core.StringPtr(dbConfig["MON_LOCKTIMEOUT"].(string)),
					MONLOCKWAIT:        core.StringPtr(dbConfig["MON_LOCKWAIT"].(string)),
					MONLWTHRESH:        core.StringPtr(dbConfig["MON_LW_THRESH"].(string)),
					MONOBJMETRICS:      core.StringPtr(dbConfig["MON_OBJ_METRICS"].(string)),
					MONPKGLISTSZ:       core.StringPtr(dbConfig["MON_PKGLIST_SZ"].(string)),
					MONREQMETRICS:      core.StringPtr(dbConfig["MON_REQ_METRICS"].(string)),
					MONRTNDATA:         core.StringPtr(dbConfig["MON_RTN_DATA"].(string)),
					MONRTNEXECLIST:     core.StringPtr(dbConfig["MON_RTN_EXECLIST"].(string)),
					MONUOWDATA:         core.StringPtr(dbConfig["MON_UOW_DATA"].(string)),
					MONUOWEXECLIST:     core.StringPtr(dbConfig["MON_UOW_EXECLIST"].(string)),
					MONUOWPKGLIST:      core.StringPtr(dbConfig["MON_UOW_PKGLIST"].(string)),
					NCHARMAPPING:       core.StringPtr(dbConfig["NCHAR_MAPPING"].(string)),
					NUMFREQVALUES:      core.StringPtr(dbConfig["NUM_FREQVALUES"].(string)),
					NUMIOCLEANERS:      core.StringPtr(dbConfig["NUM_IOCLEANERS"].(string)),
					NUMIOSERVERS:       core.StringPtr(dbConfig["NUM_IOSERVERS"].(string)),
					NUMLOGSPAN:         core.StringPtr(dbConfig["NUM_LOG_SPAN"].(string)),
					NUMQUANTILES:       core.StringPtr(dbConfig["NUM_QUANTILES"].(string)),
					OPTBUFFPAGE:        core.StringPtr(dbConfig["OPT_BUFFPAGE"].(string)),
					OPTDIRECTWRKLD:     core.StringPtr(dbConfig["OPT_DIRECT_WRKLD"].(string)),
					OPTLOCKLIST:        core.StringPtr(dbConfig["OPT_LOCKLIST"].(string)),
					OPTMAXLOCKS:        core.StringPtr(dbConfig["OPT_MAXLOCKS"].(string)),
					OPTSORTHEAP:        core.StringPtr(dbConfig["OPT_SORTHEAP"].(string)),
					PAGEAGETRGTGCR:     core.StringPtr(dbConfig["PAGE_AGE_TRGT_GCR"].(string)),
					PAGEAGETRGTMCR:     core.StringPtr(dbConfig["PAGE_AGE_TRGT_MCR"].(string)),
					PCKCACHESZ:         core.StringPtr(dbConfig["PCKCACHESZ"].(string)),
					PLSTACKTRACE:       core.StringPtr(dbConfig["PL_STACK_TRACE"].(string)),
					SELFTUNINGMEM:      core.StringPtr(dbConfig["SELF_TUNING_MEM"].(string)),
					SEQDETECT:          core.StringPtr(dbConfig["SEQDETECT"].(string)),
					SHEAPTHRESSHR:      core.StringPtr(dbConfig["SHEAPTHRES_SHR"].(string)),
					SOFTMAX:            core.StringPtr(dbConfig["SOFTMAX"].(string)),
					SORTHEAP:           core.StringPtr(dbConfig["SORTHEAP"].(string)),
					SQLCCFLAGS:         core.StringPtr(dbConfig["SQL_CCFLAGS"].(string)),
					STATHEAPSZ:         core.StringPtr(dbConfig["STAT_HEAP_SZ"].(string)),
					STMTHEAP:           core.StringPtr(dbConfig["STMTHEAP"].(string)),
					STMTCONC:           core.StringPtr(dbConfig["STMT_CONC"].(string)),
					STRINGUNITS:        core.StringPtr(dbConfig["STRING_UNITS"].(string)),
					SYSTIMEPERIODADJ:   core.StringPtr(dbConfig["SYSTIME_PERIOD_ADJ"].(string)),
					TRACKMOD:           core.StringPtr(dbConfig["TRACKMOD"].(string)),
					UTILHEAPSZ:         core.StringPtr(dbConfig["UTIL_HEAP_SZ"].(string)),
					WLMADMISSIONCTRL:   core.StringPtr(dbConfig["WLM_ADMISSION_CTRL"].(string)),
					WLMAGENTLOADTRGT:   core.StringPtr(dbConfig["WLM_AGENT_LOAD_TRGT"].(string)),
					WLMCPULIMIT:        core.StringPtr(dbConfig["WLM_CPU_LIMIT"].(string)),
					WLMCPUSHARES:       core.StringPtr(dbConfig["WLM_CPU_SHARES"].(string)),
					WLMCPUSHAREMODE:    core.StringPtr(dbConfig["WLM_CPU_SHARE_MODE"].(string)),
				}

			}

			dbmConfigRaw := customerSettingConfig["dbm"]
			if dbmConfigRaw == nil || reflect.ValueOf(dbmConfigRaw).IsNil() {
				fmt.Println("No custom setting dbm configs provided; skipping.")
			} else {
				dbmConfig = customerSettingConfig["dbm"].(map[string]interface{})

				dbm = &db2saasv1.CreateCustomSettingsDbm{
					COMMBANDWIDTH:    core.StringPtr(dbmConfig["COMM_BANDWIDTH"].(string)),
					CPUSPEED:         core.StringPtr(dbmConfig["CPUSPEED"].(string)),
					DFTMONBUFPOOL:    core.StringPtr(dbmConfig["DFT_MON_BUFPOOL"].(string)),
					DFTMONLOCK:       core.StringPtr(dbmConfig["DFT_MON_LOCK"].(string)),
					DFTMONSORT:       core.StringPtr(dbmConfig["DFT_MON_SORT"].(string)),
					DFTMONSTMT:       core.StringPtr(dbmConfig["DFT_MON_STMT"].(string)),
					DFTMONTABLE:      core.StringPtr(dbmConfig["DFT_MON_TABLE"].(string)),
					DFTMONTIMESTAMP:  core.StringPtr(dbmConfig["DFT_MON_TIMESTAMP"].(string)),
					DFTMONUOW:        core.StringPtr(dbmConfig["DFT_MON_UOW"].(string)),
					DIAGLEVEL:        core.StringPtr(dbmConfig["DIAGLEVEL"].(string)),
					FEDERATEDASYNC:   core.StringPtr(dbmConfig["FEDERATED_ASYNC"].(string)),
					INDEXREC:         core.StringPtr(dbmConfig["INDEXREC"].(string)),
					INTRAPARALLEL:    core.StringPtr(dbmConfig["INTRA_PARALLEL"].(string)),
					KEEPFENCED:       core.StringPtr(dbmConfig["KEEPFENCED"].(string)),
					MAXCONNRETRIES:   core.StringPtr(dbmConfig["MAX_CONNRETRIES"].(string)),
					MAXQUERYDEGREE:   core.StringPtr(dbmConfig["MAX_QUERYDEGREE"].(string)),
					MONHEAPSZ:        core.StringPtr(dbmConfig["MON_HEAP_SZ"].(string)),
					MULTIPARTSIZEMB:  core.StringPtr(dbmConfig["MULTIPARTSIZEMB"].(string)),
					NOTIFYLEVEL:      core.StringPtr(dbmConfig["NOTIFYLEVEL"].(string)),
					NUMINITAGENTS:    core.StringPtr(dbmConfig["NUM_INITAGENTS"].(string)),
					NUMINITFENCED:    core.StringPtr(dbmConfig["NUM_INITFENCED"].(string)),
					NUMPOOLAGENTS:    core.StringPtr(dbmConfig["NUM_POOLAGENTS"].(string)),
					RESYNCINTERVAL:   core.StringPtr(dbmConfig["RESYNC_INTERVAL"].(string)),
					RQRIOBLK:         core.StringPtr(dbmConfig["RQRIOBLK"].(string)),
					STARTSTOPTIME:    core.StringPtr(dbmConfig["START_STOP_TIME"].(string)),
					UTILIMPACTLIM:    core.StringPtr(dbmConfig["UTIL_IMPACT_LIM"].(string)),
					WLMDISPATCHER:    core.StringPtr(dbmConfig["WLM_DISPATCHER"].(string)),
					WLMDISPCONCUR:    core.StringPtr(dbmConfig["WLM_DISP_CONCUR"].(string)),
					WLMDISPCPUSHARES: core.StringPtr(dbmConfig["WLM_DISP_CPU_SHARES"].(string)),
					WLMDISPMINUTIL:   core.StringPtr(dbmConfig["WLM_DISP_MIN_UTIL"].(string)),
				}

			}

			input := &db2saasv1.PostDb2SaasDbConfigurationOptions{
				XDbProfile: core.StringPtr(*instance.CRN),
				Registry:   registry,
				Db:         db,
				Dbm:        dbm,
			}

			result, response, err := db2SaasClient.PostDb2SaasDbConfiguration(input)
			if err != nil {
				log.Printf("Error while posting DB configuration to DB2Saas: %s", err)
			} else {
				log.Printf("StatusCode of response %d", response.StatusCode)
				log.Printf("Success result %v", result)
			}
		}
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateTagsUsingCRN(oldList, newList, meta, *instance.CRN)
		if err != nil {
			log.Printf(
				"Error on create of resource instance (%s) tags: %s", d.Id(), err)
		}
	}

	return resourcecontroller.ResourceIBMResourceInstanceRead(d, meta)
}

func waitForResourceInstanceCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()
	resourceInstanceGet := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	stateConf := &retry.StateChangeConf{
		Pending: []string{RsInstanceProgressStatus, RsInstanceInactiveStatus, RsInstanceProvisioningStatus},
		Target:  []string{RsInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, resp, err := rsConClient.GetResourceInstance(&resourceInstanceGet)
			if err != nil {
				if resp != nil && resp.StatusCode == 404 {
					return nil, "", fmt.Errorf("[ERROR] The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", fmt.Errorf("[ERROR] Get the resource instance %s failed with resp code: %s, err: %v", d.Id(), resp, err)
			}
			if *instance.State == RsInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("[ERROR] The resource instance '%s' creation failed: %v", d.Id(), err)
			}
			return instance, *instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      30 * time.Second,
		MinTimeout: 30 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func validateIPAddress(ip db2saasv1.IpAddress) error {
	if ip.Address == nil || *ip.Address == "" {
		return fmt.Errorf("[ERROR] IP address is required")
	}
	if ip.Description == nil || *ip.Description == "" {
		return fmt.Errorf("[ERROR] IP address is required")
	}
	return nil
}

// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.0-d6dec9d7-20241008-212902
 */

package db2_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/db2"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmDb2TuneableParamDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDb2TuneableParamDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_db2_tuneable_param.Db2-44-test-both", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmDb2TuneableParamDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_db2_tuneable_param" "Db2-44-test-both" {
		}
	`)
}

func TestDataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		successTuneableParamsTuneableParamDbModel := make(map[string]interface{})
		successTuneableParamsTuneableParamDbModel["act_sortmem_limit"] = "'NONE', 'range(10, 100)'"
		successTuneableParamsTuneableParamDbModel["alt_collate"] = "'NULL', 'IDENTITY_16BIT'"
		successTuneableParamsTuneableParamDbModel["appgroup_mem_sz"] = "'range(1, 1000000)'"
		successTuneableParamsTuneableParamDbModel["applheapsz"] = "'AUTOMATIC', 'range(16, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["appl_memory"] = "'AUTOMATIC', 'range(128, 4294967295)'"
		successTuneableParamsTuneableParamDbModel["app_ctl_heap_sz"] = "'range(1, 64000)'"
		successTuneableParamsTuneableParamDbModel["archretrydelay"] = "'range(0, 65535)'"
		successTuneableParamsTuneableParamDbModel["authn_cache_duration"] = "'range(1,10000)'"
		successTuneableParamsTuneableParamDbModel["autorestart"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_cg_stats"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_maint"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_reorg"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_reval"] = "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'"
		successTuneableParamsTuneableParamDbModel["auto_runstats"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_sampling"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_stats_views"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_stmt_stats"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["auto_tbl_maint"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["avg_appls"] = "'-'"
		successTuneableParamsTuneableParamDbModel["catalogcache_sz"] = "'-'"
		successTuneableParamsTuneableParamDbModel["chngpgs_thresh"] = "'range(5,99)'"
		successTuneableParamsTuneableParamDbModel["cur_commit"] = "'ON, AVAILABLE, DISABLED'"
		successTuneableParamsTuneableParamDbModel["database_memory"] = "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'"
		successTuneableParamsTuneableParamDbModel["dbheap"] = "'AUTOMATIC', 'range(32, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["db_collname"] = "'-'"
		successTuneableParamsTuneableParamDbModel["db_mem_thresh"] = "'range(0, 100)'"
		successTuneableParamsTuneableParamDbModel["ddl_compression_def"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["ddl_constraint_def"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["decflt_rounding"] = "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'"
		successTuneableParamsTuneableParamDbModel["dec_arithmetic"] = "'-'"
		successTuneableParamsTuneableParamDbModel["dec_to_char_fmt"] = "'NEW', 'V95'"
		successTuneableParamsTuneableParamDbModel["dft_degree"] = "'-1', 'ANY', 'range(1, 32767)'"
		successTuneableParamsTuneableParamDbModel["dft_extent_sz"] = "'range(2, 256)'"
		successTuneableParamsTuneableParamDbModel["dft_loadrec_ses"] = "'range(1, 30000)'"
		successTuneableParamsTuneableParamDbModel["dft_mttb_types"] = "'-'"
		successTuneableParamsTuneableParamDbModel["dft_prefetch_sz"] = "'range(0, 32767)', 'AUTOMATIC'"
		successTuneableParamsTuneableParamDbModel["dft_queryopt"] = "'range(0, 9)'"
		successTuneableParamsTuneableParamDbModel["dft_refresh_age"] = "'-'"
		successTuneableParamsTuneableParamDbModel["dft_schemas_dcc"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["dft_sqlmathwarn"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["dft_table_org"] = "'COLUMN', 'ROW'"
		successTuneableParamsTuneableParamDbModel["dlchktime"] = "'range(1000, 600000)'"
		successTuneableParamsTuneableParamDbModel["enable_xmlchar"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["extended_row_sz"] = "'ENABLE', 'DISABLE'"
		successTuneableParamsTuneableParamDbModel["groupheap_ratio"] = "'range(1, 99)'"
		successTuneableParamsTuneableParamDbModel["indexrec"] = "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'"
		successTuneableParamsTuneableParamDbModel["large_aggregation"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["locklist"] = "'AUTOMATIC', 'range(4, 134217728)'"
		successTuneableParamsTuneableParamDbModel["locktimeout"] = "'-1', 'range(0, 32767)'"
		successTuneableParamsTuneableParamDbModel["logindexbuild"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["log_appl_info"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["log_ddl_stmts"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["log_disk_cap"] = "'0', '-1', 'range(1, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["maxappls"] = "'range(1, 60000)'"
		successTuneableParamsTuneableParamDbModel["maxfilop"] = "'range(64, 61440)'"
		successTuneableParamsTuneableParamDbModel["maxlocks"] = "'AUTOMATIC', 'range(1, 100)'"
		successTuneableParamsTuneableParamDbModel["min_dec_div_3"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["mon_act_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		successTuneableParamsTuneableParamDbModel["mon_deadlock"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		successTuneableParamsTuneableParamDbModel["mon_lck_msg_lvl"] = "'range(0, 3)'"
		successTuneableParamsTuneableParamDbModel["mon_locktimeout"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		successTuneableParamsTuneableParamDbModel["mon_lockwait"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		successTuneableParamsTuneableParamDbModel["mon_lw_thresh"] = "'range(1000, 4294967295)'"
		successTuneableParamsTuneableParamDbModel["mon_obj_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		successTuneableParamsTuneableParamDbModel["mon_pkglist_sz"] = "'range(0, 1024)'"
		successTuneableParamsTuneableParamDbModel["mon_req_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		successTuneableParamsTuneableParamDbModel["mon_rtn_data"] = "'NONE', 'BASE'"
		successTuneableParamsTuneableParamDbModel["mon_rtn_execlist"] = "'OFF', 'ON'"
		successTuneableParamsTuneableParamDbModel["mon_uow_data"] = "'NONE', 'BASE'"
		successTuneableParamsTuneableParamDbModel["mon_uow_execlist"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["mon_uow_pkglist"] = "'OFF', 'ON'"
		successTuneableParamsTuneableParamDbModel["nchar_mapping"] = "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'"
		successTuneableParamsTuneableParamDbModel["num_freqvalues"] = "'range(0, 32767)'"
		successTuneableParamsTuneableParamDbModel["num_iocleaners"] = "'AUTOMATIC', 'range(0, 255)'"
		successTuneableParamsTuneableParamDbModel["num_ioservers"] = "'AUTOMATIC', 'range(1, 255)'"
		successTuneableParamsTuneableParamDbModel["num_log_span"] = "'range(0, 65535)'"
		successTuneableParamsTuneableParamDbModel["num_quantiles"] = "'range(0, 32767)'"
		successTuneableParamsTuneableParamDbModel["opt_buffpage"] = "'-'"
		successTuneableParamsTuneableParamDbModel["opt_direct_wrkld"] = "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'"
		successTuneableParamsTuneableParamDbModel["opt_locklist"] = "'-'"
		successTuneableParamsTuneableParamDbModel["opt_maxlocks"] = "'-'"
		successTuneableParamsTuneableParamDbModel["opt_sortheap"] = "'-'"
		successTuneableParamsTuneableParamDbModel["page_age_trgt_gcr"] = "'range(1, 65535)'"
		successTuneableParamsTuneableParamDbModel["page_age_trgt_mcr"] = "'range(1, 65535)'"
		successTuneableParamsTuneableParamDbModel["pckcachesz"] = "'AUTOMATIC', '-1', 'range(32, 2147483646)'"
		successTuneableParamsTuneableParamDbModel["pl_stack_trace"] = "'NONE', 'ALL', 'UNHANDLED'"
		successTuneableParamsTuneableParamDbModel["self_tuning_mem"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbModel["seqdetect"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["sheapthres_shr"] = "'AUTOMATIC', 'range(250, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["softmax"] = "'-'"
		successTuneableParamsTuneableParamDbModel["sortheap"] = "'AUTOMATIC', 'range(16, 4294967295)'"
		successTuneableParamsTuneableParamDbModel["sql_ccflags"] = "'-'"
		successTuneableParamsTuneableParamDbModel["stat_heap_sz"] = "'AUTOMATIC', 'range(1096, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["stmtheap"] = "'AUTOMATIC', 'range(128, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["stmt_conc"] = "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'"
		successTuneableParamsTuneableParamDbModel["string_units"] = "'SYSTEM', 'CODEUNITS32'"
		successTuneableParamsTuneableParamDbModel["systime_period_adj"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamDbModel["trackmod"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["util_heap_sz"] = "'AUTOMATIC', 'range(16, 2147483647)'"
		successTuneableParamsTuneableParamDbModel["wlm_admission_ctrl"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbModel["wlm_agent_load_trgt"] = "'AUTOMATIC', 'range(1, 65535)'"
		successTuneableParamsTuneableParamDbModel["wlm_cpu_limit"] = "'range(0, 100)'"
		successTuneableParamsTuneableParamDbModel["wlm_cpu_shares"] = "'range(1, 65535)'"
		successTuneableParamsTuneableParamDbModel["wlm_cpu_share_mode"] = "'HARD', 'SOFT'"

		successTuneableParamsTuneableParamDbmModel := make(map[string]interface{})
		successTuneableParamsTuneableParamDbmModel["comm_bandwidth"] = "'range(0.1, 100000)', '-1'"
		successTuneableParamsTuneableParamDbmModel["cpuspeed"] = "'range(0.0000000001, 1)', '-1'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_bufpool"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_lock"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_sort"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_stmt"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_table"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_timestamp"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["dft_mon_uow"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamDbmModel["diaglevel"] = "'range(0, 4)'"
		successTuneableParamsTuneableParamDbmModel["federated_async"] = "'range(0, 32767)', '-1', 'ANY'"
		successTuneableParamsTuneableParamDbmModel["indexrec"] = "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'"
		successTuneableParamsTuneableParamDbmModel["intra_parallel"] = "'SYSTEM', 'NO', 'YES'"
		successTuneableParamsTuneableParamDbmModel["keepfenced"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbmModel["max_connretries"] = "'range(0, 100)'"
		successTuneableParamsTuneableParamDbmModel["max_querydegree"] = "'range(1, 32767)', '-1', 'ANY'"
		successTuneableParamsTuneableParamDbmModel["mon_heap_sz"] = "'range(0, 2147483647)', 'AUTOMATIC'"
		successTuneableParamsTuneableParamDbmModel["multipartsizemb"] = "'range(5, 5120)'"
		successTuneableParamsTuneableParamDbmModel["notifylevel"] = "'range(0, 4)'"
		successTuneableParamsTuneableParamDbmModel["num_initagents"] = "'range(0, 64000)'"
		successTuneableParamsTuneableParamDbmModel["num_initfenced"] = "'range(0, 64000)'"
		successTuneableParamsTuneableParamDbmModel["num_poolagents"] = "'-1', 'range(0, 64000)'"
		successTuneableParamsTuneableParamDbmModel["resync_interval"] = "'range(1, 60000)'"
		successTuneableParamsTuneableParamDbmModel["rqrioblk"] = "'range(4096, 65535)'"
		successTuneableParamsTuneableParamDbmModel["start_stop_time"] = "'range(1, 1440)'"
		successTuneableParamsTuneableParamDbmModel["util_impact_lim"] = "'range(1, 100)'"
		successTuneableParamsTuneableParamDbmModel["wlm_dispatcher"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamDbmModel["wlm_disp_concur"] = "'range(1, 32767)', 'COMPUTED'"
		successTuneableParamsTuneableParamDbmModel["wlm_disp_cpu_shares"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamDbmModel["wlm_disp_min_util"] = "'range(0, 100)'"

		successTuneableParamsTuneableParamRegistryModel := make(map[string]interface{})
		successTuneableParamsTuneableParamRegistryModel["db2_bidi"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamRegistryModel["db2_compopt"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_lock_to_rb"] = "'STATEMENT'"
		successTuneableParamsTuneableParamRegistryModel["db2_stmm"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_alternate_authz_behaviour"] = "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'"
		successTuneableParamsTuneableParamRegistryModel["db2_antijoin"] = "'YES', 'NO', 'EXTEND'"
		successTuneableParamsTuneableParamRegistryModel["db2_ats_enable"] = "'YES', 'NO'"
		successTuneableParamsTuneableParamRegistryModel["db2_deferred_prepare_semantics"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_evaluncommitted"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_extended_optimization"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_index_pctfree_default"] = "'range(0, 99)'"
		successTuneableParamsTuneableParamRegistryModel["db2_inlist_to_nljn"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_minimize_listprefetch"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_object_table_entries"] = "'range(0, 65532)'"
		successTuneableParamsTuneableParamRegistryModel["db2_optprofile"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_optstats_log"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_opt_max_temp_size"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_parallel_io"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_reduced_optimization"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_selectivity"] = "'YES', 'NO', 'ALL'"
		successTuneableParamsTuneableParamRegistryModel["db2_skipdeleted"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_skipinserted"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_sync_release_lock_attributes"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_truncate_reusestorage"] = "'IMPORT', 'LOAD', 'TRUNCATE'"
		successTuneableParamsTuneableParamRegistryModel["db2_use_alternate_page_cleaning"] = "'ON', 'OFF'"
		successTuneableParamsTuneableParamRegistryModel["db2_view_reopt_values"] = "'NO', 'YES'"
		successTuneableParamsTuneableParamRegistryModel["db2_wlm_settings"] = "'-'"
		successTuneableParamsTuneableParamRegistryModel["db2_workload"] = "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"

		model := make(map[string]interface{})
		model["db"] = []map[string]interface{}{successTuneableParamsTuneableParamDbModel}
		model["dbm"] = []map[string]interface{}{successTuneableParamsTuneableParamDbmModel}
		model["registry"] = []map[string]interface{}{successTuneableParamsTuneableParamRegistryModel}

		assert.Equal(t, result, model)
	}

	successTuneableParamsTuneableParamDbModel := new(db2saasv1.SuccessTuneableParamsTuneableParamDb)
	successTuneableParamsTuneableParamDbModel.ACTSORTMEMLIMIT = core.StringPtr("'NONE', 'range(10, 100)'")
	successTuneableParamsTuneableParamDbModel.ALTCOLLATE = core.StringPtr("'NULL', 'IDENTITY_16BIT'")
	successTuneableParamsTuneableParamDbModel.APPGROUPMEMSZ = core.StringPtr("'range(1, 1000000)'")
	successTuneableParamsTuneableParamDbModel.APPLHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(16, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.APPLMEMORY = core.StringPtr("'AUTOMATIC', 'range(128, 4294967295)'")
	successTuneableParamsTuneableParamDbModel.APPCTLHEAPSZ = core.StringPtr("'range(1, 64000)'")
	successTuneableParamsTuneableParamDbModel.ARCHRETRYDELAY = core.StringPtr("'range(0, 65535)'")
	successTuneableParamsTuneableParamDbModel.AUTHNCACHEDURATION = core.StringPtr("'range(1,10000)'")
	successTuneableParamsTuneableParamDbModel.AUTORESTART = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOCGSTATS = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOMAINT = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOREORG = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOREVAL = core.StringPtr("'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'")
	successTuneableParamsTuneableParamDbModel.AUTORUNSTATS = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOSAMPLING = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOSTATSVIEWS = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOSTMTSTATS = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AUTOTBLMAINT = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.AVGAPPLS = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.CATALOGCACHESZ = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.CHNGPGSTHRESH = core.StringPtr("'range(5,99)'")
	successTuneableParamsTuneableParamDbModel.CURCOMMIT = core.StringPtr("'ON, AVAILABLE, DISABLED'")
	successTuneableParamsTuneableParamDbModel.DATABASEMEMORY = core.StringPtr("'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'")
	successTuneableParamsTuneableParamDbModel.DBHEAP = core.StringPtr("'AUTOMATIC', 'range(32, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.DBCOLLNAME = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.DBMEMTHRESH = core.StringPtr("'range(0, 100)'")
	successTuneableParamsTuneableParamDbModel.DDLCOMPRESSIONDEF = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.DDLCONSTRAINTDEF = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.DECFLTROUNDING = core.StringPtr("'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'")
	successTuneableParamsTuneableParamDbModel.DECARITHMETIC = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.DECTOCHARFMT = core.StringPtr("'NEW', 'V95'")
	successTuneableParamsTuneableParamDbModel.DFTDEGREE = core.StringPtr("'-1', 'ANY', 'range(1, 32767)'")
	successTuneableParamsTuneableParamDbModel.DFTEXTENTSZ = core.StringPtr("'range(2, 256)'")
	successTuneableParamsTuneableParamDbModel.DFTLOADRECSES = core.StringPtr("'range(1, 30000)'")
	successTuneableParamsTuneableParamDbModel.DFTMTTBTYPES = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.DFTPREFETCHSZ = core.StringPtr("'range(0, 32767)', 'AUTOMATIC'")
	successTuneableParamsTuneableParamDbModel.DFTQUERYOPT = core.StringPtr("'range(0, 9)'")
	successTuneableParamsTuneableParamDbModel.DFTREFRESHAGE = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.DFTSCHEMASDCC = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.DFTSQLMATHWARN = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.DFTTABLEORG = core.StringPtr("'COLUMN', 'ROW'")
	successTuneableParamsTuneableParamDbModel.DLCHKTIME = core.StringPtr("'range(1000, 600000)'")
	successTuneableParamsTuneableParamDbModel.ENABLEXMLCHAR = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.EXTENDEDROWSZ = core.StringPtr("'ENABLE', 'DISABLE'")
	successTuneableParamsTuneableParamDbModel.GROUPHEAPRATIO = core.StringPtr("'range(1, 99)'")
	successTuneableParamsTuneableParamDbModel.INDEXREC = core.StringPtr("'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'")
	successTuneableParamsTuneableParamDbModel.LARGEAGGREGATION = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.LOCKLIST = core.StringPtr("'AUTOMATIC', 'range(4, 134217728)'")
	successTuneableParamsTuneableParamDbModel.LOCKTIMEOUT = core.StringPtr("'-1', 'range(0, 32767)'")
	successTuneableParamsTuneableParamDbModel.LOGINDEXBUILD = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.LOGAPPLINFO = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.LOGDDLSTMTS = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.LOGDISKCAP = core.StringPtr("'0', '-1', 'range(1, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.MAXAPPLS = core.StringPtr("'range(1, 60000)'")
	successTuneableParamsTuneableParamDbModel.MAXFILOP = core.StringPtr("'range(64, 61440)'")
	successTuneableParamsTuneableParamDbModel.MAXLOCKS = core.StringPtr("'AUTOMATIC', 'range(1, 100)'")
	successTuneableParamsTuneableParamDbModel.MINDECDIV3 = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.MONACTMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	successTuneableParamsTuneableParamDbModel.MONDEADLOCK = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	successTuneableParamsTuneableParamDbModel.MONLCKMSGLVL = core.StringPtr("'range(0, 3)'")
	successTuneableParamsTuneableParamDbModel.MONLOCKTIMEOUT = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	successTuneableParamsTuneableParamDbModel.MONLOCKWAIT = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	successTuneableParamsTuneableParamDbModel.MONLWTHRESH = core.StringPtr("'range(1000, 4294967295)'")
	successTuneableParamsTuneableParamDbModel.MONOBJMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	successTuneableParamsTuneableParamDbModel.MONPKGLISTSZ = core.StringPtr("'range(0, 1024)'")
	successTuneableParamsTuneableParamDbModel.MONREQMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	successTuneableParamsTuneableParamDbModel.MONRTNDATA = core.StringPtr("'NONE', 'BASE'")
	successTuneableParamsTuneableParamDbModel.MONRTNEXECLIST = core.StringPtr("'OFF', 'ON'")
	successTuneableParamsTuneableParamDbModel.MONUOWDATA = core.StringPtr("'NONE', 'BASE'")
	successTuneableParamsTuneableParamDbModel.MONUOWEXECLIST = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.MONUOWPKGLIST = core.StringPtr("'OFF', 'ON'")
	successTuneableParamsTuneableParamDbModel.NCHARMAPPING = core.StringPtr("'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'")
	successTuneableParamsTuneableParamDbModel.NUMFREQVALUES = core.StringPtr("'range(0, 32767)'")
	successTuneableParamsTuneableParamDbModel.NUMIOCLEANERS = core.StringPtr("'AUTOMATIC', 'range(0, 255)'")
	successTuneableParamsTuneableParamDbModel.NUMIOSERVERS = core.StringPtr("'AUTOMATIC', 'range(1, 255)'")
	successTuneableParamsTuneableParamDbModel.NUMLOGSPAN = core.StringPtr("'range(0, 65535)'")
	successTuneableParamsTuneableParamDbModel.NUMQUANTILES = core.StringPtr("'range(0, 32767)'")
	successTuneableParamsTuneableParamDbModel.OPTBUFFPAGE = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.OPTDIRECTWRKLD = core.StringPtr("'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'")
	successTuneableParamsTuneableParamDbModel.OPTLOCKLIST = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.OPTMAXLOCKS = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.OPTSORTHEAP = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.PAGEAGETRGTGCR = core.StringPtr("'range(1, 65535)'")
	successTuneableParamsTuneableParamDbModel.PAGEAGETRGTMCR = core.StringPtr("'range(1, 65535)'")
	successTuneableParamsTuneableParamDbModel.PCKCACHESZ = core.StringPtr("'AUTOMATIC', '-1', 'range(32, 2147483646)'")
	successTuneableParamsTuneableParamDbModel.PLSTACKTRACE = core.StringPtr("'NONE', 'ALL', 'UNHANDLED'")
	successTuneableParamsTuneableParamDbModel.SELFTUNINGMEM = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbModel.SEQDETECT = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.SHEAPTHRESSHR = core.StringPtr("'AUTOMATIC', 'range(250, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.SOFTMAX = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.SORTHEAP = core.StringPtr("'AUTOMATIC', 'range(16, 4294967295)'")
	successTuneableParamsTuneableParamDbModel.SQLCCFLAGS = core.StringPtr("'-'")
	successTuneableParamsTuneableParamDbModel.STATHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(1096, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.STMTHEAP = core.StringPtr("'AUTOMATIC', 'range(128, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.STMTCONC = core.StringPtr("'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'")
	successTuneableParamsTuneableParamDbModel.STRINGUNITS = core.StringPtr("'SYSTEM', 'CODEUNITS32'")
	successTuneableParamsTuneableParamDbModel.SYSTIMEPERIODADJ = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamDbModel.TRACKMOD = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.UTILHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(16, 2147483647)'")
	successTuneableParamsTuneableParamDbModel.WLMADMISSIONCTRL = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbModel.WLMAGENTLOADTRGT = core.StringPtr("'AUTOMATIC', 'range(1, 65535)'")
	successTuneableParamsTuneableParamDbModel.WLMCPULIMIT = core.StringPtr("'range(0, 100)'")
	successTuneableParamsTuneableParamDbModel.WLMCPUSHARES = core.StringPtr("'range(1, 65535)'")
	successTuneableParamsTuneableParamDbModel.WLMCPUSHAREMODE = core.StringPtr("'HARD', 'SOFT'")

	successTuneableParamsTuneableParamDbmModel := new(db2saasv1.SuccessTuneableParamsTuneableParamDbm)
	successTuneableParamsTuneableParamDbmModel.COMMBANDWIDTH = core.StringPtr("'range(0.1, 100000)', '-1'")
	successTuneableParamsTuneableParamDbmModel.CPUSPEED = core.StringPtr("'range(0.0000000001, 1)', '-1'")
	successTuneableParamsTuneableParamDbmModel.DFTMONBUFPOOL = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONLOCK = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONSORT = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONSTMT = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONTABLE = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONTIMESTAMP = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DFTMONUOW = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamDbmModel.DIAGLEVEL = core.StringPtr("'range(0, 4)'")
	successTuneableParamsTuneableParamDbmModel.FEDERATEDASYNC = core.StringPtr("'range(0, 32767)', '-1', 'ANY'")
	successTuneableParamsTuneableParamDbmModel.INDEXREC = core.StringPtr("'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'")
	successTuneableParamsTuneableParamDbmModel.INTRAPARALLEL = core.StringPtr("'SYSTEM', 'NO', 'YES'")
	successTuneableParamsTuneableParamDbmModel.KEEPFENCED = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbmModel.MAXCONNRETRIES = core.StringPtr("'range(0, 100)'")
	successTuneableParamsTuneableParamDbmModel.MAXQUERYDEGREE = core.StringPtr("'range(1, 32767)', '-1', 'ANY'")
	successTuneableParamsTuneableParamDbmModel.MONHEAPSZ = core.StringPtr("'range(0, 2147483647)', 'AUTOMATIC'")
	successTuneableParamsTuneableParamDbmModel.MULTIPARTSIZEMB = core.StringPtr("'range(5, 5120)'")
	successTuneableParamsTuneableParamDbmModel.NOTIFYLEVEL = core.StringPtr("'range(0, 4)'")
	successTuneableParamsTuneableParamDbmModel.NUMINITAGENTS = core.StringPtr("'range(0, 64000)'")
	successTuneableParamsTuneableParamDbmModel.NUMINITFENCED = core.StringPtr("'range(0, 64000)'")
	successTuneableParamsTuneableParamDbmModel.NUMPOOLAGENTS = core.StringPtr("'-1', 'range(0, 64000)'")
	successTuneableParamsTuneableParamDbmModel.RESYNCINTERVAL = core.StringPtr("'range(1, 60000)'")
	successTuneableParamsTuneableParamDbmModel.RQRIOBLK = core.StringPtr("'range(4096, 65535)'")
	successTuneableParamsTuneableParamDbmModel.STARTSTOPTIME = core.StringPtr("'range(1, 1440)'")
	successTuneableParamsTuneableParamDbmModel.UTILIMPACTLIM = core.StringPtr("'range(1, 100)'")
	successTuneableParamsTuneableParamDbmModel.WLMDISPATCHER = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamDbmModel.WLMDISPCONCUR = core.StringPtr("'range(1, 32767)', 'COMPUTED'")
	successTuneableParamsTuneableParamDbmModel.WLMDISPCPUSHARES = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamDbmModel.WLMDISPMINUTIL = core.StringPtr("'range(0, 100)'")

	successTuneableParamsTuneableParamRegistryModel := new(db2saasv1.SuccessTuneableParamsTuneableParamRegistry)
	successTuneableParamsTuneableParamRegistryModel.DB2BIDI = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamRegistryModel.DB2COMPOPT = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2LOCKTORB = core.StringPtr("'STATEMENT'")
	successTuneableParamsTuneableParamRegistryModel.DB2STMM = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'")
	successTuneableParamsTuneableParamRegistryModel.DB2ANTIJOIN = core.StringPtr("'YES', 'NO', 'EXTEND'")
	successTuneableParamsTuneableParamRegistryModel.DB2ATSENABLE = core.StringPtr("'YES', 'NO'")
	successTuneableParamsTuneableParamRegistryModel.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2EVALUNCOMMITTED = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2EXTENDEDOPTIMIZATION = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2INDEXPCTFREEDEFAULT = core.StringPtr("'range(0, 99)'")
	successTuneableParamsTuneableParamRegistryModel.DB2INLISTTONLJN = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2MINIMIZELISTPREFETCH = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2OBJECTTABLEENTRIES = core.StringPtr("'range(0, 65532)'")
	successTuneableParamsTuneableParamRegistryModel.DB2OPTPROFILE = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2OPTSTATSLOG = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2OPTMAXTEMPSIZE = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2PARALLELIO = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2REDUCEDOPTIMIZATION = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2SELECTIVITY = core.StringPtr("'YES', 'NO', 'ALL'")
	successTuneableParamsTuneableParamRegistryModel.DB2SKIPDELETED = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2SKIPINSERTED = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2TRUNCATEREUSESTORAGE = core.StringPtr("'IMPORT', 'LOAD', 'TRUNCATE'")
	successTuneableParamsTuneableParamRegistryModel.DB2USEALTERNATEPAGECLEANING = core.StringPtr("'ON', 'OFF'")
	successTuneableParamsTuneableParamRegistryModel.DB2VIEWREOPTVALUES = core.StringPtr("'NO', 'YES'")
	successTuneableParamsTuneableParamRegistryModel.DB2WLMSETTINGS = core.StringPtr("'-'")
	successTuneableParamsTuneableParamRegistryModel.DB2WORKLOAD = core.StringPtr("'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'")

	model := new(db2saasv1.SuccessTuneableParamsTuneableParam)
	model.Db = successTuneableParamsTuneableParamDbModel
	model.Dbm = successTuneableParamsTuneableParamDbmModel
	model.Registry = successTuneableParamsTuneableParamRegistryModel

	result, err := db2.DataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamDbToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["act_sortmem_limit"] = "'NONE', 'range(10, 100)'"
		model["alt_collate"] = "'NULL', 'IDENTITY_16BIT'"
		model["appgroup_mem_sz"] = "'range(1, 1000000)'"
		model["applheapsz"] = "'AUTOMATIC', 'range(16, 2147483647)'"
		model["appl_memory"] = "'AUTOMATIC', 'range(128, 4294967295)'"
		model["app_ctl_heap_sz"] = "'range(1, 64000)'"
		model["archretrydelay"] = "'range(0, 65535)'"
		model["authn_cache_duration"] = "'range(1,10000)'"
		model["autorestart"] = "'ON', 'OFF'"
		model["auto_cg_stats"] = "'ON', 'OFF'"
		model["auto_maint"] = "'ON', 'OFF'"
		model["auto_reorg"] = "'ON', 'OFF'"
		model["auto_reval"] = "'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'"
		model["auto_runstats"] = "'ON', 'OFF'"
		model["auto_sampling"] = "'ON', 'OFF'"
		model["auto_stats_views"] = "'ON', 'OFF'"
		model["auto_stmt_stats"] = "'ON', 'OFF'"
		model["auto_tbl_maint"] = "'ON', 'OFF'"
		model["avg_appls"] = "'-'"
		model["catalogcache_sz"] = "'-'"
		model["chngpgs_thresh"] = "'range(5,99)'"
		model["cur_commit"] = "'ON, AVAILABLE, DISABLED'"
		model["database_memory"] = "'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'"
		model["dbheap"] = "'AUTOMATIC', 'range(32, 2147483647)'"
		model["db_collname"] = "'-'"
		model["db_mem_thresh"] = "'range(0, 100)'"
		model["ddl_compression_def"] = "'YES', 'NO'"
		model["ddl_constraint_def"] = "'YES', 'NO'"
		model["decflt_rounding"] = "'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'"
		model["dec_arithmetic"] = "'-'"
		model["dec_to_char_fmt"] = "'NEW', 'V95'"
		model["dft_degree"] = "'-1', 'ANY', 'range(1, 32767)'"
		model["dft_extent_sz"] = "'range(2, 256)'"
		model["dft_loadrec_ses"] = "'range(1, 30000)'"
		model["dft_mttb_types"] = "'-'"
		model["dft_prefetch_sz"] = "'range(0, 32767)', 'AUTOMATIC'"
		model["dft_queryopt"] = "'range(0, 9)'"
		model["dft_refresh_age"] = "'-'"
		model["dft_schemas_dcc"] = "'YES', 'NO'"
		model["dft_sqlmathwarn"] = "'YES', 'NO'"
		model["dft_table_org"] = "'COLUMN', 'ROW'"
		model["dlchktime"] = "'range(1000, 600000)'"
		model["enable_xmlchar"] = "'YES', 'NO'"
		model["extended_row_sz"] = "'ENABLE', 'DISABLE'"
		model["groupheap_ratio"] = "'range(1, 99)'"
		model["indexrec"] = "'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'"
		model["large_aggregation"] = "'YES', 'NO'"
		model["locklist"] = "'AUTOMATIC', 'range(4, 134217728)'"
		model["locktimeout"] = "'-1', 'range(0, 32767)'"
		model["logindexbuild"] = "'ON', 'OFF'"
		model["log_appl_info"] = "'YES', 'NO'"
		model["log_ddl_stmts"] = "'YES', 'NO'"
		model["log_disk_cap"] = "'0', '-1', 'range(1, 2147483647)'"
		model["maxappls"] = "'range(1, 60000)'"
		model["maxfilop"] = "'range(64, 61440)'"
		model["maxlocks"] = "'AUTOMATIC', 'range(1, 100)'"
		model["min_dec_div_3"] = "'YES', 'NO'"
		model["mon_act_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		model["mon_deadlock"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		model["mon_lck_msg_lvl"] = "'range(0, 3)'"
		model["mon_locktimeout"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		model["mon_lockwait"] = "'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'"
		model["mon_lw_thresh"] = "'range(1000, 4294967295)'"
		model["mon_obj_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		model["mon_pkglist_sz"] = "'range(0, 1024)'"
		model["mon_req_metrics"] = "'NONE', 'BASE', 'EXTENDED'"
		model["mon_rtn_data"] = "'NONE', 'BASE'"
		model["mon_rtn_execlist"] = "'OFF', 'ON'"
		model["mon_uow_data"] = "'NONE', 'BASE'"
		model["mon_uow_execlist"] = "'ON', 'OFF'"
		model["mon_uow_pkglist"] = "'OFF', 'ON'"
		model["nchar_mapping"] = "'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'"
		model["num_freqvalues"] = "'range(0, 32767)'"
		model["num_iocleaners"] = "'AUTOMATIC', 'range(0, 255)'"
		model["num_ioservers"] = "'AUTOMATIC', 'range(1, 255)'"
		model["num_log_span"] = "'range(0, 65535)'"
		model["num_quantiles"] = "'range(0, 32767)'"
		model["opt_buffpage"] = "'-'"
		model["opt_direct_wrkld"] = "'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'"
		model["opt_locklist"] = "'-'"
		model["opt_maxlocks"] = "'-'"
		model["opt_sortheap"] = "'-'"
		model["page_age_trgt_gcr"] = "'range(1, 65535)'"
		model["page_age_trgt_mcr"] = "'range(1, 65535)'"
		model["pckcachesz"] = "'AUTOMATIC', '-1', 'range(32, 2147483646)'"
		model["pl_stack_trace"] = "'NONE', 'ALL', 'UNHANDLED'"
		model["self_tuning_mem"] = "'ON', 'OFF'"
		model["seqdetect"] = "'YES', 'NO'"
		model["sheapthres_shr"] = "'AUTOMATIC', 'range(250, 2147483647)'"
		model["softmax"] = "'-'"
		model["sortheap"] = "'AUTOMATIC', 'range(16, 4294967295)'"
		model["sql_ccflags"] = "'-'"
		model["stat_heap_sz"] = "'AUTOMATIC', 'range(1096, 2147483647)'"
		model["stmtheap"] = "'AUTOMATIC', 'range(128, 2147483647)'"
		model["stmt_conc"] = "'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'"
		model["string_units"] = "'SYSTEM', 'CODEUNITS32'"
		model["systime_period_adj"] = "'NO', 'YES'"
		model["trackmod"] = "'YES', 'NO'"
		model["util_heap_sz"] = "'AUTOMATIC', 'range(16, 2147483647)'"
		model["wlm_admission_ctrl"] = "'YES', 'NO'"
		model["wlm_agent_load_trgt"] = "'AUTOMATIC', 'range(1, 65535)'"
		model["wlm_cpu_limit"] = "'range(0, 100)'"
		model["wlm_cpu_shares"] = "'range(1, 65535)'"
		model["wlm_cpu_share_mode"] = "'HARD', 'SOFT'"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessTuneableParamsTuneableParamDb)
	model.ACTSORTMEMLIMIT = core.StringPtr("'NONE', 'range(10, 100)'")
	model.ALTCOLLATE = core.StringPtr("'NULL', 'IDENTITY_16BIT'")
	model.APPGROUPMEMSZ = core.StringPtr("'range(1, 1000000)'")
	model.APPLHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(16, 2147483647)'")
	model.APPLMEMORY = core.StringPtr("'AUTOMATIC', 'range(128, 4294967295)'")
	model.APPCTLHEAPSZ = core.StringPtr("'range(1, 64000)'")
	model.ARCHRETRYDELAY = core.StringPtr("'range(0, 65535)'")
	model.AUTHNCACHEDURATION = core.StringPtr("'range(1,10000)'")
	model.AUTORESTART = core.StringPtr("'ON', 'OFF'")
	model.AUTOCGSTATS = core.StringPtr("'ON', 'OFF'")
	model.AUTOMAINT = core.StringPtr("'ON', 'OFF'")
	model.AUTOREORG = core.StringPtr("'ON', 'OFF'")
	model.AUTOREVAL = core.StringPtr("'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'")
	model.AUTORUNSTATS = core.StringPtr("'ON', 'OFF'")
	model.AUTOSAMPLING = core.StringPtr("'ON', 'OFF'")
	model.AUTOSTATSVIEWS = core.StringPtr("'ON', 'OFF'")
	model.AUTOSTMTSTATS = core.StringPtr("'ON', 'OFF'")
	model.AUTOTBLMAINT = core.StringPtr("'ON', 'OFF'")
	model.AVGAPPLS = core.StringPtr("'-'")
	model.CATALOGCACHESZ = core.StringPtr("'-'")
	model.CHNGPGSTHRESH = core.StringPtr("'range(5,99)'")
	model.CURCOMMIT = core.StringPtr("'ON, AVAILABLE, DISABLED'")
	model.DATABASEMEMORY = core.StringPtr("'AUTOMATIC', 'COMPUTED', 'range(0, 4294967295)'")
	model.DBHEAP = core.StringPtr("'AUTOMATIC', 'range(32, 2147483647)'")
	model.DBCOLLNAME = core.StringPtr("'-'")
	model.DBMEMTHRESH = core.StringPtr("'range(0, 100)'")
	model.DDLCOMPRESSIONDEF = core.StringPtr("'YES', 'NO'")
	model.DDLCONSTRAINTDEF = core.StringPtr("'YES', 'NO'")
	model.DECFLTROUNDING = core.StringPtr("'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'")
	model.DECARITHMETIC = core.StringPtr("'-'")
	model.DECTOCHARFMT = core.StringPtr("'NEW', 'V95'")
	model.DFTDEGREE = core.StringPtr("'-1', 'ANY', 'range(1, 32767)'")
	model.DFTEXTENTSZ = core.StringPtr("'range(2, 256)'")
	model.DFTLOADRECSES = core.StringPtr("'range(1, 30000)'")
	model.DFTMTTBTYPES = core.StringPtr("'-'")
	model.DFTPREFETCHSZ = core.StringPtr("'range(0, 32767)', 'AUTOMATIC'")
	model.DFTQUERYOPT = core.StringPtr("'range(0, 9)'")
	model.DFTREFRESHAGE = core.StringPtr("'-'")
	model.DFTSCHEMASDCC = core.StringPtr("'YES', 'NO'")
	model.DFTSQLMATHWARN = core.StringPtr("'YES', 'NO'")
	model.DFTTABLEORG = core.StringPtr("'COLUMN', 'ROW'")
	model.DLCHKTIME = core.StringPtr("'range(1000, 600000)'")
	model.ENABLEXMLCHAR = core.StringPtr("'YES', 'NO'")
	model.EXTENDEDROWSZ = core.StringPtr("'ENABLE', 'DISABLE'")
	model.GROUPHEAPRATIO = core.StringPtr("'range(1, 99)'")
	model.INDEXREC = core.StringPtr("'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'")
	model.LARGEAGGREGATION = core.StringPtr("'YES', 'NO'")
	model.LOCKLIST = core.StringPtr("'AUTOMATIC', 'range(4, 134217728)'")
	model.LOCKTIMEOUT = core.StringPtr("'-1', 'range(0, 32767)'")
	model.LOGINDEXBUILD = core.StringPtr("'ON', 'OFF'")
	model.LOGAPPLINFO = core.StringPtr("'YES', 'NO'")
	model.LOGDDLSTMTS = core.StringPtr("'YES', 'NO'")
	model.LOGDISKCAP = core.StringPtr("'0', '-1', 'range(1, 2147483647)'")
	model.MAXAPPLS = core.StringPtr("'range(1, 60000)'")
	model.MAXFILOP = core.StringPtr("'range(64, 61440)'")
	model.MAXLOCKS = core.StringPtr("'AUTOMATIC', 'range(1, 100)'")
	model.MINDECDIV3 = core.StringPtr("'YES', 'NO'")
	model.MONACTMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	model.MONDEADLOCK = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	model.MONLCKMSGLVL = core.StringPtr("'range(0, 3)'")
	model.MONLOCKTIMEOUT = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	model.MONLOCKWAIT = core.StringPtr("'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'")
	model.MONLWTHRESH = core.StringPtr("'range(1000, 4294967295)'")
	model.MONOBJMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	model.MONPKGLISTSZ = core.StringPtr("'range(0, 1024)'")
	model.MONREQMETRICS = core.StringPtr("'NONE', 'BASE', 'EXTENDED'")
	model.MONRTNDATA = core.StringPtr("'NONE', 'BASE'")
	model.MONRTNEXECLIST = core.StringPtr("'OFF', 'ON'")
	model.MONUOWDATA = core.StringPtr("'NONE', 'BASE'")
	model.MONUOWEXECLIST = core.StringPtr("'ON', 'OFF'")
	model.MONUOWPKGLIST = core.StringPtr("'OFF', 'ON'")
	model.NCHARMAPPING = core.StringPtr("'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'")
	model.NUMFREQVALUES = core.StringPtr("'range(0, 32767)'")
	model.NUMIOCLEANERS = core.StringPtr("'AUTOMATIC', 'range(0, 255)'")
	model.NUMIOSERVERS = core.StringPtr("'AUTOMATIC', 'range(1, 255)'")
	model.NUMLOGSPAN = core.StringPtr("'range(0, 65535)'")
	model.NUMQUANTILES = core.StringPtr("'range(0, 32767)'")
	model.OPTBUFFPAGE = core.StringPtr("'-'")
	model.OPTDIRECTWRKLD = core.StringPtr("'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'")
	model.OPTLOCKLIST = core.StringPtr("'-'")
	model.OPTMAXLOCKS = core.StringPtr("'-'")
	model.OPTSORTHEAP = core.StringPtr("'-'")
	model.PAGEAGETRGTGCR = core.StringPtr("'range(1, 65535)'")
	model.PAGEAGETRGTMCR = core.StringPtr("'range(1, 65535)'")
	model.PCKCACHESZ = core.StringPtr("'AUTOMATIC', '-1', 'range(32, 2147483646)'")
	model.PLSTACKTRACE = core.StringPtr("'NONE', 'ALL', 'UNHANDLED'")
	model.SELFTUNINGMEM = core.StringPtr("'ON', 'OFF'")
	model.SEQDETECT = core.StringPtr("'YES', 'NO'")
	model.SHEAPTHRESSHR = core.StringPtr("'AUTOMATIC', 'range(250, 2147483647)'")
	model.SOFTMAX = core.StringPtr("'-'")
	model.SORTHEAP = core.StringPtr("'AUTOMATIC', 'range(16, 4294967295)'")
	model.SQLCCFLAGS = core.StringPtr("'-'")
	model.STATHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(1096, 2147483647)'")
	model.STMTHEAP = core.StringPtr("'AUTOMATIC', 'range(128, 2147483647)'")
	model.STMTCONC = core.StringPtr("'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'")
	model.STRINGUNITS = core.StringPtr("'SYSTEM', 'CODEUNITS32'")
	model.SYSTIMEPERIODADJ = core.StringPtr("'NO', 'YES'")
	model.TRACKMOD = core.StringPtr("'YES', 'NO'")
	model.UTILHEAPSZ = core.StringPtr("'AUTOMATIC', 'range(16, 2147483647)'")
	model.WLMADMISSIONCTRL = core.StringPtr("'YES', 'NO'")
	model.WLMAGENTLOADTRGT = core.StringPtr("'AUTOMATIC', 'range(1, 65535)'")
	model.WLMCPULIMIT = core.StringPtr("'range(0, 100)'")
	model.WLMCPUSHARES = core.StringPtr("'range(1, 65535)'")
	model.WLMCPUSHAREMODE = core.StringPtr("'HARD', 'SOFT'")

	result, err := db2.DataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamDbToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamDbmToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["comm_bandwidth"] = "'range(0.1, 100000)', '-1'"
		model["cpuspeed"] = "'range(0.0000000001, 1)', '-1'"
		model["dft_mon_bufpool"] = "'ON', 'OFF'"
		model["dft_mon_lock"] = "'ON', 'OFF'"
		model["dft_mon_sort"] = "'ON', 'OFF'"
		model["dft_mon_stmt"] = "'ON', 'OFF'"
		model["dft_mon_table"] = "'ON', 'OFF'"
		model["dft_mon_timestamp"] = "'ON', 'OFF'"
		model["dft_mon_uow"] = "'ON', 'OFF'"
		model["diaglevel"] = "'range(0, 4)'"
		model["federated_async"] = "'range(0, 32767)', '-1', 'ANY'"
		model["indexrec"] = "'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'"
		model["intra_parallel"] = "'SYSTEM', 'NO', 'YES'"
		model["keepfenced"] = "'YES', 'NO'"
		model["max_connretries"] = "'range(0, 100)'"
		model["max_querydegree"] = "'range(1, 32767)', '-1', 'ANY'"
		model["mon_heap_sz"] = "'range(0, 2147483647)', 'AUTOMATIC'"
		model["multipartsizemb"] = "'range(5, 5120)'"
		model["notifylevel"] = "'range(0, 4)'"
		model["num_initagents"] = "'range(0, 64000)'"
		model["num_initfenced"] = "'range(0, 64000)'"
		model["num_poolagents"] = "'-1', 'range(0, 64000)'"
		model["resync_interval"] = "'range(1, 60000)'"
		model["rqrioblk"] = "'range(4096, 65535)'"
		model["start_stop_time"] = "'range(1, 1440)'"
		model["util_impact_lim"] = "'range(1, 100)'"
		model["wlm_dispatcher"] = "'YES', 'NO'"
		model["wlm_disp_concur"] = "'range(1, 32767)', 'COMPUTED'"
		model["wlm_disp_cpu_shares"] = "'NO', 'YES'"
		model["wlm_disp_min_util"] = "'range(0, 100)'"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessTuneableParamsTuneableParamDbm)
	model.COMMBANDWIDTH = core.StringPtr("'range(0.1, 100000)', '-1'")
	model.CPUSPEED = core.StringPtr("'range(0.0000000001, 1)', '-1'")
	model.DFTMONBUFPOOL = core.StringPtr("'ON', 'OFF'")
	model.DFTMONLOCK = core.StringPtr("'ON', 'OFF'")
	model.DFTMONSORT = core.StringPtr("'ON', 'OFF'")
	model.DFTMONSTMT = core.StringPtr("'ON', 'OFF'")
	model.DFTMONTABLE = core.StringPtr("'ON', 'OFF'")
	model.DFTMONTIMESTAMP = core.StringPtr("'ON', 'OFF'")
	model.DFTMONUOW = core.StringPtr("'ON', 'OFF'")
	model.DIAGLEVEL = core.StringPtr("'range(0, 4)'")
	model.FEDERATEDASYNC = core.StringPtr("'range(0, 32767)', '-1', 'ANY'")
	model.INDEXREC = core.StringPtr("'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'")
	model.INTRAPARALLEL = core.StringPtr("'SYSTEM', 'NO', 'YES'")
	model.KEEPFENCED = core.StringPtr("'YES', 'NO'")
	model.MAXCONNRETRIES = core.StringPtr("'range(0, 100)'")
	model.MAXQUERYDEGREE = core.StringPtr("'range(1, 32767)', '-1', 'ANY'")
	model.MONHEAPSZ = core.StringPtr("'range(0, 2147483647)', 'AUTOMATIC'")
	model.MULTIPARTSIZEMB = core.StringPtr("'range(5, 5120)'")
	model.NOTIFYLEVEL = core.StringPtr("'range(0, 4)'")
	model.NUMINITAGENTS = core.StringPtr("'range(0, 64000)'")
	model.NUMINITFENCED = core.StringPtr("'range(0, 64000)'")
	model.NUMPOOLAGENTS = core.StringPtr("'-1', 'range(0, 64000)'")
	model.RESYNCINTERVAL = core.StringPtr("'range(1, 60000)'")
	model.RQRIOBLK = core.StringPtr("'range(4096, 65535)'")
	model.STARTSTOPTIME = core.StringPtr("'range(1, 1440)'")
	model.UTILIMPACTLIM = core.StringPtr("'range(1, 100)'")
	model.WLMDISPATCHER = core.StringPtr("'YES', 'NO'")
	model.WLMDISPCONCUR = core.StringPtr("'range(1, 32767)', 'COMPUTED'")
	model.WLMDISPCPUSHARES = core.StringPtr("'NO', 'YES'")
	model.WLMDISPMINUTIL = core.StringPtr("'range(0, 100)'")

	result, err := db2.DataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamDbmToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamRegistryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["db2_bidi"] = "'YES', 'NO'"
		model["db2_compopt"] = "'-'"
		model["db2_lock_to_rb"] = "'STATEMENT'"
		model["db2_stmm"] = "'NO', 'YES'"
		model["db2_alternate_authz_behaviour"] = "'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'"
		model["db2_antijoin"] = "'YES', 'NO', 'EXTEND'"
		model["db2_ats_enable"] = "'YES', 'NO'"
		model["db2_deferred_prepare_semantics"] = "'NO', 'YES'"
		model["db2_evaluncommitted"] = "'NO', 'YES'"
		model["db2_extended_optimization"] = "'-'"
		model["db2_index_pctfree_default"] = "'range(0, 99)'"
		model["db2_inlist_to_nljn"] = "'NO', 'YES'"
		model["db2_minimize_listprefetch"] = "'NO', 'YES'"
		model["db2_object_table_entries"] = "'range(0, 65532)'"
		model["db2_optprofile"] = "'NO', 'YES'"
		model["db2_optstats_log"] = "'-'"
		model["db2_opt_max_temp_size"] = "'-'"
		model["db2_parallel_io"] = "'-'"
		model["db2_reduced_optimization"] = "'-'"
		model["db2_selectivity"] = "'YES', 'NO', 'ALL'"
		model["db2_skipdeleted"] = "'NO', 'YES'"
		model["db2_skipinserted"] = "'NO', 'YES'"
		model["db2_sync_release_lock_attributes"] = "'NO', 'YES'"
		model["db2_truncate_reusestorage"] = "'IMPORT', 'LOAD', 'TRUNCATE'"
		model["db2_use_alternate_page_cleaning"] = "'ON', 'OFF'"
		model["db2_view_reopt_values"] = "'NO', 'YES'"
		model["db2_wlm_settings"] = "'-'"
		model["db2_workload"] = "'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'"

		assert.Equal(t, result, model)
	}

	model := new(db2saasv1.SuccessTuneableParamsTuneableParamRegistry)
	model.DB2BIDI = core.StringPtr("'YES', 'NO'")
	model.DB2COMPOPT = core.StringPtr("'-'")
	model.DB2LOCKTORB = core.StringPtr("'STATEMENT'")
	model.DB2STMM = core.StringPtr("'NO', 'YES'")
	model.DB2ALTERNATEAUTHZBEHAVIOUR = core.StringPtr("'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'")
	model.DB2ANTIJOIN = core.StringPtr("'YES', 'NO', 'EXTEND'")
	model.DB2ATSENABLE = core.StringPtr("'YES', 'NO'")
	model.DB2DEFERREDPREPARESEMANTICS = core.StringPtr("'NO', 'YES'")
	model.DB2EVALUNCOMMITTED = core.StringPtr("'NO', 'YES'")
	model.DB2EXTENDEDOPTIMIZATION = core.StringPtr("'-'")
	model.DB2INDEXPCTFREEDEFAULT = core.StringPtr("'range(0, 99)'")
	model.DB2INLISTTONLJN = core.StringPtr("'NO', 'YES'")
	model.DB2MINIMIZELISTPREFETCH = core.StringPtr("'NO', 'YES'")
	model.DB2OBJECTTABLEENTRIES = core.StringPtr("'range(0, 65532)'")
	model.DB2OPTPROFILE = core.StringPtr("'NO', 'YES'")
	model.DB2OPTSTATSLOG = core.StringPtr("'-'")
	model.DB2OPTMAXTEMPSIZE = core.StringPtr("'-'")
	model.DB2PARALLELIO = core.StringPtr("'-'")
	model.DB2REDUCEDOPTIMIZATION = core.StringPtr("'-'")
	model.DB2SELECTIVITY = core.StringPtr("'YES', 'NO', 'ALL'")
	model.DB2SKIPDELETED = core.StringPtr("'NO', 'YES'")
	model.DB2SKIPINSERTED = core.StringPtr("'NO', 'YES'")
	model.DB2SYNCRELEASELOCKATTRIBUTES = core.StringPtr("'NO', 'YES'")
	model.DB2TRUNCATEREUSESTORAGE = core.StringPtr("'IMPORT', 'LOAD', 'TRUNCATE'")
	model.DB2USEALTERNATEPAGECLEANING = core.StringPtr("'ON', 'OFF'")
	model.DB2VIEWREOPTVALUES = core.StringPtr("'NO', 'YES'")
	model.DB2WLMSETTINGS = core.StringPtr("'-'")
	model.DB2WORKLOAD = core.StringPtr("'1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'")

	result, err := db2.DataSourceIbmDb2TuneableParamSuccessTuneableParamsTuneableParamRegistryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

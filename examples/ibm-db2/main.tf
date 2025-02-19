data "ibm_resource_group" "group" {
  name = var.resource_group
}

//Db2 SaaS Instance Creation
resource "ibm_db2" "db2_instance" {
  name                         = "demo-db2-v8"
  service                      = "dashdb-for-transactions"
  plan                         = "performance"
  location                     = var.region
  resource_group_id            = data.ibm_resource_group.group.id
  service_endpoints            = "public-and-private"
  instance_type                = "bx2.4x16"
  high_availability            = "yes"
  backup_location              = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn      = "none"
  oracle_compatibility         = "no"

  autoscale_config {
    auto_scaling_enabled          = "true"
    auto_scaling_threshold        = "60"
    auto_scaling_over_time_period = "15"
    auto_scaling_pause_limit      = "70"
    auto_scaling_allow_plan_limit = "true"
  }

  // custom_setting_config {
  // 	db {
  // 	  act_sortmem_limit = "NONE"
  // 	  alt_collate = "IDENTITY_16BIT"
  // 	  appgroup_mem_sz = "1000"
  // 	  applheapsz = "AUTOMATIC"
  // 	  appl_memory = "AUTOMATIC"
  // 	  app_ctl_heap_sz = "500"
  // 	  archretrydelay = "60"
  // 	  authn_cache_duration = "300"
  // 	  autorestart = "ON"
  // 	  auto_cg_stats = "OFF"
  // 	  auto_maint = "ON"
  // 	  auto_reorg = "OFF"
  // 	  auto_reval = "IMMEDIATE"
  // 	  auto_runstats = "ON"
  // 	  auto_sampling = "OFF"
  // 	  auto_stats_views = "ON"
  // 	  auto_stmt_stats = "OFF"
  // 	  auto_tbl_maint = "ON"
  // 	  chngpgs_thresh = "50"
  // 	  cur_commit = "AVAILABLE"
  // 	  database_memory = "AUTOMATIC"
  // 	  dbheap = "AUTOMATIC"
  // 	  db_mem_thresh = "80"
  // 	  ddl_compression_def = "YES"
  // 	  ddl_constraint_def = "NO"
  // 	  decflt_rounding = "ROUND_HALF_UP"
  // 	  dec_to_char_fmt = "NEW"
  // 	  dft_degree = "ANY"
  // 	  dft_extent_sz = "32"
  // 	  dft_loadrec_ses = "1000"
  // 	  dft_prefetch_sz = "AUTOMATIC"
  // 	  dft_queryopt = "5"
  // 	  dft_schemas_dcc = "NO"
  // 	  dft_sqlmathwarn = "YES"
  // 	  dft_table_org = "ROW"
  // 	  dlchktime = "30000"
  // 	  enable_xmlchar = "NO"
  // 	  extended_row_sz = "ENABLE"
  // 	  groupheap_ratio = "75"
  // 	  indexrec = "ACCESS"
  // 	  large_aggregation = "NO"
  // 	  locklist = "AUTOMATIC"
  // 	  locktimeout = "10"
  // 	  logindexbuild = "OFF"
  // 	  log_appl_info = "YES"
  // 	  log_ddl_stmts = "NO"
  // 	  log_disk_cap = "1000000"
  // 	  maxappls = "500"
  // 	  maxfilop = "1024"
  // 	  maxlocks = "AUTOMATIC"
  // 	  min_dec_div_3 = "NO"
  // 	  mon_act_metrics = "BASE"
  // 	  mon_deadlock = "HISTORY"
  // 	  mon_lck_msg_lvl = "2"
  // 	  mon_locktimeout = "HISTORY"
  // 	  mon_lockwait = "WITHOUT_HIST"
  // 	  mon_lw_thresh = "5000"
  // 	  mon_obj_metrics = "EXTENDED"
  // 	  mon_pkglist_sz = "512"
  // 	  mon_req_metrics = "BASE"
  // 	  mon_rtn_data = "NONE"
  // 	  mon_rtn_execlist = "OFF"
  // 	  mon_uow_data = "BASE"
  // 	  mon_uow_execlist = "ON"
  // 	  mon_uow_pkglist = "OFF"
  // 	  nchar_mapping = "GRAPHIC_CU32"
  // 	  num_freqvalues = "1000"
  // 	  num_iocleaners = "AUTOMATIC"
  // 	  num_ioservers = "AUTOMATIC"
  // 	  num_log_span = "10000"
  // 	  num_quantiles = "1000"
  // 	  opt_direct_wrkld = "YES"
  // 	  page_age_trgt_gcr = "5000"
  // 	  page_age_trgt_mcr = "5000"
  // 	  pckcachesz = "AUTOMATIC"
  // 	  pl_stack_trace = "ALL"
  // 	  self_tuning_mem = "ON"
  // 	  seqdetect = "YES"
  // 	  sheapthres_shr = "AUTOMATIC"
  // 	  sortheap = "AUTOMATIC"
  // 	  stat_heap_sz = "AUTOMATIC"
  // 	  stmtheap = "AUTOMATIC"
  // 	  stmt_conc = "LITERALS"
  // 	  string_units = "SYSTEM"
  // 	  systime_period_adj = "NO"
  // 	  trackmod = "YES"
  // 	  util_heap_sz = "AUTOMATIC"
  // 	  wlm_admission_ctrl = "YES"
  // 	  wlm_agent_load_trgt = "AUTOMATIC"
  // 	  wlm_cpu_limit = "50"
  // 	  wlm_cpu_shares = "1000"
  // 	  wlm_cpu_share_mode = "SOFT"
  // 	}
  // 	dbm {
  // 	  comm_bandwidth = "1000"
  // 	  cpuspeed = "-1"
  // 	  dft_mon_bufpool = "ON"
  // 	  dft_mon_lock = "OFF"
  // 	  dft_mon_sort = "ON"
  // 	  dft_mon_stmt = "OFF"
  // 	  dft_mon_table = "ON"
  // 	  dft_mon_timestamp = "OFF"
  // 	  dft_mon_uow = "ON"
  // 	  diaglevel = "2"
  // 	  federated_async = "ANY"
  // 	  indexrec = "ACCESS"
  // 	  intra_parallel = "YES"
  // 	  keepfenced = "NO"
  // 	  max_connretries = "10"
  // 	  max_querydegree = "ANY"
  // 	  mon_heap_sz = "AUTOMATIC"
  // 	  multipartsizemb = "100"
  // 	  notifylevel = "3"
  // 	  num_initagents = "500"
  // 	  num_initfenced = "1000"
  // 	  num_poolagents = "2000"
  // 	  resync_interval = "300"
  // 	  rqrioblk = "8192"
  // 	  start_stop_time = "60"
  // 	  util_impact_lim = "50"
  // 	  wlm_dispatcher = "YES"
  // 	  wlm_disp_concur = "COMPUTED"
  // 	  wlm_disp_cpu_shares = "YES"
  // 	  wlm_disp_min_util = "75"
  // 	}
  // 	registry {
  // 	  db2_bidi = "YES"
  // 	  db2_lock_to_rb = "STATEMENT"
  // 	  db2_stmm = "YES"
  // 	  db2_alternate_authz_behaviour = "EXTERNAL_ROUTINE_DBADM"
  // 	  db2_antijoin = "EXTEND"
  // 	  db2_ats_enable = "YES"
  // 	  db2_deferred_prepare_semantics = "NO"
  // 	  db2_evaluncommitted = "YES"
  // 	  db2_index_pctfree_default = "10"
  // 	  db2_inlist_to_nljn = "YES"
  // 	  db2_minimize_listprefetch = "NO"
  // 	  db2_object_table_entries = "5000"
  // 	  db2_optprofile = "NO"
  // 	  db2_selectivity = "ALL"
  // 	  db2_skipdeleted = "YES"
  // 	  db2_skipinserted = "NO"
  // 	  db2_sync_release_lock_attributes = "YES"
  // 	  db2_truncate_reusestorage = "IMPORT"
  // 	  db2_use_alternate_page_cleaning = "ON"
  // 	  db2_view_reopt_values = "NO"
  // 	  db2_workload = "SAP"
  // 	}
  //  }

  timeouts {
    create = "720m"
    update = "60m"
    delete = "30m"
  }
}

//Db2 SaaS Connection Info
# data "ibm_db2_connection_info" "db2_connection_info" {
#     deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-east%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3Af9455c22-07af-4a86-b9df-f02fd4774471%3A%3A"
#     x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
# }
#
//Db2 SaaS Autoscale
# data "ibm_db2_autoscale" "Db2-44-test-both" {
#    deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3Af5c9359c-9a66-4087-9eda-0024c2603c92%3A%3A"
# }

//Db2 SaaS Allowed list of IPs
# data "ibm_db2_allowlist_ip" "db2_allowlistips" {
#     x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
# }

//DataSource reading existing Db2 SaaS instance
# data "ibm_db2" "db2_instance" {
#   name              = "dDb2-v0-test-public"
#   resource_group_id = data.ibm_resource_group.group.id
#   location          = var.region
#   service           = "dashdb-for-transactions"
# }

//DataSource reading tuneable params of Db2 instance
# data "ibm_db2_tuneable_param" "Db2-kj-test-pub" {
# }

//DataSource reading backups of Db2 instance
# data "ibm_db2_backup" "Db2-kj-test-pub" {
#  deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A5d673016-3dbf-428c-8e59-e6ab82028b53%3A%3A"
# }
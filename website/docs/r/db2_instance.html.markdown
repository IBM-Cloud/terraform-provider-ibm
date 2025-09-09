---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2"
description: |-
  Manages IBM Db2 SaaS instance.
---

# ibm_db2

Create or update or delete an IBM Db2 SaaS on IBM Cloud instance. The `ibmcloud_api_key` that are used by Terraform should grant IAM rights to create and modify IBM Cloud Db2 Databases and have access to the resource group the Db2 SaaS instance is associated with. For more information, see [documentation](https://cloud.ibm.com/docs/Db2onCloud?topic=Db2onCloud-getting-started) to manage Db2 SaaS instances.


Configuration of an Db2 SaaS resource requires that the `region` parameter is set for the IBM provider in the `provider.tf` to be the same as the target Db2 SaaS `location/region`. If the Terraform configuration needs to deploy resources into multiple regions, provider alias can be used. For more information, see [Terraform provider configuration](https://www.terraform.io/docs/configuration/providers.html#multiple-provider-instances).


## Example usage
To find an example for provisioning and configuring a Db2 SaaS instance , see [here](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-db2).

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }
}

```

Sample Db2 SaaS instance using `users_config` attribute

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }

  users_config  {
    email = "test_user@mycompany.com"
    iam = false
    ibmid = "test-ibm-id"
    locked = "no"
    name = "test_user"
    password = "dEkMc43@gfAPl!867^dSbu"
    role = "bluuser"
    x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
    authentication {
      method = "method"
      policy_id = "policy_id"
    }
  }
}
```

Sample Db2 SaaS instance using `allowlist_config` attribute

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }

  allowlist_config  {
    ip_addresses {
      address     = "127.0.0.1"
      description = "A sample IP address 1"
    }
  }
}
```

Sample Db2 SaaS instance using `autoscale_config` attribute

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }

  autoscale_config  {
    auto_scaling_enabled = "true"
    auto_scaling_threshold = "60"
    auto_scaling_over_time_period = "15"
    auto_scaling_pause_limit = "70"
    auto_scaling_allow_plan_limit = "true"
  }
}
```

Sample Db2 SaaS instance using `custom_setting_config` attribute

```terraform
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"
  subscription_id = "<id_of_subscription_plan>"

  timeouts {
    create = "720m"
    update = "30m"
    delete = "30m"
  }

  custom_setting_config {
      db {
        act_sortmem_limit = "NONE"
        alt_collate = "IDENTITY_16BIT"
        appgroup_mem_sz = "1000"
        applheapsz = "AUTOMATIC"
        appl_memory = "AUTOMATIC"
        app_ctl_heap_sz = "500"
        archretrydelay = "60"
        authn_cache_duration = "300"
        autorestart = "ON"
        auto_cg_stats = "OFF"
        auto_maint = "ON"
        auto_reorg = "OFF"
        auto_reval = "IMMEDIATE"
        auto_runstats = "ON"
        auto_sampling = "OFF"
        auto_stats_views = "ON"
        auto_stmt_stats = "OFF"
        auto_tbl_maint = "ON"
        chngpgs_thresh = "50"
        cur_commit = "AVAILABLE"
        database_memory = "AUTOMATIC"
        dbheap = "AUTOMATIC"
        db_mem_thresh = "80"
        ddl_compression_def = "YES"
        ddl_constraint_def = "NO"
        decflt_rounding = "ROUND_HALF_UP"
        dec_to_char_fmt = "NEW"
        dft_degree = "ANY"
        dft_extent_sz = "32"
        dft_loadrec_ses = "1000"
        dft_prefetch_sz = "AUTOMATIC"
        dft_queryopt = "5"
        dft_schemas_dcc = "NO"
        dft_sqlmathwarn = "YES"
        dft_table_org = "ROW"
        dlchktime = "30000"
        enable_xmlchar = "NO"
        extended_row_sz = "ENABLE"
        groupheap_ratio = "75"
        indexrec = "ACCESS"
        large_aggregation = "NO"
        locklist = "AUTOMATIC"
        locktimeout = "10"
        logindexbuild = "OFF"
        log_appl_info = "YES"
        log_ddl_stmts = "NO"
        log_disk_cap = "1000000"
        maxappls = "500"
        maxfilop = "1024"
        maxlocks = "AUTOMATIC"
        min_dec_div_3 = "NO"
        mon_act_metrics = "BASE"
        mon_deadlock = "HISTORY"
        mon_lck_msg_lvl = "2"
        mon_locktimeout = "HISTORY"
        mon_lockwait = "WITHOUT_HIST"
        mon_lw_thresh = "5000"
        mon_obj_metrics = "EXTENDED"
        mon_pkglist_sz = "512"
        mon_req_metrics = "BASE"
        mon_rtn_data = "NONE"
        mon_rtn_execlist = "OFF"
        mon_uow_data = "BASE"
        mon_uow_execlist = "ON"
        mon_uow_pkglist = "OFF"
        nchar_mapping = "GRAPHIC_CU32"
        num_freqvalues = "1000"
        num_iocleaners = "AUTOMATIC"
        num_ioservers = "AUTOMATIC"
        num_log_span = "10000"
        num_quantiles = "1000"
        opt_direct_wrkld = "YES"
        page_age_trgt_gcr = "5000"
        page_age_trgt_mcr = "5000"
        pckcachesz = "AUTOMATIC"
        pl_stack_trace = "ALL"
        self_tuning_mem = "ON"
        seqdetect = "YES"
        sheapthres_shr = "AUTOMATIC"
        sortheap = "AUTOMATIC"
        stat_heap_sz = "AUTOMATIC"
        stmtheap = "AUTOMATIC"
        stmt_conc = "LITERALS"
        string_units = "SYSTEM"
        systime_period_adj = "NO"
        trackmod = "YES"
        util_heap_sz = "AUTOMATIC"
        wlm_admission_ctrl = "YES"
        wlm_agent_load_trgt = "AUTOMATIC"
        wlm_cpu_limit = "50"
        wlm_cpu_shares = "1000"
        wlm_cpu_share_mode = "SOFT"
      }
      dbm {
        comm_bandwidth = "1000"
        cpuspeed = "-1"
        dft_mon_bufpool = "ON"
        dft_mon_lock = "OFF"
        dft_mon_sort = "ON"
        dft_mon_stmt = "OFF"
        dft_mon_table = "ON"
        dft_mon_timestamp = "OFF"
        dft_mon_uow = "ON"
        diaglevel = "2"
        federated_async = "ANY"
        indexrec = "ACCESS"
        intra_parallel = "YES"
        keepfenced = "NO"
        max_connretries = "10"
        max_querydegree = "ANY"
        mon_heap_sz = "AUTOMATIC"
        multipartsizemb = "100"
        notifylevel = "3"
        num_initagents = "500"
        num_initfenced = "1000"
        num_poolagents = "2000"
        resync_interval = "300"
        rqrioblk = "8192"
        start_stop_time = "60"
        util_impact_lim = "50"
        wlm_dispatcher = "YES"
        wlm_disp_concur = "COMPUTED"
        wlm_disp_cpu_shares = "YES"
        wlm_disp_min_util = "75"
      }
      registry {
        db2_bidi = "YES"
        db2_lock_to_rb = "STATEMENT"
        db2_stmm = "YES"
        db2_alternate_authz_behaviour = "EXTERNAL_ROUTINE_DBADM"
        db2_antijoin = "EXTEND"
        db2_ats_enable = "YES"
        db2_deferred_prepare_semantics = "NO"
        db2_evaluncommitted = "YES"
        db2_index_pctfree_default = "10"
        db2_inlist_to_nljn = "YES"
        db2_minimize_listprefetch = "NO"
        db2_object_table_entries = "5000"
        db2_optprofile = "NO"
        db2_selectivity = "ALL"
        db2_skipdeleted = "YES"
        db2_skipinserted = "NO"
        db2_sync_release_lock_attributes = "YES"
        db2_truncate_reusestorage = "IMPORT"
        db2_use_alternate_page_cleaning = "ON"
        db2_view_reopt_values = "NO"
        db2_workload = "SAP"
      }
  }
}
```
**provider.tf**
Please make sure to target right region in the provider block. If database is created in region other than `us-south` , please specify it in provider block.

```terraform
provider "ibm" {
  ibmcloud_api_key      = var.ibmcloud_api_key
}
```


## Timeouts
The following timeouts are defined for this resource.

* `Create` The creation of an instance is considered failed when no response is received for 720 minutes.
* `Delete` The deletion of an instance is considered failed when no response is received for 30 minutes.

Db2 SaaS create instance typically takes between 30 minutes to 45 minutes. Delete and update takes a minute. Provisioning time are unpredictable, if the apply fails due to a timeout, import the database resource once the create is completed.


## Argument reference
Review the argument reference that you can specify for your resource.


- `location` - (Required, String) The location where you want to deploy your instance. The location must match the `region` parameter that you specify in the `provider` block of your  Terraform configuration file. Currently, supported regions are `us-south`, `us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `mon01`, `br-sao`, `ca-tor`, `mil01`.
- `name` - (Required, String) A descriptive name that is used to identify the database instance. The name must not include spaces.
- `plan` - (Required, Forces new resource, String) The name of the service plan to use when provisioning.  Currently the only supported option is `performance`.
- `resource_group_id` - (Optional, Forces new resource, String)  The ID of the resource group where you want to create the instance. To retrieve this value, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `service` - (Required, Forces new resource, String) The type of Cloud Db2 SaaS that you want to create. Only the following services are currently accepted: `dashdb-for-transactions` only.
- `service_endpoints` - (Required, String) Specify whether you want to enable the public, private, or both service endpoints. Supported values are `public`, `private`, or `public-and-private`.
- `tags` - (Optional, Array of Strings) A list of tags that you want to add to your instance.
- `high_availability` - (Optional, String) By default, it is `no`.if you want please change to `yes`
- `backup_location` - (Optional, String) Cross Regional backups can be stored across multiple regions in a zone. Regional backups are stored in only one specific region.
- `instance_type` - (Optional, String) The hosting infrastructure identifier.By Default `bx2.1x4` taken automatically. With this identifier, minimum resource configurations apply. Alternatively, setting the identifier to any of the following host sizes places your database on the specified host size with no other tenants.
          - `bx2.4x16`
          - `bx2.8x32`
          - `bx2.16x64`
          - `bx2.32x128`
          - `bx2.48x192`
          - `mx2.4x32`
          - `mx2.16x128`
          - `mx2.128x1024`
- `disk_encryption_instance_crn` - (Optional, String) Please ensure Databases for Db2 has been authorized to access the selected KMS instance.
- `disk_encryption_key_crn` - (Optional, String) Warning: deleting this key will result in the loss of all data stored in this Db2 instance.
- `oracle_compatibility` - (Optional, String) If you require Oracle compatibility, please choose this option(YES/NO).
- `subscription_id` - (Optional, String) ID which is required for subscription plans, for example: PerformanceSubscription.
- `parameters_json` - (Optional, JSON) Parameters to create Db2 SaaS instance. The value must be a JSON string.

  Nested schema for `parameters_json`:
  - `backup_encryption_key_crn` -  (Optional, Forces new resource, String) The CRN of a key protect key, that you want to use for encrypting disk that holds deployment backups. A key protect CRN is in the format `crn:v1:<...>:key:`. `backup_encryption_key_crn` can be added only at the time of creation and no update support  are available.
- `users_config` - (Optional, List) Defines users configurations you want to set to Db2 SaaS instance.
    Nested schema for `users_config`
    - `email` - (Required, String) Email address of the user.
    - `iam` - (Required, Boolean) Indicates if IAM is enabled or not.
    - `ibmid` - (Required, String) IBM ID of the user.
    - `locked` - (Required, String) Account lock status for the user.
      - Constraints: Allowable values are: `yes`, `no`.
    - `name` - (Required, String) The display name of the user.
    - `password` - (Required, String) User's password.
    - `role` - (Required, String) Role assigned to the user.
      - Constraints: Allowable values are: `bluadmin`, `bluuser`.
    - `x_deployment_id` - (Required, String) CRN deployment id.
      - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^crn(:[A-Za-z0-9\\-\\.]*){9}$/`.
    - `authentication` - (Required, List) Authentication details for the user.
    Nested schema for **authentication**:
      * `method` - (Required, String) Authentication method.
      * `policy_id` - (Required, String) Policy ID of authentication.

- `allowlist_config` - (Optional, List) Defines allowlist configurations you want to set to Db2 SaaS instance.
    Nested schema for `allowlist_config`
    - `address` - (Required, String) Indicates teh IP Address to be allowed
    - `description` - (Required, String) Defines the description.

- `autoscale_config` - (Optional, List) Defines autoscale configurations you want to set to Db2 SaaS instance.
    Nested schema for `autoscale_config`
      - `auto_scaling_enabled` - (Optional, Boolean) Indicates if automatic scaling is enabled or not.
      - `auto_scaling_threshold` - (Optional, Integer) Specifies the resource utilization level that triggers an auto-scaling.
      - `auto_scaling_over_time_period` - (Optional, Integer) Defines the time period over which auto-scaling adjustments are monitored and applied.
      - `auto_scaling_pause_limit` - (Optional, Integer) Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
      - `auto_scaling_allow_plan_limit` - (Optional, Boolean) Indicates the maximum number of scaling actions that are allowed within a specified time period.
- `custom_setting_config` (Optional, Set) Indicates the custom configuration fields related to database which you want to set to Db2 SaaS instance.
    Nested schema for `custom_setting_config`
      - `db` - (Optional, List) 
          Nested scheme for `db`
            * `act_sortmem_limit` - (Optional, String) Configures the sort memory limit for DB2. Valid values: range(10, 100)
            * `alt_collate` - (Optional, String) Configures the collation sequence, Valid values: "NULL", "IDENTITY_16BIT"
            * `app_ctl_heap_sz` - (Optional, String) Configures the application control heap size. Valid values: range(1, 64000)
            * `appgroup_mem_sz` - (Optional, String) Sets the application group memory size. Valid values: range(1, 1000000)
            * `appl_memory` - (Optional, String) Configures the application memory allocation. Valid values: AUTOMATIC range(128, 4294967295)
            * `applheapsz` - (Optional, String) Configures the application heap size,Valid values: "AUTOMATIC" "range(16, 2147483647)
            * `archretrydelay` - (Optional, String) Configures the archive retry delay time. Valid values: range(0, 65535)
            * `authn_cache_duration` - (Optional, String) onfigures the authentication cache duration. Valid values: range(1,10000)
            * `auto_cg_stats` - (Optional, String) Configures whether auto collection of CG statistics is enabled ,Valid values: ON, OFF
            * `auto_maint` - (Optional, String) Configures automatic maintenance for the database,Valid values: ON ,OFF
            * `auto_reorg` - (Optional, String) Configures automatic reorganization for the database,Valid values: ON ,OFF
            * `auto_reval` - (Optional, String) Configures the auto refresh or revalidation method,Valid values: 'IMMEDIATE', 'DISABLED', 'DEFERRED', 'DEFERRED_FORCE'
            * `auto_runstats` - (Optional, String) Configures automatic collection of run-time statistics,Valid values:'ON', 'OFF'
            * `auto_sampling` - (Optional, String) Configures whether auto-sampling is enabled,Valid values: 'ON', 'OFF'
            * `auto_stats_views` - (Optional, String) Configures automatic collection of statistics on views,Valid values: 'ON', 'OFF'
            * `auto_stmt_stats` - (Optional, String) Configures automatic collection of statement-level statistics,Valid values: 'ON', 'OFF'
            * `auto_tbl_maint` - (Optional, String) Configures automatic table maintenance,Valid values: 'ON', 'OFF'
            * `autorestart` - (Optional, String) Configures whether the database will automatically restart,Valid values: 'ON', 'OFF'
            * `avg_appls` - (Optional, String) Average number of applications.
            * `catalogcache_sz` - (Optional, String) Configures the catalog cache size, 
            * `chngpgs_thresh` - (Optional, String) Configures the change pages threshold percentage, Valid values: range(5,99)
            * `cur_commit` - (Optional, String) Configures the commit behavior, Valid values: "ON" "AVAILABLE" "DISABLED
            * `database_memory` - (Optional, String) Configures the database memory management, Valid values: "AUTOMATIC" "COMPUTED" "range(0,4294967295)"
            * `db_collname` - (Optional, String) Specifies the database collation name
            * `db_mem_thresh` - (Optional, String) Configures the memory threshold percentage for database,Valid values: range(0,100)
            * `dbheap` - (Optional, String) Configures the database heap size,Valid values: "AUTOMATIC" range(32, 2147483647)
            * `ddl_compression_def` - (Optional, String) Defines the default DDL compression behavior,Valid values: YES, NO
            * `ddl_constraint_def` - (Optional, String) Defines the default constraint behavior in DDL,Valid values: YES ,NO
            * `dec_arithmetic` - (Optional, String) Configures the default arithmetic for decimal operations
            * `dec_to_char_fmt` - (Optional, String) Configures the decimal-to-character conversion format, Valid values: "NEW" "V95"
            * `decflt_rounding` - (Optional, String) Configures the decimal floating-point rounding method, Valid values: 'ROUND_HALF_EVEN', 'ROUND_CEILING', 'ROUND_FLOOR', 'ROUND_HALF_UP', 'ROUND_DOWN'
            * `dft_degree` - (Optional, String) Configures the default degree for parallelism, Valid values:'-1', 'ANY', 'range(1, 32767)'
            * `dft_extent_sz` - (Optional, String) Configures the default extent size for tables, Valid values: range(2, 256)
            * `dft_loadrec_ses` - (Optional, String) Configures the default load record session count, Valid values: range(1, 30000)
            * `dft_mttb_types` - (Optional, String) Configures the default MTTB (multi-table table scan) types
            * `dft_prefetch_sz` - (Optional, String) Configures the default prefetch size for queries, Valid values: 'range(0, 32767)', 'AUTOMATIC'
            * `dft_queryopt` - (Optional, String) onfigures the default query optimization level, Valid values: range(0, 9)
            * `dft_refresh_age` - (Optional, String) Configures the default refresh age for views
            * `dft_schemas_dcc` - (Optional, String) Configures whether DCC (database control center) is enabled for schemas,Valid values:'YES', 'NO'
            * `dft_sqlmathwarn` - (Optional, String) Configures whether SQL math warnings are enabled,Valid values: 'YES', 'NO'
            * `dft_table_org` - (Optional, String) Configures the default table organization (ROW or COLUMN),Valid values: 'COLUMN', 'ROW'
            * `dlchktime` - (Optional, String) Configures the deadlock check time in milliseconds, Valid values: range(1000, 600000)'
            * `enable_xmlchar` - (Optional, String) Configures whether XML character support is enabled,Valid values: YES', 'NO'
            * `extended_row_sz` - (Optional, String) Configures whether extended row size is enabled,Valid values: ENABLE', 'DISABLE'
            * `groupheap_ratio` - (Optional, String) Configures the heap ratio for group heap memory, Valid values: range(1, 99)'
            * `indexrec` - (Optional, String) Configures the index recovery method, Valid values:'SYSTEM', 'ACCESS', 'ACCESS_NO_REDO', 'RESTART', 'RESTART_NO_REDO'
            * `large_aggregation` - (Optional, String) Configures whether large aggregation is enabled, Valid values:'YES', 'NO'
            * `locklist` - (Optional, String) Configures the lock list memory size, Valid values:'AUTOMATIC', 'range(4, 134217728)'
            * `locktimeout` - (Optional, String) Configures the lock timeout duration,Valid values: '-1', 'range(0, 32767)'
            * `log_appl_info` - (Optional, String) Configures whether application information is logged,Valid values: 'YES', 'NO'
            * `log_ddl_stmts` - (Optional, String) Configures whether DDL statements are logged, Valid values:'YES', 'NO'
            * `log_disk_cap` - (Optional, String) Configures the disk capacity log setting,Valid values: '0', '-1', 'range(1, 2147483647)'
            * `logindexbuild` - (Optional, String) Configures whether index builds are logged, Valid values:'ON', 'OFF'
            * `maxappls` - (Optional, String) Configures the maximum number of applications,Valid values:range(1, 60000)
            * `maxfilop` - (Optional, String) Configures the maximum number of file operations,Valid values: range(64, 61440)
            * `maxlocks` - (Optional, String) Configures the maximum number of locks, Valid values:'AUTOMATIC', 'range(1, 100)'
            * `min_dec_div_3` - (Optional, String) Configures whether decimal division by 3 should be handled,Valid values: 'YES', 'NO'
            * `mon_act_metrics` - (Optional, String) Configures the level of activity metrics to be monitored,Valid values: 'NONE', 'BASE', 'EXTENDED'
            * `mon_deadlock` - (Optional, String) Configures deadlock monitoring settings,Valid values: 'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'
            * `mon_lck_msg_lvl` - (Optional, String) Configures the lock message level for monitoring, Valid values: range(0, 3)
            * `mon_locktimeout` - (Optional, String) Configures lock timeout monitoring settings, Valid values:'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'
            * `mon_lockwait` - (Optional, String) Configures lock wait monitoring settings,Valid values: 'NONE', 'WITHOUT_HIST', 'HISTORY', 'HIST_AND_VALUES'
            * `mon_lw_thresh` - (Optional, String) Configures the lightweight threshold for monitoring, Valid values: range(1000, 4294967295)
            * `mon_obj_metrics` - (Optional, String) Configures the object metrics level for monitoring,Valid values: 'NONE', 'BASE', 'EXTENDED'
            * `mon_pkglist_sz` - (Optional, String) Configures the package list size for monitoring, Valid values: range(0, 1024)
            * `mon_req_metrics` - (Optional, String) Configures the request metrics level for monitoring, Valid values: 'NONE', 'BASE', 'EXTENDED'
            * `mon_rtn_data` - (Optional, String) Configures the level of return data for monitoring,Valid values: 'NONE', 'BASE'
            * `mon_rtn_execlist` - (Optional, String) Configures whether stored procedure execution list is monitored, Valid values:'OFF', 'ON'
            * `mon_uow_data` - (Optional, String) Configures the level of unit of work (UOW) data for monitoring,Valid values: 'NONE', 'BASE'
            * `mon_uow_execlist` - (Optional, String) Configures whether UOW execution list is monitored,Valid values: 'ON', 'OFF'
            * `mon_uow_pkglist` - (Optional, String) Configures whether UOW package list is monitored,Valid values: 'OFF', 'ON'
            * `nchar_mapping` - (Optional, String) Configures the mapping of NCHAR character types, Valid values:'CHAR_CU32', 'GRAPHIC_CU32', 'GRAPHIC_CU16', 'NOT APPLICABLE'
            * `num_freqvalues` - (Optional, String) Configures the number of frequent values for optimization, Valid values: range(0, 32767)
            * `num_iocleaners` - (Optional, String) Configures the number of IO cleaners, 'AUTOMATIC', Valid values: 'range(0, 255)'
            * `num_ioservers` - (Optional, String) Configures the number of IO servers, 'AUTOMATIC', Valid values: 'range(1, 255)'
            * `num_log_span` - (Optional, String) Configures the number of log spans, Valid values: 'range(0, 65535)'
            * `num_quantiles` - (Optional, String) Configures the number of quantiles for optimizations, Valid values: 'range(0, 32767)'
            * `opt_buffpage` - (Optional, String) Configures the buffer page optimization setting
            * `opt_direct_wrkld` - (Optional, String) Configures the direct workload optimization setting,Valid values: 'ON', 'OFF', 'YES', 'NO', 'AUTOMATIC'
            * `opt_locklist` - (Optional, String) Configures the lock list optimization setting
            * `opt_maxlocks` - (Optional, String) Configures the max locks optimization setting
            * `opt_sortheap` - (Optional, String) Configures the sort heap optimization setting
            * `page_age_trgt_gcr` - (Optional, String) Configures the page age target for garbage collection, Valid values: 'range(1, 65535)'
            * `page_age_trgt_mcr` - (Optional, String) Configures the page age target for memory collection, Valid values: 'range(1, 65535)'
            * `pckcachesz` - (Optional, String) Configures the package cache size,Valid values: 'AUTOMATIC', '-1', 'range(32, 2147483646)'
            * `pl_stack_trace` - (Optional, String) Configures the level of stack trace logging for stored procedures, Valid values:'NONE', 'ALL', 'UNHANDLED'
            * `self_tuning_mem` - (Optional, String) Configures whether self-tuning memory is enabled,Valid values: 'ON', 'OFF'
            * `seqdetect` - (Optional, String) Configures sequence detection for queries,Valid values: 'YES', 'NO'
            * `sheapthres_shr` - (Optional, String) Configures the shared heap threshold size, Valid values:'AUTOMATIC', 'range(250, 2147483647)'
            * `softmax` - (Optional, String) Configures the soft max setting
            * `sortheap` - (Optional, String) Configures the sort heap memory size,Valid values: 'AUTOMATIC', 'range(16, 4294967295)'
            * `sql_ccflags` - (Optional, String) Configures the SQL compiler flags
            * `stat_heap_sz` - (Optional, String) Configures the statistics heap size, Valid values:'AUTOMATIC', 'range(1096, 2147483647)'
            * `stmt_conc` - (Optional, String) Configures the statement concurrency, Valid values:'OFF', 'LITERALS', 'COMMENTS', 'COMM_LIT'
            * `stmtheap` - (Optional, String) Configures the statement heap size,Valid values: 'AUTOMATIC', 'range(128, 2147483647)'
            * `string_units` - (Optional, String) Configures the string unit settings,Valid values: 'SYSTEM', 'CODEUNITS32'
            * `systime_period_adj` - (Optional, String) Configures whether system time period adjustments are enabled,Valid values: 'NO', 'YES'
            * `trackmod` - (Optional, String) Configures whether modifications to tracked objects are logged,Valid values: 'YES', 'NO'
            * `util_heap_sz` - (Optional, String) Configures the utility heap size, Valid values:'AUTOMATIC', 'range(16, 2147483647)'
            * `wlm_admission_ctrl` - (Optional, String) Configures whether WLM (Workload Management) admission control is enabled, Valid values:'YES', 'NO'
            * `wlm_agent_load_trgt` - (Optional, String) Configures the WLM agent load target,Valid values: 'AUTOMATIC', 'range(1, 65535)'
            * `wlm_cpu_limit` - (Optional, String) Configures the CPU limit for WLM workloads, Valid values: 'range(0, 100)'
            * `wlm_cpu_share_mode` - (Optional, String) Configures the mode of CPU shares for WLM workloads,Valid values: 'HARD', 'SOFT'
            * `wlm_cpu_shares` - (Optional, String) Configures the CPU share count for WLM workloads, Valid values: 'range(1, 65535)'
	    - `dbm` - (List) Tunable parameters related to the Db2 instance manager (dbm).
	        Nested schema for **dbm**:
            * `comm_bandwidth` - (Optional, String) Configures the communication bandwidth for the database manager, Valid values: 'range(0.1, 100000)', '-1'
            * `cpuspeed` - (Optional, String) Configures the CPU speed for the database manager, Valid values: 'range(0.0000000001, 1)', '-1'
            * `dft_mon_bufpool` - (Optional, String) Configures whether the buffer pool is monitored by default,Valid values: 'ON', 'OFF'
            * `dft_mon_lock` - (Optional, String) Configures whether lock monitoring is enabled by default,Valid values: 'ON', 'OFF'
            * `dft_mon_sort` - (Optional, String) Configures whether sort operations are monitored by default,Valid values: 'ON', 'OFF'
            * `dft_mon_stmt` - (Optional, String) Configures whether statement execution is monitored by default,Valid values: 'ON', 'OFF'
            * `dft_mon_table` - (Optional, String) Configures whether table operations are monitored by default, Valid values:'ON', 'OFF'
            * `dft_mon_timestamp` - (Optional, String) Configures whether timestamp monitoring is enabled by default,Valid values: 'ON', 'OFF'
            * `dft_mon_uow` - (Optional, String) Configures whether unit of work (UOW) monitoring is enabled by default,Valid values: 'ON', 'OFF'
            * `diaglevel` - (Optional, String) Configures the diagnostic level for the database manager, Valid values: 'range(0, 4)'
            * `federated_async` - (Optional, String) Configures whether federated asynchronous mode is enabled, Valid values: 'range(0, 32767)', '-1', 'ANY'
            * `indexrec` - (Optional, String) Configures the type of indexing to be used in the database manager, Valid values:'RESTART', 'RESTART_NO_REDO', 'ACCESS', 'ACCESS_NO_REDO'
            * `intra_parallel` - (Optional, String) Configures the parallelism settings for intra-query parallelism, Valid values:'SYSTEM', 'NO', 'YES'
            * `keepfenced` - (Optional, String) Configures whether fenced routines are kept in memory,Valid values: 'YES', 'NO'
            * `max_connretries` - (Optional, String) Configures the maximum number of connection retries, Valid values: 'range(0, 100)'
            * `max_querydegree` - (Optional, String) Configures the maximum degree of parallelism for queries, Valid values: 'range(1, 32767)', '-1', 'ANY'
            * `mon_heap_sz` - (Optional, String) Configures the size of the monitoring heap, Valid values: 'range(0, 2147483647)', 'AUTOMATIC'
            * `multipartsizemb` - (Optional, String) Configures the size of multipart queries in MB, Valid values: 'range(5, 5120)'
            * `notifylevel` - (Optional, String) Configures the level of notifications for the database manager, Valid values: 'range(0, 4)'
            * `num_initagents` - (Optional, String) Configures the number of initial agents in the database manager, Valid values: 'range(0, 64000)'
            * `num_initfenced` - (Optional, String) Configures the number of initial fenced routines, Valid values: 'range(0, 64000)'
            * `num_poolagents` - (Optional, String) Configures the number of pool agents,Valid values: '-1', 'range(0, 64000)'
            * `resync_interval` - (Optional, String) Configures the interval between resync operations, Valid values: 'range(1, 60000)'
            * `rqrioblk` - (Optional, String) Configures the request/response I/O block size, Valid values: 'range(4096, 65535)'
            * `start_stop_time` - (Optional, String) Configures the time in minutes for start/stop operations, Valid values: 'range(1, 1440)'
            * `util_impact_lim` - (Optional, String) Configures the utility impact limit, Valid values: 'range(1, 100)'
            * `wlm_disp_concur` - (Optional, String) Configures whether the WLM (Workload Management) dispatcher is enabled, Valid values:'YES', 'NO'
            * `wlm_disp_min_util` - (Optional, String) Configures the minimum utility threshold for WLM dispatcher, Valid values: 'range(0, 100)'
            * `wlm_dispatcher` - (Optional, String) Configures whether the WLM (Workload Management) dispatcher is enabled,Valid values:'YES', 'NO'
	      - `registry` - (List) Tunable parameters related to the Db2 registry.
	        Nested schema for **registry**:
            * `db2_alternate_authz_behaviour` - (Optional, String) Configures the alternate authorization behavior for Valid values:DB2, 'EXTERNAL_ROUTINE_DBADM', 'EXTERNAL_ROUTINE_DBAUTH'
            * `db2_antijoin` - (Optional, String) Configures how DB2 handles anti-joins, Valid values:'YES', 'NO', 'EXTEND'
            * `db2_ats_enable` - (Optional, String) Configures whether DB2 asynchronous table scanning (ATS) is enabled,Valid values: 'YES', 'NO'
            * `db2_bidi` - (Optional, String) Configures the bidi (bidirectional) support for DB2,Valid values: 'YES', 'NO'
            * `db2_compopt` - (Optional, String) Configures the DB2 component options (not specified in values)
            * `db2_deferred_prepare_semantics` - (Optional, String) Configures whether deferred prepare semantics are enabled in DB2,Valid values: 'NO', 'YES'
            * `db2_evaluncommitted` - (Optional, String) Configures whether uncommitted data is evaluated by DB2,Valid values: 'NO', 'YES'
            * `db2_extended_optimization` - (Optional, String) Configures extended optimization in DB2 (not specified in values)
            * `db2_index_pctfree_default` - (Optional, String) Configures the default percentage of free space for DB2 indexes, Valid values: 'range(0, 99)'
            * `db2_inlist_to_nljn` - (Optional, String) Configures whether in-list queries are converted to nested loop joins,Valid values: 'NO', 'YES'
            * `db2_lock_to_rb` - (Optional, String) Configures the DB2 lock timeout behavior, Valid values:'STATEMENT'
            * `db2_minimize_listprefetch` - (Optional, String) Configures whether DB2 minimizes list prefetching for queries, Valid values:'NO', 'YES'
            * `db2_object_table_entries` - (Optional, String) Configures the number of entries for DB2 object tables, Valid values: 'range(0, 65532)'
            * `db2_opt_max_temp_size` - (Optional, String) Configures the maximum temporary space size for DB2 optimizer
            * `db2_optprofile` - (Optional, String) Configures whether DB2's optimizer profile is enabled, Valid values:'YES','NO'
            * `db2_optstats_log` - (Optional, String) Configures the logging of optimizer statistics (not specified in values)
            * `db2_parallel_io` - (Optional, String) Configures parallel I/O behavior in DB2 (not specified in values)
            * `db2_reduced_optimization` - (Optional, String) Configures whether reduced optimization is applied in DB2 (not specified in values)
            * `db2_selectivity` - (Optional, String) Configures the selectivity behavior for DB2 queries, Valid values:'YES', 'NO', 'ALL'
            * `db2_skipdeleted` - (Optional, String) Configures whether DB2 skips deleted rows during query processing, Valid values:'NO', 'YES'
            * `db2_skipinserted` - (Optional, String) Configures whether DB2 skips inserted rows during query processing, Valid values:'NO', 'YES'
            * `db2_stmm` - (Optional, String) Configures whether DB2's self-tuning memory manager (STMM) is enabled, Valid values:'NO', 'YES'
            * `db2_sync_release_lock_attributes` - (Optional, String) Configures whether DB2 synchronizes lock release attributes,Valid values: 'NO', 'YES'
            * `db2_truncate_reusestorage` - (Optional, String) Configures the types of operations that reuse storage after truncation, Valid values:'IMPORT', 'LOAD', 'TRUNCATE'
            * `db2_use_alternate_page_cleaning` - (Optional, String) Configures whether DB2 uses alternate page cleaning methods, Valid values:'ON', 'OFF'
            * `db2_view_reopt_values` - (Optional, String) Configures whether DB2 view reoptimization values are used,Valid values: 'NO', 'YES'
            * `db2_wlm_settings` - (Optional, String) Configures the WLM (Workload Management) settings for DB2 (not specified in values)
            * `db2_workload` - (Optional, String) Configures the DB2 workload type,Valid values: '1C', 'ANALYTICS', 'CM', 'COGNOS_CS', 'FILENET_CM', 'INFOR_ERP_LN', 'MAXIMO', 'MDM', 'SAP', 'TPM', 'WAS', 'WC', 'WP'

## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The CRN of the database instance.
- `status` - (String) The status of the instance.
- `version` - (String) The database version.
- `all_clean` - (Boolean) Indicates if the user account has no issues.
- `dv_role` - (String) User's DV role.
- `formated_ibmid` - (String) Formatted IBM ID.
- `iamid` - (String) IAM ID for the user.
- `init_error_msg` - (String) Initial error message.
- `metadata` - (Map) Metadata associated with the user.
- `permitted_actions` - (List) List of allowed actions of the user.

## Import
The database instance can be imported by using the ID, that is formed from the CRN. To import the resource, you must specify the `region` parameter in the `provider` block of your  Terraform configuration file. If the region is not specified, `us-south` is used by default. An  Terraform refresh or apply fails, if the database instance is not in the same region as configured in the provider or its alias.

CRN is a 120 digit character string of the form -  `crn:v1:bluemix:public:dashdb-for-transactions:us-south:a/60970f92286548d8a64cbb45bce39bc1:deae06ff-3966-4534-bfa0-4b42281e7cef::`

**Syntax**

```
$ terraform import ibm_db2.my_db <crn>
```

**Example**

```
$ terraform import ibm_db2.my_db crn:v1:bluemix:public:dashdb-for-transactions:us-south:a/60970f92286548d8a64cbb45bce39bc1:deae06ff-3966-4534-bfa0-4b42281e7cef::
```

Import requires a minimal Terraform config file to allow importing.

```terraform
resource "ibm_db2" "<your_database>" {
  name              = "<your_database_name>"
}
```

Run `terraform state show ibm_db2.<your_database>` after import to retrieve the more values to be included in the resource config file. It does not export any more user IDs and passwords that are configured on the instance. These values must be retrieved from an alternative source.
# Terraform IBM Provider 
<!-- markdownlint-disable MD026 -->
This area is primarily for IBM provider contributors and maintainers. For information on _using_ Terraform and the IBM provider, see the links below.

## Handy Links
* [Find out about contributing](../../../CONTRIBUTING.md) to the IBM provider!
* IBM Provider Docs: [Home](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs)
* IBM Provider Docs: [One of the  resources](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/common_source_registration_request)
* IBM API Docs: [IBM API Docs for ]()
* IBM  SDK: [IBM SDK for ](https://github.com/IBM/appconfiguration-go-admin-sdk/tree/master/backuprecoveryv1)


## Known Limitations
### Resources with Incomplete CRUD Operations
This service includes certain resources that do not have fully implemented CRUD (Create, Read, Update, Delete) operations due to limitations in the underlying APIs. Specifically:

#### Protection Group Run:

***Create:*** A `ibm_protection_group_run_request` resource is available for creating new protection group run.

***Update:*** Book updates are managed through a separate `ibm_update_protection_group_run_request` resource. Note that the `ibm_protection_group_run_request` and `ibm_update_protection_group_run_request` resources must be used in tandem to manage Protection Group Runs.

***Delete:*** There is no delete operation available for the protection group run resource. If  ibm_update_protection_group_run_request or ibm_protection_group_run_request resource is removed from the `main.tf` file, Terraform will remove it from the state file but not from the backend. The resource will continue to exist in the backend system.


#### Other resources that do not support update and delete:

Some resources in this service do not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration will only remove them from the Terraform state and will not affect the actual resources in the backend.
- ibm_perform_action_on_protection_group_run_request
- ibm_recovery_download_files_folders
- ibm_recovery_cancel
- ibm_recovery_teardown
- ibm_search_indexed_object
- ibm_protection_group_state

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**
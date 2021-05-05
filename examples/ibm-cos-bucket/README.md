# IBM Cloud Object Storage example

The following example creates an instance of IBM Cloud Object Storage, IBM Cloud Activity Tracker, and IBM Cloud Monitoring with Sysdig. Then, multiple buckets are created and configured to send audit events and metrics to your service instances.

Following types of resources are supported:

* [ Cloud Object Storage Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-object-storage-resources#cos-bucket)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.7.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.29.1`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create an IBM Cloud Object Storage bucket. The bucket is used to store your data:

```hcl

data "ibm_resource_group" "cos_group" {
  name = var.resource_group_name
}
resource "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}
resource "ibm_resource_instance" "activity_tracker" {
  name              = "activity_tracker"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "logdnaat"
  plan              = "lite"
  location          = "us-south"
}
resource "ibm_resource_instance" "metrics_monitor" {
  name              = "metrics_monitor"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "sysdig-monitor"
  plan              = "graduated-tier"
  location          = "us-south"
  parameters        = {
    default_receiver = true
  }
}
resource "ibm_cos_bucket" "standard-ams03" {
  bucket_name          = var.bucket_name
  resource_instance_id = ibm_resource_instance.cos_instance.id
  single_site_location = "sjc04"
  #cross_region_location = var.region
  storage_class        = var.storage
  activity_tracking {
    read_data_events     = true
    write_data_events    = true
    activity_tracker_crn = ibm_resource_instance.activity_tracker.id
  }
  metrics_monitoring {
    usage_metrics_enabled  = true
    request_metrics_enabled = true
    metrics_monitoring_crn = ibm_resource_instance.metrics_monitor.id
  }
  allowed_ip =  ["223.196.168.27","223.196.161.38","192.168.0.1"]
}

```

```hcl
data "ibm_resource_instance" "cos_instance" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
}

data "ibm_cos_bucket" "standard-ams03" {
  bucket_name = ibm_cos_bucket.standard-ams03-firewall.bucket_name
  resource_instance_id = data.ibm_resource_instance.cos_instance.id
  bucket_type = "region_location"
  bucket_region = "us-south"
}
```

## Examples

* [ Cloud Objcet Storage  ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cos-bucket)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| bucket\_name | Name of the bucket. | `string` | yes |
| resource\_group\_name | Name of the resource group. | `string` | yes |
| storage | The storage class that you want to use for the bucket. Supported values are standard, vault, cold, flex, and smart.| `string` | no |
| region | The location for a cross-regional bucket. Supported values are us, eu, and ap. | `string` | no |
| read_data_events | Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events.. | `array` | no
| write_data_events | all object write events (i.e. uploads) will be sent to Activity Tracker. | `bool` | no
| activity_tracker_crn | Required the first time activity_tracking is configured. | `string` | yes
| usage_metrics_enabled | Specify true or false to set usage metrics (i.e. bytes_used). | `bool` | no
| request_metrics_enabled | Specify true or false to set cos request metrics (i.e. get,put,post request). | `bool` | no
| metrics_monitoring_crn | Required the first time metrics_monitoring is configured. The instance of IBM Cloud Monitoring that will receive the bucket metrics. | `string` | yes
| archive_ruleid | Unique identifier for the rule. | `string` | no
| regional_loc | The location for a regional bucket. Supported values are au-syd, eu-de, eu-gb, jp-tok,,us-east,us-south. | `string` | no
| archive_days | Specifies the number of days when the specific archive rule action takes effect. | `int` | yes
| archive_types | Specifies the archive type to which you want the object to transition. Supported values are  Glacier or Accelerated. | `string` | yes
| expire_ruleid | Unique identifier for the rule. | `string` | no
| expire_days | Specifies the number of days when the specific expire rule action takes effect. | `int` | yes
| expire_prefix | Specifies a prefix filter to apply to only a subset of objects with names that match the prefix. | `string` | no
| default | Specifies a default retention period to apply in all objects in the bucket. | `int` | yes
| maximum | Specifies maximum duration of time an object can be kept unmodified in the bucket. | `int` | yes
| minimum | Specifies minimum duration of time an object must be kept unmodified in the bucket. | `int` | yes
| permanent | Specifies a permanent retention status either enable or disable for a bucket. | `bool` | no

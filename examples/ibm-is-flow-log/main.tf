# ==========================================================================
# Flow Logs Configuration
# ==========================================================================

# ==========================================================================
# Data Sources
# ==========================================================================

# Get resource group for Cloud Object Storage
data "ibm_resource_group" "cos_group" {
  name = var.resource_group # Use resource group name from variable
}

# Get information about an existing instance
data "ibm_is_instance" "ds_instance" {
  name = "vpc1-instance" # Target instance for flow logs
}

# ==========================================================================
# Cloud Object Storage Resources
# ==========================================================================

# Create a Cloud Object Storage service instance
resource "ibm_resource_instance" "instance1" {
  name              = "cos-instance"                       # Name for the COS instance
  resource_group_id = data.ibm_resource_group.cos_group.id # Resource group for COS
  service           = "cloud-object-storage"               # IBM Cloud service type
  plan              = "standard"                           # Service plan
  location          = "global"                             # Global service
}

# Create a bucket for storing flow logs
resource "ibm_cos_bucket" "bucket1" {
  bucket_name          = "us-south-bucket-vpc1"             # Name of the bucket
  resource_instance_id = ibm_resource_instance.instance1.id # COS instance
  region_location      = var.region                         # Region for the bucket
  storage_class        = "standard"                         # Storage class (standard, vault, cold, smart)
}

# ==========================================================================
# Flow Log Resources
# ==========================================================================

# Create a flow log collector for an instance
resource "ibm_is_flow_log" "test_flowlog" {
  depends_on     = [ibm_cos_bucket.bucket1]            # Create after bucket is available
  name           = "test-instance-flow-log"            # Name for the flow log
  target         = data.ibm_is_instance.ds_instance.id # Target instance to collect logs for
  active         = true                                # Flow log collector is active
  storage_bucket = ibm_cos_bucket.bucket1.bucket_name  # Bucket where logs will be stored
}

# ==========================================================================
# Flow Log Data Sources
# ==========================================================================

# List all flow logs
data "ibm_is_flow_logs" "test_flow_logs" {
  # No filters specified - will return all flow logs
}
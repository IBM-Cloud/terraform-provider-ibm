
# Terraform IBM Local Development Guide

This guide outlines how to build, install, and test a **local build** of the IBM Terraform Provider. It supports working against the STG environment.

---

## Environment Setup (STG)

Before running Terraform, set the following environment variables:

```bash
#!/bin/sh
export IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT=https://accountmanagement.stage1.ng.bluemix.net
export IBMCLOUD_CF_API_ENDPOINT=https://api.stage1.ng.bluemix.net
export IBMCLOUD_CS_API_ENDPOINT=https://containers.test.cloud.ibm.com/global
export IBMCLOUD_CR_API_ENDPOINT=https://registry.stage1.ng.bluemix.net
export IBMCLOUD_CIS_API_ENDPOINT=https://api.cis.test.cloud.ibm.com
export IBMCLOUD_GS_API_ENDPOINT=https://api.global-search-tagging.test.cloud.ibm.com
export IBMCLOUD_IAMPAP_API_ENDPOINT=https://iam.test.cloud.ibm.com
export IBMCLOUD_ICD_API_ENDPOINT=https://api.us-south.databases.test.cloud.ibm.com
export IBMCLOUD_MCCP_API_ENDPOINT=https://mccp.us-south.cf.test.cloud.ibm.com
export IBMCLOUD_UAA_ENDPOINT=https://login.stage1.ng.bluemix.net/UAALoginServerWAR
export IBMCLOUD_COS_ENDPOINT=https://s3.us-west.cloud-object-storage.test.appdomain.cloud
export IBMCLOUD_COS_CONFIG_ENDPOINT=https://config.cloud-object-storage.test.cloud.ibm.com/v1

export IBMCLOUD_IAM_API_ENDPOINT=https://iam.test.cloud.ibm.com
export IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT=https://resource-controller.test.cloud.ibm.com
export IBMCLOUD_RESOURCE_CATALOG_API_ENDPOINT=https://globalcatalog.test.cloud.ibm.com
export IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT=https://resource-controller.test.cloud.ibm.com
export IBMCLOUD_GT_API_ENDPOINT=https://tags.global-search-tagging.test.cloud.ibm.com
export IBMCLOUD_IS_NG_API_ENDPOINT=https://us-south-stage01.iaasdev.cloud.ibm.com/v1
export IBMCLOUD_TG_API_ENDPOINT=https://transit.test.cloud.ibm.com/v1
export IBMCLOUD_API_KEY=**********
```

---

## Local Build Setup (One Time Only)

### Step 1: Build the Provider
```bash
cd ~/go/src/github.ibm.com/CloudBBS/terraform-provider-ibm
make build
```

### Step 2: Create Local Plugin Directory
```bash
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/1.77.1/darwin_amd64
```

### Step 3: Copy the Binary
```bash
cp ~/go/bin/terraform-provider-ibm ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/1.77.1/darwin_amd64/terraform-provider-ibm_v1.77.1
chmod +x ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/1.77.1/darwin_amd64/terraform-provider-ibm_v1.77.1
```

### Step 4: Update Your Terraform CLI Config
Edit `~/.terraformrc`:

```hcl
provider_installation {
  filesystem_mirror {
    path    = "/Users/<your_username>/.terraform.d/plugins/"
    include = ["ibm-cloud/ibm"]
  }
  direct {
    include = ["*/*"]
  }
}
```

---

## Rebuild 

```bash
cd ~/go/src/github.ibm.com/CloudBBS/terraform-provider-ibm
make build
cp ~/go/bin/terraform-provider-ibm ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/1.77.1/darwin_amd64/terraform-provider-ibm_v1.77.1
chmod +x ~/.terraform.d/plugins/registry.terraform.io/ibm-cloud/ibm/1.77.1/darwin_amd64/terraform-provider-ibm_v1.77.1
```

Then run Terraform:
```bash
terraform init
terraform plan
terraform apply
```

```


If your logs show your custom build debug lines, the local binary is working.

---

## Troubleshooting

**Problem:**
```text
Error while installing ibm-cloud/ibm v1.77.1: the local package doesn't match any checksums
```

**Fix:**
```bash
rm .terraform.lock.hcl
terraform init
```

---

Keep this file around to save your future self HOURS of headbanging. 🤘


# Examples: 
Create a Gateway and a Redundant GRE Connection

```
terraform {
  required_providers {
    ibm = {
      source  = "ibm-cloud/ibm"
      version = "1.77.1"
    }
    time = {
          source = "hashicorp/time"
          version = "0.9.1"
        }
   }
}

variable "region" {
  default     = "us-south"
  description = "The VPC Region that you want your VPC, Transit Gateway and PDNS to be provisioned it. To list available regions, run `ibmcloud is regions`."
}

# Provider block - Alias initialized to interact with VNF Experiment account
##############################################################################
provider "ibm" {
  ibmcloud_api_key      = "**YourAPIKey**"
  region                = var.region
  ibmcloud_timeout      = 300
}


# Create the Transit Gateway
resource "ibm_tg_gateway" "new_tg_gw" {
  name     = "gw-terraform"
  location = "us-south"
  global   = true
}

# Create RGRE connection using existing gateway
resource "ibm_tg_connection" "test_ibm_tg_connection" {
  gateway           = data.ibm_tg_gateway.existing_tg_gw.id
  network_type      = "redundant_gre"
  name              = "rgre-terraform"
  base_network_type = "classic"

  tunnels {
    local_gateway_ip  = "192.193.200.1"
    local_tunnel_ip   = "192.193.239.2"
    name              = "tunne1_terraform"
    remote_gateway_ip = "10.144.104.123"
    remote_tunnel_ip  = "192.193.239.1"
    zone              = "us-south-1"
  }

  tunnels {
    local_gateway_ip  = "192.193.201.1"
    local_tunnel_ip   = "192.193.238.2"
    name              = "tunne2_terraform"
    remote_gateway_ip = "10.144.104.123"
    remote_tunnel_ip  = "192.193.238.1"
    zone              = "us-south-1"
  }
}

```

# Example 2: Create a Redundant GRE connection on an existing Gateway
```

terraform {
  required_providers {
    ibm = {
      source  = "ibm-cloud/ibm"
      version = "1.77.1"
    }
    time = {
          source = "hashicorp/time"
          version = "0.9.1"
        }
   }
}

variable "region" {
  default     = "us-south"
  description = "The VPC Region that you want your VPC, Transit Gateway and PDNS to be provisioned it. To list available regions, run `ibmcloud is regions`."
}

# Provider block - Alias initialized to interact with VNF Experiment account
##############################################################################
provider "ibm" {
  ibmcloud_api_key      = "**YourAPIKey**"
  region                = var.region
  ibmcloud_timeout      = 300
}


# Look up existing TG gateway
data "ibm_tg_gateway" "existing_tg_gw" {
  name = "gw-terraform"
}

# Create RGRE connection using existing gateway
resource "ibm_tg_connection" "test_ibm_tg_connection" {
  gateway           = data.ibm_tg_gateway.existing_tg_gw.id
  network_type      = "redundant_gre"
  name              = "rgre-terraform"
  base_network_type = "classic"

  tunnels {
    local_gateway_ip  = "192.193.200.1"
    local_tunnel_ip   = "192.193.239.2"
    name              = "tunne1_terraform"
    remote_gateway_ip = "10.144.104.123"
    remote_tunnel_ip  = "192.193.239.1"
    zone              = "us-south-1"
  }

  tunnels {
    local_gateway_ip  = "192.193.201.1"
    local_tunnel_ip   = "192.193.238.2"
    name              = "tunne2_terraform"
    remote_gateway_ip = "10.144.104.123"
    remote_tunnel_ip  = "192.193.238.1"
    zone              = "us-south-1"
  }
}
```



## Example 3. Create a Classic connection 
```

terraform {
  required_providers {
    ibm = {
      source  = "ibm-cloud/ibm"
      version = "1.77.1"
    }
    time = {
          source = "hashicorp/time"
          version = "0.9.1"
        }
   }
}

variable "region" {
  default     = "us-south"
  description = "The VPC Region that you want your VPC, Transit Gateway and PDNS to be provisioned it. To list available regions, run `ibmcloud is regions`."
}

# Provider block - Alias initialized to interact with VNF Experiment account
##############################################################################
provider "ibm" {
  ibmcloud_api_key      = "**YourAPIKey**"
  region                = var.region
  ibmcloud_timeout      = 300
}


# Look up existing TG gateway
data "ibm_tg_gateway" "existing_tg_gw" {
  name = "gw-terraform"
}

# Create RGRE connection using existing gateway
resource "ibm_tg_connection" "test_ibm_tg_connection" {
  gateway           = data.ibm_tg_gateway.existing_tg_gw.id
  network_type      = "classic"
  name              = "classic-terraform"
  local_gateway_ip  = "192.193.200.1"
  local_tunnel_ip   = "192.193.239.2"
  remote_gateway_ip = "10.144.104.123"
  remote_tunnel_ip  = "192.193.239.1"
  zone              = "us-south-1"
}
```
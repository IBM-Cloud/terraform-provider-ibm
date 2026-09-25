
# Terraform IBM Local Development Guide

This guide outlines how to build, install, and test a **local build** of the IBM Terraform Provider. It supports working against the STG environment.

---

## Environment Setup (STG)
Before running Terraform, set the following environment variables:

```bash
#!/bin/sh
export IBMCLOUD_IAM_API_ENDPOINT=https://iam.test.cloud.ibm.com
export IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT=https://cbr.test.cloud.ibm.com
export IBMCLOUD_API_KEY=**********
```

---

## Local Build Setup (One Time Only)

### Step 1: Build the Provider
```bash
cd <local-path-to-repo>/terraform-provider-ibm
go build -o terraform-provider-ibm 
```

### Step 2: Update Your Terraform CLI Config
Edit `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "IBM-Cloud/ibm" = "<local-path-to-repo>/terraform-provider-ibm"
  }
  direct {}
}
```


# Example Network Zone:
```hcl
terraform {
  required_providers {
    ibm = { 
      source  = "IBM-Cloud/ibm"
      version = "> 1.0.0"
    }   
  }
}
resource "ibm_cbr_zone" "cbr_zone_test" {
  name        = "A terraform created network zone"
  account_id  = "<account-id>"
  description = "Network zone created via terraform"

  addresses {
    type  = "ipRange"
    value = "192.168.1.0-192.168.1.5"
  }

  addresses {
    type  = "ipAddress"
    value = "169.23.56.234"
  }
}

data "ibm_cbr_zone" "cbr_zone_data" {
  zone_id = ibm_cbr_zone.cbr_zone_test.id
}

```
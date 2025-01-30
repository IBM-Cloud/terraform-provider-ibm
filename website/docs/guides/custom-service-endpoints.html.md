---
subcategory: ""
layout: "ibm"
page_title: "IBM Cloud Provider plugin for Terraform Custom Service Endpoint Configuration"
description: |-
  Configuring the IBM Cloud Provider plugin for Terraform to connect to custom IBM service endpoints.
---

# Customizing default cloud service endpoints

The IBM Cloud Provider plug-in for Terraform can be configured to use non-default IBM Cloud service endpoints. This setup might be useful for environments with specific compliance requirements or where public network connectivity is not allowed.

**Important**: Support for using non-default IBM Cloud service endpoints with the IBM Cloud Provider plug-in for Terraform is offered as best effort. Individual Terraform resources might require compatibility updates to support the declaration of custom service endpoints. By default, the IBM Cloud Provider plug-in is tested with the default IBM Cloud service endpoints. 

<!-- TOC depthFrom:2 -->

- [Customizing default cloud service endpoints](#customizing-default-cloud-service-endpoints)
  - [Getting started with custom service endpoints](#getting-started-with-custom-service-endpoints)
  - [Supported endpoint customizations](#supported-endpoint-customizations)
  - [File structure for endpoints file](#file-structure-for-endpoints-file)
  - [Prioritisation of endpoints](#prioritisation-of-endpoints)
    - [1. Define service endpoints by using environment variables](#1-define-service-endpoints-by-using-environment-variables)
    - [2. Define service endpoints by using an endpoints file](#2-define-service-endpoints-by-using-an-endpoints-file)
    - [3. Use the default private or public service endpoint based on the `visibility` setting in the provider block](#3-use-the-default-private-or-public-service-endpoint-based-on-the-visibility-setting-in-the-provider-block)
<!-- /TOC -->

## Getting started with custom service endpoints

To configure the IBM Cloud Provider plug-in for Terraform to use custom service endpoints, you can use the `visibility` and `endpoints_file_path` arguments in your `provider` declaration as shown in the following example. 

```terraform
provider "ibm" {
  
  # ... other provider configuration ...

  visiblity="private"
  endpoints_file_path= "endpoints.json"
}
```

**Tip**: If you want to use different endpoint declarations for other services, you must add multiple provider configurations by creating a provider alias. For more information, see the [Terraform documentation](https://www.terraform.io/docs/configuration/providers.html#alias-multiple-provider-instances).

## Supported endpoint customizations 

| Service | Endpoint Variable |
|---------|-----------------|
|Account Management|IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT|
|API Gateway|IBMCLOUD_API_GATEWAY_ENDPOINT|
|App Id|IBMCLOUD_APPID_MANAGEMENT_API_ENDPOINT|
|Atracker|IBMCLOUD_ATRACKER_API_ENDPOINT|
|Catalog Management|IBMCLOUD_CATALOG_MANAGEMENT_API_ENDPOINT|
|Cloud Object Storage|IBMCLOUD_COS_CONFIG_ENDPOINT|
|Context-based Restrictions|IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT|
|Internet Services|IBMCLOUD_CIS_API_ENDPOINT|
|Cloud Shell|IBMCLOUD_CLOUD_SHELL_API_ENDPOINT|
|Compilance (Posture Management)|IBMCLOUD_COMPLIANCE_API_ENDPOINT|
|Container Registry|IBMCLOUD_CR_API_ENDPOINT|
|Cloud Logs | IBMCLOUD_LOGS_API_ENDPOINT |
|Kubernetes Service|IBMCLOUD_CS_API_ENDPOINT|
|Metrics Router| IBMCLOUD_METRICS_ROUTING_API_ENDPOINT|
|MQ on Cloud| IBMCLOUD_MQCLOUD_CONFIG_ENDPOINT|
|Direct Link|IBMCLOUD_DL_API_ENDPOINT|
|Direct Link Provider|IBMCLOUD_DL_PROVIDER_API_ENDPOINT|
|Enterprise Management|IBMCLOUD_ENTERPRISE_API_ENDPOINT|
|Cloud Functions|IBMCLOUD_FUNCTIONS_API_ENDPOINT|
|Global Tagging|IBMCLOUD_GT_API_ENDPOINT|
|Global Search|IBMCLOUD_GS_API_ENDPOINT|
|Hyper Protect Crypto Services|IBMCLOUD_HPCS_API_ENDPOINT|
|Hyper Protect Crypto Services TKE Endpoint|IBMCLOUD_HPCS_TKE_ENDPOINT|
|Identity and Access Management|IBMCLOUD_IAM_API_ENDPOINT|
|Cloud Databases|IBMCLOUD_ICD_API_ENDPOINT|
|Virtual Private Cloud (VPC)|IBMCLOUD_IS_NG_API_ENDPOINT|
|Key Management Services|IBMCLOUD_KP_API_ENDPOINT|
|Cloud Foundry|IBMCLOUD_MCCP_API_ENDPOINT|
|Push Notifications|IBMCLOUD_PUSH_API_ENDPOINT|
|Private DNS|IBMCLOUD_PRIVATE_DNS_API_ENDPOINT|
|Resource Controller|IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT|
|Resource Manager|IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT|
|Global Catalog|IBMCLOUD_RESOURCE_CATALOG_API_ENDPOINT|
|Satellite|IBMCLOUD_SATELLITE_API_ENDPOINT|
|Satellite Link|IBMCLOUD_SATELLITE_LINK_API_ENDPOINT|
|Schematics|IBMCLOUD_SCHEMATICS_API_ENDPOINT|
|Secrets Manager|IBMCLOUD_SECRETS_MANAGER_API_ENDPOINT|
|Transit Gateway|IBMCLOUD_TG_API_ENDPOINT|
|UAA|IBMCLOUD_UAA_ENDPOINT|
|User Management|IBMCLOUD_USER_MANAGEMENT_ENDPOINT|
|Event Notifications|IBMCLOUD_EVENT_NOTIFICATIONS_API_ENDPOINT|
|Logs Routing|IBMCLOUD_LOGS_ROUTING_API_ENDPOINT|

## File structure for endpoints file

To use public and private regional endpoints for a service, you must add these endpoints to a JSON file and categorize them as public or private service endpoints. 

**Syntax**: 

```json
{
    "<endpoint_variable>":{
        "<public_or_private>":{
            "<region>":"<service endpoint>"
        }
    }
}
```

**Example**:

```json
{
    "IBMCLOUD_API_GATEWAY_ENDPOINT":{
        "public":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        },
        "private":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        }
    },
    "IBMCLOUD_ACCOUNT_MANAGEMENT_API_ENDPOINT":{
        "public":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        },
        "private":{
            "us-south":"<endpoint>",
            "us-east":"<endpoint>",
            "eu-gb":"<endpoint>",
            "eu-de":"<endpoint>"
        }
    }
}
```
**Note:** 

The endpoints file accepts "public", "private" and "public-and-private" as visibility while COS resources support "public", "private" and "direct as endpoint-types. 
Since endpoints file schema does not supprt "direct", users must define the url for "direct" endpoint-type under exisiting visibility type "private" for "IBMCLOUD_COS_CONFIG_ENDPOINT" and "IBMCLOUD_COS_ENDPOINT".
The user cannot define urls for both private and direct endpoint-type simultaneously in the endpoints file under "private" field. 

**Example**:

```json
{
    "IBMCLOUD_COS_CONFIG_ENDPOINT":{
        "public":{
            "us-south":"https://config.cloud-object-storage.cloud.ibm.com/v1"
        },
        "private":{
            "us-south":"https://config.direct.cloud-object-storage.cloud.ibm.com/v1"
        }
    }
}
```

OR 

```json
{
    "IBMCLOUD_COS_CONFIG_ENDPOINT":{
        "public":{
            "us-south":"https://config.cloud-object-storage.cloud.ibm.com/v1"
        },
        "private":{
            "us-south":"https://config.private.cloud-object-storage.cloud.ibm.com/v1"
        }
    }
}
```


## Prioritisation of endpoints

The IBM Cloud Provider plug-in gives the following prioritisation 

1. Endpoints defined by using environment variables
2. Endpoints defined by using the `endpoints_file_path` argument in the provider block
3. Default private or public service endpoints based on the `visibility` argument in the provider block 

### 1. Define service endpoints by using environment variables

The IBM Cloud Provider plug-in gives highest priority to the exported environment variables. To find the environment variable name that you need to export, see **Supportd endpoint customizations**. If an environment variable is exported, the provider uses the defined endpoint URL to connect to the IBM Cloud service. Additional configurations that you made in the provider block, such as the `visibility` or `endpoints_file_path` arguments, are ignored. 

1. Specify your provider block with or without the `visibility` and `endpoints_file_path` arguments. 
   ```terraform
   provider "ibm" {
    # ... other provider configuration ...
   }
   ```

2. Export the environment variable for your IBM Cloud service and set it to the IBM Cloud service endpoint that you want to use. 
   ```text
   export IBMCLOUD_API_GATEWAY_ENDPOINT="<endpoint_url>" 
   ```
   
3. Initialize the Terraform CLI. The IBM Cloud Provider plug-in automatically loads the environment variables. 
4. Run other Terraform commands, such as `terraform plan` or `terraform apply`. 


### 2. Define service endpoints by using an endpoints file 

You can declare all your service endpoints in a JSON file and either reference this file in your provider block by using the `endpoints_file_path` argument, or export the path to your file with the `IBMCLOUD_ENDPOINTS_FILE_PATH` or `IC_ENDPOINTS_FILE_PATH` environment variable. The endpoints file can include private and public service endpoints, and you can also specify different endpoints for each region. Depending on the `visibility` and `region` settings in your provider block, the IBM Cloud Provider plug-in determines the endpoint from the endpoint file that you want to use.  

**Note:**  

- Use the `endpoints_file_path` argument to reference the endpoints file in your provider block. 
- Use the `IBMCLOUD_ENDPOINTS_FILE_PATH` or `IC_ENDPOINTS_FILE_PATH` environment variable to export the path to your endpoints file.
- Use the `visibility` argument along with the `endpoints_file_path` in the provider block to determine the `public` and `private` endpoints.
- Supported values for the `visibility` argument when the `endpoints_file_path` argument is set, include `public` and `private`. Default value: `public. 

**Syntax for referencing the endpoints file in the provider block**: 

```terraform
    provider "ibm" {
        # ... other provider configuration ...
        endpoints_file_path = "<file_path_to_endpoints_file_path>"
        visibility     = "<private_or_public>"
    }
```

**Syntax for exporting the path to the endpoints file as an environment variable**: 

1. Specify your provider block with or without the `visibility` and `endpoint_file` arguments. 
   ```terraform
    provider "ibm" {
        # ... other provider configuration ...
    }

   ```
2. Export the path to your endpoints file as an environment variable. You can also declare if you want to use the private or public service endpoint from your endpoints file by using the `IC_VISIBILITY` environment variable.  
   ```text
   export IC_ENDPOINTS_FILE_PATH="<endpoint_value>"
   export IC_VISIBILITY="<private_or_public>"
   ```

### 3. Use the default private or public service endpoint based on the `visibility` setting in the provider block 

If for a given `region` and `visibility` setting in your provider block, the IBM Cloud Provider plug-in cannot find an environment variable or an endpoint in your endpoints file, the default service endpoint that is implemented in the IBM Cloud Provider plug-in is used. 

**Note:** In order to use the private endpoint from an IBM Cloud resource, you must have a VRF-enabled IBM cloudaccount. If the service does not support private endpoints, the Terraform resource or datas ource will log an error.

- Supported values for the `visibility` argument are `public`, `private`, `public-and-private`. Default value: `public`.
  - If the visibility is set to `public`, the provider uses a regional public endpoint or the global public endpoint. The regional public endpoints has higher precedence.
  - If the visibility is set to `private`, the. provider uses a regional private endpoint or the global private endpoint. The regional private endpoint is given higher precedence.  
  - If the visibility is set to `public-and-private`, the provider uses the regional private endpoints or the global private endpoint. If the service does not support regional or global private endpoints, the provider uses the regional or global public endpoint.
- You can set the visibility by using the `IC_VISIBILITY` (higher precedence) or `IBMCLOUD_VISIBILITY` environment variable.

**Syntax for using the `visibility` argument in the provider block**:

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
        visibility     = "private"
    }
```

**Syntax for setting the visibility as an environment variable**: 

```terraform
    provider "ibm" {
        # ... potentially other provider configuration ...
    }
```

```text
export IC_VISIBILITY="private" or export IC_VISIBILITY="public-and-private"
```

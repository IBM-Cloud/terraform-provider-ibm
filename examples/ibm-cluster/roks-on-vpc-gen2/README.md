# ROKS on VPC Gen-2 Example

## Introduction
This Example enables user to create a RedHat Openshift cluster on the IBM CLOUD VPC gen-2 Infrastructure.

## Terraform Versions
Terraform v0.12. and terraform-provider-ibm version to `~> v1.7.2`. Branch - `master`.

## Running the configuration
```shell
terraform init
terraform plan
```

For apply phase

```shell
terraform apply
```
 
 For destroy phase

```shell
terraform destroy
```  

## Dependencies
- User needs to create Cloud Object Store instance before the Openshift cluster creation in VPC gen-2 infrastructure. COS instance created will be used to back up the internal registry in the OpenShift on VPC Gen 2 cluster.
- User needs to create security group rule in the Default security group of VPC to enable port 30000 to 32767(NodePorts) to get the traffic into the cluster.


## Notes
- terraform-provider-ibm provides entitlement parameter in both `ibm_container_vpc_cluster` and `ibm_container_vpc_worker_pool` resources to reduce the OCP licence cost occuring in cluster creation.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name      | Version       |
|-----------|---------------|
| terraform | ~> v0.12      |

## Providers

| Name | Version   |
|------|-----------|
| ibm  | ~> v1.7.2 |

## Configurations
The following variable needs to be set in the variables.tf
* `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.

## Inputs

### Module - cos
| Name                   | Description                                            | Type    | Required |
|------------------------|--------------------------------------------------------|---------|----------|
|service_instance_name   | COS instance name                                      |`string` |   yes    |
|plan                    | COS service plan                                       |`string` |   yes    |

### Module - cluster_and_workerpool
| Name                     | Description                                               | Type    | Required |
|--------------------------|-----------------------------------------------------------|---------|----------|
|flavor                    | Cluster node Flavor                                       |`string` |   yes    |
|kube_version              | Openshift cluster kube version                            |`string` |   yes    |
|worker_count              | Number of workers needs to be added in deafult workerpool |`int`    |   yes    |
|region                    | CLuster depolyment region                                 |`string` |   yes    |
|resource_group            | Resouurce group name                                      |`string` |   yes    |
|cluster_name              | Cluster name                                              |`string` |   yes    |
|worker_pool_name          | Worker pool name                                          |`string` |   yes    |
|worker_pool_workers_count | Number of Worker nodes in the worker pool                 |`int`    |   yes    |
|cos_instance_crn          | CRN id of the COS instance created                        |`string` |   yes    |
|entitlement               | Openshift Cluster entitlement                             |`string` |   yes    |

## Outputs

### Module - cos
| Name             | Description                               |
|------------------|-------------------------------------------|
|cos_instance_crn  | CRN id of the COS instacne created        |

### Module - cluster_and_workerpool
| Name             | Description                               |
|------------------|-------------------------------------------|
|cluster_id        | Id of the Cluster created                 |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## References

1. [Conainer Cluster Bluemix-go SDK](https://github.com/Mavrickk3/bluemix-go/tree/master/api/container/containerv2)
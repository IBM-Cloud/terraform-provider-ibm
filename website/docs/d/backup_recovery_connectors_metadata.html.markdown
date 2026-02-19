---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_connectors_metadata"
description: |-
  Get information about backup_recovery_connectors_metadata
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_connectors_metadata

Provides a read-only data source to retrieve information about a backup_recovery_connectors_metadata. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_connectors_metadata" "backup_recovery_connectors_metadata" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_connectors_metadata.
* `connector_image_metadata` - (List) Specifies information about the connector images for various platforms.
Nested schema for **connector_image_metadata**:
	* `connector_image_file_list` - (List) Specifies info about connector images for the supported platforms.
	Nested schema for **connector_image_file_list**:
		* `image_type` - (String) Specifies the platform on which the image can be deployed.
		  * Constraints: Allowable values are: `VSI`, `VMware`.
		* `url` - (String) Specifies the URL to access the file.
* `k8s_connector_info_list` - (List) k8sConnectorInfoList specifies information about supported kubernetes environments where Data-Source Connectors can be deployed. Also, specifies the helm chart location (OCI URL) for each supported Kubernetes environment and instructions for installing it.
Nested schema for **k8s_connector_info_list**:
	* `helm_chart_oci_ref` - (List) Represents the structured components of an OCI (Open Container Initiative) artifact reference. A full reference string can be constructed from these parts. See Also: https://github.com/opencontainers/distribution-spec/blob/main/spec.md.
	Nested schema for **helm_chart_oci_ref**:
		* `digest` - (String) The immutable, content-addressable digest of the artifact's manifest. If only digest is set, the artifact is fetched by its immutable reference. If both tag and digest are set, the application should verify that the tag resolves to the given digest before proceeding. This should include the algorithm prefix.
		* `namespace` - (String) The namespace or organization within the registry. For public registries like Docker Hub, this can be 'library' for official images or a user's account name. May be optional for certain registry configurations.
		* `registry_host` - (String) The address of the OCI-compliant container registry. This can be a hostname or an IP address, and may optionally include a port number.
		* `repository` - (String) The name of the repository that holds the artifact.
		* `tag` - (String) The mutable tag for the artifact.
	* `helm_install_cmd` - (String) Specifies the Helm install command for this type of k8s connector.
	* `k8s_platform_type` - (String) Enum representing the different supported Kubernetes platform types.
	  * Constraints: Allowable values are: `kRoksVpc`, `kRoksClassic`, `kIksVpc`, `kIksClassic`.
	* `ugrade_doc_url` - (String) URL for upgrade documentation for this type of k8s connector.


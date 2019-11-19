package helpers

const (
	// IBM PI Instance

	PIInstanceName          = "pi_instance_name"
	PIInstanceDate          = "pi_creation_date"
	PIInstanceSSHKeyName    = "pi_key_pair_name"
	PIInstanceImageName     = "pi_image_id"
	PIInstanceProcessors    = "pi_processors"
	PIInstanceProcType      = "pi_proc_type"
	PIInstanceMemory        = "pi_memory"
	PIInstanceSystemType    = "pi_sys_type"
	PIInstanceId            = "pi_instance_id"
	PIInstanceDiskSize      = "pi_disk_size"
	PIInstanceStatus        = "pi_instance_status"
	PIInstanceMinProc       = "pi_minproc"
	PIInstanceVolumeIds     = "pi_volume_ids"
	PIInstanceNetworkIds    = "pi_network_ids"
	PIInstancePublicNetwork = "pi_public_network"
	PIInstanceMigratable    = "pi_migratable"
	PICloudInstanceId       = "pi_cloud_instance_id"

	PIInstanceHealthStatus      = "pi_health_status"
	PIInstanceReplicants        = "pi_replicants"
	PIInstanceReplicationPolicy = "pi_replication_policy"
	PIInstanceReplicationScheme = "pi_replication_scheme"
	PIInstanceProgress          = "pi_progress"
	PIInstanceUserData          = "pi_user_data"

	// IBM PI Volume
	PIVolumeName      = "pi_volume_name"
	PIVolumeSize      = "pi_volume_size"
	PIVolumeType      = "pi_volume_type"
	PIVolumeShareable = "pi_volume_shareable"
	PIVolumeId        = "pi_volume_id"

	// IBM PI Image

	PIImageName         = "pi_image_name"
	PIImageQueStatus    = "queued"
	PIImageActiveStatus = "active"

	// IBM PI Key

	PIKeyName = "pi_key_name"
	PIKey     = "pi_ssh_key"
	PIKeyDate = "pi_creation_date"
	PIKeyId   = "pi_key_id"

	// IBM PI Network

	PINetworkReady          = "ready"
	PINetworkID             = "pi_networkid"
	PINetworkName           = "pi_network_name"
	PINetworkCidr           = "pi_cidr"
	PINetworkDNS            = "pi_dns"
	PINetworkType           = "pi_network_type"
	PINetworkGateway        = "pi_gateway"
	PINetworkIPAddressRange = "pi_ipaddress_range"
	PINetworkVlanId         = "pi_vlan_id"
	PINetworkProvisioning   = "build"

	// IBM PI Operations
	PIInstanceOperationType       = "pi_operation"
	PIInstanceOperationProgress   = "pi_progress"
	PIInstanceOperationStatus     = "pi_status"
	PIInstanceOperationServerName = "pi_instance_name"

	// IBM PI Volume Attach
	PIVolumeAttachName                = "pi_volume_attach_name"
	PIVolumeAllowableAttachStatus     = "in-use"
	PIVolumeAttachStatus              = "status"
	PowerVolumeAttachDeleting         = "deleting"
	PowerVolumeAttachProvisioning     = "creating"
	PowerVolumeAttachProvisioningDone = "available"

	// IBM PI Instance Capture
	PIInstanceCaptureName                  = "pi_capture_name"
	PIInstanceCaptureDestination           = "pi_capture_destination"
	PIInstanceCaptureVolumeIds             = "pi_capture_volume_ids"
	PIInstanceCaptureCloudStorageImagePath = "pi_capture_storage_image_path"
	PIInstanceCaptureCloudStorageRegion    = "pi_capture_cloud_storage_region"
	PIInstanceCaptureCloudStorageAccessKey = "pi_capture_cloud_storage_access_key"
	PIInstanceCaptureCloudStorageSecretKey = "pi_capture_cloud_storage_secret_key"

	// Status For all the resources

	PIVolumeStatus           = "pi_volume_status"
	PIVolumeDeleting         = "deleting"
	PIVolumeDeleted          = "done"
	PIVolumeProvisioning     = "creating"
	PIVolumeProvisioningDone = "available"
	PIInstanceAvailable      = "ACTIVE"
	PIInstanceHealthOk       = "OK"
	PIInstanceHealthWarning  = "WARNING"
	PIInstanceBuilding       = "BUILD"
	PIInstanceDeleting       = "DELETING"
	PIInstanceNotFound       = "Not Found"
)

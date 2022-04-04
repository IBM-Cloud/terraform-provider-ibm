package power

import "time"

const (

	// Keys
	PIKeys    = "keys"
	PIKeyName = "name"
	PIKey     = "ssh_key"
	PIKeyDate = "creation_date"

	// SAP Profile
	PISAPProfiles         = "profiles"
	PISAPProfileCertified = "certified"
	PISAPProfileCores     = "cores"
	PISAPProfileMemory    = "memory"
	PISAPProfileID        = "profile_id"
	PISAPProfileType      = "type"

	// DHCP
	PIDhcpStatusBuilding = "Building"
	PIDhcpStatusActive   = "ACTIVE"
	PIDhcpDeleting       = "Deleting"
	PIDhcpDeleted        = "Deleted"
	PIDhcpId             = "dhcp_id"
	PIDhcpStatus         = "status"
	PIDhcpNetwork        = "network"
	PIDhcpLeases         = "leases"
	PIDhcpInstanceIp     = "instance_ip"
	PIDhcpInstanceMac    = "instance_mac"

	// Instance
	//Added timeout values for warning  and active status
	warningTimeOut = 60 * time.Second
	activeTimeOut  = 2 * time.Minute
	// power service instance capabilities
	CUSTOM_VIRTUAL_CORES          = "custom-virtualcores"
	PIInstanceNetwork             = "pi_network"
	PIInstanceStoragePool         = "pi_storage_pool"
	PISAPInstanceProfileID        = "pi_sap_profile_id"
	PIInstanceStoragePoolAffinity = "pi_storage_pool_affinity"

	// Placement Group
	PIPlacementGroupID      = "placement_group_id"
	PIPlacementGroupMembers = "members"

	// Volume
	PIAffinityPolicy        = "pi_affinity_policy"
	PIAffinityVolume        = "pi_affinity_volume"
	PIAffinityInstance      = "pi_affinity_instance"
	PIAntiAffinityInstances = "pi_anti_affinity_instances"
	PIAntiAffinityVolumes   = "pi_anti_affinity_volumes"

	// VPN
	PIVPNConnectionId                         = "connection_id"
	PIVPNConnectionStatus                     = "connection_status"
	PIVPNConnectionDeadPeerDetection          = "dead_peer_detections"
	PIVPNConnectionDeadPeerDetectionAction    = "action"
	PIVPNConnectionDeadPeerDetectionInterval  = "interval"
	PIVPNConnectionDeadPeerDetectionThreshold = "threshold"
	PIVPNConnectionLocalGatewayAddress        = "local_gateway_address"
	PIVPNConnectionVpnGatewayAddress          = "gateway_address"

	PICloudInstanceID = "pi_cloud_instance_id"

	// ---------- Image ----------
	// Arguments
	PIImageName                  = "pi_image_name"
	PIImageAffinityInstance      = "pi_affinity_instance"
	PIImageAffinityPolicy        = "pi_affinity_policy"
	PIImageAffinityVolume        = "pi_affinity_volume"
	PIImageAntiAffinityInstances = "pi_anti_affinity_instances"
	PIImageAntiAffinityVolumes   = "pi_anti_affinity_volumes"
	PIImageID                    = "pi_image_id"
	PIImageBucketName            = "pi_image_bucket_name"
	PIImageAccessKey             = "pi_image_access_key"
	PIImageBucketAccess          = "pi_image_bucket_access"
	PIImageBucketFile            = "pi_image_bucket_file_name"
	PIImageBucketRegion          = "pi_image_bucket_region"
	PIImageSecretKey             = "pi_image_secret_key"
	PIImageStoragePool           = "pi_image_storage_pool"
	PIImageStorageType           = "pi_image_storage_type"
	CatalogImagesSAP             = "sap"
	CatalogImagesVTL             = "vtl"

	// Attributes
	Images               = "image_info"
	CatalogImages        = "images"
	ImagesID             = "id"
	ImageName            = "name"
	ImageID              = "image_id"
	ImageArchitecture    = "architecture"
	ImageOperatingSystem = "operatingsystem"
	ImageSize            = "size"
	ImageState           = "state"
	ImageHyperVisor      = "hypervisor"
	ImageStorageType     = "storage_type"
	ImageStoragePool     = "storage_pool"
	ImageType            = "image_type"
	ImageCreationDate    = "creation_date"
	ImageDescription     = "description"
	ImageDiskFormat      = "disk_format"
	ImageEndianness      = "endianness"
	ImageHref            = "href"
	ImageLastUpdateDate  = "last_update_date"
	ImageContainerFormat = "container_format"

	// Misc
	ImageRetry  = "retry"
	ImageQueued = "queued"
	ImageActive = "active"

	// Attributes need to fix
	ImageHypervisorType         = "hypervisor_type"
	CatalogImageOperatingSystem = "operating_system"
	// ---------------------------
)

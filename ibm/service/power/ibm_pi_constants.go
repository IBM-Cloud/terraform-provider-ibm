package power

import "time"

const (

	// -- Capture -------------------------------------------------------------
	// Arguments
	Arg_CaptureName             = "pi_capture_name"
	Arg_CaptureDestination      = "pi_capture_destination"
	Arg_CaptureVolumeIDs        = "pi_capture_volume_ids"
	Arg_CaptureStorageRegion    = "pi_capture_cloud_storage_region"
	Arg_CaptureStorageAccessKey = "pi_caputre_cloud_storage_access_key"
	Arg_CaptureStorageSecretKey = "pi_caputre_cloud_storage_secret_key"
	Arg_CaptureStorageImagePath = "pi_capture_storage_image_path"

	// Attributes
	Attr_CaptureImageID = "image_id"

	// Misc
	CaptureDestinationBoth  = "both"
	CaptureDestinationCloud = "cloud-storage"
	CaptureDestinationImage = "image-catalog"

	// -- Cloud Connection ----------------------------------------------------

	// -- Cloud Instance ------------------------------------------------------
	Arg_CloudInstanceID = "pi_cloud_instance_id"

	// -- Console Language ----------------------------------------------------

	// -- DHCP ----------------------------------------------------------------
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

	// -- Image ---------------------------------------------------------------

	// -- Instance ------------------------------------------------------------
	Arg_InstanceName = "pi_instance_name"

	//Added timeout values for warning  and active status
	warningTimeOut = 60 * time.Second
	activeTimeOut  = 2 * time.Minute
	// power service instance capabilities
	CUSTOM_VIRTUAL_CORES          = "custom-virtualcores"
	PIInstanceNetwork             = "pi_network"
	PIInstanceStoragePool         = "pi_storage_pool"
	PISAPInstanceProfileID        = "pi_sap_profile_id"
	PIInstanceStoragePoolAffinity = "pi_storage_pool_affinity"

	// -- Key -----------------------------------------------------------------
	PIKeys    = "keys"
	PIKeyName = "name"
	PIKey     = "ssh_key"
	PIKeyDate = "creation_date"

	// -- Network -------------------------------------------------------------

	// -- Operations ----------------------------------------------------------

	// -- Placement Group -----------------------------------------------------
	PIPlacementGroupID      = "placement_group_id"
	PIPlacementGroupMembers = "members"

	// -- SAP -----------------------------------------------------------------
	PISAPProfiles         = "profiles"
	PISAPProfileCertified = "certified"
	PISAPProfileCores     = "cores"
	PISAPProfileMemory    = "memory"
	PISAPProfileID        = "profile_id"
	PISAPProfileType      = "type"

	// -- Snapshot ------------------------------------------------------------

	// -- Storage Pool --------------------------------------------------------

	// -- Storage Type --------------------------------------------------------

	// -- Tenant --------------------------------------------------------------

	// -- Volume --------------------------------------------------------------
	PIAffinityPolicy        = "pi_affinity_policy"
	PIAffinityVolume        = "pi_affinity_volume"
	PIAffinityInstance      = "pi_affinity_instance"
	PIAntiAffinityInstances = "pi_anti_affinity_instances"
	PIAntiAffinityVolumes   = "pi_anti_affinity_volumes"

	// -- VPN -----------------------------------------------------------------
	PIVPNConnectionId                         = "connection_id"
	PIVPNConnectionStatus                     = "connection_status"
	PIVPNConnectionDeadPeerDetection          = "dead_peer_detections"
	PIVPNConnectionDeadPeerDetectionAction    = "action"
	PIVPNConnectionDeadPeerDetectionInterval  = "interval"
	PIVPNConnectionDeadPeerDetectionThreshold = "threshold"
	PIVPNConnectionLocalGatewayAddress        = "local_gateway_address"
	PIVPNConnectionVpnGatewayAddress          = "gateway_address"

	// -- VPN Policy ----------------------------------------------------------

	// Health
	HealthOk = "OK"
)

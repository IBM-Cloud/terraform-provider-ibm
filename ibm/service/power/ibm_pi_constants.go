package power

import "time"

const (

	// -- Capture -------------------------------------------------------------

	// -- Cloud Connection ----------------------------------------------------

	// -- Cloud Instance ------------------------------------------------------
	// Required Arguments
	Arg_CloudInstanceID = "pi_cloud_instance_id"

	// Attributes
	Attr_CloudInstanceEnabled         = "enabled"
	Attr_CloudInstanceTenant          = "tenant_id"
	Attr_CloudInstanceRegion          = "region"
	Attr_CloudInstanceCapabilities    = "capabilities"
	Attr_CloudInstanceTotalProcessors = "total_processors_consumed"
	Attr_CloudInstanceTotalInstances  = "total_instances"
	Attr_CloudInstanceTotalMemory     = "total_memory_consumed"
	Attr_CloudInstanceTotalSSD        = "total_ssd_storage_consumed"
	Attr_CloudInstanceTotalStorage    = "total_standard_storage_consumed"
	Attr_CloudInstanceInstances       = "pvm_instances"

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

	//Added timeout values for warning  and active status
	warningTimeOut = 60 * time.Second
	activeTimeOut  = 2 * time.Minute
	// power service instance capabilities
	CUSTOM_VIRTUAL_CORES          = "custom-virtualcores"
	PIInstanceNetwork             = "pi_network"
	PIInstanceStoragePool         = "pi_storage_pool"
	PISAPInstanceProfileID        = "pi_sap_profile_id"
	PIInstanceStoragePoolAffinity = "pi_storage_pool_affinity"

	// Attributes
	Attr_InstanceID           = "id"
	Attr_InstanceName         = "name"
	Attr_InstanceHref         = "href"
	Attr_InstanceStatus       = "status"
	Attr_InstanceSysType      = "systype"
	Attr_InstanceCreationDate = "creation_date"

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

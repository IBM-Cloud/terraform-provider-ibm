package power

import "time"

const (

	// -- Capture -------------------------------------------------------------

	// -- Cloud Connection ----------------------------------------------------
	// Required Arguments
	Arg_CloudConnectionID        = "pi_cloud_connection_id"
	Arg_CloudConnectionName      = "pi_cloud_connection_name"
	Arg_CloudConnectionSpeed     = "pi_cloud_connection_speed"
	Arg_CloudConnectionNetworkID = "pi_network_id"

	// Optional Arguments
	Arg_CloudConnectionRouting  = "pi_cloud_connection_global_routing"
	Arg_CloudConnectionMetered  = "pi_cloud_connection_metered"
	Arg_CloudConnectionNetworks = "pi_cloud_connection_networks"
	Arg_CloudConnectionClassic  = "pi_cloud_connection_classic_enabled"
	Arg_CloudConnectionGreCIDR  = "pi_cloud_connection_gre_cidr"
	Arg_CloudConnectionGreDest  = "pi_cloud_connection_gre_destination_address"
	Arg_CloudConnectionVPC      = "pi_cloud_connection_vpc_enabled"
	Arg_CloudConnectionVPCCrns  = "pi_cloud_connection_vpc_crns"
	Arg_CloudConnection         = "pi_cloud_connection_"

	// Attributes
	Attr_CloudConnectionSpeed         = "speed"
	Attr_CloudConnectionID            = "cloud_connection_id"
	Attr_CloudConnectionStatus        = "status"
	Attr_CloudConnectionIbmIP         = "ibm_ip_address"
	Attr_CloudConnectionUserIP        = "user_ip_address"
	Attr_CloudConnectionPort          = "port"
	Attr_CloudConnectionSourceGreAddr = "gre_source_address"
	AttrCloudConnectionGreDestAddr    = "gre_destination_address"
	Attr_CloudConnectionRouting       = "global_routing"
	Attr_CloudConnectionMetered       = "metered"
	Attr_CloudConnectionNetworks      = "networks"
	Attr_CloudConnectionClassic       = "classic_enabled"
	Attr_CloudConnectionVPC           = "vpc_enabled"
	Attr_CloudConnectionVPCCrns       = "vpc_crns"
	Attr_CloudConnectionName          = "name"

	// need to fix

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

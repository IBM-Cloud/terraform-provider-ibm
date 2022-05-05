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
)

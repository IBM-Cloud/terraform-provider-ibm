package ibm

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var cfOrganization string
var cfSpace string
var ibmid1 string
var ibmid2 string
var IAMUser string
var datacenter string
var machineType string
var trustedMachineType string
var publicVlanID string
var privateVlanID string
var privateSubnetID string
var publicSubnetID string
var subnetID string
var lbaasDatacenter string
var lbaasSubnetId string
var dedicatedHostName string
var dedicatedHostID string
var kubeVersion string
var kubeUpdateVersion string
var zone string
var zonePrivateVlan string
var zonePublicVlan string
var zoneUpdatePrivateVlan string
var zoneUpdatePublicVlan string
var csRegion string
var extendedHardwareTesting bool
var err error

func init() {
	cfOrganization = os.Getenv("IBM_ORG")
	if cfOrganization == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ORG for testing ibm_org  resource Some tests for that resource will fail if this is not set correctly")
	}
	cfSpace = os.Getenv("IBM_SPACE")
	if cfSpace == "" {
		fmt.Println("[WARN] Set the environment variable IBM_SPACE for testing ibm_space  resource Some tests for that resource will fail if this is not set correctly")
	}
	ibmid1 = os.Getenv("IBM_ID1")
	if ibmid1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID1 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	ibmid2 = os.Getenv("IBM_ID2")
	if ibmid2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID2 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	IAMUser = os.Getenv("IBM_IAMUSER")
	if IAMUser == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAMUSER for testing ibm_iam_user_policy resource Some tests for that resource will fail if this is not set correctly")
	}

	datacenter = os.Getenv("IBM_DATACENTER")
	if datacenter == "" {
		datacenter = "par01"
		fmt.Println("[WARN] Set the environment variable IBM_DATACENTER for testing ibm_container_cluster resource else it is set to default value 'par01'")
	}

	machineType = os.Getenv("IBM_MACHINE_TYPE")
	if machineType == "" {
		machineType = "u1c.2x4"
		fmt.Println("[WARN] Set the environment variable IBM_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'u1c.2x4'")
	}

	csRegion = os.Getenv("IBM_CONTAINER_REGION")
	if csRegion == "" {
		csRegion = "eu-de"
		fmt.Println("[WARN] Set the environment variable IBM_CONTAINER_REGION for testing ibm_container resources else it is set to default value 'eu-de'")
	}

	trustedMachineType = os.Getenv("IBM_TRUSTED_MACHINE_TYPE")
	if trustedMachineType == "" {
		trustedMachineType = "mb1c.16x64"
		fmt.Println("[WARN] Set the environment variable IBM_TRUSTED_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'mb1c.16x64'")
	}

	extendedHardwareTesting, err = strconv.ParseBool(os.Getenv("IBM_BM_EXTENDED_HW_TESTING"))
	if err != nil {
		extendedHardwareTesting = false
		fmt.Println("[WARN] Set the environment variable IBM_BM_EXTENDED_HW_TESTING to true/false for testing ibm_compute_bare_metal resource else it is set to default value 'false'")
	}

	publicVlanID = os.Getenv("IBM_PUBLIC_VLAN_ID")
	if publicVlanID == "" {
		publicVlanID = "2393319"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393319'")
	}

	privateVlanID = os.Getenv("IBM_PRIVATE_VLAN_ID")
	if privateVlanID == "" {
		privateVlanID = "2393321"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393321'")
	}

	kubeVersion = os.Getenv("IBM_KUBE_VERSION")
	if kubeVersion == "" {
		kubeVersion = "1.9.11"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.9.11'")
	}

	kubeUpdateVersion = os.Getenv("IBM_KUBE_UPDATE_VERSION")
	if kubeUpdateVersion == "" {
		kubeUpdateVersion = "1.10.8"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_UPDATE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.10.8'")
	}

	privateSubnetID = os.Getenv("IBM_PRIVATE_SUBNET_ID")
	if privateSubnetID == "" {
		privateSubnetID = "1636107"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1636107'")
	}

	publicSubnetID = os.Getenv("IBM_PUBLIC_SUBNET_ID")
	if publicSubnetID == "" {
		publicSubnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	subnetID = os.Getenv("IBM_SUBNET_ID")
	if subnetID == "" {
		subnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	lbaasDatacenter = os.Getenv("IBM_LBAAS_DATACENTER")
	if lbaasDatacenter == "" {
		lbaasDatacenter = "wdc04"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_DATACENTER for testing ibm_lbaas resource else it is set to default value 'wdc04'")
	}

	lbaasSubnetId = os.Getenv("IBM_LBAAS_SUBNETID")
	if lbaasSubnetId == "" {
		lbaasSubnetId = "1511875"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_SUBNETID for testing ibm_lbaas resource else it is set to default value '1511875'")
	}

	dedicatedHostName = os.Getenv("IBM_DEDICATED_HOSTNAME")
	if dedicatedHostName == "" {
		dedicatedHostName = "terraform-dedicatedhost"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOSTNAME for testing ibm_compute_vm_instance resource else it is set to default value 'terraform-dedicatedhost'")
	}

	dedicatedHostID = os.Getenv("IBM_DEDICATED_HOST_ID")
	if dedicatedHostID == "" {
		dedicatedHostID = "30301"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOST_ID for testing ibm_compute_vm_instance resource else it is set to default value '30301'")
	}

	zone = os.Getenv("IBM_WORKER_POOL_ZONE")
	if zone == "" {
		zone = "ams03"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value 'ams03'")
	}

	zonePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PRIVATE_VLAN")
	if zonePrivateVlan == "" {
		zonePrivateVlan = "1764491"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '1764491'")
	}

	zonePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PUBLIC_VLAN")
	if zonePublicVlan == "" {
		zonePublicVlan = "1764435"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '1764435'")
	}

	zoneUpdatePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN")
	if zoneUpdatePrivateVlan == "" {
		zoneUpdatePrivateVlan = "2388377"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388377'")
	}

	zoneUpdatePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN")
	if zoneUpdatePublicVlan == "" {
		zoneUpdatePublicVlan = "2388375"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388375'")
	}

}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ibm": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("BM_API_KEY"); v == "" {
		t.Fatal("BM_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("SL_API_KEY"); v == "" {
		t.Fatal("SL_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("SL_USERNAME"); v == "" {
		t.Fatal("SL_USERNAME must be set for acceptance tests")
	}
}

package ibm

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func TestAccIBMComputeBareMetal_Basic(t *testing.T) {
	var bareMetal datatypes.Hardware
	configName := "ibm_compute_bare_metal.terraform-acceptance-test-1"
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMComputeBareMetalConfig_basic(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configName, &bareMetal),
					resource.TestCheckResourceAttr(
						configName, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configName, "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						configName, "os_reference_code", "UBUNTU_16_64"),
					resource.TestCheckResourceAttr(
						configName, "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						configName, "network_speed", "100"),
					resource.TestCheckResourceAttr(
						configName, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configName, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configName, "user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						configName, "fixed_config_preset", "S1270_32GB_1X1TBSATA_NORAID"),
					resource.TestCheckResourceAttr(
						configName, "notes", "baremetal notes"),
					CheckStringSet(
						configName,
						"tags", []string{"collectd"},
					),
				),
			},

			{
				Config:  testAccCheckIBMComputeBareMetalConfig_update(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configName, &bareMetal),
					CheckStringSet(
						configName,
						"tags", []string{"mesos-master"},
					),
				),
			},
		},
	})
}

func TestAccIBMComputeBareMetal_With_IPV6(t *testing.T) {
	var bareMetal datatypes.Hardware
	configName := "ibm_compute_bare_metal.terraform-acceptance-test-1"
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMComputeBareMetalConfig_with_ipv6(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configName, &bareMetal),
					resource.TestCheckResourceAttr(
						configName, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configName, "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						configName, "os_reference_code", "UBUNTU_16_64"),
					resource.TestCheckResourceAttr(
						configName, "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						configName, "network_speed", "100"),
					resource.TestCheckResourceAttr(
						configName, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configName, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configName, "ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						configName, "ipv6_static_enabled", "true"),
					resource.TestCheckResourceAttr(
						configName, "secondary_ip_count", "4"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.1"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.2"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.3"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.4"),
					resource.TestCheckResourceAttr(
						configName, "user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						configName, "fixed_config_preset", "S1270_32GB_1X1TBSATA_NORAID"),
					resource.TestCheckResourceAttr(
						configName, "notes", "baremetal notes"),
					CheckStringSet(
						configName,
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func TestAccIBMComputeBareMetal_With_Unbonded_Port_Speed(t *testing.T) {
	var bareMetal datatypes.Hardware
	configName := "ibm_compute_bare_metal.terraform-acceptance-test-1"
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMComputeBareMetalConfig_with_unbonded_port_speed(hostname),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configName, &bareMetal),
					resource.TestCheckResourceAttr(
						configName, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configName, "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						configName, "os_reference_code", "UBUNTU_16_64"),
					resource.TestCheckResourceAttr(
						configName, "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						configName, "network_speed", "1000"),
					resource.TestCheckResourceAttr(
						configName, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configName, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configName, "user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						configName, "fixed_config_preset", "S1270_32GB_1X1TBSATA_NORAID"),
					resource.TestCheckResourceAttr(
						configName, "notes", "baremetal notes"),
					CheckStringSet(
						configName,
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func TestAccIBMComputeBareMetal_With_Network_Storage_Access(t *testing.T) {
	var bareMetal datatypes.Hardware
	hostname := acctest.RandString(16)
	domain := "storage.tfbmuat.ibm.com"

	configInstance := "ibm_compute_bare_metal.terraform-bm-storage-access"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config: testBareMetalAccessToStoragesBasic(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configInstance, &bareMetal),
					resource.TestCheckResourceAttr(
						configInstance, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configInstance, "domain", domain),
					resource.TestCheckResourceAttr(
						configInstance, "datacenter", "dal10"),
					resource.TestCheckResourceAttr(
						configInstance, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configInstance, "file_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "block_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "redundant_power_supply", "true"),
				),
			},
			{
				Config: testBareMetalAccessToStoragesUpdate(hostname, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists(configInstance, &bareMetal),
					resource.TestCheckResourceAttr(
						configInstance, "file_storage_ids.#", "1"),
					resource.TestCheckResourceAttr(
						configInstance, "block_storage_ids.#", "0"),
				),
			},
		},
	})
}

func TestAccSoftLayerBareMetalQuote_Basic(t *testing.T) {
	var bareMetal datatypes.Hardware
	hostname := acctest.RandString(16)
	domain := "bm.quote.tfuat.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testBareMetalQuoteConfigBasic(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists("ibm_compute_bare_metal.bm-quote", &bareMetal),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-quote", "hostname", hostname),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-quote", "domain", domain),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-quote", "user_metadata", "{\"value\":\"newvalue\"}"),
					CheckStringSet(
						"ibm_compute_bare_metal.bm-quote",
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func TestAccSoftLayerBareMetalCustom_Basic(t *testing.T) {
	var bareMetal datatypes.Hardware
	hostname := acctest.RandString(14)
	domain := "bm.custom.tfuat.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testBareMetalCustomConfig(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists("ibm_compute_bare_metal.bm-custom", &bareMetal),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "memory", "64"),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "network_speed", "1000"),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "public_bandwidth", "500"),
				),
			},
		},
	})
}

func TestAccSoftLayerBareMetalCustom_with_gpus(t *testing.T) {
	var bareMetal datatypes.Hardware
	hostname := acctest.RandString(14)
	domain := "bm.custom.tfuat.gpus.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testBareMetalCustomConfigWithGpus(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists("ibm_compute_bare_metal.bm-custom", &bareMetal),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "memory", "64"),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "network_speed", "1000"),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "public_bandwidth", "500"),
				),
			},
		},
	})
}

func TestAccSoftLayerBareMetalCustom_with_monitoring_none(t *testing.T) {
	var bareMetal datatypes.Hardware
	hostname := acctest.RandString(14)
	domain := "bm.custom.tfuat.com"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeBareMetalDestroy,
		Steps: []resource.TestStep{
			{
				Config:  testBareMetalCustomConfigWithMonitoringNone(hostname, domain),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeBareMetalExists("ibm_compute_bare_metal.bm-custom", &bareMetal),
					resource.TestCheckResourceAttr(
						"ibm_compute_bare_metal.bm-custom", "memory", "128"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeBareMetalDestroy(s *terraform.State) error {
	service := services.GetHardwareService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_bare_metal" {
			continue
		}

		id, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the bare metal
		_, err := service.Id(id).GetObject()

		// Wait
		if err != nil {
			if apiErr, ok := err.(sl.Error); !ok || apiErr.StatusCode != 404 {
				return fmt.Errorf(
					"Error waiting for bare metal (%d) to be destroyed: %s",
					id, err)
			}
		}
	}

	return nil
}

func testAccCheckIBMComputeBareMetalExists(n string, bareMetal *datatypes.Hardware) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No bare metal ID is set")
		}

		id, err := strconv.Atoi(rs.Primary.ID)

		if err != nil {
			return err
		}

		service := services.GetHardwareService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		bm, err := service.Id(id).GetObject()
		if err != nil {
			return err
		}

		fmt.Printf("The ID is %d", *bm.Id)

		if *bm.Id != id {
			return errors.New("Bare metal not found")
		}

		*bareMetal = bm

		return nil
	}
}

func testAccCheckIBMComputeBareMetalConfig_basic(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
  hostname                  = "%s"
  domain                    = "terraformuat.ibm.com"
  os_reference_code         = "UBUNTU_16_64"
  datacenter                = "dal10"
  network_speed             = 100
  hourly_billing            = true
  private_network_only      = false
  extended_hardware_testing = "%t"
  user_metadata             = "{\"value\":\"newvalue\"}"
  fixed_config_preset       = "S1270_32GB_1X1TBSATA_NORAID"
  tags                      = ["collectd"]
  notes                     = "baremetal notes"
}
`, hostname, extendedHardwareTesting)
}

func testAccCheckIBMComputeBareMetalConfig_update(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
	hostname                  = "%s"
	domain                    = "terraformuat.ibm.com"
	os_reference_code         = "UBUNTU_16_64"
	datacenter                = "dal10"
	network_speed             = 100
	hourly_billing            = true
	private_network_only      = false
	extended_hardware_testing = "%t"
	user_metadata             = "{\"value\":\"newvalue\"}"
	fixed_config_preset       = "S1270_32GB_1X1TBSATA_NORAID"
	tags                      = ["mesos-master"]
}
`, hostname, extendedHardwareTesting)
}

func testBareMetalAccessToStoragesBasic(hostname, domain string) string {
	config := fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-bm-storage-access" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "UBUNTU_16_64"
    datacenter = "dal10"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    user_metadata = "{\"value\":\"newvalue\"}"
    fixed_config_preset = "S1270_32GB_1X1TBSATA_NORAID"
	
    tags = ["mesos-master"]
	file_storage_ids = ["${ibm_storage_file.fs1.id}"]
	block_storage_ids = ["${ibm_storage_block.bs.id}"]
	redundant_power_supply = true
}
%s
%s

`, hostname, domain, fsConfig1, bsConfig1)
	return config
}

func testBareMetalAccessToStoragesUpdate(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-bm-storage-access" {
    hostname = "%s"
    domain = "%s"
    os_reference_code = "UBUNTU_16_64"
    datacenter = "dal10"
    network_speed = 100
    hourly_billing = true
    private_network_only = false
    user_metadata = "{\"value\":\"newvalue\"}"
    fixed_config_preset = "S1270_32GB_1X1TBSATA_NORAID"
	file_storage_ids = ["${ibm_storage_file.fs2.id}"]
	block_storage_ids = []
	
    tags = ["mesos-master"]
	file_storage_ids = ["${ibm_storage_file.fs2.id}"]
	redundant_power_supply = true

}

%s

`, hostname, domain, fsConfig2)

}

func testBareMetalQuoteConfigBasic(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "bm-quote" {
    hostname = "%s"
    domain = "%s"
    user_metadata = "{\"value\":\"newvalue\"}"
    quote_id = 2282869
    tags = ["collectd"]
}
`, hostname, domain)
}

func testBareMetalCustomConfig(hostname, domain string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "bm-custom" {
    package_key_name = "DUAL_E52600_V4_12_DRIVES"
    process_key_name = "INTEL_INTEL_XEON_E52620_V4_2_10"
    memory = 64
    os_key_name = "OS_WINDOWS_2012_R2_FULL_DC_64_BIT_2"
    hostname = "%s"
    domain = "%s"
    datacenter = "dal05"
    network_speed = 1000
    public_bandwidth = 500
    disk_key_names = [ "HARD_DRIVE_1_00_TB_SATA_2", "HARD_DRIVE_1_00_TB_SATA_2" ]
	hourly_billing = false
    redundant_power_supply = true
}
`, hostname, domain)
}

func testBareMetalCustomConfigWithGpus(hostname, domain string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_bare_metal" "bm-custom" {
		package_key_name       = "DUAL_E52600_V4_12_DRIVES"
		process_key_name       = "INTEL_INTEL_XEON_E52620_V4_2_10"
		gpu_key_name           = "GPU_NVIDIA_TESLA_K80"
		gpu_secondary_key_name = "GPU_NVIDIA_TESLA_K80"
		memory                 = 64
		os_key_name            = "OS_WINDOWS_2012_R2_FULL_DC_64_BIT_2"
		hostname               = "%s"
		domain                 = "%s"
		datacenter             = "dal05"
		network_speed          = 1000
		public_bandwidth       = 500
		disk_key_names         = ["HARD_DRIVE_1_00_TB_SATA_2", "HARD_DRIVE_1_00_TB_SATA_2"]
		hourly_billing         = false
		}
`, hostname, domain)
}

func testAccCheckIBMComputeBareMetalConfig_with_unbonded_port_speed(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
	hostname               = "%s"
	domain                 = "terraformuat.ibm.com"
	os_reference_code      = "UBUNTU_16_64"
	datacenter             = "dal10"
	network_speed          = 1000
	unbonded_network       = true
	hourly_billing         = true
	private_network_only   = false
	user_metadata          = "{\"value\":\"newvalue\"}"
	fixed_config_preset    = "S1270_32GB_1X1TBSATA_NORAID"
	tags                   = ["collectd"]
	notes                  = "baremetal notes"
	}
`, hostname)
}

func testAccCheckIBMComputeBareMetalConfig_with_ipv6(hostname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
	hostname               = "%s"
	domain                 = "terraformuat.ibm.com"
	os_reference_code      = "UBUNTU_16_64"
	datacenter             = "dal10"
	ipv6_enabled           = true
	ipv6_static_enabled    = true
	secondary_ip_count     = 4
	network_speed          = 100
	hourly_billing         = true
	private_network_only   = false
	user_metadata          = "{\"value\":\"newvalue\"}"
	fixed_config_preset    = "S1270_32GB_1X1TBSATA_NORAID"
	tags                   = ["collectd"]
	notes                  = "baremetal notes"
	}
`, hostname)
}

func testBareMetalCustomConfigWithMonitoringNone(hostname, domain string) string {
	return fmt.Sprintf(`
	resource "ibm_compute_bare_metal" "bm-custom" {
		package_key_name       = "BI_S1_NW128_VMWARE_VIRTUALIZATION"
		process_key_name       = "INTEL_XEON_2690_2_60"
		memory                 = 128
		os_key_name            = "OS_VMWARE_SERVER_VIRTUALIZATION_6_5"
		hostname               = "%s"
		domain                 = "%s"
		datacenter             = "sao01"
		disk_key_names         = ["HARD_DRIVE_1_2_TB_SSD_10_DWPD"]
		hourly_billing         = false
		public_bandwidth       = 500
		network_speed          = 1000
		unbonded_network       = true
		}
`, hostname, domain)
}

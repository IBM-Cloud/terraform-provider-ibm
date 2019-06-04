package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMISZonesDataSource_1(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISZonesDataSourceConfig1(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_zones.testacc_ds_zones1", "region", regionName),
					resource.TestCheckResourceAttrSet("data.ibm_is_zones.testacc_ds_zones1", "zones.0"),
				),
			},
		},
	})
}

func testAccCheckIBMISZonesDataSourceConfig1() string {
	// status filter defaults to empty
	return fmt.Sprintf(`

data "ibm_is_zones" "testacc_ds_zones1" {
	region = "%s",
    }`, regionName)
}

func TestAccIBMISZonesDataSource_2(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISZonesDataSourceConfig2(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_zones.testacc_ds_zones2", "region", regionName),
					resource.TestCheckResourceAttrSet("data.ibm_is_zones.testacc_ds_zones2", "zones.0"),
				),
			},
		},
	})
}

func testAccCheckIBMISZonesDataSourceConfig2() string {
	// status filter is explicitly set to zero value
	return fmt.Sprintf(`

data "ibm_is_zones" "testacc_ds_zones2" {
	region = "%s",
	status = "",
    }`, regionName)
}

func TestAccIBMISZonesDataSource_3(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISZonesDataSourceConfig3(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_zones.testacc_ds_zones3", "region", regionName),
					resource.TestCheckNoResourceAttr("data.ibm_is_zones.testacc_ds_zones3", "zones.0"),
				),
			},
		},
	})
}
func testAccCheckIBMISZonesDataSourceConfig3() string {
	// status filter is set to a value that will never occur
	return fmt.Sprintf(`

data "ibm_is_zones" "testacc_ds_zones3" {
	region = "%s",
	status = "no.status.matches.this",
    }`, regionName)
}

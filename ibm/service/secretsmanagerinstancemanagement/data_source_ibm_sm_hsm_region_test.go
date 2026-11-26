// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

// These tests cover the "region" argument. They need the same environment as the other HSM tests.
//
//   - RegionDefault and NoDrift leave the region out, so run them with IC_REGION set to the real region
//     of the instance.
//   - RegionExplicit sets the region on every resource and data source. It needs
//     SECRETS_MANAGER_INSTANCE_REGION (the real region). Run it with IC_REGION set to a different
//     region to prove that the argument overrides the provider region.

const (
	hsmDestinationsDS = "data.ibm_sm_hsm_destinations.test"
	hsmHaGroupsDS     = "data.ibm_sm_hsm_ha_groups.test"
)

func hsmInstanceRegion(t *testing.T) string {
	if acc.SecretsManagerInstanceRegion == "" {
		t.Skip("set SECRETS_MANAGER_INSTANCE_REGION to run this region test")
	}
	return acc.SecretsManagerInstanceRegion
}

// hsmRegionLine returns the region argument for a configuration, or nothing when region is empty.
func hsmRegionLine(region string) string {
	if region == "" {
		return ""
	}
	return fmt.Sprintf("region = \"%s\"", region)
}

// hsmParentConfig is the parent destination every test builds on. When region is set, it is set on
// the parent too, so the whole configuration can run with a provider region that is not the real one.
func hsmParentConfig(instanceID, region string) string {
	return withRegion(hsmDestinationConfig(instanceID), "ibm_sm_hsm_destination", region)
}

// withRegion adds the region argument to every resource of the given type in config.
func withRegion(config, resourceType, region string) string {
	if region == "" {
		return config
	}
	header := fmt.Sprintf("resource \"%s\" \"", resourceType)
	var out strings.Builder
	for _, line := range strings.SplitAfter(config, "\n") {
		out.WriteString(line)
		if strings.Contains(line, header) {
			out.WriteString(fmt.Sprintf("\t\t\tregion = \"%s\"\n", region))
		}
	}
	return out.String()
}

func hsmDestinationsDSConfig(instanceID, region string) string {
	return hsmParentConfig(instanceID, region) + fmt.Sprintf(`
		data "ibm_sm_hsm_destinations" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			%s
		}
	`, hsmRegionLine(region))
}

func hsmHaGroupConfig(instanceID, name, region string) string {
	return hsmParentConfig(instanceID, region) + withRegion(hsmHaGroupBlock(name, ""), "ibm_sm_hsm_ha_group", region)
}

func hsmHaGroupsDSConfig(instanceID, name, region string) string {
	return hsmHaGroupConfig(instanceID, name, region) + fmt.Sprintf(`
		data "ibm_sm_hsm_ha_groups" "test" {
			instance_id = ibm_sm_hsm_ha_group.test.instance_id
			destination_id = ibm_sm_hsm_ha_group.test.destination_id
			%s
		}
	`, hsmRegionLine(region))
}

// Without the argument the region is taken from the provider configuration and written to the state.
func TestAccIbmSmHsmDestinationsDataSourceRegionDefault(t *testing.T) {
	instanceID := hsmInstanceID(t)

	testAccHsmDataSource(t, hsmDestinationsDSConfig(instanceID, ""), hsmDestinationsDS,
		[]string{"id", "instance_id", "region", "hsm_destinations.#"}, nil)
}

// With the argument the given region is used and is kept in the state.
func TestAccIbmSmHsmDestinationsDataSourceRegionExplicit(t *testing.T) {
	instanceID := hsmInstanceID(t)
	region := hsmInstanceRegion(t)

	testAccHsmDataSource(t, hsmDestinationsDSConfig(instanceID, region), hsmDestinationsDS,
		[]string{"id", "instance_id", "hsm_destinations.#"},
		map[string]string{"region": region})
}

func TestAccIbmSmHsmHaGroupsDataSourceRegionDefault(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")

	testAccHsmDataSource(t, hsmHaGroupsDSConfig(instanceID, hsmRandName(), ""), hsmHaGroupsDS,
		[]string{"id", "instance_id", "destination_id", "region", "ha_groups.#"}, nil)
}

func TestAccIbmSmHsmHaGroupsDataSourceRegionExplicit(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	region := hsmInstanceRegion(t)

	testAccHsmDataSource(t, hsmHaGroupsDSConfig(instanceID, hsmRandName(), region), hsmHaGroupsDS,
		[]string{"id", "instance_id", "destination_id", "ha_groups.#"},
		map[string]string{"region": region})
}

// The resources the data sources depend on must keep the region they were created with. A second plan
// with no changes shows that the region in the state does not cause a diff or a replacement.
func TestAccIbmSmHsmHaGroupRegionNoDrift(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	config := hsmHaGroupConfig(instanceID, hsmRandName(), "")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_sm_hsm_destination.parent", "region"),
					resource.TestCheckResourceAttrSet("ibm_sm_hsm_ha_group.test", "region"),
				),
			},
			{
				Config:   config,
				PlanOnly: true,
			},
		},
	})
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

// hsmResourceSpec describes one HSM resource for the shared create, update and import test.
type hsmResourceSpec struct {
	// resourceType is the Terraform resource type, for example ibm_sm_hsm_ha_group.
	resourceType string
	// config returns the whole configuration, including the parents the resource needs.
	// The description is omitted from the resource when it is empty.
	config func(instanceID, name, description string) string
	// extraEnv lists environment variables that must be set to run the test.
	extraEnv []string
	// extraChecks are attribute checks on the resource besides instance_id, name and description.
	extraChecks []resource.TestCheckFunc
	// importIgnore lists attributes that Read does not set because they only exist in the resource ID.
	// region is always ignored, because it is set on create and cannot be read back on import.
	importIgnore []string
	// get reads the resource by the parts of its ID. It returns the HTTP response for the 404 check.
	get func(client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, parts []string) (*core.DetailedResponse, error)
}

func (s hsmResourceSpec) address() string { return s.resourceType + ".test" }

func hsmInstanceID(t *testing.T, extraEnv ...string) string {
	instanceID := acc.SecretsManagerDedicatedInstanceID
	if instanceID == "" {
		instanceID = acc.SecretsManagerInstanceID
	}
	if instanceID == "" {
		t.Skip("set SECRETS_MANAGER_DEDICATED_INSTANCE_ID or SECRETS_MANAGER_INSTANCE_ID to run the HSM resource tests")
	}
	for _, env := range extraEnv {
		if env == "SM_HSM_MEMBER_SERIALS" {
			if acc.SecretsManagerHsmMemberSerials == "" {
				t.Skip("set SM_HSM_MEMBER_SERIALS or SM_HSM_HA_GROUP_MEMBER_SERIALS to run this HSM resource test")
			}
			continue
		}
		if os.Getenv(env) == "" {
			t.Skipf("set %s to run this HSM resource test", env)
		}
	}
	return instanceID
}

// hsmDestinationConfig is the destination every dependent resource test builds on.
func hsmDestinationConfig(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_destination" "parent" {
			instance_id = "%s"
			name = "tf_parent_dest_%d"
			vendor = "thales_luna"
		}
	`, instanceID, acctest.RandIntRange(10, 1000))
}

func hsmDescriptionLine(description string) string {
	if description == "" {
		return ""
	}
	return fmt.Sprintf("description = \"%s\"", description)
}

func hsmRandName() string {
	return fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 1000))
}

func hsmGroupLabel() string {
	if v := os.Getenv("SM_HSM_GROUP_LABEL"); v != "" {
		return v
	}
	if v := os.Getenv("SM_HSM_HA_GROUP_LABEL"); v != "" {
		return v
	}
	return "production-ha"
}

func hsmGroupSerial() string {
	if v := os.Getenv("SM_HSM_GROUP_SERIAL"); v != "" {
		return v
	}
	if v := os.Getenv("SM_HSM_HA_GROUP_SERIAL"); v != "" {
		return v
	}
	return "11323465009870"
}

func hsmMemberSerials() []string {
	v := acc.SecretsManagerHsmMemberSerials
	if v == "" {
		return []string{"1323465009870", "1323465009871"}
	}
	var result []string
	v = strings.TrimSpace(v)
	if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
		if err := json.Unmarshal([]byte(v), &result); err == nil && len(result) > 0 {
			return result
		}
	}
	for _, p := range strings.Split(v, ",") {
		p = strings.TrimSpace(p)
		p = strings.Trim(p, "\"'`")
		if p != "" {
			result = append(result, p)
		}
	}
	if len(result) > 0 {
		return result
	}
	return []string{"1323465009870", "1323465009871"}
}

func hsmMemberSerialsHCL() string {
	serials := hsmMemberSerials()
	b, _ := json.Marshal(serials)
	return string(b)
}

func hsmHaGroupBlock(name, description string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_ha_group" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
			group_label = "%s"
			group_serial = "%s"
			member_serials = %s
		}
	`, name, hsmDescriptionLine(description), hsmGroupLabel(), hsmGroupSerial(), hsmMemberSerialsHCL())
}

// hsmPathBlock needs SM_HSM_PPS_GATEWAY_CRN; pass it to hsmInstanceID so the test skips when it is missing.
func hsmPathBlock(name, description string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_path" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
			crn = "%s"
		}
	`, name, hsmDescriptionLine(description), acc.SecretsManagerHsmPpsGatewayCrn)
}

// hsmServerCaBlock needs SM_HSM_PPS_GATEWAY_CRN and SM_HSM_SERVER_CA_PEM_PATH.
func hsmServerCaBlock(name, description string) string {
	return hsmPathBlock("tf_parent_path", "") + fmt.Sprintf(`
		resource "ibm_sm_hsm_server_ca" "test" {
			instance_id = ibm_sm_hsm_path.test.instance_id
			destination_id = ibm_sm_hsm_path.test.destination_id
			path_id = ibm_sm_hsm_path.test.path_id
			name = "%s"
			%s
			certificate = file("%s")
		}
	`, name, hsmDescriptionLine(description), acc.SecretsManagerHsmServerCaPemPath)
}

func hsmClientCertBlock(name, description string) string {
	return fmt.Sprintf(`
		resource "ibm_sm_hsm_client_certificate" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
		}
	`, name, hsmDescriptionLine(description))
}

// testAccHsmDataSource applies the config and checks the attributes of the data source at addr.
func testAccHsmDataSource(t *testing.T, config, addr string, set []string, equal map[string]string) {
	var checks []resource.TestCheckFunc
	for _, attr := range set {
		checks = append(checks, resource.TestCheckResourceAttrSet(addr, attr))
	}
	for attr, want := range equal {
		checks = append(checks, resource.TestCheckResourceAttr(addr, attr, want))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{Config: config, Check: resource.ComposeAggregateTestCheckFunc(checks...)},
		},
	})
}

// testAccHsmResource runs create, rename and (when withDescription is set) description update, and import.
func testAccHsmResource(t *testing.T, spec hsmResourceSpec, withDescription bool) {
	instanceID := hsmInstanceID(t, spec.extraEnv...)
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 1000))
	nameUpdate := name + "_updated"
	description, descriptionUpdate := "", ""
	if withDescription {
		description = fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 1000))
		descriptionUpdate = description + "_updated"
	}
	addr := spec.address()

	checks := func(wantName, wantDescription string) resource.TestCheckFunc {
		all := []resource.TestCheckFunc{
			testAccHsmExists(addr, spec.get),
			resource.TestCheckResourceAttr(addr, "instance_id", instanceID),
			resource.TestCheckResourceAttr(addr, "name", wantName),
		}
		if withDescription {
			all = append(all, resource.TestCheckResourceAttr(addr, "description", wantDescription))
		}
		return resource.ComposeAggregateTestCheckFunc(append(all, spec.extraChecks...)...)
	}

	steps := []resource.TestStep{
		{Config: spec.config(instanceID, name, description), Check: checks(name, description)},
		{Config: spec.config(instanceID, nameUpdate, descriptionUpdate), Check: checks(nameUpdate, descriptionUpdate)},
	}
	if withDescription {
		steps = append(steps, resource.TestStep{
			ResourceName:            addr,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: append([]string{"region"}, spec.importIgnore...),
		})
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccHsmDestroy(spec.resourceType, spec.get),
		Steps:        steps,
	})
}

func testAccHsmExists(n string, get func(*secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, []string) (*core.DetailedResponse, error)) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		client, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}
		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}
		_, err = get(client, parts)
		return err
	}
}

func testAccHsmDestroy(resourceType string, get func(*secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, []string) (*core.DetailedResponse, error)) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}
		for _, rs := range s.RootModule().Resources {
			if rs.Type != resourceType {
				continue
			}
			parts, err := flex.SepIdParts(rs.Primary.ID, "/")
			if err != nil {
				return err
			}
			response, err := get(client, parts)
			if err == nil {
				return fmt.Errorf("%s still exists: %s", resourceType, rs.Primary.ID)
			}
			if response == nil || response.StatusCode != 404 {
				return fmt.Errorf("error checking that %s (%s) was destroyed: %s", resourceType, rs.Primary.ID, err)
			}
		}
		return nil
	}
}

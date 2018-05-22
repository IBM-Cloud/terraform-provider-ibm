package ibm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func TestAccIBMComputeAutoScalePolicy_Basic(t *testing.T) {
	var scalepolicy datatypes.Scale_Policy
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandInt())
	hostname := acctest.RandString(16)
	policyname := acctest.RandString(16)
	updatedpolicyname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScalePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScalePolicyConfig_basic(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					testAccCheckIBMComputeAutoScalePolicyAttributes(&scalepolicy, policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "append_triggers_to_existing", "true"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_type", "RELATIVE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_amount", "1"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "cooldown", "30"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "triggers.#", "3"),
					testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(&scalepolicy, 2, "0 1 ? * MON,WED *"),
					testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(&scalepolicy, 120, "80"),
					testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(&scalepolicy, testOnetimeTriggerDate),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScalePolicyConfig_updated(groupname, hostname, updatedpolicyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", updatedpolicyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_type", "ABSOLUTE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "scale_amount", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "cooldown", "35"),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "triggers.#", "3"),
					testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(&scalepolicy, 2, "0 1 ? * MON,WED,SAT *"),
					testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(&scalepolicy, 130, "90"),
					testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(&scalepolicy, testOnetimeTriggerUpdatedDate),
				),
			},
		},
	})
}

func TestAccIBMComputeAutoScaleWithTag(t *testing.T) {
	var scalepolicy datatypes.Scale_Policy
	groupname := fmt.Sprintf("terraformuat_%d", acctest.RandInt())
	hostname := acctest.RandString(16)
	policyname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMComputeAutoScalePolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScalePolicyWithTag(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					testAccCheckIBMComputeAutoScalePolicyAttributes(&scalepolicy, policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "tags.#", "2"),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMComputeAutoScalePolicyWithUpdatedTag(groupname, hostname, policyname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeAutoScalePolicyExists("ibm_compute_autoscale_policy.sample-http-cluster-policy", &scalepolicy),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "name", policyname),
					resource.TestCheckResourceAttr(
						"ibm_compute_autoscale_policy.sample-http-cluster-policy", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeAutoScalePolicyDestroy(s *terraform.State) error {
	service := services.GetScalePolicyService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_autoscale_policy" {
			continue
		}

		scalepolicyId, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := service.Id(scalepolicyId).GetObject()

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf("Error waiting for Auto Scale Policy (%s) to be destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func testAccCheckIBMComputeAutoScalePolicyContainsResourceUseTriggers(scalePolicy *datatypes.Scale_Policy, period int, value string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false

		for _, scaleResourceUseTrigger := range scalePolicy.ResourceUseTriggers {
			for _, scaleResourceUseWatch := range scaleResourceUseTrigger.Watches {
				if *scaleResourceUseWatch.Metric == "host.cpu.percent" && *scaleResourceUseWatch.Operator == ">" &&
					*scaleResourceUseWatch.Period == period && *scaleResourceUseWatch.Value == value {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("Resource use trigger not found in scale policy")

		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyContainsRepeatingTriggers(scalePolicy *datatypes.Scale_Policy, typeId int, schedule string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false

		for _, scaleRepeatingTrigger := range scalePolicy.RepeatingTriggers {
			if *scaleRepeatingTrigger.TypeId == typeId && *scaleRepeatingTrigger.Schedule == schedule {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("Repeating trigger %d with schedule %s not found in scale policy", typeId, schedule)

		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyContainsOneTimeTriggers(scalePolicy *datatypes.Scale_Policy, testOnetimeTriggerDate string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		found := false
		const IBMComputeTimeFormat = "2006-01-02T15:04:05-07:00"
		utcLoc, _ := time.LoadLocation("UTC")

		for _, scaleOneTimeTrigger := range scalePolicy.OneTimeTriggers {
			if scaleOneTimeTrigger.Date.In(utcLoc).Format(IBMComputeTimeFormat) == testOnetimeTriggerDate {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("One time trigger with date %s not found in scale policy", testOnetimeTriggerDate)
		}

		return nil

	}
}

func testAccCheckIBMComputeAutoScalePolicyAttributes(scalepolicy *datatypes.Scale_Policy, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if *scalepolicy.Name != policyname {
			return fmt.Errorf("Bad name: %s", *scalepolicy.Name)
		}

		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyExists(n string, scalepolicy *datatypes.Scale_Policy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		scalepolicyId, _ := strconv.Atoi(rs.Primary.ID)

		service := services.GetScalePolicyService(testAccProvider.Meta().(ClientSession).SoftLayerSession())
		foundScalePolicy, err := service.Id(scalepolicyId).Mask(strings.Join(IBMComputeAutoScalePolicyObjectMask, ",")).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundScalePolicy.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*scalepolicy = foundScalePolicy
		return nil
	}
}

func testAccCheckIBMComputeAutoScalePolicyConfig_basic(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_7_64"
        local_disk = false
        datacenter = "dal09"
    }
}

resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "RELATIVE"
    scale_amount = 1
    cooldown = 30
	scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
	append_triggers_to_existing=true
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "80"
                    period = 120
        }
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED *"
    }

}`, groupname, hostname, policyname, testOnetimeTriggerDate)
}

const IBMComputeTestTimeFormat = string("2006-01-02T15:04:05-07:00")

var utcLoc, _ = time.LoadLocation("UTC")

var testOnetimeTriggerDate = time.Now().In(utcLoc).AddDate(0, 0, 1).Format(IBMComputeTestTimeFormat)

func testAccCheckIBMComputeAutoScalePolicyConfig_updated(groupname, hostname, updatedpolicyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_7_64"
        local_disk = false
        datacenter = "dal09"
    }
}
resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "ABSOLUTE"
    scale_amount = 2
    cooldown = 35
	scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
	append_triggers_to_existing=true
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "90"
                    period = 130
        }
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED,SAT *"
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
}`, groupname, hostname, updatedpolicyname, testOnetimeTriggerUpdatedDate)
}

func testAccCheckIBMComputeAutoScalePolicyWithTag(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_7_64"
        local_disk = false
        datacenter = "dal09"
    }
}

resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "RELATIVE"
    scale_amount = 1
    cooldown = 30
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "80"
                    period = 120
        }
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED *"
	}
	tags = ["one", "two"]

}`, groupname, hostname, policyname, testOnetimeTriggerDate)
}

func testAccCheckIBMComputeAutoScalePolicyWithUpdatedTag(groupname, hostname, policyname string) string {
	return fmt.Sprintf(`
resource "ibm_compute_autoscale_group" "sample-http-cluster-with-policy" {
    name = "%s"
    regional_group = "na-usa-central-1"
    cooldown = 30
    minimum_member_count = 1
    maximum_member_count = 10
    termination_policy = "CLOSEST_TO_NEXT_CHARGE"
    virtual_guest_member_template = {
        hostname = "%s"
        domain = "terraformuat.ibm.com"
        cores = 1
        memory = 4096
        network_speed = 1000
        hourly_billing = true
        os_reference_code = "DEBIAN_7_64"
        local_disk = false
        datacenter = "dal09"
    }
}
resource "ibm_compute_autoscale_policy" "sample-http-cluster-policy" {
    name = "%s"
    scale_type = "ABSOLUTE"
    scale_amount = 2
    cooldown = 35
    scale_group_id = "${ibm_compute_autoscale_group.sample-http-cluster-with-policy.id}"
    triggers = {
        type = "RESOURCE_USE"
        watches = {

                    metric = "host.cpu.percent"
                    operator = ">"
                    value = "90"
                    period = 130
        }
    }
    triggers = {
        type = "REPEATING"
        schedule = "0 1 ? * MON,WED,SAT *"
    }
    triggers = {
        type = "ONE_TIME"
        date = "%s"
	}
	tags = ["one", "two", "three"]
}`, groupname, hostname, policyname, testOnetimeTriggerUpdatedDate)
}

var testOnetimeTriggerUpdatedDate = time.Now().In(utcLoc).AddDate(0, 0, 2).Format(IBMComputeTestTimeFormat)

func Test_resourceIBMComputeAutoScalePolicy(t *testing.T) {
	tests := []struct {
		name string
		want *schema.Resource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := resourceIBMComputeAutoScalePolicy(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicy() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyCreate(t *testing.T) {
	type args struct {
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := resourceIBMComputeAutoScalePolicyCreate(tt.args.d, tt.args.meta); (err != nil) != tt.wantErr {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyCreate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyRead(t *testing.T) {
	type args struct {
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := resourceIBMComputeAutoScalePolicyRead(tt.args.d, tt.args.meta); (err != nil) != tt.wantErr {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyRead() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyUpdate(t *testing.T) {
	type args struct {
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := resourceIBMComputeAutoScalePolicyUpdate(tt.args.d, tt.args.meta); (err != nil) != tt.wantErr {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyUpdate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyDelete(t *testing.T) {
	type args struct {
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := resourceIBMComputeAutoScalePolicyDelete(tt.args.d, tt.args.meta); (err != nil) != tt.wantErr {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyDelete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyExists(t *testing.T) {
	type args struct {
		d    *schema.ResourceData
		meta interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := resourceIBMComputeAutoScalePolicyExists(tt.args.d, tt.args.meta)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyExists() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyExists() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_validateTriggerTypes(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := validateTriggerTypes(tt.args.d); (err != nil) != tt.wantErr {
			t.Errorf("%q. validateTriggerTypes() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Test_prepareOneTimeTriggers(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    []datatypes.Scale_Policy_Trigger_OneTime
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := prepareOneTimeTriggers(tt.args.d)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. prepareOneTimeTriggers() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. prepareOneTimeTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_prepareRepeatingTriggers(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    []datatypes.Scale_Policy_Trigger_Repeating
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := prepareRepeatingTriggers(tt.args.d)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. prepareRepeatingTriggers() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. prepareRepeatingTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_prepareResourceUseTriggers(t *testing.T) {
	type args struct {
		d *schema.ResourceData
	}
	tests := []struct {
		name    string
		args    args
		want    []datatypes.Scale_Policy_Trigger_ResourceUse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := prepareResourceUseTriggers(tt.args.d)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. prepareResourceUseTriggers() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. prepareResourceUseTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_prepareWatches(t *testing.T) {
	type args struct {
		d *schema.Set
	}
	tests := []struct {
		name    string
		args    args
		want    []datatypes.Scale_Policy_Trigger_ResourceUse_Watch
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := prepareWatches(tt.args.d)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. prepareWatches() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. prepareWatches() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_readOneTimeTriggers(t *testing.T) {
	type args struct {
		list []datatypes.Scale_Policy_Trigger_OneTime
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := readOneTimeTriggers(tt.args.list); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. readOneTimeTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_readRepeatingTriggers(t *testing.T) {
	type args struct {
		list []datatypes.Scale_Policy_Trigger_Repeating
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := readRepeatingTriggers(tt.args.list); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. readRepeatingTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_readResourceUseTriggers(t *testing.T) {
	type args struct {
		list []datatypes.Scale_Policy_Trigger_ResourceUse
	}
	tests := []struct {
		name string
		args args
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := readResourceUseTriggers(tt.args.list); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. readResourceUseTriggers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_readResourceUseWatches(t *testing.T) {
	type args struct {
		list []datatypes.Scale_Policy_Trigger_ResourceUse_Watch
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := readResourceUseWatches(tt.args.list); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. readResourceUseWatches() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyTriggerHash(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := resourceIBMComputeAutoScalePolicyTriggerHash(tt.args.v); got != tt.want {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyTriggerHash() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_resourceIBMComputeAutoScalePolicyHandlerHash(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := resourceIBMComputeAutoScalePolicyHandlerHash(tt.args.v); got != tt.want {
			t.Errorf("%q. resourceIBMComputeAutoScalePolicyHandlerHash() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

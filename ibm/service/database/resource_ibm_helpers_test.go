package database_test

import (
	"fmt"
	"reflect"
	"strings"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
)

const (
	databaseInstanceSuccessStatus      = "active"
	databaseInstanceProvisioningStatus = "provisioning"
	databaseInstanceProgressStatus     = "in progress"
	databaseInstanceInactiveStatus     = "inactive"
	databaseInstanceFailStatus         = "failed"
	databaseInstanceRemovedStatus      = "removed"
	databaseInstanceReclamation        = "pending_reclamation"
)

func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_database" {
			continue
		}

		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsContClient.GetResourceInstance(&rsInst)
		if err == nil {
			if !reflect.DeepEqual(instance, models.ServiceInstance{}) && *instance.State == "active" {
				return fmt.Errorf("Database still exists: %s", rs.Primary.ID)
			}
		} else {
			if !strings.Contains(err.Error(), "404") {
				return fmt.Errorf("[ERROR] Error checking if database (%s) has been destroyed: %s %s", rs.Primary.ID, err, response)
			}
		}
	}
	return nil
}

func testAccDatabaseInstanceManuallyDelete(tfDatabaseID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_ = testAccDatabaseInstanceManuallyDeleteUnwrapped(s, tfDatabaseID)
		return nil
	}
}

func testAccDatabaseInstanceManuallyDeleteUnwrapped(s *terraform.State, tfDatabaseID *string) error {
	rsConClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}
	instance := *tfDatabaseID
	var instanceID string
	if strings.HasPrefix(instance, "crn") {
		instanceID = instance
	} else {
		_, instanceID, _ = flex.ConvertTftoCisTwoVar(instance)
	}
	recursive := true
	deleteReq := rc.DeleteResourceInstanceOptions{
		ID:        &instanceID,
		Recursive: &recursive,
	}
	response, err := rsConClient.DeleteResourceInstance(&deleteReq)
	if err != nil {
		return fmt.Errorf("[ERROR] Error deleting resource instance: %s %s", err, response)
	}

	_ = &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus, databaseInstanceSuccessStatus},
		Target:  []string{databaseInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			rsInst := rc.GetResourceInstanceOptions{
				ID: &instanceID,
			}
			instance, response, err := rsConClient.GetResourceInstance(&rsInst)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, databaseInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if *instance.State == databaseInstanceFailStatus {
				return instance, *instance.State, fmt.Errorf("[ERROR] The resource instance %s failed to delete: %v %s", instanceID, err, response)
			}
			return instance, *instance.State, nil
		},
		Timeout:    90 * time.Second,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}
	if err != nil {
		return fmt.Errorf("[ERROR] Error waiting for resource instance (%s) to be deleted: %s", instanceID, err)
	}
	return nil
}

func testAccCheckIBMDatabaseInstanceExists(n string, tfDatabaseID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			return err
		}
		instanceID := rs.Primary.ID

		rsInst := rc.GetResourceInstanceOptions{
			ID: &instanceID,
		}
		instance, response, err := rsContClient.GetResourceInstance(&rsInst)
		if err != nil {
			if strings.Contains(err.Error(), "Object not found") ||
				strings.Contains(err.Error(), "status code: 404") {
				*tfDatabaseID = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response)
		}
		if strings.Contains(*instance.State, "removed") {
			*tfDatabaseID = ""
			return nil
		}

		*tfDatabaseID = instanceID
		return nil
	}
}

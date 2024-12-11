// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/session"
)

const userCustomerCancelStatus = 1021

func TestAccIBMComputeUser_Basic(t *testing.T) {
	t.Skip()
	var user datatypes.User_Customer

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeUserConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeUserExists("ibm_compute_user.testuser", &user),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "first_name", "first_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "last_name", "last_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "email", testAccRandomEmail),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "company_name", "company_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "address1", "1 Main St."),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "address2", "Suite 345"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "city", "Atlanta"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "state", "GA"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "country", "US"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "timezone", "EST"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "user_status", "ACTIVE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "password", hash(testAccUserPassword)),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "permissions.#", "2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "has_api_key", "true"),
					resource.TestMatchResourceAttr(
						"ibm_compute_user.testuser", "api_key", apiKeyRegexp),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_user.testuser", "username"),
				),
			},

			{
				Config: testAccCheckIBMComputeUserConfig_updated,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "first_name", "new_first_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "last_name", "new_last_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "email", "new"+testAccRandomEmail),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "company_name", "new_company_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "address1", "1 1st Avenue"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "address2", "Apartment 2"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "city", "Montreal"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "state", "QC"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "country", "CA"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "timezone", "MST"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "user_status", "INACTIVE"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "password", hash(testAccUserPassword)),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "permissions.#", "3"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "has_api_key", "false"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "api_key", ""),
					resource.TestCheckResourceAttrSet(
						"ibm_compute_user.testuser", "username"),
				),
			},
		},
	})
}

func TestAccIBMComputeUserWithTag(t *testing.T) {
	t.Skip()

	var user datatypes.User_Customer

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMComputeUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMComputeUserWithTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMComputeUserExists("ibm_compute_user.testuser", &user),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "first_name", "first_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "last_name", "last_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "email", testAccRandomEmail),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "tags.#", "2"),
				),
			},

			{
				Config: testAccCheckIBMComputeUserWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "first_name", "first_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "last_name", "last_name"),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "email", testAccRandomEmail),
					resource.TestCheckResourceAttr(
						"ibm_compute_user.testuser", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMComputeUserDestroy(s *terraform.State) error {
	client := services.GetUserCustomerService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_user" {
			continue
		}

		userID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the user
		user, err := client.Id(userID).Mask("userStatusId").GetObject()

		// Users are not immediately deleted, but rather placed into a 'cancel_pending' (1021) status
		if err != nil || *user.UserStatusId != userCustomerCancelStatus {
			return fmt.Errorf("IBM Cloud User still exists")
		}
	}

	return nil
}

func testAccCheckIBMComputeUserExists(n string, user *datatypes.User_Customer) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		userID, _ := strconv.Atoi(rs.Primary.ID)

		client := services.GetUserCustomerService(acc.TestAccProvider.Meta().(conns.ClientSession).SoftLayerSession())
		foundUser, err := client.Id(userID).GetObject()

		if err != nil {
			return err
		}

		if strconv.Itoa(int(*foundUser.Id)) != rs.Primary.ID {
			return fmt.Errorf("Record not found")
		}

		*user = foundUser

		return nil
	}
}

// Use session.New() to get a new session because the function should be called before testAccProvider is configured.
func testGetAccountId() string {
	service := services.GetAccountService(session.New())
	account, err := service.Mask("id").GetObject()
	if err != nil {
		return ""
	} else {
		return strconv.Itoa(*account.Id)
	}
}

var testAccCheckIBMComputeUserConfig_basic = fmt.Sprintf(`
resource "ibm_compute_user" "testuser" {
    first_name = "first_name"
    last_name = "last_name"
    email = "%s"
    company_name = "company_name"
    address1 = "1 Main St."
    address2 = "Suite 345"
    city = "Atlanta"
    state = "GA"
    country = "US"
    timezone = "EST"
    username = "%s"
    password = "%s"
    permissions = [
        "SERVER_ADD",
        "ACCESS_ALL_GUEST"
    ]
    has_api_key = true
}`, testAccRandomEmail, testAccRandomUser, testAccUserPassword)

var testAccCheckIBMComputeUserConfig_updated = fmt.Sprintf(`
resource "ibm_compute_user" "testuser" {
    first_name = "new_first_name"
    last_name = "new_last_name"
    email = "new%s"
    company_name = "new_company_name"
    address1 = "1 1st Avenue"
    address2 = "Apartment 2"
    city = "Montreal"
    state = "QC"
    country = "CA"
    timezone = "MST"
    user_status = "INACTIVE"
    username = "%s"
    password = "%s"
    permissions = [
        "SERVER_ADD",
        "ACCESS_ALL_HARDWARE",
        "TICKET_EDIT"
    ]
    has_api_key = false
}`, testAccRandomEmail, testAccRandomUser, testAccUserPassword)

var testAccRandomEmail = id.UniqueId() + "@example.com"
var testAccRandomUser = testGetAccountId() + "_" + testAccRandomEmail
var testAccUserPassword = "Softlayer2017!"
var apiKeyRegexp, _ = regexp.Compile(`\w+`)

// Function used by provider for hashing passwords
func hash(v interface{}) string {
	hash := sha1.Sum([]byte(v.(string)))
	return hex.EncodeToString(hash[:])
}

var testAccCheckIBMComputeUserWithTag = fmt.Sprintf(`
	resource "ibm_compute_user" "testuser" {
		first_name = "first_name"
		last_name = "last_name"
		email = "%s"
		company_name = "company_name"
		address1 = "1 Main St."
		address2 = "Suite 345"
		city = "Atlanta"
		state = "GA"
		country = "US"
		timezone = "EST"
		username = "%s"
		password = "%s"
		permissions = [
			"SERVER_ADD",
			"ACCESS_ALL_GUEST"
		]
		has_api_key = true
		tags = ["one", "two"]
	}`, testAccRandomEmail, testAccRandomUser, testAccUserPassword)

var testAccCheckIBMComputeUserWithUpdatedTag = fmt.Sprintf(`
		resource "ibm_compute_user" "testuser" {
			first_name = "first_name"
			last_name = "last_name"
			email = "%s"
			company_name = "company_name"
			address1 = "1 Main St."
			address2 = "Suite 345"
			city = "Atlanta"
			state = "GA"
			country = "US"
			timezone = "EST"
			username = "%s"
			password = "%s"
			permissions = [
				"SERVER_ADD",
				"ACCESS_ALL_GUEST"
			]
			has_api_key = true
			tags = ["one", "two", "three"]
		}`, testAccRandomEmail, testAccRandomUser, testAccUserPassword)

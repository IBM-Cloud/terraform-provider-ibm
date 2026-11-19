// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

var (
	tenantIdRegister = "wkk1yqrdce/"
)

func TestAccIbmBackupRecoverySourceRegistrationBasic(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationResponseParams

	environment := "kPhysical"
	connectionId := "5128356219792164864"
	endpoint := "172.26.202.5"
	hostType := "kLinux"
	physicalType := "kHost"
	applications := ""

	applicationsUpdate := `applications = ["kSQL"]`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoverySourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applications, endpoint, hostType, physicalType, connectionId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoverySourceRegistrationExists("ibm_backup_recovery_source_registration.baas_source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantIdRegister),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "environment", environment),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.baas_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.baas_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "connection_id", connectionId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.name", endpoint),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.os_type", hostType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.object_type", physicalType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.endpoint", endpoint),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.host_type", hostType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.physical_type", physicalType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applicationsUpdate, endpoint, hostType, physicalType, connectionId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantIdRegister),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "environment", environment),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applications, endpoint, hostType, physicalType string, connectionId string) string {
	return fmt.Sprintf(`
			resource "ibm_backup_recovery_source_registration" "baas_source_registration_instance" {
				
				x_ibm_tenant_id = "%s"
				environment = "%s"
				connection_id = "%s"
				physical_params {
				endpoint = "%s"
				host_type = "%s"
				physical_type = "%s"
				%s
				}
			}
	`, tenantIdRegister, environment, connectionId, endpoint, hostType, physicalType, applications)
}

func TestAccIbmBackupRecoveryKubernetesSourceRegistrationBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryKubernetesSourceRegistrationConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.ibm_backup_recovery_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.ibm_backup_recovery_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttrPair("ibm_backup_recovery_source_registration.ibm_backup_recovery_source_registration_instance", "auto_protection_group_id", "ibm_backup_recovery_source_registration.ibm_backup_recovery_source_registration_instance", "auto_proetction_group_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryKubernetesSourceRegistrationConfigBasic() string {
	return fmt.Sprintf(`
	resource "ibm_backup_recovery_source_registration" "ibm_backup_recovery_source_registration_instance" {
		x_ibm_tenant_id = "wkk1yqrdce/"
		environment = "kKubernetes"
		connection_id = "932952619841170816"
		name = "terra-iks-registration-1"
		kubernetes_params {
		  endpoint = "https://c104-e.us-east.containers.cloud.ibm.com:30587"
		  kubernetes_distribution = "kIKS"
		  cohesity_dataprotect_plugin_image_location = "icr.io/ext/brs/cohesity-dataprotect-plugin:7.2.15-p2"
		  velero_openshift_plugin_image_location = "icr.io/ext/brs/velero-plugin-for-openshift:7.2.15-p2"
		  velero_image_location = "icr.io/ext/brs/velero:7.2.15-p2"
		  velero_aws_plugin_image_location = "icr.io/ext/brs/velero-plugin-for-aws:7.2.15-p2"
		  data_mover_image_location = "icr.io/ext/brs/cohesity-datamover:7.2.15-p2"
		  client_private_key = ""
		}
	  }
	`)
}

func testAccCheckIbmBackupRecoverySourceRegistrationExists(n string, obj backuprecoveryv1.SourceRegistrationResponseParams) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}
		num, _ := strconv.Atoi(rs.Primary.Attributes["source_id"])
		getProtectionSourceRegistrationOptions.SetID(int64(num))
		getProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantIdRegister)

		sourceRegistrationReponseParams, _, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *sourceRegistrationReponseParams
		return nil
	}
}

func testAccCheckIbmBackupRecoverySourceRegistrationDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_source_registration" {
			continue
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

		num, _ := strconv.Atoi(rs.Primary.Attributes["source_id"])

		getProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantIdRegister)
		getProtectionSourceRegistrationOptions.SetID(int64(num))

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)

		if err == nil {
			return fmt.Errorf("baas_source_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_source_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestKubernetesSourceRegistrationParamsDatamoverHostportNumber(t *testing.T) {
	cases := []struct {
		name     string
		hcl      map[string]interface{}
		expected *int64
	}{
		{
			name: "hostport number omitted with kNodePort",
			hcl: map[string]interface{}{
				"endpoint":                "https://example.com:443",
				"kubernetes_distribution": "kROKS",
				"client_private_key":      "token", // pragma: allowlist secret
				"datamover_service_type":  "kNodePort",
			},
			expected: nil,
		},
		{
			name: "hostport number set with kHostPort",
			hcl: map[string]interface{}{
				"endpoint":                  "https://example.com:443",
				"kubernetes_distribution":   "kROKS",
				"client_private_key":        "token", // pragma: allowlist secret
				"datamover_service_type":    "kHostPort",
				"datamover_hostport_number": 31000,
			},
			expected: core.Int64Ptr(31000),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := backuprecovery.ResourceIbmBackupRecoverySourceRegistration()
			d := schema.TestResourceDataRaw(t, res.Schema, map[string]interface{}{
				"x_ibm_tenant_id":   "tenant/",
				"environment":       "kKubernetes",
				"connection_id":     "123",
				"kubernetes_params": []interface{}{tc.hcl},
			})

			modelMap := d.Get("kubernetes_params").([]interface{})[0].(map[string]interface{})
			model, err := backuprecovery.ResourceIbmBackupRecoverySourceRegistrationMapToKubernetesSourceRegistrationParams(modelMap)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if tc.expected == nil {
				if model.DatamoverHostportNumber != nil {
					t.Fatalf("expected DatamoverHostportNumber to be nil, got %d", *model.DatamoverHostportNumber)
				}
				return
			}
			if model.DatamoverHostportNumber == nil || *model.DatamoverHostportNumber != *tc.expected {
				t.Fatalf("expected DatamoverHostportNumber %d, got %v", *tc.expected, model.DatamoverHostportNumber)
			}
		})
	}
}

// Verifies registering a Kubernetes source with a non-kHostPort datamover service type
// and no datamover_hostport_number succeeds (datamoverHostportNumber must not be sent).
func TestAccIbmBackupRecoveryKubernetesSourceRegistrationNodePort(t *testing.T) {
	tenantId := os.Getenv("BRS_TENANT_ID")
	connectionId := os.Getenv("BRS_CONNECTION_ID")
	instanceId := os.Getenv("BRS_INSTANCE_ID")
	region := os.Getenv("BRS_REGION")
	clusterEndpoint := os.Getenv("BRS_K8S_ENDPOINT")
	clusterToken := os.Getenv("BRS_K8S_TOKEN")
	serviceName := os.Getenv("BRS_SERVICE_NAME")
	if serviceName == "" {
		serviceName = "backup-recovery"
	}
	distribution := os.Getenv("BRS_K8S_DISTRIBUTION")
	if distribution == "" {
		distribution = "kROKS"
	}
	if tenantId == "" || connectionId == "" || instanceId == "" || region == "" || clusterEndpoint == "" || clusterToken == "" {
		t.Skip("set BRS_TENANT_ID, BRS_CONNECTION_ID, BRS_INSTANCE_ID, BRS_REGION, BRS_K8S_ENDPOINT, BRS_K8S_TOKEN to run")
	}
	name := "ibm_backup_recovery_source_registration.k8s_nodeport"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				resource "ibm_backup_recovery_source_registration" "k8s_nodeport" {
					x_ibm_tenant_id = "%s"
					environment     = "kKubernetes"
					connection_id   = "%s"
					endpoint_type   = "public"
					instance_id     = "%s"
					region          = "%s"
					service_name    = "%s"
					kubernetes_params {
						endpoint                = "%s"
						kubernetes_distribution = "%s"
						kubernetes_type         = "kCluster"
						client_private_key      = "%s" # pragma: allowlist secret
						datamover_service_type  = "kNodePort"
					}
				}`, tenantId, connectionId, instanceId, region, serviceName, clusterEndpoint, distribution, clusterToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
					resource.TestCheckResourceAttr(name, "kubernetes_params.0.datamover_service_type", "kNodePort"),
				),
			},
		},
	})
}

// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"

	// "github.com/IBM/dra-go-sdk/drautomationservicev1"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func ResourceIBMPdrManagedr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPdrManagedrCreate,
		ReadContext:   resourceIBMPdrManagedrRead,
		DeleteContext: resourceIBMPdrManagedrDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "instance_id"),
				Description: "Service Instance ID.",
			},
			"stand_by_redeploy": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					allowed := []string{"true", "false"}
					for _, a := range allowed {
						if v == a {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%q must be one of %v, got: %s", key, allowed, v))
					return
				},
				Description: "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\").",
			},
			"accept_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "accept_language"),
				Description: "The language requested for the return document.",
			},
			"accepts_incomplete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.",
			},
			"orchestrator_location_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "orchestrator_location_type"),
				Description: "The cloud location where your orchestator need to be created.",
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "location_id"),
				Description: "The location or data center identifier where the service instance is deployed.",
			},
			"ssh_key_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "ssh_key_name"),
				Description: "The name of the SSH key used for deploying the orchestator.",
			},
			"standby_ssh_key_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "standby_ssh_key_name"),
				Description: "The name of the SSH key used for deploying the standby orchestator.",
			},
			"orchestrator_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "orchestrator_name"),
				Description: "The username used for the orchestrator.",
			},
			"orchestrator_workspace_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "orchestrator_workspace_id"),
				Description: "The unique identifier orchestrator workspace.",
			},
			"standby_orchestrator_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "standby_orchestrator_name"),
				Description: "The username for the standby orchestrator management interface.",
			},
			"standby_orchestrator_workspace_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "standby_orchestrator_workspace_id"),
				Description: "The unique identifier of the standby orchestrator workspace.",
			},
			"orchestrator_ha": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether the orchestrator High Availability (HA) is enabled for the service instance.",
			},
			"resource_instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "resource_instance"),
				Description: "The uniquie identifier of the associated IBM Cloud resource instance.",
			},
			"secret_group": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "secret_group"),
				Description: "The secret group name in IBM Cloud Secrets Manager containing sensitive data for the service instance.",
			},
			"secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "secret"),
				Description: "The secret name or identifier used for retrieving credentials from secrets manager.",
			},
			"region_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "region_id"),
				Description: "The power virtual server region where the service instance is deployed.",
			},
			"guid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "guid"),
				Description: "The global unique identifier of the service instance.",
			},
			"machine_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "machine_type"),
				Description: "The machine type used for deploying orchestrator.",
			},
			"tier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "tier"),
				Description: "The storage tier used for deploying primary orchestrator.",
			},
			"standby_tier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "standby_tier"),
				Description: "The storage tier used for deploying standby orchestrator.",
			},
			"standby_machine_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "standby_machine_type"),
				Description: "The machine type used for deploying standby virtual machines.",
			},
			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "tenant_name"),
				Description: "The tenant name for MFA authentication API.",
			},
			"proxy_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "proxy_ip"),
				Description: "Proxy IP for the Communication between Orchestrator and Service broker.",
			},
			"primary_orch_ca_certs_server_cert": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Encoded content of pem file of primary orchestrator.",
			},

			"primary_orch_ca_certs_server_key": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Encoded content of key file of primary orchestrator.",
			},

			"standby_orch_ca_certs_server_cert": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Encoded content of pem file of standby orchestrator.",
			},

			"standby_orch_ca_certs_server_key": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "Encoded content of key file of standby orchestrator.",
			},

			"api_key": {
				Type:             schema.TypeString,
				Sensitive:        true,
				ForceNew:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Optional:         true,
				Description:      "The API key associated with the IBM Cloud service instance.",
			},

			"orchestrator_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				ForceNew:  true,
				// Required:         true,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "The password that you can use to access your orchestrator.",
			},

			"client_id": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "The Client ID created for MFA authentication API.",
			},

			"client_secret": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				Sensitive:        true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "The client secret created for MFA authentication API.",
			},

			"managed_apikey": &schema.Schema{
				Type:             schema.TypeString,
				Sensitive:        true,
				ForceNew:         true,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				Description:      "API key used to manage the workloads by adding the PowerVS instances to the orchestrator.",
			},
			"orchestrator_network_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of network IDs for primary orchestrator VM.",
			},
			"standby_orchestrator_network_ids": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of network IDs for standby orchestrator VM.",
			},
			"primary_orch_ca_certs_secrets": &schema.Schema{
				Type:             schema.TypeList,
				ForceNew:         true,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_certificate_secret_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "ID of the primary CA certificate secret.",
						},
						"ca_secret_manager_guid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "ID of the primary CA certificate secret.",
						},
					},
				},
			},
			"standby_orch_ca_certs_secrets": &schema.Schema{
				Type:             schema.TypeList,
				ForceNew:         true,
				Optional:         true,
				DiffSuppressFunc: flex.ApplyOnce,
				MaxItems:         1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ca_certificate_secret_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "ID of the standby CA certificate secret.",
						},
						"ca_secret_manager_guid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "ID of the standby CA certificate secret.",
						},
					},
				},
			},
			"dashboard_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the dashboard for managing the DR service instance in IBM Cloud.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN (Cloud Resource Name) of the DR service instance.",
			},
		},
	}
}

func ResourceIBMPdrManagedrValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_:\/.]+$`,
			MinValueLength:             1,
			MaxValueLength:             512,
		},
		validate.ValidateSchema{
			Identifier:                 "stand_by_redeploy",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(true|false)$`,
			MinValueLength:             1,
			MaxValueLength:             5,
		},
		validate.ValidateSchema{
			Identifier:                 "accept_language",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
		validate.ValidateSchema{
			Identifier:                 "orchestrator_location_type",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "location_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "ssh_key_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "standby_ssh_key_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "orchestrator_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "orchestrator_workspace_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "standby_orchestrator_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "standby_orchestrator_workspace_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "resource_instance",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^crn:v1:[a-zA-Z0-9\-_]+:public:resource-controller:[a-zA-Z0-9\-_]+:[a-zA-Z0-9\-_\/]+:[a-zA-Z0-9\-_]+::$`,
			MinValueLength:             1,
			MaxValueLength:             512,
		},
		validate.ValidateSchema{
			Identifier:                 "secret_group",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "secret",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "region_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "guid",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "machine_type",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "tier",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "standby_tier",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "standby_machine_type",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_]+$`,
			MinValueLength:             1,
			MaxValueLength:             32,
		},
		validate.ValidateSchema{
			Identifier:                 "tenant_name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9.\-]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "proxy_ip",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9.\-_:]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pdr_managedr", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPdrManagedrCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createManageDrOptions := &drautomationservicev1.CreateManageDrOptions{}

	createManageDrOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("primary_orch_ca_certs_server_cert"); ok {
		createManageDrOptions.SetPrimaryOrchCaCertsServerCert(d.Get("primary_orch_ca_certs_server_cert").(string))
	}
	if _, ok := d.GetOk("primary_orch_ca_certs_server_key"); ok {
		createManageDrOptions.SetPrimaryOrchCaCertsServerKey(d.Get("primary_orch_ca_certs_server_key").(string))
	}
	if _, ok := d.GetOk("standby_orch_ca_certs_server_cert"); ok {
		createManageDrOptions.SetStandbyOrchCaCertsServerCert(d.Get("standby_orch_ca_certs_server_cert").(string))
	}
	if _, ok := d.GetOk("standby_orch_ca_certs_server_key"); ok {
		createManageDrOptions.SetStandbyOrchCaCertsServerKey(d.Get("standby_orch_ca_certs_server_key").(string))
	}
	if _, ok := d.GetOk("api_key"); ok {
		createManageDrOptions.SetAPIKey(d.Get("api_key").(string))
	}
	if _, ok := d.GetOk("orchestrator_location_type"); ok {
		createManageDrOptions.SetOrchestratorLocationType(d.Get("orchestrator_location_type").(string))
	}
	if _, ok := d.GetOk("location_id"); ok {
		createManageDrOptions.SetLocationID(d.Get("location_id").(string))
	}
	if _, ok := d.GetOk("ssh_key_name"); ok {
		createManageDrOptions.SetSSHKeyName(d.Get("ssh_key_name").(string))
	}
	if _, ok := d.GetOk("standby_ssh_key_name"); ok {
		createManageDrOptions.SetStandbySSHKeyName(d.Get("standby_ssh_key_name").(string))
	}
	if _, ok := d.GetOk("orchestrator_name"); ok {
		createManageDrOptions.SetOrchestratorName(d.Get("orchestrator_name").(string))
	}
	if _, ok := d.GetOk("orchestrator_password"); ok {
		createManageDrOptions.SetOrchestratorPassword(d.Get("orchestrator_password").(string))
	}
	if _, ok := d.GetOk("orchestrator_workspace_id"); ok {
		createManageDrOptions.SetOrchestratorWorkspaceID(d.Get("orchestrator_workspace_id").(string))
	}
	if _, ok := d.GetOk("standby_orchestrator_name"); ok {
		createManageDrOptions.SetStandbyOrchestratorName(d.Get("standby_orchestrator_name").(string))
	}
	if _, ok := d.GetOk("standby_orchestrator_workspace_id"); ok {
		createManageDrOptions.SetStandbyOrchestratorWorkspaceID(d.Get("standby_orchestrator_workspace_id").(string))
	}
	if _, ok := d.GetOk("orchestrator_ha"); ok {
		createManageDrOptions.SetOrchestratorHa(d.Get("orchestrator_ha").(bool))
	}
	if _, ok := d.GetOk("resource_instance"); ok {
		createManageDrOptions.SetResourceInstance(d.Get("resource_instance").(string))
	}
	if _, ok := d.GetOk("secret_group"); ok {
		createManageDrOptions.SetSecretGroup(d.Get("secret_group").(string))
	}
	if _, ok := d.GetOk("secret"); ok {
		createManageDrOptions.SetSecret(d.Get("secret").(string))
	}
	if _, ok := d.GetOk("region_id"); ok {
		createManageDrOptions.SetRegionID(d.Get("region_id").(string))
	}
	if _, ok := d.GetOk("guid"); ok {
		createManageDrOptions.SetGUID(d.Get("guid").(string))
	}
	if _, ok := d.GetOk("machine_type"); ok {
		createManageDrOptions.SetMachineType(d.Get("machine_type").(string))
	}
	if _, ok := d.GetOk("tier"); ok {
		createManageDrOptions.SetTier(d.Get("tier").(string))
	}
	if _, ok := d.GetOk("standby_tier"); ok {
		createManageDrOptions.SetStandbyTier(d.Get("standby_tier").(string))
	}
	if _, ok := d.GetOk("standby_machine_type"); ok {
		createManageDrOptions.SetStandbyMachineType(d.Get("standby_machine_type").(string))
	}
	if _, ok := d.GetOk("client_id"); ok {
		createManageDrOptions.SetClientID(d.Get("client_id").(string))
	}
	if _, ok := d.GetOk("client_secret"); ok {
		createManageDrOptions.SetClientSecret(d.Get("client_secret").(string))
	}
	if _, ok := d.GetOk("tenant_name"); ok {
		createManageDrOptions.SetTenantName(d.Get("tenant_name").(string))
	}
	if _, ok := d.GetOk("proxy_ip"); ok {
		createManageDrOptions.SetProxyIP(d.Get("proxy_ip").(string))
	}
	if _, ok := d.GetOk("orchestrator_network_ids"); ok {
		var orchestratorNetworkIds []string
		for _, v := range d.Get("orchestrator_network_ids").([]interface{}) {
			orchestratorNetworkIdsItem := v.(string)
			orchestratorNetworkIds = append(orchestratorNetworkIds, orchestratorNetworkIdsItem)
		}
		createManageDrOptions.SetOrchestratorNetworkIds(orchestratorNetworkIds)
	}
	if _, ok := d.GetOk("standby_orchestrator_network_ids"); ok {
		var standbyOrchestratorNetworkIds []string
		for _, v := range d.Get("standby_orchestrator_network_ids").([]interface{}) {
			standbyOrchestratorNetworkIdsItem := v.(string)
			standbyOrchestratorNetworkIds = append(standbyOrchestratorNetworkIds, standbyOrchestratorNetworkIdsItem)
		}
		createManageDrOptions.SetStandbyOrchestratorNetworkIds(standbyOrchestratorNetworkIds)
	}
	if _, ok := d.GetOk("managed_apikey"); ok {
		createManageDrOptions.SetManagedApikey(d.Get("managed_apikey").(string))
	}
	if _, ok := d.GetOk("primary_orch_ca_certs_secrets"); ok {
		primaryOrchCaCertsSecretsModel, err := ResourceIBMPdrManagedrMapToPrimaryOrchCaSecrets(d.Get("primary_orch_ca_certs_secrets.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "parse-primary_orch_ca_certs_secrets").GetDiag()
		}
		createManageDrOptions.SetPrimaryOrchCaCertsSecrets(primaryOrchCaCertsSecretsModel)
	}
	if _, ok := d.GetOk("standby_orch_ca_certs_secrets"); ok {
		standbyOrchCaCertsSecretsModel, err := ResourceIBMPdrManagedrMapToStandbyOrchCaSecrets(d.Get("standby_orch_ca_certs_secrets.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "parse-standby_orch_ca_certs_secrets").GetDiag()
		}
		createManageDrOptions.SetStandbyOrchCaCertsSecrets(standbyOrchCaCertsSecretsModel)
	}
	if _, ok := d.GetOk("stand_by_redeploy"); ok {
		createManageDrOptions.SetStandByRedeploy(d.Get("stand_by_redeploy").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("accepts_incomplete"); ok {
		createManageDrOptions.SetAcceptsIncomplete(d.Get("accepts_incomplete").(bool))
	}

	_, response, err := drAutomationServiceClient.CreateManageDrWithContext(context, createManageDrOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateManageDrWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateManageDrWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_managedr", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s", *createManageDrOptions.InstanceID))
	instanceID := *createManageDrOptions.InstanceID
	const (
		pollInterval = 1 * time.Minute
		maxWaitTime  = 90 * time.Minute // optional, can extend as needed
	)
	timeout := time.After(maxWaitTime)
	ticker := time.NewTicker(pollInterval)
	// defer ticker.Stop()

	log.Printf("[INFO] Started polling last operation status for instance %s every %s", instanceID, pollInterval)
	enableha, _ := d.GetOk("orchestrator_ha")

	if !enableha.(bool) {
		for {
			select {
			case <-timeout:
				errMsg := fmt.Sprintf("Timeout exceeded while waiting for Manage DR to become Active (instance_id: %s)", instanceID)
				tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")
				log.Printf("[ERROR] %s", errMsg)
				ticker.Stop()
				return tfErr.GetDiag()

			case <-ticker.C:
				status, _, _, statusErr, _ := checkLastOperationStatus(context, drAutomationServiceClient, instanceID)
				if statusErr != nil {
					tfErr := flex.TerraformErrorf(statusErr, fmt.Sprintf("GetLastOperation failed: %s", statusErr.Error()), "ibm_pdr_managedr", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					ticker.Stop()
					return tfErr.GetDiag()
				}

				log.Printf("[INFO] Current Last Operation status for instance %s: %s", instanceID, *status.Status)

				switch strings.ToLower(*status.Status) {
				case "active":
					log.Printf("[INFO] Manage DR operation completed successfully for instance %s", instanceID)
					ticker.Stop()
					return resourceIBMPdrManagedrRead(context, d, meta)

				case "fail", "failed", "error":
					errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.PrimaryDescription)
					tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

					log.Printf("[ERROR] %s", errMsg)
					ticker.Stop()
					return tfErr.GetDiag()

				default:
					log.Printf("[DEBUG] Manage DR still in progress... retrying in %v", pollInterval)
				}
			}
		}
	} else {
		for {
			select {
			case <-timeout:
				errMsg := fmt.Sprintf("Timeout exceeded while waiting for Manage DR to become Active (instance_id: %s)", instanceID)
				tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")
				log.Printf("[ERROR] %s", errMsg)
				ticker.Stop()
				return tfErr.GetDiag()

			case <-ticker.C:
				status, _, standbyStatus, statusErr, _ := checkLastOperationStatus(context, drAutomationServiceClient, instanceID)
				if statusErr != nil {
					tfErr := flex.TerraformErrorf(statusErr, fmt.Sprintf("GetLastOperation failed: %s", statusErr.Error()), "ibm_pdr_managedr", "create")
					log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
					ticker.Stop()
					return tfErr.GetDiag()
				}

				switch strings.ToLower(*status.Status) {
				case "active":
					if strings.ToLower(standbyStatus) == "active" {
						log.Printf("[INFO] Manage DR operation completed successfully for instance %s (both primary and standby active)", instanceID)
						ticker.Stop()
						return resourceIBMPdrManagedrRead(context, d, meta)
					}
					if strings.ToLower(standbyStatus) == "failed" {
						errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.StandbyDescription)
						tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

						log.Printf("[ERROR] %s", errMsg)
						ticker.Stop()
						return tfErr.GetDiag()
					}

					// If standby still initializing
					log.Printf("[INFO] Manage DR overall status is Active, but standby orchestrator still in progress (status: %s). Retrying in %v...",
						standbyStatus, pollInterval)
					continue

				case "fail", "failed", "error":
					errMsg := fmt.Sprintf("Manage DR operation failed for instance %s and error message: %s", instanceID, *status.PrimaryDescription)
					tfErr := flex.TerraformErrorf(fmt.Errorf("%s", errMsg), errMsg, "ibm_pdr_managedr", "create")

					log.Printf("[ERROR] %s", errMsg)
					ticker.Stop()
					return tfErr.GetDiag()

				default:
					log.Printf("[DEBUG] Manage DR still in progress ... retrying in %v", pollInterval)
				}
			}
		}
	}

	// return resourceIBMPdrManagedrRead(context, d, meta)
}

func resourceIBMPdrManagedrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getManageDrOptions := &drautomationservicev1.GetManageDrOptions{}

	instanceID := d.Id()

	log.Printf("[DEBUG] Read operation using instance ID from resource: %s", instanceID)

	getManageDrOptions.SetInstanceID(instanceID)

	// if _, ok := d.GetOk("accept_language"); ok {
	// 	getManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	// }

	serviceInstanceManageDr, response, err := drAutomationServiceClient.GetManageDrWithContext(context, getManageDrOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetManageDrWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetManageDrWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_managedr", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	if !core.IsNil(serviceInstanceManageDr.OrchestratorLocationType) {
		if err = d.Set("orchestrator_location_type", serviceInstanceManageDr.OrchestratorLocationType); err != nil {
			err = fmt.Errorf("Error setting orchestrator_location_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-orchestrator_location_type").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.LocationID) {
		if err = d.Set("location_id", serviceInstanceManageDr.LocationID); err != nil {
			err = fmt.Errorf("Error setting location_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-location_id").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.SSHKeyName) {
		if err = d.Set("ssh_key_name", serviceInstanceManageDr.SSHKeyName); err != nil {
			err = fmt.Errorf("Error setting ssh_key_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-ssh_key_name").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.StandbySSHKeyName) {
		if err = d.Set("standby_ssh_key_name", serviceInstanceManageDr.StandbySSHKeyName); err != nil {
			err = fmt.Errorf("Error setting standby_ssh_key_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-standby_ssh_key_name").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.OrchestratorName) {
		if err = d.Set("orchestrator_name", serviceInstanceManageDr.OrchestratorName); err != nil {
			err = fmt.Errorf("Error setting orchestrator_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-orchestrator_name").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.OrchestratorWorkspaceID) {
		if err = d.Set("orchestrator_workspace_id", serviceInstanceManageDr.OrchestratorWorkspaceID); err != nil {
			err = fmt.Errorf("Error setting orchestrator_workspace_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-orchestrator_workspace_id").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.StandbyOrchestratorName) {
		if err = d.Set("standby_orchestrator_name", serviceInstanceManageDr.StandbyOrchestratorName); err != nil {
			err = fmt.Errorf("Error setting standby_orchestrator_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-standby_orchestrator_name").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.StandbyOrchestratorWorkspaceID) {
		if err = d.Set("standby_orchestrator_workspace_id", serviceInstanceManageDr.StandbyOrchestratorWorkspaceID); err != nil {
			err = fmt.Errorf("Error setting standby_orchestrator_workspace_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-standby_orchestrator_workspace_id").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.OrchestratorHa) {
		if err = d.Set("orchestrator_ha", serviceInstanceManageDr.OrchestratorHa); err != nil {
			err = fmt.Errorf("Error setting orchestrator_ha: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-orchestrator_ha").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ResourceInstance) {
		if err = d.Set("resource_instance", serviceInstanceManageDr.ResourceInstance); err != nil {
			err = fmt.Errorf("Error setting resource_instance: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-resource_instance").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.SecretGroup) {
		if err = d.Set("secret_group", serviceInstanceManageDr.SecretGroup); err != nil {
			err = fmt.Errorf("Error setting secret_group: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-secret_group").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.Secret) {
		if err = d.Set("secret", serviceInstanceManageDr.Secret); err != nil {
			err = fmt.Errorf("Error setting secret: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-secret").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.RegionID) {
		if err = d.Set("region_id", serviceInstanceManageDr.RegionID); err != nil {
			err = fmt.Errorf("Error setting region_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-region_id").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.GUID) {
		if err = d.Set("guid", serviceInstanceManageDr.GUID); err != nil {
			err = fmt.Errorf("Error setting guid: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-guid").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.MachineType) {
		if err = d.Set("machine_type", serviceInstanceManageDr.MachineType); err != nil {
			err = fmt.Errorf("Error setting machine_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-machine_type").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.Tier) {
		if err = d.Set("tier", serviceInstanceManageDr.Tier); err != nil {
			err = fmt.Errorf("Error setting tier: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-tier").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.StandbyTier) {
		if err = d.Set("standby_tier", serviceInstanceManageDr.StandbyTier); err != nil {
			err = fmt.Errorf("Error setting standby_tier: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-standby_tier").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.StandbyMachineType) {
		if err = d.Set("standby_machine_type", serviceInstanceManageDr.StandbyMachineType); err != nil {
			err = fmt.Errorf("Error setting standby_machine_type: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-standby_machine_type").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.TenantName) {
		if err = d.Set("tenant_name", serviceInstanceManageDr.TenantName); err != nil {
			err = fmt.Errorf("Error setting tenant_name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-tenant_name").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ProxyIP) {
		if err = d.Set("proxy_ip", serviceInstanceManageDr.ProxyIP); err != nil {
			err = fmt.Errorf("Error setting proxy_ip: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-proxy_ip").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.DashboardURL) {
		if err = d.Set("dashboard_url", serviceInstanceManageDr.DashboardURL); err != nil {
			err = fmt.Errorf("Error setting dashboard_url: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-dashboard_url").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ID) {
		if err = d.Set("instance_id", serviceInstanceManageDr.ID); err != nil {
			err = fmt.Errorf("Error setting instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-instance_id").GetDiag()
		}
	}

	return nil
}

func resourceIBMPdrManagedrDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIBMPdrManagedrMapToPrimaryOrchCaSecrets(modelMap map[string]interface{}) (*drautomationservicev1.PrimaryOrchCaSecrets, error) {
	model := &drautomationservicev1.PrimaryOrchCaSecrets{}
	if modelMap["ca_certificate_secret_id"] != nil && modelMap["ca_certificate_secret_id"].(string) != "" {
		model.CaCertificateSecretID = core.StringPtr(modelMap["ca_certificate_secret_id"].(string))
	}
	if modelMap["ca_secret_manager_guid"] != nil && modelMap["ca_secret_manager_guid"].(string) != "" {
		model.CaSecretManagerGUID = core.StringPtr(modelMap["ca_secret_manager_guid"].(string))
	}
	return model, nil
}

func ResourceIBMPdrManagedrMapToStandbyOrchCaSecrets(modelMap map[string]interface{}) (*drautomationservicev1.StandbyOrchCaSecrets, error) {
	model := &drautomationservicev1.StandbyOrchCaSecrets{}
	if modelMap["ca_certificate_secret_id"] != nil && modelMap["ca_certificate_secret_id"].(string) != "" {
		model.CaCertificateSecretID = core.StringPtr(modelMap["ca_certificate_secret_id"].(string))
	}
	return model, nil
}

func checkLastOperationStatus(ctx context.Context, client *drautomationservicev1.DrAutomationServiceV1, instanceID string) (*drautomationservicev1.ServiceInstanceStatus, string, string, error, error) {
	opts := &drautomationservicev1.GetLastOperationOptions{}
	opts.SetInstanceID(instanceID)

	statusResponse, _, err := client.GetLastOperationWithContext(ctx, opts)
	if err != nil {
		return nil, "", "", err, nil
	}

	if statusResponse.Status == nil {
		return statusResponse, "", "", fmt.Errorf("received nil status for instance %s", instanceID), nil
	}

	status := strings.ToLower(*statusResponse.Status)
	primaryStatus := strings.ToLower(*statusResponse.PrimaryOrchestratorStatus)
	standbyStatus := strings.ToLower(*statusResponse.StandbyStatus)

	// --- Custom error logic based on your conditions ---
	if status == "failed" {
		switch {
		case primaryStatus == "failed" && standbyStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s \n %s", *statusResponse.PrimaryDescription, *statusResponse.StandbyDescription)
		case primaryStatus == "failed" && (standbyStatus == "" || standbyStatus == "na"):
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s", *statusResponse.PrimaryDescription)
		case primaryStatus == "active" && (standbyStatus != "" || standbyStatus == "failed"):
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("%s \n %s", *statusResponse.PrimaryDescription, *statusResponse.StandbyDescription)
		case primaryStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("primary orchestrator failed for instance %s", instanceID)
		case standbyStatus == "failed":
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("standby orchestrator failed for instance %s", instanceID)
		default:
			return statusResponse, primaryStatus, standbyStatus, nil, fmt.Errorf("operation failed for instance %s with unknown cause", instanceID)
		}
	}

	return statusResponse, primaryStatus, standbyStatus, nil, nil
}

// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMContainerIngressSecret() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerIngressSecretCreate,
		Read:     resourceIBMContainerIngressSecretRead,
		Update:   resourceIBMContainerIngressSecretUpdate,
		Delete:   resourceIBMContainerIngressSecretDelete,
		Exists:   resourceIBMContainerIngressSecretExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID",
				ValidateFunc: validate.InvokeValidator(
					"ibm_container_ingress_secret",
					"cluster"),
			},
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secret name",
				ForceNew:    true,
			},
			"secret_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Secret namespace",
			},
			"secret_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type TLS or opaque",
			},
			"tls_secret": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "TLS secret",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cert_crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Certificate CRN",
						},
						"persistence": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Persistence of secret",
						},
						"domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name",
						},
						"expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate expires on date",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secret Status",
						},
						"user_managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the secret was created by the user",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of Secret Manager secret",
						},

						"last_updated_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp secret was last updated",
						},
					},
				},
			},
			"opaque_secret_fields": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Fields of an opaque secret",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Secret CRN corresponding to the field",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Field name",
						},
						"prefix": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Prefix field name with Secrets Manager secret name",
						},
						"expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field expires on date",
						},
						"secret_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secrets manager secret type",
						},
						"last_updated_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field last updated timestamp",
						},
					},
				},
			},
		},
	}
}

func ResourceIBMContainerIngressSecretValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cluster",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			Required:                   true,
			CloudDataType:              "cluster",
			CloudDataRange:             []string{"resolved_to:id"}})

	iBMContainerIngressInstanceValidator := validate.ResourceValidator{ResourceName: "ibm_container_ingress_secret", Schema: validateSchema}
	return &iBMContainerIngressInstanceValidator
}

func resourceIBMContainerIngressSecretCreate(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	cluster := d.Get("cluster").(string)
	secretName := d.Get("secret_name").(string)
	secretNamespace := d.Get("secret_namespace").(string)
	secretType := d.Get("secret_type").(string)

	params := v2.SecretCreateConfig{
		Cluster:   cluster,
		Name:      secretName,
		Namespace: secretNamespace,
		Type:      secretType,
	}

	if tlsSecret, ok := d.GetOk("tls_secret"); ok {
		tlsSecretList := tlsSecret.(*schema.Set).List()

		tlsSecretMap := tlsSecretList[0].(map[string]interface{})

		if certCRN, ok := tlsSecretMap["cert_crn"]; ok {
			params.CRN = certCRN.(string)
		}

		if persistence, ok := tlsSecretMap["persistence"]; ok {
			params.Persistence = persistence.(bool)
		}
	}

	ingressAPI := ingressClient.Ingresses()
	response, err := ingressAPI.CreateIngressSecret(params)

	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", cluster, response.Name, response.Namespace))

	return resourceIBMContainerIngressSecretRead(d, meta)
}

func resourceIBMContainerIngressSecretRead(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	secretName := parts[1]
	secretNamespace := parts[1]

	ingressAPI := ingressClient.Ingresses()
	ingressSecretConfig, err := ingressAPI.GetIngressSecret(cluster, secretName, secretNamespace)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", cluster, secretName, secretNamespace))
	d.Set("cluster", ingressSecretConfig.Cluster)
	d.Set("secret_name", ingressSecretConfig.Name)
	d.Set("secret_namespace", ingressSecretConfig.Namespace)
	d.Set("secret_type", ingressSecretConfig.SecretType)

	if ingressSecretConfig.Type == "TLS" && ingressSecretConfig.CRN != "" {
		tlsSecret := make(map[string]interface{})
		tlsSecret["cert_crn"] = ingressSecretConfig.CRN
		tlsSecret["persistence"] = ingressSecretConfig.Persistence
		tlsSecret["domain"] = ingressSecretConfig.Domain
		tlsSecret["expires_on"] = ingressSecretConfig.ExpiresOn
		tlsSecret["status"] = ingressSecretConfig.Status
		tlsSecret["user_managed"] = ingressSecretConfig.UserManaged
		tlsSecret["type"] = ingressSecretConfig.ExpiresOn
		tlsSecret["last_updated_timestamp"] = ingressSecretConfig.ExpiresOn
		d.Set("tls_secret", []map[string]interface{}{tlsSecret})
	}

	return nil
}

func resourceIBMContainerIngressSecretDelete(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	ingressAPI := ingressClient.Ingresses()

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	secretName := parts[1]
	secretNamespace := parts[1]

	params := v2.SecretDeleteConfig{
		Cluster:   cluster,
		Name:      secretName,
		Namespace: secretNamespace,
	}

	err = ingressAPI.DeleteIngressSecret(params)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMContainerIngressSecretUpdate(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	secretName := parts[1]
	secretNamespace := parts[1]

	params := v2.SecretUpdateConfig{
		Cluster:   cluster,
		Name:      secretName,
		Namespace: secretNamespace,
	}

	if d.HasChange("tls_secret") {
		tlsSecretList := d.Get("tls_secret").(*schema.Set).List()
		if len(tlsSecretList) != 1 {
			return fmt.Errorf("[ERROR] Only one TLS ingress secret accepted")
		}

		tlsSecretMap := tlsSecretList[0].(map[string]interface{})

		if certCRN, ok := tlsSecretMap["cert_crn"]; ok {
			params.CRN = certCRN.(string)
		}
	}

	ingressAPI := ingressClient.Ingresses()
	response, err := ingressAPI.UpdateIngressSecret(params)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", cluster, response.Name, response.Namespace))
	return resourceIBMContainerIngressSecretRead(d, meta)
}

func resourceIBMContainerIngressSecretExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}

	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return false, err
	}
	cluster := parts[0]
	secretName := parts[1]
	secretNamespace := parts[1]

	ingressAPI := ingressClient.Ingresses()
	ingressSecretConfig, err := ingressAPI.GetIngressSecret(cluster, secretName, secretNamespace)

	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("[ERROR] Error getting ingress secret: %s", err)
	}

	return ingressSecretConfig.Cluster == cluster && ingressSecretConfig.Name == secretName && ingressSecretConfig.Namespace == secretNamespace, nil
}

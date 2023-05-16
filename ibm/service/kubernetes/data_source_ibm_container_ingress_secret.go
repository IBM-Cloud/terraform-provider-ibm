// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMContainerIngressSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerIngressSecretRead,
		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster ID",
				ValidateFunc: validate.InvokeValidator(
					"ibm_container_ingress_secret",
					"cluster"),
			},
			"secret_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secret name",
			},
			"secret_namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secret namespace",
			},
			"secret_type": {
				Type:        schema.TypeString,
				Computed:    true,
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
							Computed:    true,
							Description: "Certificate CRN",
						},
						"persistence": {
							Type:        schema.TypeBool,
							Computed:    true,
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
			"fields": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Fields of the secret",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Field name",
						},
						"crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Field crn",
						},
						"expires_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field expires on date",
						},
						"secret_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field secret type",
						},
						"last_updated_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field expires on date",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMContainerIngressSecretValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cluster",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			Required:                   true,
			CloudDataType:              "cluster",
			CloudDataRange:             []string{"resolved_to:id"}})

	iBMContainerIngressSecretValidator := validate.ResourceValidator{ResourceName: "ibm_container_ingress_secret", Schema: validateSchema}
	return &iBMContainerIngressSecretValidator
}

func dataSourceIBMContainerIngressSecretRead(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	clusterID := d.Get("cluster").(string)
	name := d.Get("secret_name").(string)
	namespace := d.Get("secret_namespace").(string)

	ingressAPI := ingressClient.Ingresses()
	ingressSecretConfig, err := ingressAPI.GetIngressSecret(clusterID, name, namespace)
	if err != nil {
		return err
	}

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

	d.SetId(fmt.Sprintf("%s/%s/%s", clusterID, name, namespace))

	return nil
}

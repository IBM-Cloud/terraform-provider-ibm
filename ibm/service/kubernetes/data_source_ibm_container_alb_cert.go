// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMContainerALBCert() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMContainerALBCertRead,

		Schema: map[string]*schema.Schema{
			"cert_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate CRN id",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID",
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_container_alb_cert",
					"cluster_id"),
			},
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secret name",
			},
			"namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ibm-cert-store",
				Description: "Namespace of the secret",
			},
			"persistence": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Persistence of secret",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secret Status",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain name",
			},
			"expires_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate expaire on date",
			},
			"issuer_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "certificate issuer name",
				Deprecated:  "This field is depricated and is not available in v2 version of ingress api",
			},
			"cluster_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cluster CRN",
				Deprecated:  "This field is depricated and is not available in v2 version of ingress api",
			},
			"cloud_cert_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cloud cert instance ID",
			},
		},
	}
}

func DataSourceIBMContainerALBCertValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cluster_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			Required:                   true,
			CloudDataType:              "cluster",
			CloudDataRange:             []string{"resolved_to:id"}})

	iBMContainerALBCertValidator := validate.ResourceValidator{ResourceName: "ibm_container_alb_cert", Schema: validateSchema}
	return &iBMContainerALBCertValidator
}
func dataSourceIBMContainerALBCertRead(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	clusterID := d.Get("cluster_id").(string)
	secretName := d.Get("secret_name").(string)
	namespace := d.Get("namespace").(string)

	ingressAPI := ingressClient.Ingresses()
	ingressSecretConfig, err := ingressAPI.GetIngressSecret(clusterID, secretName, namespace)
	if err != nil {
		return err
	}

	d.Set("cluster_id", ingressSecretConfig.Cluster)
	d.Set("secret_name", ingressSecretConfig.Name)
	d.Set("cert_crn", ingressSecretConfig.CRN)
	d.Set("namespace", ingressSecretConfig.Namespace)
	instancecrn := strings.Split(ingressSecretConfig.CRN, ":certificate:")
	d.Set("cloud_cert_instance_id", fmt.Sprintf("%s::", instancecrn[0]))
	// d.Set("cluster_crn", ingressSecretConfig.ClusterCrn)
	d.Set("domain_name", ingressSecretConfig.Domain)
	d.Set("expires_on", ingressSecretConfig.ExpiresOn)
	d.Set("status", ingressSecretConfig.Status)
	d.Set("persistence", ingressSecretConfig.Persistence)
	// d.Set("issuer_name", ingressSecretConfig.IssuerName)
	d.SetId(fmt.Sprintf("%s/%s/%s", clusterID, secretName, ingressSecretConfig.Namespace))

	return nil
}

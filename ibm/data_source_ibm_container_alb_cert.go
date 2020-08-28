package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMContainerALBCert() *schema.Resource {
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
			},
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Secret name",
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
			},
			"cluster_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cluster CRN",
			},
			"cloud_cert_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cloud cert instance ID",
			},
		},
	}
}

func dataSourceIBMContainerALBCertRead(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	clusterID := d.Get("cluster_id").(string)
	secretName := d.Get("secret_name").(string)

	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return err
	}

	albAPI := albClient.Albs()
	albSecretConfig, err := albAPI.GetClusterALBCertBySecretName(clusterID, secretName, targetEnv)
	if err != nil {
		return err
	}

	d.Set("cluster_id", albSecretConfig.ClusterID)
	d.Set("secret_name", albSecretConfig.SecretName)
	d.Set("cert_crn", albSecretConfig.CertCrn)
	d.Set("cloud_cert_instance_id", albSecretConfig.CloudCertInstanceID)
	d.Set("cluster_crn", albSecretConfig.ClusterCrn)
	d.Set("domain_name", albSecretConfig.DomainName)
	d.Set("expires_on", albSecretConfig.ExpiresOn)
	d.Set("issuer_name", albSecretConfig.IssuerName)
	d.SetId(fmt.Sprintf("%s/%s", clusterID, secretName))

	return nil
}

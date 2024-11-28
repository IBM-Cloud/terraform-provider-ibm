package db2

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func DataSourceIBMDB2ConnectionInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMDB2ConnectionInfoRead,
		Schema: map[string]*schema.Schema{
			"deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the deployment to which IBMDB2 connection info will be used",
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_db2_connection_info",
					"deployment_id"),
			},

			"public": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Public connection info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname of the public connection info",
						},
						"databaseName": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Name of the public connection info",
						},
						"host_ros": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname of the public connection info",
						},
						"certificateBase64": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Base64 encoded public connection info",
						},
						"sslPort": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port for SSL connection info",
						},
						"ssl": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SSL connection info is enabled",
						},
						"databaseVersion": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version of the public connection info",
						},
					},
				},
			},

			"private": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Private connection info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname of the public connection info",
						},
						"databaseName": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the public connection info",
						},
						"host_ros": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname of the public connection info",
						},
						"certificateBase64": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Base64 encoded public connection info",
						},
						"sslPort": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port for SSL connection info",
						},
						"ssl": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SSL connection info is enabled",
						},
						"databaseVersion": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Version of the public connection info",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMDB2ConnectionInfoRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2SaasV1Client, err := meta.(conns.ClientSession).DB2SaasV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_connection_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getConnectionInfoOptions := &db2saasv1.GetDb2SaasConnectionInfoOptions{}

	getConnectionInfoOptions.SetDeploymentID(d.Get("deployment_id").(string))
	getConnectionInfoOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	result, response, err := db2SaasV1Client.GetDb2SaasConnectionInfoWithContext(context, getConnectionInfoOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDb2SaasConnectionInfoWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_db2_connection_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	public := []map[string]interface{}{}
	if result.Public != nil {
		modelMap, err := DataSourceIBMDB2ConnectionInfoPublicToMap(*result.Public)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_connection_info", "read")
			return tfErr.GetDiag()
		}
		public = append(public, modelMap)
	}

	if err := d.Set("public", public); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting public connection info: %s", err), "(Data) ibm_database_connection_info", "read")
		return tfErr.GetDiag()
	}

	private := []map[string]interface{}{}
	if result.Private != nil {
		modelMap, err := DataSourceIBMDB2ConnectionInfoPrivateToMap(*result.Private)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_connection_info", "read")
			return tfErr.GetDiag()
		}
		private = append(private, modelMap)
	}

	if err := d.Set("private", public); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting private connection info: %s", err), "(Data) ibm_database_connection_info", "read")
		return tfErr.GetDiag()
	}

	return nil
}

func DataSourceIBMDB2ConnectionInfoPublicToMap(model db2saasv1.SuccessConnectionInfoPublic) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	if model.Hostname != nil {
		modelMap["hostname"] = *model.Hostname
	}

	if model.DatabaseName != nil {
		modelMap["databaseName"] = *model.DatabaseName
	}

	if model.HostRos != nil {
		modelMap["host_ros"] = *model.HostRos
	}

	if model.CertificateBase64 != nil {
		modelMap["certificateBase64"] = *model.CertificateBase64
	}

	if model.SslPort != nil {
		modelMap["sslPort"] = *model.SslPort
	}

	if model.Ssl != nil {
		modelMap["ssl"] = *model.Ssl
	}

	if model.DatabaseVersion != nil {
		modelMap["databaseVersion"] = *model.DatabaseVersion
	}

	return modelMap, nil
}

func DataSourceIBMDB2ConnectionInfoPrivateToMap(model db2saasv1.SuccessConnectionInfoPrivate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	if model.Hostname != nil {
		modelMap["hostname"] = *model.Hostname
	}

	if model.DatabaseName != nil {
		modelMap["databaseName"] = *model.DatabaseName
	}

	if model.HostRos != nil {
		modelMap["host_ros"] = *model.HostRos
	}

	if model.CertificateBase64 != nil {
		modelMap["certificateBase64"] = *model.CertificateBase64
	}

	if model.SslPort != nil {
		modelMap["sslPort"] = *model.SslPort
	}

	if model.Ssl != nil {
		modelMap["ssl"] = *model.Ssl
	}

	if model.DatabaseVersion != nil {
		modelMap["databaseVersion"] = *model.DatabaseVersion
	}

	return modelMap, nil
}

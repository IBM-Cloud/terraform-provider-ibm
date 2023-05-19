// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes

import (
	"fmt"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMContainerIngressSecretOpaque() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerIngressSecretOpaqueCreate,
		Read:     resourceIBMContainerIngressSecretOpaqueRead,
		Update:   resourceIBMContainerIngressSecretOpaqueUpdate,
		Delete:   resourceIBMContainerIngressSecretOpaqueDelete,
		Exists:   resourceIBMContainerIngressSecretOpaqueExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cluster": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID or name",
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
				Required:    true,
				Description: "Secret namespace",
				ForceNew:    true,
			},
			"secret_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Opaque secret type",
			},
			"persistence": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Persistence of secret",
			},
			"user_managed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "If the secret was created by the user",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the secret",
			},
			"fields": {
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

func ResourceIBMContainerIngressSecretOpaqueValidator() *validate.ResourceValidator {
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

func resourceIBMContainerIngressSecretOpaqueCreate(d *schema.ResourceData, meta interface{}) error {
	ingressClient, err := meta.(conns.ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	cluster := d.Get("cluster").(string)
	secretName := d.Get("secret_name").(string)
	secretNamespace := d.Get("secret_namespace").(string)
	secretType := "opaque"

	params := v2.SecretCreateConfig{
		Cluster:   cluster,
		Name:      secretName,
		Namespace: secretNamespace,
		Type:      secretType,
	}

	if opaque_secret_fields, ok := d.GetOk("opaque_secret_fields"); ok {
		opaqueFieldsList := opaque_secret_fields.(*schema.Set).List()

		fieldsToAdd := []containerv2.FieldAdd{}
		for _, opaqueField := range opaqueFieldsList {
			var fieldToAdd containerv2.FieldAdd

			opaqueFieldMap := opaqueField.(map[string]interface{})

			if name, ok := opaqueFieldMap["name"]; ok {
				fieldToAdd.Name = name.(string)
			}

			if crn, ok := opaqueFieldMap["crn"]; ok {
				fieldToAdd.CRN = crn.(string)
			}

			if prefix, ok := opaqueFieldMap["prefix"]; ok {
				fieldToAdd.AppendPrefix = prefix.(bool)
			}

			fieldsToAdd = append(fieldsToAdd, fieldToAdd)
		}

		params.FieldsToAdd = fieldsToAdd
	}

	ingressAPI := ingressClient.Ingresses()
	response, err := ingressAPI.CreateIngressSecret(params)

	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", response.Cluster, response.Name, response.Namespace))

	return resourceIBMContainerIngressSecretOpaqueRead(d, meta)
}

func resourceIBMContainerIngressSecretOpaqueRead(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("secret_type", ingressSecretConfig.Type)
	d.Set("persistence", ingressSecretConfig.Persistence)
	d.Set("user_managed", ingressSecretConfig.UserManaged)
	d.Set("status", ingressSecretConfig.Status)

	if len(ingressSecretConfig.Fields) > 0 {
		d.Set("fields", flex.FlattenOpaqueSecret(ingressSecretConfig.Fields))
	}
	return nil
}

func resourceIBMContainerIngressSecretOpaqueDelete(d *schema.ResourceData, meta interface{}) error {
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

func resourceIBMContainerIngressSecretOpaqueUpdate(d *schema.ResourceData, meta interface{}) error {
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

	if d.HasChange("opaque_secret_fields") {
		oldList, newList := d.GetChange("opaque_secret_fields")

		if oldList == nil {
			oldList = new(schema.Set)
		}
		if newList == nil {
			newList = new(schema.Set)
		}
		os := oldList.(*schema.Set)
		ns := newList.(*schema.Set)

		remove := os.Difference(ns).List()
		add := ns.Difference(os).List()

		if len(remove) > 0 {
			var fieldsToRemove []containerv2.FieldRemove
			for _, removeField := range remove {
				removeFieldMap := removeField.(map[string]interface{})
				var fieldToRemove containerv2.FieldRemove

				if name, ok := removeFieldMap["name"]; ok {
					fieldToRemove.Name = name.(string)
				}

				fieldsToRemove = append(fieldsToRemove, fieldToRemove)
			}
			params.FieldsToRemove = fieldsToRemove
		}

		if len(add) > 0 {
			var fieldsToAdd []containerv2.FieldAdd
			for _, addField := range add {
				var fieldToAdd containerv2.FieldAdd

				addFieldMap := addField.(map[string]interface{})

				if name, ok := addFieldMap["name"]; ok {
					fieldToAdd.Name = name.(string)
				}

				if crn, ok := addFieldMap["crn"]; ok {
					fieldToAdd.CRN = crn.(string)
				}

				if prefix, ok := addFieldMap["prefix"]; ok {
					fieldToAdd.AppendPrefix = prefix.(bool)
				}

				fieldsToAdd = append(fieldsToAdd, fieldToAdd)
			}
			params.FieldsToAdd = fieldsToAdd
		}
		ingressAPI := ingressClient.Ingresses()

		_, err := ingressAPI.UpdateIngressSecret(params)
		if err != nil {
			return err
		}
	}

	return resourceIBMContainerIngressSecretOpaqueRead(d, meta)
}

func resourceIBMContainerIngressSecretOpaqueExists(d *schema.ResourceData, meta interface{}) (bool, error) {
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

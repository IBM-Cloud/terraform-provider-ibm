package ibm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
)

func resourceIBMCertificateManagerImport() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCertificateManagerImportCertificate,
		Read:     resourceIBMCertificateManagerGet,
		Update:   resourceIBMCertificateManagerUpdate,
		Importer: &schema.ResourceImporter{},
		Delete:   resourceIBMCertificateManagerDelete,
		Exists:   resourceIBMCertificateManagerExists,
		Schema: map[string]*schema.Schema{
			"certificate_manager_instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:     schema.TypeString,
							Required: true,
						},
						"priv_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"intermediate": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"begins_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"expires_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"imported": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"has_previous": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMCertificateManagerImportCertificate(d *schema.ResourceData, meta interface{}) error {

	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Get("certificate_manager_instance_id").(string)
	importData := models.Data{}
	name := d.Get("name").(string)
	description := d.Get("description").(string)

	if certificateimpdata, ok := d.GetOk("data"); ok && certificateimpdata != nil {
		datainfo := certificateimpdata.(map[string]interface{})
		if content, ok := datainfo["content"]; ok && content != nil {
			importData.Content = content.(string)
		}
		if privkey, ok := datainfo["priv_key"]; ok && privkey != nil {
			importData.Privatekey = privkey.(string)
		}
		if intermediate, ok := datainfo["intermediate"]; ok && intermediate != nil {
			importData.IntermediateCertificate = intermediate.(string)
		}
	}

	client := cmService.Certificate()
	payload := models.CertificateImportData{Name: name, Description: description, Data: importData}

	result, importCertError := client.ImportCertificate(instanceID, payload)
	if importCertError != nil {
		return importCertError
	}
	d.SetId(result.ID)
	return resourceIBMCertificateManagerUpdate(d, meta)
}
func resourceIBMCertificateManagerGet(d *schema.ResourceData, meta interface{}) error {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}
	certID := d.Id()
	certificatedata, err := cmService.Certificate().GetCertData(certID)

	cminstanceid := strings.Split(certID, ":certificate:")
	d.Set("certificate_manager_instance_id", cminstanceid[0]+"::")
	d.Set("name", certificatedata.Name)
	d.Set("description", certificatedata.Description)
	if certificatedata.Data != nil {
		data := map[string]interface{}{
			"content": certificatedata.Data.Content,
		}
		if certificatedata.Data.Privatekey != "" {
			data["priv_key"] = certificatedata.Data.Privatekey
		}
		if certificatedata.Data.IntermediateCertificate != "" {
			data["intermediate"] = certificatedata.Data.IntermediateCertificate
		}
		d.Set("data", data)
	}
	d.Set("begins_on", certificatedata.BeginsOn)
	d.Set("expires_on", certificatedata.ExpiresOn)
	d.Set("status", certificatedata.Status)
	d.Set("issuer", certificatedata.Issuer)
	d.Set("imported", certificatedata.Imported)
	d.Set("has_previous", certificatedata.HasPrevious)
	d.Set("key_algorithm", certificatedata.KeyAlgorithm)
	d.Set("algorithm", certificatedata.Algorithm)

	return nil
}

func resourceIBMCertificateManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}
	certID := d.Id()
	client := cmService.Certificate()
	if d.HasChange("name") || d.HasChange("description") {
		name := d.Get("name").(string)
		description := d.Get("description").(string)
		payload := models.CertificateMetadataUpdate{Name: name, Description: description}

		importCertError := client.UpdateCertificateMetaData(certID, payload)
		if importCertError != nil {
			return importCertError
		}
	}
	if d.HasChange("data") {
		importData := models.Data{}
		if certificateimpdata, ok := d.GetOk("data"); ok && certificateimpdata != nil {
			datainfo := certificateimpdata.(map[string]interface{})
			if content, ok := datainfo["content"]; ok && content != nil {
				importData.Content = content.(string)
			}
			if privkey, ok := datainfo["priv_key"]; ok && privkey != nil {
				importData.Privatekey = privkey.(string)
			}
			if intermediate, ok := datainfo["intermediate"]; ok && intermediate != nil {
				importData.IntermediateCertificate = intermediate.(string)
			}
		}
		payload := models.CertificateReimportData{Content: importData.Content, Privatekey: importData.Privatekey, IntermediateCertificate: importData.IntermediateCertificate}
		_, reImportCertError := client.ReimportCertificate(certID, payload)
		if reImportCertError != nil {
			return reImportCertError
		}
	}
	return resourceIBMCertificateManagerGet(d, meta)
}
func resourceIBMCertificateManagerDelete(d *schema.ResourceData, meta interface{}) error {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return err
	}
	certID := d.Id()
	err = cmService.Certificate().DeleteCertificate(certID)
	if err != nil {
		return fmt.Errorf("Error deleting Certificate: %s", err)
	}
	d.SetId("")

	return nil
}

func resourceIBMCertificateManagerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cmService, err := meta.(ClientSession).CertificateManagerAPI()
	if err != nil {
		return false, err
	}
	client := cmService.Certificate()
	certID := d.Id()

	_, err = client.GetCertData(certID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return true, nil
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"

	kp "github.com/IBM/keyprotect-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMKmsKMIPClientCertificate() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMKmsKMIPClientCertCreate,
		Read:     resourceIBMKmsKMIPClientCertRead,
		Delete:   resourceIBMKmsKMIPClientCertDelete,
		Exists:   resourceIBMKmsKMIPClientCertExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Key protect Instance GUID",
				ForceNew:         true,
				DiffSuppressFunc: suppressKMSInstanceIDDiff,
			},
			"adapter_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name or UUID of the KMIP adapter that contains the cert",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The name of the KMIP client certificate",
			},
			"certificate": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The PEM-encoded contents of the certificate",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier that is associated with the entity that created the adapter.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date when a resource was created. The date format follows RFC 3339.",
			},
		},
	}
}

func resourceIBMKmsKMIPClientCertCreate(d *schema.ResourceData, meta interface{}) error {
	certToCreate, adapterID, instanceID, err := ExtractAndValidateKMIPClientCertDataFromSchema(d)
	if err != nil {
		return err
	}
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	ctx := context.Background()
	cert, err := kpAPI.CreateKMIPClientCertificate(ctx,
		adapterID,
		certToCreate.Certificate,
		kp.WithKMIPClientCertName(certToCreate.Name),
	)
	if err != nil {
		return fmt.Errorf("[ERROR] Error while creating KMIP client certificate: %s", err)
	}
	return populateKMIPClientCertSchemaDataFromStruct(d, *cert)
}

func resourceIBMKmsKMIPClientCertRead(d *schema.ResourceData, meta interface{}) error {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Get("adapter_id").(string)
	certID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	ctx := context.Background()
	cert, err := kpAPI.GetKMIPClientCertificate(ctx, adapterID, certID)
	if err != nil {
		return err
	}
	err = populateKMIPClientCertSchemaDataFromStruct(d, *cert)
	if err != nil {
		return err
	}
	return nil
}

func resourceIBMKmsKMIPClientCertUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMKmsKMIPClientCertRead(d, meta)
}

func resourceIBMKmsKMIPClientCertDelete(d *schema.ResourceData, meta interface{}) error {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Get("adapter_id").(string)
	certID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return err
	}
	err = kpAPI.DeleteKMIPClientCertificate(context.Background(), adapterID, certID)
	return err
}

func resourceIBMKmsKMIPClientCertExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	instanceID := d.Get("instance_id").(string)
	adapterID := d.Get("adapter_id").(string)
	certID := d.Id()
	kpAPI, _, err := populateKPClient(d, meta, instanceID)
	if err != nil {
		return false, err
	}
	ctx := context.Background()
	_, err = kpAPI.GetKMIPClientCertificate(ctx, adapterID, certID)
	if err != nil {
		kpError := err.(*kp.Error)
		if kpError.StatusCode == 404 {
			return false, nil
		}
		return false, wrapError(err, "Error checking KMIP Client Certificate existence")
	}
	return true, nil
}

func ExtractAndValidateKMIPClientCertDataFromSchema(d *schema.ResourceData) (cert kp.KMIPClientCertificate, adapterIDStr string, instanceID string, err error) {
	err = nil
	instanceID = getInstanceIDFromResourceData(d, "instance_id")

	cert = kp.KMIPClientCertificate{}
	if name, ok := d.GetOk("name"); ok {
		nameStr, ok2 := name.(string)
		if !ok2 {
			err = fmt.Errorf("[ERROR] Error converting name to string")
			return
		}
		cert.Name = nameStr
	}
	if certPayload, ok := d.GetOk("certificate"); ok {
		certStr, ok2 := certPayload.(string)
		if !ok2 {
			err = fmt.Errorf("[ERROR] Error converting certificate to string")
			return
		}
		cert.Certificate = certStr
	}

	adapterID := d.Get("adapter_id")
	adapterIDStr = adapterID.(string)
	return
}

func populateKMIPClientCertSchemaDataFromStruct(d *schema.ResourceData, cert kp.KMIPClientCertificate) (err error) {
	d.SetId(cert.ID)
	if err = d.Set("name", cert.Name); err != nil {
		return fmt.Errorf("[ERROR] Error setting name: %s", err)
	}
	if err = d.Set("certificate", cert.Certificate); err != nil {
		return fmt.Errorf("[ERROR] Error setting certificate: %s", err)
	}
	if cert.CreatedAt != nil {
		if err = d.Set("created_at", cert.CreatedAt.String()); err != nil {
			return fmt.Errorf("[ERROR] Error setting created_at: %s", err)
		}
		if err = d.Set("created_by", cert.CreatedBy); err != nil {
			return fmt.Errorf("[ERROR] Error setting created_by: %s", err)
		}
	}
	return nil
}

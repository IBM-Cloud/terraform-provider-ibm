// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager

import (
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-go-sdk/secretsmanagerv2"
)

func DataSourceIbmSmPrivateCertificateConfigurationActionSignCsr() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrRead,

		Schema: map[string]*schema.Schema{
			"action_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of configuration action.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the configuration.",
			},
			"common_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Common Name (AKA CN) represents the server name that is protected by the SSL certificate.",
			},
			"alt_names": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ip_sans": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IP Subject Alternative Names to define for the CA certificate, in a comma-delimited list.",
			},
			"uri_sans": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URI Subject Alternative Names to define for the CA certificate, in a comma-delimited list.",
			},
			"other_sans": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The custom Object Identifier (OID) or UTF8-string Subject Alternative Names to define for the CA certificate.The alternative names must match the values that are specified in the `allowed_other_sans` field in the associated certificate template. The format is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ttl": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The time-to-live (TTL) to assign to a private certificate.The value can be supplied as a string representation of a duration in hours, for example '12h'. The value can't exceed the `max_ttl` that is defined in the associated certificate template.",
			},
			"format": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The format of the returned data.",
			},
			"max_path_length": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum path length to encode in the generated certificate. `-1` means no limit.If the signing certificate has a maximum path length set, the path length is set to one less than that of the signing certificate. A limit of `0` means a literal path length of zero.",
			},
			"exclude_cn_from_sans": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Controls whether the common name is excluded from Subject Alternative Names (SANs).If the common name set to `true`, it is not included in DNS or Email SANs if they apply. This field can be useful if the common name is a human-readable identifier, instead of a hostname or an email address.",
			},
			"permitted_dns_domains": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The allowed DNS domains or subdomains for the certificates that are to be signed and issued by this CA certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_csr_values": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines whether to use values from a certificate signing request (CSR) to complete a `private_cert_configuration_action_sign_csr` action. If it is set to `true`, then:1) Subject information, including names and alternate names, are preserved from the CSR rather than by using the values that are provided in the other parameters to this operation.2) Any key usage, for example, non-repudiation, that are requested in the CSR are added to the basic set of key usages used for CA certificates that are signed by the intermediate authority.3) Extensions that are requested in the CSR are copied into the issued private certificate.",
			},
			"ou": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Organizational Unit (OU) values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"organization": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Organization (O) values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"country": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Country (C) values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"locality": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Locality (L) values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"province": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Province (ST) values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"street_address": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The street address values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"postal_code": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The postal code values to define in the subject field of the resulting certificate.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"serial_number": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The serial number to assign to the generated certificate. To assign a random serial number, you can omit this field.",
			},
			"csr": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The certificate signing request.",
			},
			"data": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The data that is associated with the root certificate authority.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The PEM-encoded contents of your certificate.",
						},
						"issuing_ca": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Sensitive:   true,
							Description: "The PEM-encoded certificate of the certificate authority that signed and issued this certificate.",
						},
						"ca_chain": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Sensitive:   true,
							Description: "The chain of certificate authorities that are associated with the certificate.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"expiration": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The certificate expiration time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerClient, err := meta.(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	secretsManagerClient = getClientWithInstanceEndpoint(secretsManagerClient, d)

	createConfigurationActionOptions := &secretsmanagerv2.CreateConfigurationActionOptions{}

	createConfigurationActionOptions.SetName(d.Get("name").(string))
	createConfigurationActionOptions.SetXSmAcceptConfigurationType("private_cert_configuration_root_ca")
	configurationActionPrototypeModel, err := ibmSmPrivateCertificateConfigurationActionSignCsrToConfigurationActionPrototype(d)
	if err != nil {
		return diag.FromErr(err)
	}
	createConfigurationActionOptions.SetConfigActionPrototype(configurationActionPrototypeModel)

	privateCertificateConfigurationActionSignCSRIntf, response, err := secretsManagerClient.CreateConfigurationActionWithContext(context, createConfigurationActionOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigurationActionWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateConfigurationActionWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrID(d))

	privateCertificateConfigurationActionSignCSR := privateCertificateConfigurationActionSignCSRIntf.(*secretsmanagerv2.PrivateCertificateConfigurationActionSignCSR)
	if err = d.Set("common_name", privateCertificateConfigurationActionSignCSR.CommonName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting common_name: %s", err))
	}
	if privateCertificateConfigurationActionSignCSR.AltNames != nil {
		if err = d.Set("alt_names", privateCertificateConfigurationActionSignCSR.AltNames); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting alt_names: %s", err))
		}
	}
	if err = d.Set("ip_sans", privateCertificateConfigurationActionSignCSR.IpSans); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting ip_sans: %s", err))
	}
	if err = d.Set("uri_sans", privateCertificateConfigurationActionSignCSR.UriSans); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting uri_sans: %s", err))
	}
	if privateCertificateConfigurationActionSignCSR.OtherSans != nil {
		if err = d.Set("other_sans", privateCertificateConfigurationActionSignCSR.OtherSans); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting other_sans: %s", err))
		}
	}
	if err = d.Set("ttl", privateCertificateConfigurationActionSignCSR.TTL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting ttl: %s", err))
	}
	if err = d.Set("format", privateCertificateConfigurationActionSignCSR.Format); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting format: %s", err))
	}
	if err = d.Set("max_path_length", flex.IntValue(privateCertificateConfigurationActionSignCSR.MaxPathLength)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting max_path_length: %s", err))
	}
	if err = d.Set("exclude_cn_from_sans", privateCertificateConfigurationActionSignCSR.ExcludeCnFromSans); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting exclude_cn_from_sans: %s", err))
	}
	if privateCertificateConfigurationActionSignCSR.PermittedDnsDomains != nil {
		if err = d.Set("permitted_dns_domains", privateCertificateConfigurationActionSignCSR.PermittedDnsDomains); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting permitted_dns_domains: %s", err))
		}
	}
	if err = d.Set("use_csr_values", privateCertificateConfigurationActionSignCSR.UseCsrValues); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting use_csr_values: %s", err))
	}
	if privateCertificateConfigurationActionSignCSR.Ou != nil {
		if err = d.Set("ou", privateCertificateConfigurationActionSignCSR.Ou); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting ou: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.Organization != nil {
		if err = d.Set("organization", privateCertificateConfigurationActionSignCSR.Organization); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting organization: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.Country != nil {
		if err = d.Set("country", privateCertificateConfigurationActionSignCSR.Country); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting country: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.Locality != nil {
		if err = d.Set("locality", privateCertificateConfigurationActionSignCSR.Locality); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting locality: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.Province != nil {
		if err = d.Set("province", privateCertificateConfigurationActionSignCSR.Province); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting province: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.StreetAddress != nil {
		if err = d.Set("street_address", privateCertificateConfigurationActionSignCSR.Ou); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting street_address: %s", err))
		}
	}
	if privateCertificateConfigurationActionSignCSR.PostalCode != nil {
		if err = d.Set("postal_code", privateCertificateConfigurationActionSignCSR.PostalCode); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting postal_code: %s", err))
		}
	}
	if err = d.Set("serial_number", privateCertificateConfigurationActionSignCSR.SerialNumber); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting serial_number: %s", err))
	}
	if err = d.Set("action_type", privateCertificateConfigurationActionSignCSR.ActionType); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action_type: %s", err))
	}
	if err = d.Set("csr", privateCertificateConfigurationActionSignCSR.Csr); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting csr: %s", err))
	}
	data := []map[string]interface{}{}
	if privateCertificateConfigurationActionSignCSR.Data != nil {
		modelMap, err := dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrPrivateCertificateConfigurationCACertificateToMap(privateCertificateConfigurationActionSignCSR.Data)
		if err != nil {
			return diag.FromErr(err)
		}
		data = append(data, modelMap)
	}
	if err = d.Set("data", data); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting data %s", err))
	}

	return nil
}

func ibmSmPrivateCertificateConfigurationActionSignCsrToConfigurationActionPrototype(d *schema.ResourceData) (*secretsmanagerv2.PrivateCertificateConfigurationActionSignCSRPrototype, error) {
	model := &secretsmanagerv2.PrivateCertificateConfigurationActionSignCSRPrototype{
		ActionType: core.StringPtr("private_cert_configuration_action_sign_csr"),
	}
	if _, ok := d.GetOk("common_name"); ok {
		model.CommonName = core.StringPtr(d.Get("common_name").(string))
	}
	if _, ok := d.GetOk("alt_names"); ok {
		altNames := d.Get("alt_names").([]interface{})
		altNamesParsed := make([]string, len(altNames))
		for i, v := range altNames {
			altNamesParsed[i] = fmt.Sprint(v)
		}
		model.AltNames = altNamesParsed
	}
	if _, ok := d.GetOk("ip_sans"); ok {
		model.IpSans = core.StringPtr(d.Get("ip_sans").(string))
	}
	if _, ok := d.GetOk("uri_sans"); ok {
		model.UriSans = core.StringPtr(d.Get("uri_sans").(string))
	}
	if _, ok := d.GetOk("other_sans"); ok {
		otherSans := d.Get("other_sans").([]interface{})
		otherSansParsed := make([]string, len(otherSans))
		for i, v := range otherSans {
			otherSansParsed[i] = fmt.Sprint(v)
		}
		model.OtherSans = otherSansParsed
	}
	if _, ok := d.GetOk("ttl"); ok {
		model.TTL = core.StringPtr(d.Get("ttl").(string))
	}
	if _, ok := d.GetOk("format"); ok {
		model.Format = core.StringPtr(d.Get("format").(string))
	}
	if _, ok := d.GetOk("max_path_length"); ok {
		model.MaxPathLength = core.Int64Ptr(d.Get("max_path_length").(int64))
	}
	if _, ok := d.GetOk("exclude_cn_from_sans"); ok {
		model.ExcludeCnFromSans = core.BoolPtr(d.Get("exclude_cn_from_sans").(bool))
	}
	if _, ok := d.GetOk("permitted_dns_domains"); ok {
		permittedDnsDomains := d.Get("permitted_dns_domains").([]interface{})
		permittedDnsDomainsParsed := make([]string, len(permittedDnsDomains))
		for i, v := range permittedDnsDomains {
			permittedDnsDomainsParsed[i] = fmt.Sprint(v)
		}
		model.PermittedDnsDomains = permittedDnsDomainsParsed
	}
	if _, ok := d.GetOk("use_csr_values"); ok {
		model.UseCsrValues = core.BoolPtr(d.Get("use_csr_values").(bool))
	}
	if _, ok := d.GetOk("ou"); ok {
		ou := d.Get("ou").([]interface{})
		ouParsed := make([]string, len(ou))
		for i, v := range ou {
			ouParsed[i] = fmt.Sprint(v)
		}
		model.Ou = ouParsed
	}
	if _, ok := d.GetOk("organization"); ok {
		organization := d.Get("organization").([]interface{})
		organizationParsed := make([]string, len(organization))
		for i, v := range organization {
			organizationParsed[i] = fmt.Sprint(v)
		}
		model.Organization = organizationParsed
	}
	if _, ok := d.GetOk("country"); ok {
		country := d.Get("country").([]interface{})
		countryParsed := make([]string, len(country))
		for i, v := range country {
			countryParsed[i] = fmt.Sprint(v)
		}
		model.Country = countryParsed
	}
	if _, ok := d.GetOk("locality"); ok {
		locality := d.Get("locality").([]interface{})
		localityParsed := make([]string, len(locality))
		for i, v := range locality {
			localityParsed[i] = fmt.Sprint(v)
		}
		model.Locality = localityParsed
	}
	if _, ok := d.GetOk("province"); ok {
		province := d.Get("province").([]interface{})
		provinceParsed := make([]string, len(province))
		for i, v := range province {
			provinceParsed[i] = fmt.Sprint(v)
		}
		model.Province = provinceParsed
	}
	if _, ok := d.GetOk("street_address"); ok {
		streetAddress := d.Get("street_address").([]interface{})
		streetAddressParsed := make([]string, len(streetAddress))
		for i, v := range streetAddress {
			streetAddressParsed[i] = fmt.Sprint(v)
		}
		model.StreetAddress = streetAddressParsed
	}
	if _, ok := d.GetOk("postal_code"); ok {
		postalCode := d.Get("postal_code").([]interface{})
		postalCodeParsed := make([]string, len(postalCode))
		for i, v := range postalCode {
			postalCodeParsed[i] = fmt.Sprint(v)
		}
		model.PostalCode = postalCodeParsed
	}
	if _, ok := d.GetOk("serial_number"); ok {
		model.SerialNumber = core.StringPtr(d.Get("serial_number").(string))
	}
	if _, ok := d.GetOk("csr"); ok {
		model.Csr = core.StringPtr(d.Get("csr").(string))
	}
	return model, nil
}

// dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrID returns a reasonable ID for the list.
func dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmSmPrivateCertificateConfigurationActionSignCsrPrivateCertificateConfigurationCACertificateToMap(model *secretsmanagerv2.PrivateCertificateConfigurationCACertificate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Certificate != nil {
		modelMap["certificate"] = *model.Certificate
	}
	if model.IssuingCa != nil {
		modelMap["issuing_ca"] = *model.IssuingCa
	}
	if model.CaChain != nil {
		modelMap["ca_chain"] = model.CaChain
	}
	if model.Expiration != nil {
		modelMap["expiration"] = *model.Expiration
	}
	return modelMap, nil
}

// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceTemplate_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "availability_policy_host_failure"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_QoS(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	qosMode := "weighted"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDQoSConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, qosMode),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "volume_bandwidth_qos_mode"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplateDatasourceCluster(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	// templateName := fmt.Sprintf("testtemplate%d", randInt)
	templateName := "eu-de-test-cluster-it"
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDatasourceClusterConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data_name", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "boot_volume_attachment.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "cluster_network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "cluster_network_attachments.0.cluster_network_interface.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "enable_secure_boot"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "id"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "image"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "keys.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "profile"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "resource_group"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "vpc"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data_name", "zone"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_dataconcom(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	ccmode := "sgx"
	esb := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDconcomConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, ccmode, esb),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "confidential_compute_mode"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "enable_secure_boot"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_dataVni(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDVniConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.allow_ip_spoofing"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.auto_delete"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.enable_infrastructure_nat"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.ips.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.security_groups.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_attachment.0.virtual_network_interface.0.subnet.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_data_catalog(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateDCatalogConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "catalog_offering.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "catalog_offering.0.plan_crn"),
					resource.TestCheckNoResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "image"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_ReservedIp_Basic(t *testing.T) {
	randInt := acctest.RandIntRange(600, 700)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDQ+WiiUR1Jg3oGSmB/2//GJ3XnotriBiGN6t3iwGces6sUsvRkza1t0Mf05DKZxC/zp0WvDTvbit2gTkF9sD37OZSn5aCJk1F5URk/JNPmz25ZogkICFL4OUfhrE3mnyKio6Bk1JIEIypR5PRtGxY9vFDUfruADDLfRi+dGwHF6U9RpvrDRo3FNtI8T0GwvWwFE7bg63vLz65CjYY5XqH9z/YWz/asH6BKumkwiphLGhuGn03+DV6DkIZqr3Oh13UDjMnTdgv1y/Kou5UM3CK1dVsmLRXPEf2KUWUq1EwRfrJXkPOrBwn8to+Yydo57FgrRM9Qw8uzvKmnVxfKW6iG3oSGA0L6ROuCq1lq0MD8ySLd56+d1ftSDaUq+0/Yt9vK3olzVP0/iZobD7chbGqTLMCzL4/CaIUR/UmX08EA0Oh0DdyAdj3UUNETAj3W8gBrV6xLR7fZAJ8roX2BKb4K8Ed3YqzgiY0zgjqvpBYl9xZl0jgVX0qMFaEa6+CeGI8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateRipDataConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "primary_network_interface.0.primary_ip.0.reserved_ip"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_AllowedUse_Basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	volAttachName := fmt.Sprintf("tf-testvolattach%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	userTag := "tag-0"
	bandwidth := int64(2000)
	apiVersion := "2025-07-02"
	bareMetalServer := "true"
	instanceval := "true"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateAllowedUseConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag, apiVersion, bareMetalServer, instanceval, bandwidth),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_template.instance_template_data", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.api_version"),
					resource.TestCheckResourceAttr("data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.bare_metal_server", bareMetalServer),
					resource.TestCheckResourceAttr("data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.instance", instanceval),
					resource.TestCheckResourceAttr("data.ibm_is_instance_template.instance_template_data", "boot_volume_attachment.0.allowed_use.0.api_version", apiVersion),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceTemplateDQoSConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, qosMode string) string {
	return testAccCheckIBMISInstanceTemplateConfig_QoSMode(vpcName, subnetName, sshKeyName, publicKey, templateName, qosMode) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}

func testAccCheckIBMISInstanceTemplateDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}
func testAccCheckIBMISInstanceTemplateDatasourceClusterConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data_name" {
			name = "eu-de-test-cluster-it"
		}
		data "ibm_is_instance_template" "instance_template_data_identifier" {
			identifier = "02c7-4a2d29da-429a-4355-9354-31af7c2e6627"
		}
	`)
}
func testAccCheckIBMISInstanceTemplateDconcomConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, ccmode string, esb bool) string {
	return testAccCheckIBMISInstanceTemplateConComConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, ccmode, esb) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}
func testAccCheckIBMISInstanceTemplateDVniConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return testAccCheckIBMISInstanceTemplateVniConfig(vpcName, subnetName, sshKeyName, publicKey, templateName) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}
func testAccCheckIBMISInstanceTemplateDCatalogConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return testAccCheckIBMISInstanceTemplateCatalogConfig(vpcName, subnetName, sshKeyName, publicKey, templateName) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}
func testAccCheckIBMISInstanceTemplateRipDataConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return testAccCheckIBMISInstanceTemplateRipConfig(vpcName, subnetName, sshKeyName, publicKey, templateName) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}

func testAccCheckIBMISInstanceTemplateAllowedUseConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag, apiVersion, bareMetalServer, instanceval string, bandwidth int64) string {
	return testAccCheckIBMISInstanceTemplateWithBoot_AllowedUse(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag, apiVersion, bareMetalServer, instanceval, bandwidth) + fmt.Sprintf(`
		data "ibm_is_instance_template" "instance_template_data" {
			name = ibm_is_instance_template.instancetemplate1.name
		}
	`)
}

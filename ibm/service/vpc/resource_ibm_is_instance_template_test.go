// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceTemplate_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_concom(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_concom1(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	esb := true
	ccmode := "sgx"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateConComConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, ccmode, esb),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "enable_secure_boot"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "confidential_compute_mode"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "confidential_compute_mode", ccmode),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "enable_secure_boot", fmt.Sprintf("%t", esb)),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_vni(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateVniConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_attachment"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_attachment.0.virtual_network_interface.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_vniPSFM(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	protocolStateFilteringMode := "auto"

	pNacName := fmt.Sprintf("tf-testvpc-pnac%d", randInt)
	sNacName := fmt.Sprintf("tf-testvpc-snac%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateVniPSFMConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, pNacName, protocolStateFilteringMode, sNacName, protocolStateFilteringMode),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_attachment.0.virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_attachment.0.virtual_network_interface.0.protocol_state_filtering_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "network_attachments.0.virtual_network_interface.0.protocol_state_filtering_mode"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_catalog_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateCatalogConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "catalog_offering.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "catalog_offering.0.plan_crn"),
					resource.TestCheckNoResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "image"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_Reservation(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateReservationConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "reservation_affinity.0.policy", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "reservation_affinity.0.pool"),
					resource.TestCheckNoResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "image"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_Reserved_IP_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateRipConfig(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "primary_network_interface.0.primary_ip.0.reserved_ip"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_metadata_service(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceMetadataServiceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, true, "https", 10),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_instance_template.instancetemplate1", "metadata_service.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_is_instance_template.instancetemplate1", "metadata_service.0.protocol", "https"),
					resource.TestCheckResourceAttr("ibm_is_instance_template.instancetemplate1", "metadata_service.0.response_hop_limit", "10"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceTemplate_withAvailabilityPolicy(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateConfigAvailablePolicyHostFailure_Default(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttr("ibm_is_instance_template.instancetemplate1", "availability_policy_host_failure", "stop"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceTemplateConfigAvailablePolicyHostFailure_Updated(vpcName, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttr("ibm_is_instance_template.instancetemplate1", "availability_policy_host_failure", "stop"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_WithVolumeAttachment(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("tf-testvpc%d", randInt)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	volAttachName := fmt.Sprintf("tf-testvolattach%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateWithVolume(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_WithVolumeAttachmentUserTag(t *testing.T) {
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
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateWithVolumeUserTag(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "boot_volume.0.tags.0", userTag),
				),
			},
		},
	})
}

func TestAccIBMISInstanceTemplate_WithBootBandwidth(t *testing.T) {
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
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateWithBootBandwidth(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag, bandwidth),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "name", templateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance_template.instancetemplate1", "profile"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "boot_volume.0.tags.0", userTag),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "boot_volume.0.size", "250"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.instancetemplate1", "boot_volume.0.bandwidth", fmt.Sprintf("%d", bandwidth)),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceTemplateDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_template" {
			continue
		}

		getInstanceTemplateOptions := vpcv1.GetInstanceTemplateOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetInstanceTemplate(&getInstanceTemplateOptions)

		if err == nil {
			return fmt.Errorf("instance template still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}
func testAccCheckIBMISInstanceTemplateConComConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, ccmode string, esb bool) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "%s"
	   confidential_compute_mode = "%s"
	   enable_secure_boot = %t
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, acc.InstanceProfileName, ccmode, esb)

}
func testAccCheckIBMISInstanceTemplateVniConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_attachment {
			name = "vni-2-test"
			virtual_network_interface {
				primary_ip {
					auto_delete 	= true
				}
		 		subnet = ibm_is_subnet.subnet2.id
			}
	   }
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}

func testAccCheckIBMISInstanceTemplateVniPSFMConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, pNacName, ppsfm, sNacName, spsfm string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_attachment {
			name = "%s"
			virtual_network_interface {
				primary_ip {
					auto_delete 	= true
				}
		 		subnet = ibm_is_subnet.subnet2.id
				 protocol_state_filtering_mode = "%s"
			}
	   }
	   network_attachments {
		name = "%s"
		virtual_network_interface { 
			primary_ip {
				auto_delete 	= true
			}
			subnet = ibm_is_subnet.subnet2.id
			protocol_state_filtering_mode = "%s"
        }
	}
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, pNacName, ppsfm, sNacName, spsfm)

}

func testAccCheckIBMISInstanceTemplateCatalogConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
		catalog_managed = true
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   catalog_offering {
		version_crn = data.ibm_is_images.is_images.images.0.catalog_offering.0.version.0.crn
		plan_crn = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
	   }
	   profile = "bx2-2x8"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}

func testAccCheckIBMISInstanceTemplateReservationConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
		catalog_managed = true
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   profile = "bx2-2x8"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   reservation_affinity {
			policy = "manual"
			pool {
				id = "0735-b4a78f50-33bd-44f9-a3ff-4c33f444459d"
			}
		}
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}
func testAccCheckIBMISInstanceTemplateRipConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
		name            = "%s"
		vpc             = ibm_is_vpc.vpc2.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_subnet_reserved_ip" "testacc_rip" {
		subnet = ibm_is_subnet.subnet2.id
		name = "test-instance-template-rip"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	data "ibm_is_images" "is_images" {
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
		 primary_ip {
			reserved_ip = ibm_is_subnet_reserved_ip.testacc_rip.reserved_ip
			}
	   }
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "%s"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, acc.ISZoneName, acc.ISCIDR, sshKeyName, publicKey, templateName, acc.ISZoneName)

}

func testAccCheckIBMISInstanceMetadataServiceTemplateConfig(vpcName, subnetName, sshKeyName, publicKey, templateName string, metadataService bool, protocol string, hop_limit int) string {
	return fmt.Sprintf(`
	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	   metadata_service {
		enabled = %t
		protocol = "%s"
		response_hop_limit = %d
	  }
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, acc.IsImage, metadataService, protocol, hop_limit)

}

func testAccCheckIBMISInstanceTemplateWithVolume(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   volume_attachments {
        delete_volume_on_instance_delete = true
        name                             = "%s"
			volume_prototype {
				iops = 6000
				profile = "custom"
				capacity = 100
			}   
    	}
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, acc.IsImage, volAttachName)

}

func testAccCheckIBMISInstanceTemplateWithVolumeUserTag(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   boot_volume{
		tags = ["%s"]
	   }
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, acc.IsImage, userTag)

}

func testAccCheckIBMISInstanceTemplateConfigAvailablePolicyHostFailure_Default(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}
	data "ibm_is_images" "is_images" {
	}
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}
func testAccCheckIBMISInstanceTemplateConfigAvailablePolicyHostFailure_Updated(vpcName, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}
	data "ibm_is_images" "is_images" {
	}
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = data.ibm_is_images.is_images.images.0.id
	   profile = "bx2-8x32"
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   availability_policy_host_failure = "stop"
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName)

}

func TestAccIBMISInstanceTemplate_clusternetworkbasic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	clustersubnetname := fmt.Sprintf("tf-clustersubnet-%d", acctest.RandIntRange(10, 100))
	clustersubnetreservedipname := fmt.Sprintf("tf-clustersubnet-reservedip-%d", acctest.RandIntRange(10, 100))
	clusterinterfacename := fmt.Sprintf("tf-clusterinterface-%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	subnetName := fmt.Sprintf("tf-testsubnet%d", randInt)
	templateName := fmt.Sprintf("tf-testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceTemplateClusterNetworkConfig(vpcname, clustersubnetname, clustersubnetreservedipname, clusterinterfacename, subnetName, sshKeyName, publicKey, templateName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name", clustersubnetname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "name", clustersubnetreservedipname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name", clusterinterfacename),
					resource.TestCheckResourceAttrSet("ibm_is_subnet.is_subnet", "id"),
					resource.TestCheckResourceAttr("ibm_is_subnet.is_subnet", "name", subnetName),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key.is_sshkey", "id"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key.is_sshkey", "name", sshKeyName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_template.is_instance_template", "name", templateName),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "profile"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "image"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "keys.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "vpc"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "boot_volume.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance_template.is_instance_template", "cluster_network_attachments.#"),
					resource.TestCheckResourceAttr("ibm_is_instance_template.is_instance_template", "cluster_network_attachments.#", "8"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceTemplateClusterNetworkConfig(vpcname, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename, subnetName, sshKeyName, publicKey, templateName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "is_vpc" {
  			name = "%s"
		}
		resource "ibm_is_cluster_network" "is_cluster_network_instance" {
			profile = "%s"
			vpc {
				id = ibm_is_vpc.is_vpc.id
			}
			zone  = "%s"
		}
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			total_ipv4_address_count = 64
		}
		resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
			cluster_network_id 			= ibm_is_cluster_network.is_cluster_network_instance.id
			cluster_network_subnet_id 	= ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			address 					= "${replace(ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.ipv4_cidr_block, "0/26", "11")}"
  			name						= "%s"
		}
		resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			primary_ip {
				id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
			}
			subnet {
				id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			}
		}
	
		resource "ibm_is_subnet" "is_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.is_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 64
		}
		
		resource "ibm_is_ssh_key" "is_sshkey" {
			name       = "%s"
			public_key = "%s"
		}

		resource "ibm_is_instance_template" "is_instance_template" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			
			primary_network_interface {
				subnet = ibm_is_subnet.is_subnet.id
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}
			cluster_network_attachments {
				cluster_network_interface{
					id = ibm_is_cluster_network_interface.is_cluster_network_interface_instance.cluster_network_interface_id
				}
			}			
			vpc       = ibm_is_vpc.is_vpc.id
			zone      = "%s"
			keys      = [ibm_is_ssh_key.is_sshkey.id]
		}
			
		
		`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename, subnetName, acc.ISZoneName, sshKeyName, publicKey, templateName, acc.IsImage, acc.ISInstanceGPUProfileName, acc.ISZoneName)

}

func testAccCheckIBMISInstanceTemplateWithBootBandwidth(vpcName, subnetName, sshKeyName, publicKey, templateName, volAttachName, userTag string, bandwidth int64) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            			= "%s"
	  vpc             			= ibm_is_vpc.vpc2.id
	  zone            			= "%s"
	  total_ipv4_address_count 	= 256
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}

	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
	   image   = "%s"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	   boot_volume{
		tags 			= ["%s"]
		bandwidth 		= %d
		profile 		= "sdp"
		size			= 250
	   }
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "%s"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	
	`, vpcName, subnetName, acc.ISZoneName, sshKeyName, publicKey, templateName, acc.IsImage, userTag, bandwidth, acc.ISZoneName)

}

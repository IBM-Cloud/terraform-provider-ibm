// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package functions_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
)

func TestAccCFFunctionPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionPackageCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageNameUpdate(updatedName, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageWithAnnotations(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageWithAnnotationsUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageWithParameters(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageWithParametersUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},
			{
				Config: testAccCheckCFFunctionPackageUpdatePublish(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "true"),
				),
			},
		},
	})
}

func TestAccIAMFunctionPackage_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_package_updated_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionPackageCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageNameUpdate(updatedName, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageWithAnnotations(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageWithAnnotationsUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageWithParameters(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageWithParametersUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
				),
			},
			{
				Config: testAccCheckIAMFunctionPackageUpdatePublish(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "true"),
				),
			},
		},
	})
}

func TestAccCFFunctionPackage_Bind_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_package_updated_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionPackageBindCreate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "bind_package_name", bindName),
				),
			},
			{
				Config: testAccCheckCFFunctionPackageNameBindUpdate(updatedName, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageBindWithAnnotations(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageBindWithAnnotationsUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageBindWithParameters(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckCFFunctionPackageBindWithParametersUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},
			{
				Config: testAccCheckCFFunctionPackageBindUpdatePublish(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "true"),
				),
			},
		},
	})
}

func TestAccIAMFunctionPackage_Bind_Basic(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	updatedName := fmt.Sprintf("terraform_updated_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	bindName := "/whisk.system/alarms"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionPackageBindCreate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "bind_package_name", bindName),
				),
			},
			{
				Config: testAccCheckIAMFunctionPackageNameBindUpdate(updatedName, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", updatedName),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageBindWithAnnotations(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageBindWithAnnotationsUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageBindWithParameters(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.3"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},

			{
				Config: testAccCheckIAMFunctionPackageBindWithParametersUpdate(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.bindpackage", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.4"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "false"),
				),
			},
			{
				Config: testAccCheckIAMFunctionPackageBindUpdatePublish(name, namespace, bindName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "version", "0.0.5"),
					resource.TestCheckResourceAttr("ibm_function_package.bindpackage", "publish", "true"),
				),
			},
		},
	})
}

func TestAccCFFunctionPackage_Import(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionPackageImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				ResourceName:      "ibm_function_package.package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIAMFunctionPackage_Import(t *testing.T) {
	var conf whisk.Package
	name := fmt.Sprintf("terraform_package_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionPackageDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionPackageImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionPackageExists("ibm_function_package.package", &conf),
					resource.TestCheckResourceAttr("ibm_function_package.package", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_package.package", "name", name),
					resource.TestCheckResourceAttr("ibm_function_package.package", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "annotations", "[]"),
					resource.TestCheckResourceAttr("ibm_function_package.package", "parameters", "[]"),
				),
			},

			{
				ResourceName:      "ibm_function_package.package",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionPackageExists(n string, obj *whisk.Package) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := flex.CfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
		if err != nil {
			return err
		}

		bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}
		client, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil {
			return err

		}

		pkg, _, err := client.Packages.Get(name)
		if err != nil {
			return err
		}

		*obj = *pkg
		return nil
	}
}

func testAccCheckFunctionPackageDestroy(s *terraform.State) error {
	functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_package" {
			continue
		}

		parts, err := flex.CfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		wskClient, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil && strings.Contains(err.Error(), "is not in the list of entitled namespaces") {
			return nil
		}
		if err != nil {
			return err
		}

		_, _, err = wskClient.Packages.Get(name)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for IBM Cloud Function Package (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIAMFunctionPackageCreate(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}	
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
}`, namespace, name)

}

func testAccCheckCFFunctionPackageCreate(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name = "%s"
		namespace = "%s"
}`, name, namespace)

}

func testAccCheckIAMFunctionPackageNameUpdate(updatedName string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}		
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
}`, namespace, updatedName)
}

func testAccCheckCFFunctionPackageNameUpdate(updatedName string, namespace string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "package" {
		name = "%s"
		namespace = "%s"
}`, updatedName, namespace)
}

func testAccCheckIAMFunctionPackageWithAnnotations(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}		
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name                     = "%s"
		namespace                = ibm_function_namespace.namespace.name
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  },
		  {
			  "key":"sampleOutput",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "A string",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name)

}

func testAccCheckCFFunctionPackageWithAnnotations(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                     = "%s"
		namespace                = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  },
		  {
			  "key":"sampleOutput",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "A string",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace)

}

func testAccCheckIAMFunctionPackageWithAnnotationsUpdate(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "package" {
		name                     = "%s"
		namespace                = ibm_function_namespace.namespace.name
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF 
	  }
`, namespace, name)

}

func testAccCheckCFFunctionPackageWithAnnotationsUpdate(name string, namespace string) string {
	return fmt.Sprintf(`

	
	resource "ibm_function_package" "package" {
		name                     = "%s"
		namespace                = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF 
	  }
`, name, namespace)

}

func testAccCheckIAMFunctionPackageWithParameters(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name                    = "%s"
		namespace               = ibm_function_namespace.namespace.name
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"place",
			  "value":"city"
		  },
		  {
			  "key":"parameter",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "Set of Values",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
	  
`, namespace, name)

}

func testAccCheckCFFunctionPackageWithParameters(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"place",
			  "value":"city"
		  },
		  {
			  "key":"parameter",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "Set of Values",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
	  
`, name, namespace)

}

func testAccCheckIAMFunctionPackageWithParametersUpdate(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name                    = "%s"
		namespace               = ibm_function_namespace.namespace.name
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"name",
			  "value":"utils"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name)

}

func testAccCheckCFFunctionPackageWithParametersUpdate(name string, namespace string) string {
	return fmt.Sprintf(`


	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"name",
			  "value":"utils"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace)

}

func testAccCheckIAMFunctionPackageImport(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
   		name = "%s"
		namespace = ibm_function_namespace.namespace.name
	}
`, namespace, name)

}

func testAccCheckCFFunctionPackageImport(name string, namespace string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "package" {
   		name = "%s"
		namespace = "%s"
	}
`, name, namespace)

}

func testAccCheckIAMFunctionPackageUpdatePublish(name string, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "package" {
		depends_on = [ibm_function_namespace.namespace]
		name                    = "%s"
		namespace               = ibm_function_namespace.namespace.name
		publish                 = true
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"name",
			  "value":"utils"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name)
}

func testAccCheckCFFunctionPackageUpdatePublish(name string, namespace string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "package" {
		name                    = "%s"
		namespace               = "%s"
		publish                 = true
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"name",
			  "value":"utils"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace)
}

func testAccCheckIAMFunctionPackageBindCreate(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	
	resource "ibm_function_package" "bindpackage" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		bind_package_name = "%s"
}`, namespace, name, bind)

}

func testAccCheckCFFunctionPackageBindCreate(name, namespace, bind string) string {
	return fmt.Sprintf(`
	
	resource "ibm_function_package" "bindpackage" {
		name = "%s"
		namespace = "%s"
		bind_package_name = "%s"
}`, name, namespace, bind)

}

func testAccCheckIAMFunctionPackageNameBindUpdate(updatedName, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   namespace = ibm_function_namespace.namespace.name
	   bind_package_name = "%s"
}`, namespace, updatedName, bind)
}

func testAccCheckCFFunctionPackageNameBindUpdate(updatedName, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
	   name = "%s"
	   namespace = "%s"
	   bind_package_name = "%s"
}`, updatedName, namespace, bind)
}

func testAccCheckIAMFunctionPackageBindWithAnnotations(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
		depends_on = [ibm_function_namespace.namespace]
		name                     = "%s"
		namespace                = ibm_function_namespace.namespace.name
		bind_package_name        = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"binded alaram package"
		  },
		  {
			  "key":"sampleOutput",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "A string",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name, bind)

}

func testAccCheckCFFunctionPackageBindWithAnnotations(name, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
		name                     = "%s"
		namespace                = "%s"
		bind_package_name        = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"binded alaram package"
		  },
		  {
			  "key":"sampleOutput",
			  "value": {
							  "count": 3
					  }
		  },
		  {
			  "key":"final",
			  "value": [
							  {
									  "description": "A string",
									  "name": "payload",
									  "required": true
							  }
					  ]
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace, bind)

}

func testAccCheckIAMFunctionPackageBindWithAnnotationsUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
		depends_on = [ibm_function_namespace.namespace]
		name                     = "%s"
		namespace                = ibm_function_namespace.namespace.name
		bind_package_name        = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"binded alaram package"
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name, bind)

}

func testAccCheckCFFunctionPackageBindWithAnnotationsUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
		name                     = "%s"
		namespace                = "%s"
		bind_package_name        = "%s"
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"binded alaram package"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace, bind)

}

func testAccCheckIAMFunctionPackageBindWithParameters(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
		depends_on = [ibm_function_namespace.namespace]
		name                    = "%s"
		namespace               = ibm_function_namespace.namespace.name
		bind_package_name       = "%s"
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"cron",
			  "value":"0 0 1 0 *"
		  },
		  {
			  "key":"trigger_payload ",
			  "value":"{'message':'bye old Year!'}"
		  },
		  {
			  "key":"maxTriggers",
			  "value":1
		  },
		  {
			  "key":"userdefined",
			  "value":"test"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name, bind)

}

func testAccCheckCFFunctionPackageBindWithParameters(name, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace               = "%s"
		bind_package_name       = "%s"
		user_defined_parameters = <<EOF
			  [
		  {
			  "key":"cron",
			  "value":"0 0 1 0 *"
		  },
		  {
			  "key":"trigger_payload ",
			  "value":"{'message':'bye old Year!'}"
		  },
		  {
			  "key":"maxTriggers",
			  "value":1
		  },
		  {
			  "key":"userdefined",
			  "value":"test"
		  }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace, bind)

}

func testAccCheckIAMFunctionPackageBindWithParametersUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
		depends_on = [ibm_function_namespace.namespace]
		name                    = "%s"
		namespace				= ibm_function_namespace.namespace.name
		bind_package_name       = "%s"
		user_defined_parameters = <<EOF
				 [
		 {
				 "key":"cron",
				 "value":"0 0 1 0 *"
		 }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, namespace, name, bind)

}

func testAccCheckCFFunctionPackageBindWithParametersUpdate(name, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace				= "%s"
		bind_package_name       = "%s"
		user_defined_parameters = <<EOF
				 [
		 {
				 "key":"cron",
				 "value":"0 0 1 0 *"
		 }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
`, name, namespace, bind)

}

func testAccCheckIAMFunctionPackageBindUpdatePublish(name, namespace, bind string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace				= ibm_function_namespace.namespace.name 
		bind_package_name       = "%s"
		publish                 = true
		user_defined_parameters = <<EOF
				 [
		 {
				 "key":"cron",
				 "value":"0 0 1 0 *"
		 }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
	  
`, namespace, name, bind)
}

func testAccCheckCFFunctionPackageBindUpdatePublish(name, namespace, bind string) string {
	return fmt.Sprintf(`

	resource "ibm_function_package" "bindpackage" {
		name                    = "%s"
		namespace				= "%s"
		bind_package_name       = "%s"
		publish                 = true
		user_defined_parameters = <<EOF
				 [
		 {
				 "key":"cron",
				 "value":"0 0 1 0 *"
		 }
	  ]
	  EOF
	  
	  
		user_defined_annotations = <<EOF
			  [
		  {
			  "key":"description",
			  "value":"Count words in a string"
		  }
	  ]
	  EOF
	  
	  }
	  
`, name, namespace, bind)
}

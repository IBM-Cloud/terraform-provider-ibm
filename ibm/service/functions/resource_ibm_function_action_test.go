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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/Mavrickk3/bluemix-go/bmxerror"
)

func TestAccIAMFunctionAction_NodeJS(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionNodeJS(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_NodeJSWithParams(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionNodeJSWithParams(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehellowithparameter", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_NodeJSZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionNodeJSZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodezip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_Python(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionPython(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonhello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "exec.0.kind", "python:3"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_PythonZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionPythonZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "exec.0.kind", "python:3"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_PHP(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionPHP(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phphello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "exec.0.kind", "php:7.3"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_PHPZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionPHPZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phpzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "exec.0.kind", "php:7.3"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_Swift(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionSwift(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.swifthello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "exec.0.kind", "swift:4.2"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionSequence(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.sequence", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "exec.0.kind", "sequence"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},

			{
				Config: testAccCheckIAMFunctionActionUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "true"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "5"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "50000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccIAMFunctionAction_Import(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := fmt.Sprintf("namespace_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIAMFunctionActionImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.import", "publish", "false"),
				),
			},

			{
				ResourceName:      "ibm_function_action.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCFFunctionAction_NodeJS(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionNodeJS(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_NodeJSWithParams(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCFFunctionActionNodeJSWithParams(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodehellowithparameter", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodehellowithparameter", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_NodeJSZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionNodeJSZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.nodezip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "exec.0.kind", "nodejs:10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.nodezip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_Python(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionPython(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonhello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "exec.0.kind", "python:3"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonhello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_PythonZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionPythonZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.pythonzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "exec.0.kind", "python:3"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.pythonzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_PHP(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionPHP(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phphello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "exec.0.kind", "php:7.3"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phphello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_PHPZip(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionPHPZip(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.phpzip", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "exec.0.kind", "php:7.3"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.phpzip", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_Swift(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionSwift(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.swifthello", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "exec.0.kind", "swift:4.2"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.swifthello", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_Sequence(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionSequence(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.sequence", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "exec.0.kind", "sequence"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.sequence", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_Basic(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionCreate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "false"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "10"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "60000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},

			{
				Config: testAccCheckCFFunctionActionUpdate(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.action", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.action", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.action", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.action", "version", "0.0.2"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "publish", "true"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.log_size", "5"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.timeout", "50000"),
					resource.TestCheckResourceAttr("ibm_function_action.action", "limits.0.memory", "256"),
				),
			},
		},
	})
}

func TestAccCFFunctionAction_Import(t *testing.T) {
	var conf whisk.Action
	name := fmt.Sprintf("terraform_action_%d", acctest.RandIntRange(10, 100))
	namespace := os.Getenv("IBM_FUNCTION_NAMESPACE")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckFunctionActionDestroy,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckCFFunctionActionImport(name, namespace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckFunctionActionExists("ibm_function_action.import", &conf),
					resource.TestCheckResourceAttr("ibm_function_action.import", "name", name),
					resource.TestCheckResourceAttr("ibm_function_action.import", "namespace", namespace),
					resource.TestCheckResourceAttr("ibm_function_action.import", "version", "0.0.1"),
					resource.TestCheckResourceAttr("ibm_function_action.import", "publish", "false"),
				),
			},

			{
				ResourceName:      "ibm_function_action.import",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckFunctionActionExists(n string, obj *whisk.Action) resource.TestCheckFunc {

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

		action, _, err := client.Actions.Get(name, true)
		if err != nil {
			return err
		}

		*obj = *action
		return nil
	}
}

func testAccCheckFunctionActionDestroy(s *terraform.State) error {
	functionNamespaceAPI, err := acc.TestAccProvider.Meta().(conns.ClientSession).FunctionIAMNamespaceAPI()
	if err != nil {
		return err
	}

	bxSession, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_function_action" {
			continue
		}

		parts, err := flex.CfIdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		namespace := parts[0]
		name := parts[1]

		client, err := conns.SetupOpenWhiskClientConfig(namespace, bxSession, functionNamespaceAPI)
		if err != nil && strings.Contains(err.Error(), "is not in the list of entitled namespaces") {
			return nil
		}
		if err != nil {
			return err
		}

		_, _, err = client.Actions.Get(name, true)

		if err != nil {
			if apierr, ok := err.(bmxerror.RequestFailure); ok && apierr.StatusCode() != 404 {
				return fmt.Errorf("[ERROR] Error waiting for IBM Cloud Function Action (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
	}
	return nil
}

func testAccCheckIAMFunctionActionNodeJS(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "nodehello" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
	  }
	
`, namespace, name)

}

func testAccCheckIAMFunctionActionNodeJSWithParams(name, namespace string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}
	
	resource "ibm_function_action" "nodehellowithparameter" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		user_defined_parameters = <<EOF
							   [
									   {
											  "key":"place",
											   "value":"India"
									  }
							  ]
	  
	  EOF
	  
	  }
	  
`, namespace, name)

}

func testAccCheckIAMFunctionActionNodeJSZip(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "nodezip" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = base64encode("../../test-fixtures/nodeaction.zip")
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionPython(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "pythonhello" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "python:3"
		  code = file("../../test-fixtures/helloPython.py")
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionPythonZip(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "pythonzip" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "python:3"
		  code_path = "../../test-fixtures/pythonaction.zip"
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionPHP(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "phphello" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "php:7.3"
		  code = file("../../test-fixtures/hellophp.php")
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionPHPZip(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "phpzip" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind 		= "php:7.3"
		  code_path = "../../test-fixtures/phpaction.zip"
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionSwift(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "swifthello" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "swift:4.2"
		  code = file("../../test-fixtures/helloSwift.swift")
		}
	  }
	
`, namespace, name)

}

func testAccCheckIAMFunctionActionSequence(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "sequence" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind       = "sequence"
		  components = ["/whisk.system/utils/split", "/whisk.system/utils/sort"]
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionCreate(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "action" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
		limits {
		}
	  }
`, namespace, name)

}

func testAccCheckIAMFunctionActionUpdate(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "action" {
		depends_on = [ibm_function_namespace.namespace]
		name    = "%s"
		namespace = ibm_function_namespace.namespace.name
		publish = "true"
		limits {
		  log_size = 5
		  timeout  = 50000
		}
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
										  "value":"mub"
								 }
						 ]
	  
	       EOF	
		  
	 }
	  
`, namespace, name)

}

func testAccCheckIAMFunctionActionImport(name, namespace string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default=true
	}

	resource "ibm_function_namespace" "namespace" {
		name                = "%s"
		resource_group_id   = data.ibm_resource_group.test_acc.id
	}

	resource "ibm_function_action" "import" {
		depends_on = [ibm_function_namespace.namespace]
		name = "%s"
		namespace = ibm_function_namespace.namespace.name
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, namespace, name)

}

func testAccCheckCFFunctionAction(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "nodehello" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
	  }
	
`, name, namespace)

}

func testAccCheckCFFunctionActionNodeJSWithParams(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "nodehellowithparameter" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		user_defined_parameters = <<EOF
							   [
									   {
											  "key":"place",
											   "value":"India"
									  }
							  ]
	  
	  EOF
	  
	  }
	  
`, name, namespace)

}

func testAccCheckCFFunctionActionNodeJSZip(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "nodezip" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = base64encode("../../test-fixtures/nodeaction.zip")
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionPython(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "pythonhello" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "python:3"
		  code = file("../../test-fixtures/helloPython.py")
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionPythonZip(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "pythonzip" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "python:3"
		  code = base64encode("../../test-fixtures/pythonaction.zip")
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionPHP(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "phphello" {
		name = "%s"
		namespace = "%s"	
		exec {
		  kind = "php:7.3"
		  code = file("../../test-fixtures/hellophp.php")
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionPHPZip(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "phpzip" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "php:7.3"
		  code = base64encode("../../test-fixtures/phpaction.zip")
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionSwift(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "swifthello" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "swift:4.2"
		  code = file("../../test-fixtures/helloSwift.swift")
		}
	  }
	
`, name, namespace)

}

func testAccCheckCFFunctionActionSequence(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "sequence" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind       = "sequence"
		  components = ["/whisk.system/utils/split", "/whisk.system/utils/sort"]
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionCreate(name, namespace string) string {
	return fmt.Sprintf(`

	resource "ibm_function_action" "action" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
		limits {
		}
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionUpdate(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "action" {
		name    = "%s"
		namespace = "%s"
		publish = "true"
		limits {
		  log_size = 5
		  timeout  = 50000
		}
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
										  "value":"mub"
								 }
						 ]
	  
	       EOF	
		  
	 }
	  
`, name, namespace)

}

func testAccCheckCFFunctionActionImport(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "import" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonodewithparameter.js")
		}
		user_defined_parameters = <<EOF
							  [
									  {
										 "key":"place",
											  "value":"India"
								 }
						 ]
	  
	  EOF
	  
	  
		user_defined_annotations = <<EOF
				 [
						 {
								"key":"Description",
								 "value":"Sample code to display hello"
						}
				]
	  
	  EOF
	  
	  }
`, name, namespace)

}

func testAccCheckCFFunctionActionNodeJS(name, namespace string) string {
	return fmt.Sprintf(`
	resource "ibm_function_action" "nodehello" {
		name = "%s"
		namespace = "%s"
		exec {
		  kind = "nodejs:10"
		  code = file("../../test-fixtures/hellonode.js")
		}
	  }
	
`, name, namespace)

}

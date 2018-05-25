# Contributing to IBMCloud Terraform Provider

**First:** if you're unsure  _anything_, just ask or submit the issue or pull request anyways. We appreciate any sort of contributions.

However, for those individuals who want a bit more guidance on the best way to contribute to the project, read on. This document will cover what we're looking for. By addressing all the points we're looking for, it raises the chances we can quickly merge or address your contributions.

Specifically, we have provided checklists below for each type of issue and pull request that can happen on the project. These checklists represent everything we need to be able to review and respond quickly.

## Issues

### Issue Reporting Checklists

We welcome issues of all kinds including feature requests, bug reports, and general questions. Below you'll find checklists with guidelines for well-formed issues of each type.

#### Bug Reports

 - [ ] __Test against latest release__: Make sure you test against the latest released version. It is possible we already fixed the bug you're experiencing.

 - [ ] __Search for possible duplicate reports__: It's helpful to keep bug reports consolidated to one thread, so do a quick search on existing bug reports to check if anybody else has reported the same thing. You can scope searches by the label "bug" to help narrow things down.

 - [ ] __Include steps to reproduce__: Provide steps to reproduce the issue, along with your `.tf` files, with secrets removed, so we can try to reproduce it. Without this, it makes it much harder to fix the issue.

 - [ ] __For panics, include `crash.log`__: If you experienced a panic, please create a [gist](https://gist.github.com) of the *entire* generated crash log for us to look at. Double check no sensitive items were in the log.

#### Feature Requests

 - [ ] __Search for possible duplicate requests__: It's helpful to keep requests consolidated to one thread, so do a quick search on existing requests to check if anybody else has reported the same thing. You can scope searches by the label "enhancement" to help narrow things down.

 - [ ] __Include a use case description__: In addition to describing the behavior of the feature you'd like to see added, it's helpful to also lay out the reason why the feature would be important and how it would benefit Terraform users.

#### Questions

 - [ ] __Search for answers in Terraform documentation__: We're happy to answer questions in GitHub Issues, but it helps reduce issue churn and maintainer workload if you work to find answers to common questions in the documentation. Often times Question issues result in documentation updates to help future users, so if you don't find an answer, you can give us pointers for where you'd expect to see it in the docs.

## Pull Requests

Thank you for contributing! Here you'll find information on what to include in your Pull Request to ensure it is accepted quickly.

 * For pull requests that follow the guidelines, we expect to be able to review and merge very quickly.
 * Pull requests that don't follow the guidelines will be annotated with what they're missing. A community or core team member may be able to swing around and help finish up the work, but these PRs will generally hang out much longer until they can be completed and merged.

### Checklists for Contribution

There are several different kinds of contribution, each of which has its own standards for a speedy review. The following sections describe guidelines for each type of contribution.


#### Enhancement/Bugfix to a Resource

Working on existing resources is a great way to get started as a Terraform contributor because you can work within existing code and tests to get a feel for what to do.

 - [ ] __Acceptance test coverage of new behavior__: Existing resources each have a set of [acceptance tests][acctests] covering their functionality. These tests should exercise all the behavior of the resource. Whether you are adding something or fixing a bug, the idea is to have an acceptance test that fails if your code were to be removed. Sometimes it is sufficient to "enhance" an existing test by adding an assertion or tweaking the config that is used, but often a new test is better to add. You can copy/paste an existing test and follow the conventions you see there, modifying the test to exercise the behavior of your code.

 - [ ] __Documentation updates__: If your code makes any changes that need to be documented, you should include those doc updates in the same PR. 
   
 - [ ] __Well-formed Code__: Do your best to follow existing conventions you see in the codebase, and ensure your code is formatted with `go fmt`. (The Travis CI build will fail if `go fmt` has not been run on incoming code.) The PR reviewers can help out on this front, and may provide comments with suggestions on how to improve the code.

#### New Resource

Implementing a new resource is a good way to learn more about how Terraform interacts with upstream APIs. There are plenty of examples to draw from in the existing resources, but you still get to implement something completely new.

 - [ ] __Minimal LOC__: It can be inefficient for both the reviewer and author to go through long feedback cycles on a big PR with many resources. We therefore encourage you to only submit **1 resource at a time**.
 - [ ] __Acceptance tests__: New resources should include acceptance tests covering their behavior. See [Writing Acceptance Tests](#writing-acceptance-tests) below for a detailed guide on how to approach these.
 - [ ] __Documentation__: Each resource gets a page in the Terraform documentation. The [Terraform website][website] source is in this repo and includes instructions for getting a local copy of the site up and running if you'd like to preview your changes. For a resource, you'll want to add a new file in the appropriate place and add a link to the sidebar for that page.
 - [ ] __Well-formed Code__: Do your best to follow existing conventions you see in the codebase, and ensure your code is formatted with `go fmt`. (The Travis CI build will fail if `go fmt` has not been run on incoming code.) The PR reviewers can help out on this front, and may provide comments with suggestions on how to improve the code.


### Writing Acceptance Tests

Terraform includes an acceptance test harness that does most of the repetitive work involved in testing a resource.

#### Acceptance Tests Often Cost Money to Run

Because acceptance tests create real resources, they often cost money to run. Because the resources only exist for a short period of time, the total amount of money required is usually a relatively small. Nevertheless, we don't want financial limitations to be a barrier to contribution, so if you are unable to pay to run acceptance tests for your contribution, simply mention this in your pull request. We will happily accept "best effort" implementations of acceptance tests and run them for you on our side. This might mean that your PR takes a bit longer to merge, but it most definitely is not a blocker for contributions.

#### Running an Acceptance Test

Acceptance tests can be run using the `testacc` target in the `Makefile`. The individual tests to run can be controlled using a regular expression. Prior to running the tests provider configuration details such as access keys must be made available as environment variables.

For example, to run an acceptance test, the following environment variables must be set:

```sh
export BM_API_KEY=...
export SL_API_KEY=...
export SL_USERNAME=...
```

For certain tests, the following values may also needs to be set:

```sh
export IBM_ORG=...
export IBM_SPACE=...
export IBM_ID1=...
export IBM_ID2=...
export IBM_IAMUSER=...
```

You can enable the terraform logs by setting the following environment variable:
```sh
export TF_LOG=DEBUG
```

Tests can then be run by specifying the target provider and a regular expression defining the tests to run:

```sh
$ make testacc TEST=./ibm TESTARGS='-run=TestAccIBMComputeVmInstance_basic'
==> Checking that code complies with gofmt requirements...
go generate ./...
TF_ACC=1 go test ./ibm -v -run=TestAccIBMComputeVmInstance_basic -timeout 700m
=== RUN   TestAccIBMComputeVmInstance_basic
--- PASS: TestAccIBMComputeVmInstance_basic (177.48s)
PASS
ok      github.com/terraform-providers/terraform-provider-ibm/ibm   177.504s
```

Entire resource test suites can be targeted by using the naming convention to write the regular expression. For example, to run all tests of the `ibm_compute_vm_instance` resource rather than just the update test, you can start testing like this:

```sh
$ make testacc TEST=./ibm TESTARGS='-run=TestAccIBMComputeVmInstance'
==> Checking that code complies with gofmt requirements...
go generate ./...
TF_ACC=1 go test ./builtin/providers/azurerm -v -run=TestAccIBMComputeVmInstance -timeout 700m
=== RUN   TestAccIBMComputeVmInstance_basic
--- PASS: TestAccIBMComputeVmInstance_basic (137.74s)
=== RUN   TestAccIBMComputeVmInstance_basic_import
--- PASS: TestAccIBMComputeVmInstance_basic_import (180.63s)
PASS
ok      github.com/terraform-providers/terraform-provider-ibm/ibm   318.392s
```

#### Writing an Acceptance Test

Terraform has a framework for writing acceptance tests which minimises the amount of boilerplate code necessary to use common testing patterns. The entry point to the framework is the `resource.Test()` function.

Tests are divided into `TestSteps`. Each `TestStep` proceeds by applying some
Terraform configuration using the provider under test, and then verifying that results are as expected by making assertions using the provider API. It is common for a single test function to exercise both the creation of and updates to a single resource. Most tests follow a similar structure.

1. Pre-flight checks are made to ensure that sufficient provider configuration is available to be able to proceed - for example in an acceptance test `SL_API_KEY` , `SL_USERNAME` and `BM_API_KEY` must be set prior to running acceptance tests. This is common to all tests exercising a single provider.

Each `TestStep` is defined in the call to `resource.Test()`. Most assertion functions are defined out of band with the tests. This keeps the tests readable, and allows reuse of assertion functions across different tests of the same type of resource. The definition of a complete test looks like this:

```go
func TestAccIBMComputeVmInstance_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccIBMComputeVmInstanceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIBMComputeVmInstanceConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccIBMComputeVmInstanceExists("ibm_compute_vm_instance.terraform-acceptance-test-1"),
				),
			},
        },
    })
}
```

When executing the test, the following steps are taken for each `TestStep`:

1. The Terraform configuration required for the test is applied. This is responsible for configuring the resource under test, and any dependencies it may have. For example, to test the `ibm_compute_vm_instance` resource. This results in configuration which looks like this:

```hcl
resource "ibm_compute_vm_instance" "terraform-acceptance-test-1" {
   hostname = "terraform-sample-blockDeviceTemplateGroup"
   domain = "bar.example.com"
   datacenter = "ams01"
   public_network_speed = 10
   hourly_billing = false
   cores = 1
   memory = 1024
   local_disk = false
   image_id = 12345
   tags = [
     "collectd",
     "mesos-master"
   ]
   public_subnet = "50.97.46.160/28"
   private_subnet = "10.56.109.128/26"
} 
```

2. Assertions are run using the provider API. These use the provider API directly rather than asserting against the resource state. For example, to verify that the `ibm_compute_vm_instance` described above was created successfully, a test function like this is used:

```go
    func resourceIBMComputeVmInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	service := services.GetVirtualGuestService(meta.(ClientSession).SoftLayerSession())
	guestID, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}

	result, err := service.Id(guestID).GetObject()
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return result.Id != nil && *result.Id == guestID, nil
	}
```

Notice that the only information used from the Terraform state is the ID of the resource - though in this case it is necessary to split the ID into constituent parts in order to use the provider API. For computed properties, we instead assert that the value saved in the Terraform state was the expected value if possible. The testing framework provides helper functions for several common types of check - for example:

```go
    resource.TestCheckResourceAttr("ibm_compute_vm_instance.terraform-test-1", "hourly_billing", "true"),
```

1. The resources created by the test are destroyed. This step happens automatically, and is the equivalent of calling `terraform destroy`.

2. Assertions are made against the provider API to verify that the resources have indeed been removed. If these checks fail, the test fails and reports "dangling resources". The code to ensure that the `ibm_compute_vm_instance` shown above looks like this:

```go
    go func testAccIBMComputeVmInstanceDestroy(s *terraform.State) error {
	service := services.GetVirtualGuestService(testAccProvider.Meta().(ClientSession).SoftLayerSession())

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_compute_vm_instance" {
			continue
		}

		guestID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the guest
		_, err := service.Id(guestID).GetObject()

		// Wait

		if err != nil && !strings.Contains(err.Error(), "404") {
			return fmt.Errorf(
				"Error waiting for virtual guest (%s) to be destroyed: %s",
				rs.Primary.ID, err)
		}
	}
	return nil
	}
```

These functions usually test only for the resource directly under test.

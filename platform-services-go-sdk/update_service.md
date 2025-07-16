# How to update a service
This document describes the steps needed to update a service contained in this SDK project.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i update_service.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Initial project setup](#initial-project-setup)
- [Steps to update a service](#steps-to-update-a-service)
  * [1. Validate the API definition](#1-validate-the-api-definition)
  * [2. Create feature branch](#2-create-feature-branch)
  * [3. Re-generate the SDK code](#3-re-generate-the-sdk-code)
  * [4. Inspect new generated SDK code](#4-inspect-new-generated-sdk-code)
  * [5. Run unit tests](#5-run-unit-tests)
  * [6. Modify integration tests and examples](#6-modify-integration-tests-and-examples)
  * [7. Open PR with your changes](#7-open-pr-with-your-changes)
- [Appendix](#appendix)
  * [Running Integration Tests/Examples](#running-integration-testsexamples)
  * [Updating Integration Tests/Examples](#updating-integration-testsexamples)
- [References](#references)

<!-- tocstop -->

## Overview 
It is a good practice to keep the SDK code for each service updated so that it is in sync
with the most recent production version of its API definition.
So, when a service's API definition is changed, the SDK code for the service should be updated
(re-generated) in each SDK project in which it exists.
This could be a change such as (a) editorial changes made to various descriptions, (b) the addition of
a new parameter to an existing operation, or (c) the addition of one or more new operations.

## Prerequisites
1. If you are an IBMer, make sure that your
[Annual Open Source Training](https://w3.ibm.com/developer/docs/open-source/training/) is current.
2. Make sure that your internal github.ibm.com id is [linked](https://gh-user-map.dal1a.cirrus.ibm.com/)
to your external github.com id. The id linking step will also result in an invitation to join the
`github.com/IBM` org. Accept that invitation.
3. If you do not yet have "push" access to the SDK project, contact the project maintainer to request push access
(you must be a member of the github.com/IBM org).
4. Make sure that you have installed the [tools required to build the project](CONTRIBUTING.md#prerequisites).
5. To update a service, make sure the following additional tools are installed:
* The [IBM OpenAPI Validator](https://github.com/IBM/openapi-validator)
* The [IBM OpenAPI SDK Generator](github.ibm.com/CloudEngineering/openapi-sdkgen)

## Initial project setup
1. Clone/fork the repo.  If you have push access (see above), you can clone the repo directly (no fork).  
Example:  
```sh
git clone git@github.com:IBM/platform-services-go-sdk.git
```
2. If you do not have push access, then you'll need to first create a fork and then clone your fork in your
local sandbox environment.  
Example:  
```sh
git clone git@github.com:my-git-id/platform-services-go-sdk.git
```
3. Make sure that your local sandbox is in sync with the remote and then build/test the project. If you're
using a fork, you'd need to first make sure that your fork is in sync with the primary repo.  
Example:    
```sh
cd <project-root>
git checkout main
git pull
make all               # This runs unit tests and the linter
```
4. Make sure that the integration tests and working examples run clean for your service.
See [Running Integration Tests/Examples](#running-integration-testsexamples) for details.


Before proceeding to make any changes, make sure the above steps complete cleanly.  This is your "baseline".

## Steps to update a service

### 1. Validate the API definition
Prior to re-generating the SDK code for your service, be sure to validate the updated version of the API definition
using the [IBM OpenAPI Validator](https://github.com/IBM/openapi-validator).  
Example:  
```sh
lint-openapi example-service.yaml
```
This command will display a list of errors and warnings found in the API definition
as well as a summary at the end.
It's not required that you fix all errors and warnings before trying to use the SDK generator, but
this step should identify any critical errors that will need to be fixed prior to the generation step.

Video: [Getting Started With The OpenAPI Validator](https://secure.video.ibm.com/channel/23887899/playlist/651457/video/131770428)

### 2. Create feature branch
After validating the API definition, you're ready to generate new SDK code for your service.
However, before you do that, you should probably create a new feature branch in which to deliver your updates:  
```sh
cd <project-root>
git checkout -b update-example-service
```


### 3. Re-generate the SDK code
Next, run the [IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen) to
process your API definition and generate new service and unit test code for the service:  
```sh
cd <project-root>
openapi-sdkgen.sh generate -g ibm-go -i example-service.json -o .
```
The generated service and unit test code is written to the service's package directory within the SDK project
(e.g. `./exampleservicev1`).

Video: [Getting Started With The SDK Generator](https://secure.video.ibm.com/channel/23887899/playlist/651457/video/131770438)


### 4. Inspect new generated SDK code
Next, it is recommended that you inspect the differences between the previous and new generated code to
get an overall view of the changes caused by the re-generation step. The changes that you see in the
generated SDK code should align with the API definition changes that have occurred since you last
generated the SDK code.  
Example:  
```
git diff     # alternative: use the "source control" view within vscode
```


### 5. Run unit tests
Next, run the unit tests. You can run the unit tests for all the services like this:  
```sh
cd <project-root>
make all
```
or you can run the unit tests for your particular service like this:  
```sh
cd <project-root>/exampleservicev1
go test
```
The unit tests should run clean.  If not, then any test failures should be diagnosed and resolved
before proceeding.


### 6. Modify integration tests and examples
After ensuring that your service's unit tests run clean, the next step would be to modify
your service's integration tests and working examples code to reflect the updated version of
your API definition.  See [Updating Integration Tests/Examples](#updating-integration-testsexamples)
for more information on this topic.

Even if no changes are needed (perhaps only very minor updates were made to the generated
SDK code), at a mininum you should make sure that the integration tests and examples run clean after you
re-generate the service and unit test code.

For instructions on running the integration tests and examples code,
see [Running Integration Tests/Examples](#running-integration-testsexamples).


### 7. Open PR with your changes
After completing the previous steps to update the service, unit test, integration test, and working examples
code, commit your changes.  
Example:  
```sh
cd <project-root>
git add .
git commit -s -m "feat(Example Service): re-gen service after recent API changes"
git push
```
Note: be sure to sign off on your commits (git commit `-s` option) as that is a required PR check within the
github.com/IBM org.

Finally, open a pull request (PR) and tag the project maintainer for approval.

## Appendix
### Running Integration Tests/Examples
To run the integration tests and working examples for a particular service, follow these steps.  We'll use
the mythical "Example Service" within the examples below, but you can make the necessary adjustments for your
own service.

1. Make sure you have the required .env file in your project root directory. Each service's integration
test and working examples code assumes that external configuration properties (service URL, IAM ApiKey, etc.)
are stored in a .env file located in the project's root directory. The name of the file can be found in
the integration test and examples code.  
Example:  
```go
        const externalConfigFile = "../example_service_v1.env"
```
The precise set of configuration properties required by each service will vary somewhat among the services,
but there are a minimal set of properties that are commonly required by every service.  The integration tests
and examples code for certain services might require additional service-specific configuration properties as well.
Typically these are documented in the working examples code.

2. Make sure that you have built/unit-tested the project successfully before trying to run the integration tests
and/or examples:
```sh
make all
```

3. To run the integration tests and examples for a service, follow these steps:  
```sh
cd <project-root>/exampleservicev1
go test -tags=integration      # Runs unit and integration tests
go test -tags=examples         # Runs unit tests and examples code
```
For each of the `go` commands above, you should see 100% clean test results
with no tests being skipped.

### Updating Integration Tests/Examples
Certain types of API changes will require that the integration tests and examples code
are also updated along with the re-generated SDK service and unit test code.
For example, perhaps a new operation was introduced or a new parameter was added
to an existing operation and you'd like to incorporate it in the integration tests
and examples.

Keep in mind that the integration tests are used to verify that the
generated SDK code interacts correctly with the service implementation, so any non-trivial changes
made to the API definition (and hence the generated service code) should probably result in updates
to the integration tests.  At a minimum, the integration tests for a service should include a
testcase for EACH operation.

While modifying the integration tests, also consider if you should make any changes to the service's
working examples code.  We want the working examples to provide a good example for users
to follow when writing their own application code which uses your service, so consider whether or not
the examples code should be updated to reflect the changes made to the API.

The integration tests and examples code for each service were initially
generated by the SDK generator, then (most likely) manual changes were made
so that the tests and examples run cleanly using realistic values for
various parameters and properties.  The amount of manual changes required will vary from
one service to the next, but usually depends on the degree to which your API definition:
1. Includes good, realistic example values for operation parameters, request bodies, and responses.
2. Includes links that capture any inter-operation dependencies (e.g. the `create_cloud` operation's `id`
response property's value should be used as the `get_cloud` operation's `cloud_id` path parameter).

Regardless, it is likely that the integration tests and examples code have **some** manual
changes which will need to be retained as you apply updates to them to reflect the current
changes being made to the API.

Therefore, it is not recommended that you simply re-generate the integration tests and examples code
such that the existing files are overwritten.   Instead, we recommend that you generate new integration tests
and examples off to the side, then manually copy fragments from the newly-generated files to the existing 
files located in the SDK project.  
Example:
```sh
openapi-sdkgen.sh generate -g ibm-go -i example-service.json --genITs --genExamples -o /tmp/code
```
The newly-generated integration tests and examples would be found in `/tmp/code/exampleservicev1`
(the SDK generator automatically adds the service's package directory name),
You could then copy fragments from there as needed to modify the corresponding files in the SDK project.
This is not ideal, but you can minimize the amount of manual changes by improving your API definition
as mentioned above.

## References
- [IBM OpenAPI Validator](https://github.com/IBM/openapi-validator)
- [IBM OpenAPI SDK Generator](https://github.ibm.com/CloudEngineering/openapi-sdkgen)
- [Effective Go - The Go Programming Language](https://golang.org/doc/effective_go)
- [Go Documentation: Download and install](https://go.dev/doc/install)

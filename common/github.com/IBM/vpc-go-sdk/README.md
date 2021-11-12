[![Build Status](https://travis-ci.com/IBM/vpc-go-sdk.svg?branch=master)](https://travis-ci.com/IBM/vpc-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/vpc-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud VPC Go SDK Version 0.4.2
Go client library to interact with the various [IBM Cloud VPC Services APIs](https://cloud.ibm.com/apidocs?category=vpc).

**Note** As IBM continues to invest and innovate on the IBM Cloud Virtual Private Cloud (gen 2 compute) infrastructure, we're focusing on delivering maximum value in a single VPC Infrastructure platform. To support this effort, generation 1 compute infrastructure is being deprecated. The end of service date is 26 February 2021. For more information, see the [Start your migration](https://www.ibm.com/cloud/blog/announcements/start-your-vpc-gen1-to-vpc-gen2-migration) blog.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
    + [`go get` command](#go-get-command)
    + [Go modules](#go-modules)
    + [`dep` dependency manager](#dep-dependency-manager)
- [Using the SDK](#using-the-sdk)
- [Setting up VPC service](#setting-up-vpc-service)
- [Setting up VPC on Classic service](#setting-up-vpc-on-classic-service)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud VPC Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name
--- | ---
[VPC](https://cloud.ibm.com/apidocs/vpc) | vpcv1
[VPC Gen 1](https://cloud.ibm.com/apidocs/vpc-on-classic) | vpcclassicv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
There are a few different ways to download and install the VPC Go SDK services for use by your
Go application:

#### `go get` command
Use this command to download and install the VPC Classic Go SDK service to allow your Go application to
use it:

```
go get -u github.com/IBM/vpc-go-sdk/vpcv1
```

To install VPC Classic Go SDK service, use the following.

```
go get -u github.com/IBM/vpc-go-sdk/vpcclassicv1
```

#### Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:


```go
import (
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.


#### `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/vpc-go-sdk/"
  version = "0.4.2"
```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Setting up VPC service

A quick example to get you up and running with VPC Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC](https://cloud.ibm.com/apidocs/vpc#api-endpoint)  and update the `URL` variable accordingly.


Refer to the [VPC Release Notes](https://cloud.ibm.com/docs/vpc?topic=vpc-release-notes) document to find out latest version release.

```go
package main

import (
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func main() {
	apiKey := os.Getenv("IBMCLOUD_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	// Instantiate the service with an API key based IAM authenticator
	vpcService, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		log.Fatal("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	regions, detailedResponse, err := vpcService.ListRegions(&vpcv1.ListRegionsOptions{})
	if err != nil {
		log.Fatalf("Failed to list the regions: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("Regions: %#v", regions)

	// Retrieve the list of vpcs for your account.
	vpcs, detailedResponse, err := vpcService.ListVpcs(&vpcv1.ListVpcsOptions{})
	if err != nil {
		log.Fatalf("Failed to list vpcs: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("VPCs: %#v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcv1.CreateKeyOptions{
		Name: core.StringPtr("my-ssh-key"),
	}
	// Setters also exist to set fields are the struct has been created
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	detailedResponse, err = vpcService.DeleteKey(&vpcv1.DeleteKeyOptions{
		ID: key.ID,
	})
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %v and the response is: %s", err, detailedResponse)
	}

	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)
}
```

## Setting up VPC on Classic service

A quick example to get you up and running with VPC on Classic Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC on Classic](https://cloud.ibm.com/apidocs/vpc-on-classic#api-endpoint) and update the `URL` variable accordingly.

Refer to the [VPC on Classic Release Notes](https://cloud.ibm.com/docs/vpc-on-classic?topic=vpc-on-classic-release-notes) document to find out latest version release.

```go
package main

import (
	"log"
	"os"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
)

func main() {
	apiKey := os.Getenv("IBMCLOUD_API_KEY")
	if apiKey == "" {
		log.Fatal("No API key set")
	}

	// Instantiate the service with an api key based IAM autenticator
	vpcService, err := vpcclassicv1.NewVpcClassicV1(&vpcclassicv1.VpcClassicV1Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apiKey,
		},
	})
	if err != nil {
		log.Fatal("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	regions, detailedResponse, err := vpcService.ListRegions(&vpcclassicv1.ListRegionsOptions{})
	if err != nil {
		log.Fatalf("Failed to list the regions: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("Regions: %#v", regions)

	// Retrieve the list of vpcs for your account.
	vpcs, detailedResponse, err := vpcService.ListVpcs(&vpcclassicv1.ListVpcsOptions{})
	if err != nil {
		log.Fatalf("Failed to list vpcs: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("VPCs: %#v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcclassicv1.CreateKeyOptions{
		Name: core.StringPtr("my-ssh-key"),
	}
	// Setters also exist to set fields are the struct has been created
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	detailedResponse, err = vpcService.DeleteKey(&vpcclassicv1.DeleteKeyOptions{
		ID: key.ID,
	})
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %v and the response is: %s", err, detailedResponse)
	}
	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)
}
```

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](<github-repo-url>/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

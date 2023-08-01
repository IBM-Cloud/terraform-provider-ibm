[![Build Status](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)

# IBM Cloud VPC Go SDK Version 0.0.1
Go client library to interact with the various [IBM Cloud VPC Services APIs](https://cloud.ibm.com/apidocs?category=vpc).

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
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Example set up for VPC service](#example-set-up-for-VPC-service)
- [Example set up for VPC on Classic service](#example-set-up-for-VPC-on-Classic-service)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->


## Overview

The IBM Cloud VPC Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name
--- | ---
[vpcclassicv1](https://cloud.ibm.com/apidocs/vpc-on-classic) | Virtual Private Cloud on Classic
[vpcv1](https://cloud.ibm.com/apidocs/vpc) | Virtual Private Cloud

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
go get -u github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1
```

To install VPC Go SDK service, use the following.

```
go get -u github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1
```

#### Go modules
If your application is using Go modules, you can add a suitable import to your
Go application, like this:


```go
import (
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.


#### `dep` dependency manager
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
  version = "0.0.1"
[[constraint]]
  name = "github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
  version = "0.0.1"
```

then run `dep ensure`.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

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


## Example set up for VPC service

A quick example to get you up and running with VPC Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC](https://cloud.ibm.com/apidocs/vpc#api-endpoint)  and update the `URL` variable accordingly.


Refer to the [VPC Release Notes](https://cloud.ibm.com/docs/vpc?topic=vpc-release-notes) document to find out latest version release.

```go
package main

import (
	"log"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

var APIKey = "YOUR_KEY_HERE"                            // required, Add a valid API key here

func main() {
	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: APIKey,
	}

	// Create the service options struct.
	options := &vpcv1.VpcV1Options{
		Authenticator: authenticator,
	}

	// Instantiate the service.
	vpcService, vpcServiceErr := vpcv1.NewVpcV1(options)

	if vpcServiceErr != nil {
		log.Fatalf("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	listRegionsOptions := &vpcv1.ListRegionsOptions{}
	regions, detailedResponse, err := vpcService.ListRegions(listRegionsOptions)
	if err != nil {
		log.Fatalf("Failed to list the regions: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("Regions: %+v", regions)

	// Retrieve the list of vpcs for your account.
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	vpcs, detailedResponse, err := vpcService.ListVpcs(listVpcsOptions)
	if err != nil {
		log.Fatalf("Failed to list vpcs: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("VPCs: %+v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcv1.CreateKeyOptions{}
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	deleteKeyOptions := &vpcv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(*key.ID)
	detailedResponse, err = vpcService.DeleteKey(deleteKeyOptions)
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)
}
```

## Example set up for VPC on Classic service

A quick example to get you up and running with VPC on Classic Go SDK service in Dallas (us-south) region.

For other regions, Refer [API Endpoints for VPC on Classic](https://cloud.ibm.com/apidocs/vpc-on-classic#api-endpoint) and update the `URL` variable accordingly.

Refer to the [VPC on Classic Release Notes](https://cloud.ibm.com/docs/vpc-on-classic?topic=vpc-on-classic-release-notes) document to find out latest version release.

```go
package main

import (
	"fmt"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
)

var APIKey = "YOUR_KEY_HERE" 							// required, Add a valid API key here

func main() {
	// Create the IAM authenticator.
	authenticator := &core.IamAuthenticator{
		ApiKey: APIKey,
	}

	// Create the service options struct.
	options := &vpcclassicv1.VpcClassicV1Options{
		Authenticator: authenticator,
	}

	// Instantiate the service.
	vpcService, vpcServiceErr := vpcclassicv1.NewVpcClassicV1(options)

	if vpcServiceErr != nil {
		log.Fatalf("Error creating VPC Service.")
	}

	// Retrieve the list of regions for your account.
	listRegionsOptions := &vpcclassicv1.ListRegionsOptions{}
	regions, detailedResponse, err := vpcService.ListRegions(listRegionsOptions)
	if err != nil {
		log.Fatalf("Failed to list the regions: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("Regions: %+v", regions)

	// Retrieve the list of vpcs for your account.
	listVpcsOptions := &vpcclassicv1.ListVpcsOptions{}
	vpcs, detailedResponse, err := vpcService.ListVpcs(listVpcsOptions)
	if err != nil {
		log.Fatalf("Failed to list vpcs: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("VPCs: %+v", vpcs)

	// Create an SSH key
	sshKeyOptions := &vpcclassicv1.CreateKeyOptions{}
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetName("my-ssh-key")
	sshKeyOptions.SetPublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDsnrSAe8eBi8mS576Z96UtYgUzDR9Sbw/s1ELxsa1KUK82JQ0Ejmz31N6sHyiT/l5533JgGL6rKamLFziMY2VX2bdyuF5YzyHhmapT+e21kuTatB50UsXzxlYEWpCmFdnd4LhwFn6AycJWOV0k3e0ePpVxgHc+pVfE89322cbmfuppeHxvxc+KSzQNYC59A+A2vhucbuWppyL3EIF4YgLwOr5iDISm1IR0+EEL3yJQIG4M2WKu526anI85QBcIWyFwQXOpdcX2eZRcd6WW2EgAM3fIOaezkm0CFrsz8rQ0MPYZI4BS2CWwg5d4Bj7SU2sjXz62gfQkQGTYWSqhizVb root@localhost")
	key, detailedResponse, err := vpcService.CreateKey(sshKeyOptions)
	if err != nil {
		log.Fatalf("Failed to create the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s created with ID: %s", *key.Name, *key.ID)

	// Delete SSH key
	deleteKeyOptions := &vpcclassicv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(*key.ID)
	detailedResponse, err = vpcService.DeleteKey(deleteKeyOptions)
	if err != nil {
		log.Fatalf("Failed to delete the ssh key: %s and the response is: %s", err.Error(), detailedResponse)
	}
	log.Printf("SSH key: %s deleted with ID: %s", *key.Name, *key.ID)

}
```


## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
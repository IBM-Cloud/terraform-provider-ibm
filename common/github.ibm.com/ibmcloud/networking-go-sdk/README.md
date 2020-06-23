[![Build Status](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)](https://travis.ibm.com/CloudEngineering/go-sdk-template.svg?token=eW5FVD71iyte6tTby8gr&branch=master)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Cloud Networking Go SDK Version 0.0.1
Go client library to interact with the various [IBM Cloud Networking Service APIs](https://cloud.ibm.com/apidocs?category=<networking>).

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
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Networking Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Transit Gateway Service](https://cloud.ibm.com/apidocs/example-service) | transitgatewayapisv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.1

There are a few different ways to download and install the Networking Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.ibm.com/ibmcloud/networking-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.ibm.com/ibmcloud/networking-go-sdk"
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
[bug report](github.ibm.com/ibmcloud/networking-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

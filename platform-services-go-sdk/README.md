[![Build Status](https://github.com/IBM/platform-services-go-sdk/actions/workflows/build.yaml/badge.svg)](https://github.com/IBM/platform-services-go-sdk/actions/workflows/build.yaml)
[![Release](https://img.shields.io/github/v/release/IBM/platform-services-go-sdk)](https://github.com/IBM/platform-services-go-sdk/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/IBM/platform-services-go-sdk.svg)](https://pkg.go.dev/github.com/IBM/platform-services-go-sdk)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/IBM/platform-services-go-sdk)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)
[![CLA assistant](https://cla-assistant.io/readme/badge/IBM/platform-services-go-sdk)](https://cla-assistant.io/IBM/platform-services-go-sdk)


# IBM Cloud Platform Services Go SDK Version 0.84.0

Go client library to interact with various
[IBM Cloud Platform Service APIs](https://cloud.ibm.com/docs?tab=api-docs&category=platform_services).

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
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Platform Services Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Activity Tracker](https://cloud.ibm.com/apidocs/atracker/atracker-v2) | atrackerv2
[Case Management](https://cloud.ibm.com/apidocs/case-management?code=go) | casemanagementv1
[Catalog Management](https://cloud.ibm.com/apidocs/resource-catalog/private-catalog?code=go) | catalogmanagementv1
[Context Based Restrictions](https://cloud.ibm.com/apidocs/context-based-restrictions?code=go) | contextbasedrestrictionsv1
[Enterprise Billing Units](https://cloud.ibm.com/apidocs/enterprise-apis/billing-unit?code=go) | enterprisebillingunitsv1
[Enterprise Management](https://cloud.ibm.com/apidocs/enterprise-apis/enterprise?code=go) | enterprisemanagementv1
[Enterprise Usage Reports](https://cloud.ibm.com/apidocs/enterprise-apis/resource-usage-reports?code=go) | enterpriseusagereportsv1
[Global Catalog](https://cloud.ibm.com/apidocs/resource-catalog/global-catalog?code=go) | globalcatalogv1
[Global Search](https://cloud.ibm.com/apidocs/search?code=go) | globalsearchv2
[Global Tagging](https://cloud.ibm.com/apidocs/tagging?code=go) | globaltaggingv1
[IAM Access Groups](https://cloud.ibm.com/apidocs/iam-access-groups?code=go) | iamaccessgroupsv2
[IAM Identity Service](https://cloud.ibm.com/apidocs/iam-identity-token-api?code=go) | iamidentityv1
[IAM Policy Management](https://cloud.ibm.com/apidocs/iam-policy-management?code=go) | iampolicymanagementv1
[IBM Cloud Shell](https://cloud.ibm.com/apidocs/cloudshell?code=go) | ibmcloudshellv1
[Metrics Router](https://cloud.ibm.com/apidocs/metrics-router/metrics-router-v3) | metricsrouterv3
[Open Service Broker](https://cloud.ibm.com/apidocs/resource-controller/ibm-cloud-osb-api?code=go) | openservicebrokerv1
[Partner Center Sell](https://test.cloud.ibm.com/apidocs/partner-center-sell?code=go) | partnercentersellv1
[Partner Management APIs](https://cloud.ibm.com/apidocs/partner-apis/partner?code=go) | partnermanagementv1
[Resource Controller](https://cloud.ibm.com/apidocs/resource-controller/resource-controller?code=go) | resourcecontrollerv2
[Resource Manager](https://cloud.ibm.com/apidocs/resource-controller/resource-manager?code=go) | resourcemanagerv2
[Usage Metering](https://cloud.ibm.com/apidocs/usage-metering?code=go) | usagemeteringv4
[Usage Reports](https://cloud.ibm.com/apidocs/metering-reporting?code=go) | usagereportsv4
[User Management](https://cloud.ibm.com/apidocs/user-management?code=go) | usermanagementv1

The following services have been relocated to a different SDK project.
Please consult the documentation for each service to determine the new location:

Service Name | Package name 
--- | --- 
[Configuration Governance](https://cloud.ibm.com/apidocs/security-compliance/config?code=go) | configurationgovernancev1
[Posture Management](https://cloud.ibm.com/apidocs/security-compliance/posture?code=go) | posturemanagementv1

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one
[here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.23 or above.

## Installation
The current version of this SDK: 0.84.0

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/platform-services-go-sdk/globalsearchv2"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `globalsearchv2` part of the import path is the package name
associated with the Global Search service.
See the service table above to find the approprate package name for the services used by your application.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/platform-services-go-sdk/globalsearchv2
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see
[this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/platform-services-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

The IBM Cloud Platform Services Go SDK is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).

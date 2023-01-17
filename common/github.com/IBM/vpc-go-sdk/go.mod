// module github.com/IBM-Cloud/terraform-provider-ibm/vendor/github.com/softlayer/softlayer-go
module github.com/IBM/vpc-go-sdk

go 1.12

// replace github.com/IBM/vpc-go-sdk/common => ./common/github.com/IBM/vpc-go-sdk/common

require (
	github.com/IBM/go-sdk-core/v5 v5.9.5
	github.com/go-openapi/strfmt v0.21.2
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.19.0
)

module github.com/IBM/vpc-go-sdk

go 1.16

require (
	github.com/IBM/go-sdk-core/v5 v5.9.5
	github.com/go-openapi/strfmt v0.21.1
	github.com/onsi/ginkgo v1.14.2
	github.com/onsi/gomega v1.10.5
	github.com/stretchr/testify v1.7.0
)

retract (
	v1.0.2
	v1.0.1
	v1.0.0
)
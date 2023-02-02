module github.com/IBM/vpc-go-sdk

go 1.16

require (
	github.com/IBM/go-sdk-core/v5 v5.10.2
	github.com/go-openapi/strfmt v0.21.3
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.20.0
	github.com/stretchr/testify v1.8.0
)

retract (
	v1.0.2
	v1.0.1
	v1.0.0
)

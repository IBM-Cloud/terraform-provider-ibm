module github.com/IBM/vpc-go-sdk

go 1.16

require (
	github.com/IBM/go-sdk-core/v5 v5.16.0
	github.com/go-openapi/strfmt v0.22.1
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.8.4
)

retract (
	v1.0.2
	v1.0.1
	v1.0.0
)

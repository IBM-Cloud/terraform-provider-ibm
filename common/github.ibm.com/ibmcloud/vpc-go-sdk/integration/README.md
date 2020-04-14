# IBM Cloud Go SDK Test Harness
Test harness serve as a test suite for VPC's Go SDK.

## How to use this test harness to test Go SDK

### Setup
Add IBM Cloud VPC on Classic API key to the config.
Config file is set up for integration environment.

### Testing
Tests use testing package from Golang. Get the overview of this package [here](https://golang.org/pkg/testing/).

Our test suite is divided into multiple tests that can run individually.

`TestVPCAccessControlLists`

`TestVPCLoadBalancers`

`TestVPCPublicGateways`

`TestVPCResources`

`TestVPCSecurityGroups`

`TestVPCVpn`

##### Run the test
Golang test usually run for 10 mins. If test run for more than 10 mins, timeout is set.

Pre-req
1. Make sure to add URL, IAMURL and APIKEY for your account in config.json file.
2. Make sure you are not hitting account limits. Usually 5 VPCs per account is the limit.
3. There is atleast one of the following resources -  ACL, VPC, Subnet, Instance already exist your account.

Run entire test suite
```bash
go test -run TestVPC -v -timeout 40m
```

Run individual test suite
```bash
go test -run TestVPCResources -v -timeout 20m
```

Additional flags

`-detailed `
Use this flag to view detailed responses from API.
```bash
go test -run TestVPCResources -v -detailed
```

`-testCount`
Use this flag to count number of APIs called.
```bash
go test -run TestVPCResources -v -testCount
```


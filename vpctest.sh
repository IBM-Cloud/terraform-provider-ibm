#! /bin/bash

if [ $# -ne 1 ]
then
    echo "Please Provide Resource Name"
    echo "For VPC -> vpc_all, vpc, vpc_routing_table,vpc_routing_table_route supported"
    exit 1
fi

TEST=$1

# Comment: 

if [ $TEST == "all" ]
then
    TF_ACC=1 go test ./ibm/service/vpc -v '-run=TestAccIBMIS' -timeout 700m
fi

if [ $TEST == "vpc" ]
then
    TF_ACC=1 go test ./ibm/service/vpc -v '-run=TestAccIBMISVPC_' -timeout 700m
fi

if [ $TEST == "vpc_routing_table" ]
then
    TF_ACC=1 go test ./ibm/service/vpc -v '-run=TestAccIBMISVPCRoutingTable_' -timeout 700m
fi

if [ $TEST == "vpc_routing_table_route" ]
then
    TF_ACC=1 go test ./ibm/service/vpc -v '-run=TestAccIBMISVPCRoutingTableRoute_' -timeout 700m
fi

if [ $TEST == "vpc_all" ]
then
    TF_ACC=1 go test ./ibm/service/vpc -v '-run=TestAccIBMISVPC' -timeout 700m
fi



# TESTARGS = ""
# make testacc TEST=./ibm/service/vpc TESTARGS='-run=TestAccIBMISVPC_'
	# TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout $(TEST_TIMEOUT)
# APIDOCS=("TestAccIBMISVPC_basic","TestAccIBMISVPC_basic_apm","TestAccIBMISVPC_securityGroups")

# for APIDOC in $APIDOCS; do
#     JSON=$DOCS_DIR/$APIDOC
#     # generateruby
#     # generatejava
#     # generatenode
#     generatego
#     # generatepython
# done
# echo 'Generated SDK for' $APIDOCS


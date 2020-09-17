#!/bin/bash

ibmcloud login -a https://test.cloud.ibm.com -r $region --apikey=$API_KEY

ibmcloud ks init --host https://containers.test.cloud.ibm.com

ibmcloud plugin repo-add stage https://plugins.test.cloud.ibm.com

ibmcloud plugin install container-service -r stage

#!/bin/bash
MAX_ACCEPTABLE_ISSUES_BLOCKER_CRITICAL_COUNT=320
MAX_ACCEPTABLE_VULNERABILITY_BLOCKER_CRITICAL_COUNT=7
export PATH=/tmp/$SQ_VERSION/bin:$PATH

echo "########## GOSEC ##########"
gosec -fmt=sonarqube -out gosec-report.json ./... || true
echo "sonar.externalIssuesReportPaths=gosec-report.json" >> sonar-project.properties

CLEAN_REPO_SLUG=${TRAVIS_REPO_SLUG////:}
GIT_ORG=$(echo ${TRAVIS_REPO_SLUG////:} | cut -f1 -d:)
GIT_REPO=$(echo ${TRAVIS_REPO_SLUG////:} | cut -f2 -d:)

if [ "${TRAVIS_BRANCH}" == "master" ] && [ "${TRAVIS_PULL_REQUEST}" == "false" ]; then
    REPO_VERSION=$(git describe --tags --abbrev=0 ${TRAVIS_BRANCH} 2>/dev/null)
    #managing fatal: No names found, cannot describe anything. 
    if [ $? -eq 128 ]; then
      REPO_VERSION="0.1"
    fi
    sed -i -e "s/__PROJECT_KEY__/${CLEAN_REPO_SLUG}/g" -e "s/__PROJECT_NAME__/${CLEAN_REPO_SLUG}/g" -e "s/__PROJECT_VERSION__/${REPO_VERSION}/g" -e "s/__LOGIN__/${SQ_TOKEN}/g" -e "s/__PASSWORD__//g" sonar-project.properties
    CURRENT_DATE="$(date -u '+%Y-%m-%dT%H:%M:%S')"
    echo "sonar.go.coverage.reportPaths=$GOPATH/src/github.ibm.com/IBM-Cloud/terraform-provider-ibm/coverage.out" >> sonar-project.properties
    echo "########## SONAR ##########"
    sonar-scanner
elif [[ ! "${TRAVIS_PULL_REQUEST}" == "false" ]]; then
    git checkout -b test-branch ${TRAVIS_COMMIT}
    REPO_VERSION="PR${TRAVIS_PULL_REQUEST}"
    sed -i -e "s/__PROJECT_KEY__/${CLEAN_REPO_SLUG}/g" -e "s/__PROJECT_NAME__/${CLEAN_REPO_SLUG}/g" -e "s/__PROJECT_VERSION__/${REPO_VERSION}/g" -e "s/__LOGIN__/${SQ_TOKEN}/g" -e "s/__PASSWORD__//g" sonar-project.properties
    CURRENT_DATE="$(date -u '+%Y-%m-%dT%H:%M:%S')"
    echo "sonar.go.coverage.reportPaths=$GOPATH/src/github.ibm.com/IBM-Cloud/terraform-provider-ibm/coverage.out" >> sonar-project.properties
    echo "########## SONAR ##########"
    sonar-scanner
    git checkout master
    git branch -D test-branch    
fi
#checking server background task
sleep 30
API_CALL_TASK="${SQ_URL}/api/ce/activity?minSubmittedAt=${CURRENT_DATE}%2B0000"
API_CALL_VULN="${SQ_URL}/api/issues/search?componentKeys=${CLEAN_REPO_SLUG}"'&types=VULNERABILITY&statuses=OPEN&severities=BLOCKER,CRITICAL'
SCANNER_RESULT_URL="${SQ_URL}/dashboard?id=${CLEAN_REPO_SLUG}"
API_CALL_ISSUES_BLOCKER_CRITICAL="${SQ_URL}/api/issues/search?componentKeys=${CLEAN_REPO_SLUG}"'&statuses=OPEN&severities=BLOCKER,CRITICAL'
echo "API_CALL_ISSUES_BLOCKER_CRITICAL"
echo ${API_CALL_ISSUES_BLOCKER_CRITICAL}
echo $SCANNER_RESULT_URL
VULNERABILITY_COUNT=$(curl -k -s -u ${SQ_TOKEN}: ${API_CALL_VULN} | jq '.total')
echo "VULNERABILITY_COUNT=${VULNERABILITY_COUNT}"
ISSUES_BLOCKER_CRITICAL_COUNT=$(curl -k -s -u ${SQ_TOKEN}: ${API_CALL_ISSUES_BLOCKER_CRITICAL} | jq '.total')
echo "ISSUES_BLOCKER_CRITICAL_COUNT=${ISSUES_BLOCKER_CRITICAL_COUNT}"
if [ "${TRAVIS_BRANCH}" == "master" ]; then
    if [ ! -z ${GIT_TOKEN} ]; then
        if [ "${TRAVIS_BRANCH}" != "" ] && [ "${TRAVIS_PULL_REQUEST}" == "false" ]; then
            if [ "${VULNERABILITY_COUNT}" -ge ${MAX_ACCEPTABLE_VULNERABILITY_BLOCKER_CRITICAL_COUNT} ]; then
                #Adding comment to PR if vuln are detected
                curl -H "Authorization: token ${GIT_TOKEN}" -X POST -d "{\"body\": \"Commit rejected due to ${VULNERABILITY_COUNT} vulnerabilities detected. Check here the scanner result ${SCANNER_RESULT_URL}\"}" https://github.ibm.com/api/v3/repos/${GIT_ORG}/${GIT_REPO}/commits/${TRAVIS_COMMIT}/comments
            fi
            if [ "${ISSUES_BLOCKER_CRITICAL_COUNT}" -ge ${MAX_ACCEPTABLE_ISSUES_BLOCKER_CRITICAL_COUNT} ]; then
                #Adding comment to PR if blocker and critical issues are detected
                curl -H "Authorization: token ${GIT_TOKEN}" -X POST -d "{\"body\": \"Commit rejected due to ${ISSUES_BLOCKER_CRITICAL_COUNT} blocker and critical issues detected. Check here the scanner result ${SCANNER_RESULT_URL}\"}" https://github.ibm.com/api/v3/repos/${GIT_ORG}/${GIT_REPO}/commits/${TRAVIS_COMMIT}/comments
            fi
        elif [[ ! "${TRAVIS_PULL_REQUEST}" == "false" ]]; then
            if [ "${VULNERABILITY_COUNT}" -ge ${MAX_ACCEPTABLE_VULNERABILITY_BLOCKER_CRITICAL_COUNT} ]; then
                #Adding comment to PR if vuln are detected
                curl -H "Authorization: token ${GIT_TOKEN}" -X POST -d "{\"body\": \"PR rejected due to ${VULNERABILITY_COUNT} vulnerabilities detected. Check here the scanner result ${SCANNER_RESULT_URL}\"}" https://github.ibm.com/api/v3/repos/${GIT_ORG}/${GIT_REPO}/issues/${TRAVIS_PULL_REQUEST}/comments
            fi
            if [ "${ISSUES_BLOCKER_CRITICAL_COUNT}" -ge ${MAX_ACCEPTABLE_ISSUES_BLOCKER_CRITICAL_COUNT} ]; then
                #Adding comment to PR if blocker and critical issues are detected
                curl -H "Authorization: token ${GIT_TOKEN}" -X POST -d "{\"body\": \"Commit rejected due to ${ISSUES_BLOCKER_CRITICAL_COUNT} blocker and critical issues detected. Check here the scanner result ${SCANNER_RESULT_URL}\"}" https://github.ibm.com/api/v3/repos/${GIT_ORG}/${GIT_REPO}/issues/${TRAVIS_PULL_REQUEST}/comments
            fi
        fi
    fi
    if [ "${VULNERABILITY_COUNT}" -ge ${MAX_ACCEPTABLE_VULNERABILITY_BLOCKER_CRITICAL_COUNT} ]; then
        echo "Found ${VULNERABILITY_COUNT} vuln, exiting"
        exit 1
    fi

    if [ "${ISSUES_BLOCKER_CRITICAL_COUNT}" -ge ${MAX_ACCEPTABLE_ISSUES_BLOCKER_CRITICAL_COUNT} ]; then
        echo "Found ${ISSUES_BLOCKER_CRITICAL_COUNT} blocker and critical issues are detected, exiting"
        exit 1
    fi
fi
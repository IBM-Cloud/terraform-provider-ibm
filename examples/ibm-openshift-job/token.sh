#!/usr/bin/env sh
####################
INSTALL_CMD="apk install -q"


  test -f $(which jq) || ${INSTALL_CMD} jq
  eval "$(jq -r '@sh "export SERVER_URL=\(.server_url)"')"
  apikey1=$(ic iam api-key-create OC_TEMP -d "temp api key for OC login" | grep "API Key" | tr -s " " | cut -d" " -f3)
  TOKEN=$(curl -v -u apikey:$apikey1 "${SERVER_URL}/oauth/authorize?client_id=openshift-challenging-client&response_type=token" \
    -skv -H "X-CSRF-Token: xxx" --stderr - |  grep Location: | sed -E 's/(.*)(access_token=)([^&]*)(&.*)/\3/')
  jq -n \
    --arg token "$TOKEN" \
    '{"token":$token}'

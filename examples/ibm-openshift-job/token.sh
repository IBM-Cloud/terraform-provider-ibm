#!/usr/bin/env sh
####################
INSTALL_CMD="apk install -q"


  test -f $(which jq) || ${INSTALL_CMD} jq
  eval "$(jq -r '@sh "export SERVER_URL=\(.server_url)"')"
  TOKEN=$(curl -v -u apikey:<your api_key> "https://c100-e.us-east.containers.cloud.ibm.com:30129/oauth/authorize?client_id=openshift-challenging-client&response_type=token" \
    -skv -H "X-CSRF-Token: xxx" --stderr - |  grep Location: | sed -E 's/(.*)(access_token=)([^&]*)(&.*)/\3/')
  jq -n \
    --arg token "$TOKEN" \
    '{"token":$token}'

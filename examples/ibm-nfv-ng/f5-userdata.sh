#!/bin/bash
# Description:
# Example userdata consisting of simple bash script to 
# deploy a Standalone 1nic BIG-IP w/:
#   * iApp 
#     * HTTPS Virtual 
#     * ClientSSL profile using default certificate/key
#     * w/ WAF policy
# MGMT GUI Port will be on 8443

# generic init utils
function checkStatus() {
        count=1
        sleep 10;
        STATUS=`cat /var/prompt/ps1`;
        while [[ ${STATUS}x != 'Active'x ]]; do
          echo -n '.';
          sleep 5;
          count=$(($count+1));
          STATUS=`cat /var/prompt/ps1`;
          if [[ $count -eq 60 ]]; then
            checkretstatus="restart";
            return;
          fi
        done
        checkretstatus="run";
}
function checkF5Ready {
    sleep 5
    while [[ ! -e '/var/prompt/ps1' ]]; do
      echo -n '.'
      sleep 5
    done

    sleep 5

    STATUS=`cat /var/prompt/ps1`
    while [[ ${STATUS}x != 'NO LICENSE'x ]]; do
      echo -n '.'
      sleep 5
      STATUS=`cat /var/prompt/ps1`
    done

    echo -n ' '

    while [[ ! -e '/var/prompt/cmiSyncStatus' ]]; do
      echo -n '.'
      sleep 5
    done

    STATUS=`cat /var/prompt/cmiSyncStatus`
    while [[ ${STATUS}x != 'Standalone'x ]]; do
      echo -n '.'
      sleep 5
      STATUS=`cat /var/prompt/cmiSyncStatus`
    done
}
function checkStatusnoret {
  sleep 10
  STATUS=`cat /var/prompt/ps1`
  while [[ ${STATUS}x != 'Active'x ]]; do
    echo -n '.'
    sleep 5
    STATUS=`cat /var/prompt/ps1`
  done
}
function networkUp {

    # usage:
    #   networkUp <num_attempts> <url>
    #   networkUp 120
    #   networkUp 120 https://aws.amazon.com

    NETWORK_UP=FALSE

    if [[ "$2" ]]; then
       NETWORK_UP_CMD="curl --output /dev/null --silent --fail -H 'Cache-Control: no-cache' ${2}"
    else
       NETWORK_UP_CMD="curl --output /dev/null --silent --fail --head -H 'Cache-Control: no-cache' https://activate.f5.com/license/index.jsp"
    fi

    for ((i=1;i<=$1;i++)); do
        if ${NETWORK_UP_CMD}; then
            NETWORK_UP=TRUE
            break
        else
            echo "Test network reachability attempt # ${i} failed. Trying again in 5 secs"
            sleep 5
        fi
    done
}
function LicenseBigIP {

    # usage
    # LicenseBigIP <regkey>

    for ((i=1;i<=5;i++)); do
        LICENSE_RETURN=$( tmsh install /sys license registration-key ${1} )
        if [ "${LICENSE_RETURN}" == "New license installed" ]; then
          break
        else
          echo "License attempt # ${i} failed. Trying again in 5 secs"
          sleep 5
        fi
    done

}
FILE=/var/log/onboard.log
if [ ! -e $FILE ]
then
     touch $FILE
     nohup $0 0<&- &>/dev/null &
     exit
fi
exec 1<&-
exec 2<&-
exec 1<>$FILE
exec 2>&1
checkF5Ready
### ONBOARD_CONFIG_VARS
echo "Setting Some Config Vars..."
# Base images should not have default username/passwords
# In untrusted enviornments, passwords should not be passed via user-data
# Only placing admin password examples for testing convenience 
# BIGIP_ADMIN_USERNAME=admin
# BIGIP_ADMIN_PASSWORD=yourpassword
# BIGIP_ROOT_PASSWORD=yourpassword
HOSTNAME=demo.example.com
BIGIP_LICENSE_KEY=BRETR-JALLY-JQCMW-NWRLV-BRSSPTN
POLICY_LEVEL=high 
APP_NAME=demoService1
VS_PORT=443
POOL_DNS=www.f5.com
POOL_MEMBER_PORT=80
MGMT_ADDR=$(tmsh list sys management-ip | awk '/management-ip/ {print $3}')
MGMT_IP=${MGMT_ADDR%/*}
### START ONBOARDING
echo "Starting Onboarding"
# Insert SSH public key if not automatically injected from provider
echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQClW+UyY2eWczwnEGcEtwR/ISURqmdQIpgicgVvUvZTilE5KstuyBXznpxYT3m2H/7uh5g5syAmS7rX8wSsrbtRjkFgWmDIRPaj3Dqlqqq9N+3TI3mUhMPPWuFZxhW2rK7T6OrWUw5cnJstb89OCQjH4ptqzxIV135re3nT1cJx9JZKxBeYM/tMqZHjAmCwBlj8ndbaidg/f4P0cXa3BS8etcuFGoMwnzACNtkpf6/juodedHbOW9mjamdIoOEVawHiuZNry4emxgT8x9KzBnKHAwRKhLMY/JSc+5z7n21JfDdUIa78Vv3yM3LIaZmpbBPQ7tpJpt4SmYfbhWIUXXXXXX my-ssh-key" >> /home/admin/.ssh/authorized_keys
# tmsh modify auth user admin shell bash password "${BIGIP_ADMIN_PASSWORD}"
tmsh modify sys ntp timezone UTC
tmsh modify sys ntp servers add { 0.pool.ntp.org 1.pool.ntp.org }
tmsh modify sys dns name-servers add { 8.8.8.8 }
tmsh modify sys global-settings gui-setup disabled
tmsh modify sys global-settings hostname "${HOSTNAME}"
tmsh save /sys config
tmsh modify sys httpd ssl-port 8443
tmsh modify net self-allow defaults add { tcp:8443 tcp:6123 tcp:6124 tcp:6125 tcp:6126 tcp:6127 tcp:6128 }
tmsh modify net self-allow defaults delete { tcp:443 }
tmsh mv cm device bigip1 "${HOSTNAME}"
tmsh save /sys config
echo 'start install byol license'
LicenseBigIP ${BIGIP_LICENSE_KEY}
checkStatusnoret
sleep 20 
tmsh save /sys config
echo 'provisioning asm'
tmsh modify /sys provision asm level nominal
checkretstatus='stop'
while [[ $checkretstatus != "run" ]]; do
     checkStatus
     if [[ $checkretstatus == "restart" ]]; then
         echo restarting
         tmsh modify /sys provision asm level none
         checkStatusnoret
         checkretstatus='stop'
         tmsh modify /sys provision asm level nominal
     fi
done
echo 'done provisioning asm'
# START CUSTOM HIGH LEVEL CONFIG HERE #####
mkdir -p /config/cloud/
curl --silent --fail --retry 20 -o /config/cloud/f5.http.v1.2.0rc4.tmpl https://raw.githubusercontent.com/f5devcentral/f5-cloud-init-examples/master/files/f5.http.v1.2.0rc4.tmpl
curl --silent --fail --retry 20 -o /config/cloud/asm-policy-linux.tar.gz https://raw.githubusercontent.com/f5devcentral/f5-cloud-init-examples/master/files/asm-policy-linux.tar.gz
tar xvzf /config/cloud/asm-policy-linux.tar.gz -C /config/cloud/
tmsh create ltm node ${APP_NAME} fqdn { name ${POOL_DNS} }
# BEGIN CUSTOMIZE:  Policy Name/Policy URL, etc.
# tmsh modify asm policy names below (ex. /Common/linux-${POLICY_LEVEL}) to match policy name in the xml file
tmsh load sys application template /config/cloud/f5.http.v1.2.0rc4.tmpl
tmsh load asm policy file /config/cloud/asm-policy-linux-${POLICY_LEVEL}.xml
tmsh modify asm policy /Common/linux-${POLICY_LEVEL} active
tmsh create ltm policy app-ltm-policy strategy first-match legacy
tmsh modify ltm policy app-ltm-policy controls add { asm }
tmsh modify ltm policy app-ltm-policy rules add { associate-asm-policy { actions replace-all-with { 0 { asm request enable policy /Common/linux-${POLICY_LEVEL} } } } }
tmsh create sys application service ${APP_NAME} { template f5.http.v1.2.0rc4 tables add { pool__members { column-names { addr port connection_limit } rows {{ row { ${APP_NAME} ${POOL_MEMBER_PORT} 0 }}}}} variables add { asm__use_asm { value app-ltm-policy } pool__addr { value 0.0.0.0 } pool__mask { value 0.0.0.0 } pool__port { value ${VS_PORT} } pool__port_secure { value ${VS_PORT} } ssl__cert { value /Common/default.crt } ssl__key { value /Common/default.key } ssl__mode { value client_ssl } ssl_encryption_questions__advanced { value no } ssl_encryption_questions__help { value hide } monitor__http_version { value http10 } }}

tmsh save /sys config
date

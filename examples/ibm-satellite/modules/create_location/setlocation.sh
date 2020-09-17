
#!/bin/bash

# LOCATION="testloc7m"
# ZONE="dal10"
# COS_KEY="ce80781bec96a472e5b2102ad138edcc389e7fd9eadd236"
# COS_KEY_ID="d28e72294ded4ac093912d4a6ef8af53"
# LABEL="aa=aa"
ibmcloud sat location create --managed-from $ZONE --name $LOCATION
status='prov'
echo $status
while [ "$status" != "action" ]
do
   if [[ $(ibmcloud sat location get --location $LOCATION | grep State:) == *"action required"* ]]; then
    echo provisioning
    status="action"
  fi
done

path_out=`ibmcloud sat host attach --location $LOCATION -l $LABEL`
path=$(echo $path_out| cut -d' ' -f 21)
echo path= $path
cp $path ../../scripts/addhost.sh



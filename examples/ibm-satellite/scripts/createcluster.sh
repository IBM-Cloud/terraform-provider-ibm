
out=`ibmcloud sat location get --location $location | grep ID`
location_id=$(echo $out| cut -d' ' -f 2)
echo $location_id

status="action_required"
while [ "$status" != "normal" ]
do
    echo *************hosts not ready*****************
    sleep 10
   if [[ $(ibmcloud sat location get --location $location | grep State:) == *"normal"* ]]; then
    echo location $location is normal
    status="normal"
  fi
done

ibmcloud ks cluster create satellite --name $cluster_name --location $location_id --version 4.3_openshift
state="deploying"
while [ "$status" != "warning" ]
do
    echo *************cluster not ready*****************
    sleep 10
   if [[ $(ibmcloud ks cluster get --cluster $cluster_name | grep State:) == *"warning"* ]]; then
    echo location $cluster_name is warning
    status="warning"
  fi
done

echo ****************done
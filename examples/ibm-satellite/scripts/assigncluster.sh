
# host_out=`ibmcloud sat host ls --location kavya | grep kavyavm3`
# HOST_ID=$(echo $host_out| cut -d' ' -f 2)


# ibmcloud sat host assign --cluster bsbtui220ta839u4723g --location kavya --host bsbhnhl20jusnntfnh2g --zone "us-south-1"

host_out=`ibmcloud sat host ls --location $location | grep $hostname`
HOST_ID=$(echo $host_out| cut -d' ' -f 2)


echo hostout= $host_out
echo hostid= $HOST_ID
echo hostname= $hostname
echo location= $location

echo cluster= $cluster_name
ibmcloud sat host assign --cluster $cluster_name --location $location --host $HOST_ID --zone $zone

status='Not Ready'
echo $status
while [ "$status" != "Ready" ]
do
   if [[ $(ibmcloud sat host ls --location $location | grep $hostname) == *"Ready"* ]]; then
    echo host $hostname Ready
    status="Ready"
  fi
    echo *************hosts Not ready*****************
    sleep 10
done

state='warning'
echo $state
while [ "$state" != "normal" ]
do
   if [[ $(ibmcloud ks cluster get --cluster $cluster_name | grep State:) == *"normal"* ]]; then
    echo location $cluster_name is in normal state
    state="normal"
  fi
    echo *************cluster is normal*****************
    sleep 10
done


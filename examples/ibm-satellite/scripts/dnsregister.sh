#!/bin/bash


status="action_required"
while [ "$status" != "normal" ]
do
    echo *************location not normal*****************
    sleep 10
   if [[ $(ibmcloud sat location get --location $location | grep State:) == *"normal"* ]]; then
    echo location $location is normal
    status="normal"
  fi
done

out=`ibmcloud sat location get --location $location | grep ID`
location_id=$(echo $out| cut -d' ' -f 2)
echo $location_id
ibmcloud sat location dns register --location $location_id --ip $ip0 --ip $ip1 --ip $ip2
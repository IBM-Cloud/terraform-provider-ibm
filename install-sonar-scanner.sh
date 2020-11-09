#!/bin/bash
wgetbase="wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/"
$wgetbase

wgetprefix="sonar-scanner-"
wgetpostfix="-linux"

declare -a linearray
declare -a versionarray
re3="linux\.zip<"
re4="([0-9]\.[0-9]\.[0-9]\.[0-9]{3,4})"
i=0; j=0

while read -r row; do
   if [[ $row =~ $re3 ]]; then
      linearray[$i]=$row
      ((i++))
   fi
done < index.html

for str in "${linearray[@]}"; do
   if [[ $str =~ $re4 ]]; then
      versionarray[$j]="${BASH_REMATCH[1]}"
      ((j++))
   fi
done

IFS=$'\n'
arrAsc=($(for l in ${versionarray[@]}; do echo $l; done | sort -r))
unset IFS

rm ./index.html

wgetversion=${arrAsc[0]}
wgetcmd=${wgetbase}${wgetprefix}cli-${wgetversion}${wgetpostfix}.zip
$wgetcmd

unzip ${wgetprefix}cli-${wgetversion}${wgetpostfix}.zip -d /tmp 
rm -f ${wgetprefix}cli-${wgetversion}${wgetpostfix}.zip
export SQ_VERSION=${wgetprefix}${wgetversion}${wgetpostfix}
echo "sonar.host.url=${1}" >> /tmp/$SQ_VERSION/conf/sonar-scanner.properties

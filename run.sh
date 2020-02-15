#!/usr/bin/env bash
#curpath=$(pwd)
#echo ${curpath}
#
#cd ${curpath} && cd ./gateserver8501/
#./GateServer
#
#cd ${curpath} && cd ./gateserver8502/
#./GateServer

curpath=$(pwd)

echo ${curpath}

for i in {1..12}
do
    if [[ $i -lt 10 ]]; then
       cd ${curpath} && cd ./gateserver_850${i}
       ./GateServer
    else
       cd ${curpath} && cd ./gateserver_85${i}
       ./GateServer
    fi
done
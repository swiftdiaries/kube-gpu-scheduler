#!/bin/bash
userid="user1 user2";
jobid="alexnet-gpu-job alexnet-cpu-job vgg-gpu-job vgg-cpu-job";
echo $userid
isNamespace=false #Set to true for first time
if $isNamespace
then	
	for user in $userid; do
		kubectl create -f ./namespaces/$user.yaml
	done	
	isNamespace=false
fi
for user in $userid; do
	for job in $jobid; do
		kubectl --namespace="$user" create -f ./"jobs"/$job.yaml
		randid=$(python -S -c "import random; print random.randrange(20,90)")
		sleep $randid
	done	
done    

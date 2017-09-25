#!/bin/bash
userid="user1 user2";
jobid="alexnet-gpu-job alexnet-cpu-job vgg-gpu-job vgg-cpu-job";
randid=$(python -S -c "import random; print random.randrange(1,3)")
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
	kubectl --namespace="$user" delete pods --all
done  
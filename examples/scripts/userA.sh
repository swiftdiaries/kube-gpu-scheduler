#!/bin/bash
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user1" create -f ./"jobsA"/vgg-gpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user1" create -f ./"jobsA"/vgg-cpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user1" create -f ./"jobsA"/alexnet-gpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user1" create -f ./"jobsA"/alexnet-cpu-job.yaml

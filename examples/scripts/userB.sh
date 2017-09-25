#!/bin/bash
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user2" create -f ./"jobsB"/alexnet-gpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user2" create -f ./"jobsB"/alexnet-cpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user2" create -f ./"jobsB"/vgg-gpu-job.yaml
sleep $randid
randid=$(python -S -c "import random; print random.randrange(20,90)")
kubectl --namespace="user2" create -f ./"jobsB"/vgg-cpu-job.yaml

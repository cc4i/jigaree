# Jigaree

## Description

Jigaree is microservices based application and running on Kubernetes, as an example to demostrate some technical details for development lifecycle, toolings, CI/CD, day 2 operation, and lesson & learn. 

## Architecture of Jigaree


## CI/CD

### Cloud Build & Cloud Deploy

```bash

git clone git@github.com:cc4i/jigaree.git
cd jigaree
gcloud builds submit --config=cloudbuild.yaml --async

```
### Tekton & ArgoCD

## Kubernetes

### GKE

```bash

# Project ID to host your cluster
export PROJECT_ID=play-with-anthos-340801
# Which region to place your cluster
export LOCATION=asia-east2-b
# Cluster name of your GKE
export CLUSTER=jigaree-k8s-cluster

gcloud container --project ${PROJECT_ID} clusters create ${CLUSTER} \
    --zone ${LOCATION} \
    --no-enable-basic-auth \
    --machine-type "n2d-standard-2" \
    --scopes "https://www.googleapis.com/auth/cloud-platform" \
    --num-nodes "2" \
    --enable-ip-alias \
    --enable-dataplane-v2 \
    --workload-pool "${PROJECT_ID}.svc.id.goog" \
    --node-locations ${LOCATION}


```

### GKE

### K83

## Service Mesh

### Istio/Anthos Service Mesh

### Linkerd


## Migrate onto Cloud Run






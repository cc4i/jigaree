# Jigaree

## Description

Jigaree is microservices based application and running on Kubernetes, as an example to demostrate some technical details for development lifecycle, toolings, CI/CD, day 2 operation, and lesson & learn. 

## Architecture of Jigaree


## CI/CD

### Cloud Build & Cloud Deploy

```bash
source ../hack/example-env.sh

# Create building pipeline
gcloud builds submit --config=cloudbuild.yaml --async

# Create deployment pipeline
gcloud deploy apply --file=clouddeploy.yaml --region=${REGION} --project=${PROJECT_ID} --async

# Create trigger for CI/CD
gcloud beta builds triggers create cloud-source-repositories \
    --repo=REPO_NAME \
    --branch-pattern=BRANCH_PATTERN \ # or --tag-pattern=TAG_PATTERN
    --build-config=BUILD_CONFIG_FILE \
    --service-account=SERVICE_ACCOUNT \
    --require-approval

```
### Tekton & ArgoCD

## Kubernetes

### GKE

```bash
source ../hack/example-env.sh

gcloud container --project ${PROJECT_ID} clusters create ${CLUSTER} \
    --zone ${ZONE} \
    --no-enable-basic-auth \
    --machine-type "n2d-standard-2" \
    --scopes "https://www.googleapis.com/auth/cloud-platform" \
    --num-nodes "2" \
    --enable-ip-alias \
    --enable-dataplane-v2 \
    --workload-pool "${PROJECT_ID}.svc.id.goog" \
    --node-locations ${ZONE}

```

### GKE

### K83

## Service Mesh

### Istio/Anthos Service Mesh

### Linkerd


## Migrate onto Cloud Run






# k8s-manifests

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] k8s-manifests`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree k8s-manifests`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init k8s-manifests
kpt live apply k8s-manifests --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/

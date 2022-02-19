# gate

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] gate`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree gate`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init gate
kpt live apply gate --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/

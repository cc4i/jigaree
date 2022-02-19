# website

## Description
sample description

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] website`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree website`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init website
kpt live apply website --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/

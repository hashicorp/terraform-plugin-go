---
name: 🐛 Bug report
about: Let us know about an unexpected error, a crash, or an incorrect behavior.
labels: bug
---

### terraform-plugin-go version
<!---
Inspect your go.mod as below to find the version, and paste the result between
the ``` marks below.

go list -m github.com/hashicorp/terraform-plugin-go/...

If you are not running the latest version of terraform-plugin-go, please try
upgrading because your bug may have already been fixed.
-->

```
insert version here
```

### Relevant provider source code

<!--
Paste any Go code that you believe to be relevant to the bug
When in doubt, a minimal reproduction is best
-->
```go

// insert code here

```

### Terraform Configuration Files
<!--
Paste the relevant parts of your Terraform configuration between the ``` marks below.
-->

```tf
# insert config here
```


### Expected Behavior
<!--
What should have happened?
-->

### Actual Behavior
<!--
What actually happened?
-->

### Steps to Reproduce
<!--
Please list the full steps required to reproduce the issue, for example:
1. `terraform init`
2. `terraform apply`
-->

### References
<!--
Are there any other GitHub issues (open or closed) or Pull Requests that should be linked here? For example:

- #6017
-->

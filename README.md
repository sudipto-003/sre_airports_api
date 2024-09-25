# Airport API

<!-- My thought process and decisions goes here -->

---
_For tasks, checkout [tasks.md](tasks.md)_

1. Provisioning Bucket with IaC: Using terraform IaC tool and AWS cloud for reference. I have created a **bucket** module, this module will create a bucket and attach policies to it. Initially Get, List and Put object actions will be allowed in the new bucket, for the principals given as input values. I have declared *bucket_policy_allow_principals* variable as type *list(map(any))*, so that the access can be given to multiple principals(service, AWS user, Federated users); using a dynamic block that will go through the list maps/objects and add principals to the same policy statement. Allowed field for *bucket_policy_allow_principals* objects's are:
```
{
    type: string,
    identity: string
}
```

clientgofix is a tool for adjusting `k8s.io/client-go` invocations:
* Inserts a `context.TODO()` argument where required
* Inserts an empty options struct argument where required (like `metav1.UpdateOptions{}`)
* Dereferences existing `*metav1.DeleteOptions` arguments
* Replaces existing nil `*metav1.DeleteOptions` arguments with `metav1.DeleteOptions{}`
* Adds `context` and `metav1` imports where required (deconflicting with existing import and declaration names)

**To install:**

```sh
git clone https://github.com/liggitt/clientgofix.git
cd clientgofix
go install ./
```

**To use:**

```sh
cd /path/to/project
clientgofix ./pkg/to/fix
```

**Options:**

* Multiple packages or a recursive package spec can be provided:
  * `clientgofix ./pkg1 ./pkg2`
  * `clientgofix ./...`

* Files are rewritten in place by default.
  Set `-overwrite=false` to write to peer tmp files.

* If errors are encountered processing a file, writing is skipped by default.
  Set `-write-on-error=true` to write the file even when errors are encountered.

`clientgofix` is a tool for adjusting `k8s.io/client-go` invocations for `k8s.io/client-go` v0.18.0+ versions.
It can be run on a codebase before or after updating the `k8s.io/client-go` dependency.

See https://git.k8s.io/enhancements/keps/sig-api-machinery/20200123-client-go-ctx.md for more context.

:warning: This tool is under active development and rewrites source files in place by default.
It is strongly recommended to run this on a version-controlled source tree with all unrelated work committed.

Automatic adjustments:
* Inserts a `context.TODO()` argument where required
* Inserts an empty options struct argument where required (like `metav1.UpdateOptions{}`)
* Dereferences existing `*metav1.DeleteOptions` arguments
* Replaces existing nil `*metav1.DeleteOptions` arguments with `metav1.DeleteOptions{}`
* Adds `context` and `metav1` imports where required (deconflicting with existing import and declaration names)

Manual adjustments required:
* Context parameters passed to `rest.Request#Context` must be manually moved to the request `Do()` invocation.

**To install:**

```sh
git clone https://github.com/kubernetes-sigs/clientgofix.git
cd clientgofix
make install
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

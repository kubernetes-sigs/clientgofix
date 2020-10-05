/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"sigs.k8s.io/clientgofix/pkg"
)

func main() {
	version := false
	flag.BoolVar(&version, "version", version, "display version information")

	o := pkg.DefaultFixOptions()
	flag.BoolVar(&o.Overwrite, "overwrite", o.Overwrite, "overwrite files in place (defaults to true; when false, results are written to peer tmp files)")
	flag.BoolVar(&o.WriteOnError, "write-on-error", o.WriteOnError, "write files even when errors are encountered (defaults to false)")
	flag.StringVar(&o.CustomClientset, "custom-clientset", o.CustomClientset, "package path of generated client sets that need to be transformed, multiple paths separated by commas, eg: github.com/xx/xxrepo/pkg/clientset/versioned/typed/")

	flag.Parse()

	if version {
		fmt.Printf("clientgofix version %s\n", pkg.Version)
		os.Exit(0)
	}

	o.Packages = flag.CommandLine.Args()
	if len(o.Packages) == 0 {
		fmt.Println("Usage: clientgofix [-overwrite=false] [-write-on-error=true] [packages]")
		fmt.Println()
		fmt.Println("Example: clientgofix ./...")
		flag.CommandLine.Usage()
		os.Exit(1)
	}
	if err := o.Complete(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := o.Validate(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := o.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

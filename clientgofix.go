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

	"github.com/liggitt/clientgofix/pkg"
)

func main() {
	o := pkg.DefaultFixOptions()
	flag.StringVar(&o.Dir, "dir", o.Dir, "run in the specified directory (defaults to current directory)")
	flag.BoolVar(&o.Verbose, "verbose", o.Verbose, "verbose messages")
	flag.BoolVar(&o.WriteOnError, "write-on-error", o.WriteOnError, "write files even when errors are encountered")
	flag.BoolVar(&o.Overwrite, "overwrite", o.Overwrite, "overwrite files in place (when false, results are written to peer tmp files)")
	flag.Parse()
	o.Packages = flag.CommandLine.Args()
	if len(o.Packages) == 0 {
		fmt.Println("Usage: clientgofix [packages]")
		fmt.Println()
		fmt.Println("Example: clientgofix ./...")
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

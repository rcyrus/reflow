// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package tool

import (
	"context"
	"flag"
	"fmt"
	"runtime"
)

func (c *Cmd) versionCmd(ctx context.Context, args ...string) {
	var (
		flags = flag.NewFlagSet("offers", flag.ExitOnError)
		help  = "Version displays this binary's version (datestamp) and git hash from which it was built."
	)
	c.Parse(flags, args, help, "version")
	if len(args) != 0 {
		flags.Usage()
	}
	c.Println(c.version())
}

func (c *Cmd) version() string {
	if c.Version == "" {
		c.Version = "broken"
	}
	if c.Variant != "" {
		return fmt.Sprintf("%s (%s, %s)", c.Version, c.Variant, runtime.Version())
	}
	return fmt.Sprintf("%s (%s)", c.Version, runtime.Version())
}

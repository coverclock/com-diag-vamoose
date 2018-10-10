/* vi: set ts=4 expandtab shiftwidth=4: */

// Copyright 2018 Digital Aggregates Corporation, Colorado, USA
// Licensed under the terms in LICENSE.txt
// Chip Overclock <coverclock@diag.com>
// https://github.com/coverclock/com-diag-vamoose
//
// ABSTRACT
//
// Prints build information about Vamoose.
//
// USAGE
//
// vintage
//
package main 

import (
    "fmt"
    "github.com/coverclock/com-diag-vamoose/Vamoose/pkg/vamoose"
)

func main() {
    fmt.Println("ARCH=" + vamoose.ARCH)
	fmt.Println("BRANCH=" + vamoose.BRANCH)
	fmt.Println("CONTACT=" + vamoose.CONTACT)
	fmt.Println("COPYRIGHT=" + vamoose.COPYRIGHT)
	fmt.Println("HOMEPAGE=" + vamoose.HOMEPAGE)
	fmt.Println("HOST=" + vamoose.HOST)
	fmt.Println("KERNEL=" + vamoose.KERNEL)
	fmt.Println("LICENSE=" + vamoose.LICENSE)
	fmt.Println("MODIFIED=" + vamoose.MODIFIED)
	fmt.Println("OS=" + vamoose.OS)
	fmt.Println("PLATFORM=" + vamoose.PLATFORM)
	fmt.Println("RELEASE=" + vamoose.RELEASE)
	fmt.Println("REPOSITORY=" + vamoose.REPOSITORY)
	fmt.Println("REVISION=" + vamoose.REVISION)
	fmt.Println("ROOT=" + vamoose.ROOT)
	fmt.Println("TARGET=" + vamoose.TARGET)
	fmt.Println("TITLE=" + vamoose.TITLE)
	fmt.Println("TOOLCHAIN=" + vamoose.TOOLCHAIN)
	fmt.Println("USER=" + vamoose.USER)
	fmt.Println("VINTAGE=" + vamoose.VINTAGE)
}

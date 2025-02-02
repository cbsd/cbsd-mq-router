package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// get VM properties via cbsd Xget command
func vmget(jname string, emulator string, properties string) string {
	var result string
	var getPrefix string
	// todo: rewrite to SQLite3

	switch emulator {
		case "jail": getPrefix="jget"
		case "qemu": getPrefix="qget"
		case "virtualbox": getPrefix="vget"
		case "xen": getPrefix="xget"
		default: getPrefix="bget"
	}

	cmdStr := fmt.Sprintf("/usr/local/bin/cbsd %s jname=%s mode=quiet %s", getPrefix, jname, properties)
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		Infof("vmget cmd.Run() failed (cbsd %s jname=%s mode=quiet %s) with %s\n", getPrefix, jname, properties, err)
		return ""
	}

	result=(string(out))
	fmt.Printf("vmget str: [%s]\n", result)

	return result
}

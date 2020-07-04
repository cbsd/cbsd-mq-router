package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// get bhyve properties via cbsd bget command
func bget(jname string, properties string) string {
	var result string
	// todo: rewrite to SQLite3
	cmdStr := fmt.Sprintf("/usr/local/bin/cbsd bget jname=%s mode=quiet %s", jname, properties)
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		Infof("bget cmd.Run() failed (cbsd bget jname=%s mode=quiet %s) with %s\n", jname, properties, err)
		return ""
	}

	result=(string(out))
	fmt.Printf("bget str: [%s]\n", result)

	return result
}

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// get bhyve properties via cbsd bhyve-dsk command
func bhyvedsk(jname string, properties string) string {
	var result string
	// todo: rewrite to SQLite3
	cmdStr := fmt.Sprintf("/usr/local/bin/cbsd bhyve-dsk mode=get jname=%s %s", jname, properties)
	cmdArgs := strings.Fields(cmdStr)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		Infof("bhyve-dsk cmd.Run() failed (cbsd bhyve-dsk mode=get jname=%s %s) with %s\n", jname, properties, err)
		return ""
	}
	result=(string(out))
	fmt.Printf("bhyve-dsk str: [%s]\n", result)

	return result
}

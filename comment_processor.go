package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
	"syscall"
	"io/ioutil"
	"reflect"
)

func createKeyValuePairs(m map[string]interface{}) (string,string) {
	jname := new(bytes.Buffer)
	b := new(bytes.Buffer)
	var err error
	for key, value := range m {
		switch v := reflect.ValueOf(value); v.Kind() {
			case reflect.String:
				_, err = fmt.Fprintf(b, " %s=\"%s\"", key, value)
				switch key {
					case "jname": fmt.Fprintf(jname,"%s",value)
//					default: fmt.Printf("PAIRS: %s, %s\n",key,value)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				_, err = fmt.Fprintf(b, " %s=\"%d\"", key, value)
			case reflect.Float32, reflect.Float64:
				_, err = fmt.Fprintf(b, " %s=\"%f\"", key, value)
			default:
				_, err = fmt.Fprintf(b, " %s=\"%s\"", key, value)
				//loogging!
//				_, err = fmt.Fprintf(b, " %s=\"%s\"", key, v.Kind(), value)
//				fmt.Printf("unhandled kind %s", v.Kind())
		}

//		_, err := fmt.Fprintf(b, " %s=\"%s\"", key, value)

		if err != nil {
			return "", ""
		}

//		switch key {
//			case "jname": fmt.Fprintf(jname, "%s", value)
//		}

	}
	return b.String(), jname.String()
}


func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func DoProcess(comment *Comment, logdir string) (error, *CbsdTask) {
	var vm_guid string
	var dsk_guid string
	dt := time.Now()

	Infof("broker log did: %s\n", logdir)
	CreateDirIfNotExist(logdir)

	filePath := fmt.Sprintf("%s/%s_%s_%d.txt", logdir, dt.Format(time.RFC3339), comment.Command, comment.JobID)
	commentFile, err := os.Create(filePath)
	if err != nil {
		return err, nil
	}

	defer func() {
		_ = commentFile.Close()
	}()

	fmt.Printf("JobID %d\n", comment.JobID)

	cbsdArgs, jname := createKeyValuePairs(comment.CommandArgs)

	cmdstr := fmt.Sprintf("/usr/local/bin/cbsd %s %s", comment.Command, cbsdArgs)
	_, err = fmt.Fprintf(commentFile, "%s\n", cmdstr)
	if err != nil {
		return err, nil
	}

	cmd := exec.Command("/bin/sh", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	filePath = fmt.Sprintf("%s/%d.txt", logdir, comment.JobID)

	stdoutFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = stdoutFile.Close()
	}()

	cmd.Stdout = stdoutFile

	if err := cmd.Start(); err != nil {
		fmt.Printf("\ncmd.Start: %v\n")
	}

	if err != nil {
		fmt.Printf("\n%v\n", err)
	}

	cmdStatus := 0

	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				cmdStatus = status.ExitStatus()
			}
		} else {
			fmt.Printf("\ncmd.Wait error: %v\n", err)
			cmdStatus = 0
		}
	} else {
		cmdStatus = 0
	}

	cbsdTask := CbsdTask{}

	cbsdTask.ErrCode = cmdStatus
	// progress always 100 for completed/failed command
	cbsdTask.Progress = 100

	if len(jname)>0 {
		vm_guid = bget(jname,"vm_zfs_guid")
		if len(vm_guid) > 0 {
			Infof("GUID found [%s]\n", vm_guid)
		} else {
			vm_guid = "0"
		}
		// get zfs guid for first disk
		dsk_guid = bhyvedsk(jname,"dsk_path=dsk1 dsk_zfs_guid")
		if len(dsk_guid) > 0 {
			Infof("DSK GUID found [%s]\n", dsk_guid)
		} else {
			dsk_guid = "0"
		}
	} else {
		Infof("no GUID detected [%s]\n", vm_guid)
		vm_guid = "0"
		dsk_guid = "0"
	}

	cbsdTask.Guid = string(vm_guid)
	cbsdTask.DskGuid = string(dsk_guid)
	fileLogPath := fmt.Sprintf("%s/%d.txt", logdir, comment.JobID)

	b, err := ioutil.ReadFile(fileLogPath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	cbsdTask.Message = string(b) // convert content to a 'string'

	Infof("CMD: [%s]\n", comment.Command)

	// add extra field for VNC when bstart
	switch comment.Command {
		case "bstart":
			var vnc string
			vnc = bget(jname,"vnc")

			if len(vnc) > 0 {
				Infof("get VNC [%s]\n", vnc)
				cbsdTask.Vnc = vnc
			} else {
				Infof("no VNC detected\n")
			}
	}

	return nil, &cbsdTask
}

package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const whichBin = "which"

// Creds holds the credentials of a MySQL DB Server
type Creds struct {
	User     string
	Password string
	Host     string
}

// Which calls the `which` cmd
func Which(arg string) string {
	return RunCmd(whichBin, arg)
}

// RunCmd runs a comand with args and returns the output
func RunCmd(cmdString string, args string) string {
	cmd := exec.Command(cmdString, args)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return ""
	}
	return strings.TrimSpace(out.String())
}

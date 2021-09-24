package main

import (
        "bytes"
        "fmt"
        "os"
        "os/exec"
	"strings"
)

func main() {
	// Main executin command
	cmd := exec.Command("bash", "-c", "/usr/bin/vmware-rpctool 'info-get guestinfo.ovfEnv'")

        cmdOutput := &bytes.Buffer{}
        cmd.Stdout = cmdOutput
	printCommand(cmd)

        err := cmd.Run()
        printError(err)

        output := cmdOutput.Bytes()
        printOutput(output)

	// Parsing command output and getting network vairables
        guestinfo := GuestInfo{}
        parseOutput(output, &guestinfo)

	fmt.Println(guestinfo)

        // Rendering netplan configuration
	Render(guestinfo.ip, guestinfo.prefix, guestinfo.gateway, guestinfo.dns1, guestinfo.dns2, guestinfo.suffix)

	// Apply network configuration
	netplanApply()

	// Set hostname
	setHostname(guestinfo.hostname)
}

func printCommand(cmd *exec.Cmd) {
        fmt.Printf("[INFO] Command Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
        if err != nil {
                os.Stderr.WriteString(fmt.Sprintf("[ERROR] Command Error: %s\n", err.Error()))
        }
}

func printOutput(outs []byte) {
        if len(outs) > 0 {
                fmt.Printf("[INFO] Command Output: \n%s\n", string(outs))
        }
}

func setHostname(hostname string) {
	main := "/usr/bin/hostnamectl set-hostname"
	command := fmt.Sprintf("%s %s", main, hostname)

	cmd := exec.Command("bash", "-c", command)

        cmdOutput := &bytes.Buffer{}
        cmd.Stdout = cmdOutput
	printCommand(cmd)

        err := cmd.Run()
        printError(err)
}

func netplanApply() {
	cmd := exec.Command("bash", "-c", "/usr/sbin/netplan apply")

        cmdOutput := &bytes.Buffer{}
        cmd.Stdout = cmdOutput
	printCommand(cmd)

        err := cmd.Run()
        printError(err)
}

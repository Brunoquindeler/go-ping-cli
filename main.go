package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	usage = `
Usage: ping [count](Optional, Default: 4) [ip]
Example: ping 4 1.1.1.1

Type "exit" to leave.`

	waitResult = "wait for the result..."

	successMsg = `Successful!
Type "exit" to leave or make a new "ping".`
)

var errOutput = fmt.Sprintf("\nCommand or args not supported by the CLI. \n %s \n=> ", usage)

func main() {
	fmt.Println("Ping CLI - Check ping by IP address")
	fmt.Println("---------")
	fmt.Println(usage)
	fmt.Print("\n=> ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()

		if strings.Compare(userInput, "exit") == 0 {
			os.Exit(0)
		}

		var commandString []string = strings.Split(userInput, " ")

		CallClear()
		fmt.Print(validation(commandString))

	}

}

func validation(commandString []string) string {
	if strings.Compare(commandString[0], "ping") == 0 {

		if len(commandString) == 1 {
			return errOutput
		}

		if len(commandString) == 2 && checkIPAddress(commandString[1]) {
			fmt.Println(waitResult)
			commandOutput := execPINGCommand("4", commandString[1])
			CallClear()
			fmt.Println(commandOutput)
			fmt.Println(successMsg)
			fmt.Print("=> ")
			return ""
		}

		if _, err := strconv.Atoi(commandString[1]); err == nil && checkIPAddress(commandString[2]) && len(commandString) == 3 {
			fmt.Println(waitResult)
			commandOutput := execPINGCommand(commandString[1], commandString[2])
			CallClear()
			fmt.Println(commandOutput)
			fmt.Println(successMsg)
			fmt.Print("=> ")
			return ""
		}
	}

	return errOutput
}

func execPINGCommand(count, host string) string {
	cmdWindows, argsWindows := "cmd", []string{"/c", "ping", "-n", count, host}
	cmdLinux, argsLinux := "ping", []string{"-c", count, host}

	var (
		cmd  string
		args []string
	)

	switch runtime.GOOS {
	case "windows":
		cmd, args = cmdWindows, argsWindows
	case "linux":
		cmd, args = cmdLinux, argsLinux
	default:
		panic("unsupported platform")
	}

	if output, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		return "Error: " + err.Error()
	} else {
		return string(output)
	}
}

func checkIPAddress(ip string) bool {
	return net.ParseIP(ip) != nil
}

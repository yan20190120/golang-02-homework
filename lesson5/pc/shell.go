package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	host, _ := os.Hostname()
	prompt := fmt.Sprintf("[icexin@%s]$ ", host)
	r := bufio.NewScanner(os.Stdin)
	//r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		if !r.Scan() {
			break
		}
		line := r.Text()
		// line, _ := r.ReadString('\n')
		// line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		var line_list []string
		var cmd_line, out_file string
		if strings.Contains(line, ">") {
			line_list = strings.Split(line, ">")
			cmd_line = line_list[0]
			out_file = strings.TrimSpace(line_list[1])
		} else {
			cmd_line = line
		}
		args := strings.Fields(cmd_line)
		// ls > a.txt
		//fmt.Println(args)
		cmd := exec.Command(args[0], args[1:]...)
		var out_fd *os.File
		if len(out_file) != 0 {
			out_fd, _ = os.OpenFile(out_file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			cmd.Stdout = out_fd
		} else {
			cmd.Stdout = os.Stdout
		}

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
		if out_fd != nil {
			out_fd.Close()
		}
	}
}

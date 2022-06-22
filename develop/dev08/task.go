package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cvenkman/wbL2/develop/dev08/builtins"
	"github.com/cvenkman/wbL2/develop/dev08/nc"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var colorPurple = "\033[35m"
var colorReset = "\033[0m"

func main() {
	user := os.Getenv("USER")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(colorPurple, user, "-> ", colorReset)

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		run(input)
	}
}

func run(input string) {
	if input == "\n" {
		return
	}
	if err := execInput(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

// execute input command
func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	// input = strings.TrimSpace(input)

	cmds := strings.Split(input, "|")
	for i := range cmds {
		cmds[i] = strings.TrimSpace(cmds[i])
	}

	if strings.Contains(input, "|") {
		return execPipe(cmds)
	}

	// Split the input to separate the command and the arguments
	args := strings.Split(cmds[0], " ")
	err := exeCmd(args)
	if err != nil {
		return err
	}

	return nil
}

func exeCmd(args []string) error {
	var err error

	// Check for built-in commands
	if args[0] == "cd" {
		_, err = builtins.CD(args)
		return err
	} else if args[0] == "echo" {
		_, err = builtins.Echo(args)
		if err != nil {
			return err
		}
	} else if args[0] == "exit" {
		os.Exit(0)
	} else if args[0] == "nc" {
		if len(args) < 3 {
			return errors.New("nc: usage: nc [host] [port]")
		}
		return nc.Netcat(args[1], args[2])
	}

	//Exec executes binary files
	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// // Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()
}

func execPipe(strCommands []string) error {
	var err error
	commands := make([]*exec.Cmd, 0)

	for _, cmd := range strCommands {
		cmdArgs := strings.Split(cmd, " ")
		if cmdArgs[0] == "echo" {
			cmdArgs, err = builtins.Echo(cmdArgs)
			if err != nil {
				return err
			}
		}
		commands = append(commands, exec.Command(cmdArgs[0], cmdArgs[1:]...))
	}

	var output, stderr bytes.Buffer

	for i, cmd := range commands[:len(commands)-1] {
		if commands[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return err
		}
		cmd.Stderr = &stderr
	}

	commands[len(commands)-1].Stdout, commands[len(commands)-1].Stderr = &output, &stderr

	for _, cmd := range commands {

		if err = cmd.Start(); err != nil {
			return err
		}
	}

	for _, cmd := range commands {
		if err = cmd.Wait(); err != nil {
			return err
		}
	}

	if len(output.Bytes()) > 0 {
		fmt.Fprint(os.Stdout, output.String())
	}
	if len(stderr.Bytes()) > 0 {
		fmt.Fprint(os.Stderr, output.String())
	}

	return nil
}

//В Linux запущенный экземпляр программы называется процессом.
// ps для получения списка запущенных в данный момент процессов
//ps выводит четыре столбца информации как минимум для двух процессов, запущенных
//в текущей оболочке, самой оболочки и процессов, запущенных в оболочке при вызове команды.
// PID - Идентификатор процесса
// TTY - Имя управляющего терминала для процесса.
// TIME - Совокупное время ЦП процесса, показанное в минутах и ​​секундах.
// CMD - Имя команды, которая использовалась для запуска процесса.

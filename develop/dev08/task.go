package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

)
//echo $USER | cat -e
//echo "dddd" | cat -e
//ls | pwd
// ls | ct -e
//ls | cat -e

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

func execPipe(strCommands []string) error {
	var err error
	commands := make([]*exec.Cmd, 0)

	for _, cmd := range strCommands {
		cmdArgs := strings.Split(cmd, " ")
		if cmdArgs[0] == "echo" {
			cmdArgs, err = echo(cmdArgs)
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
		fmt.Fprint(os.Stdout, string(output.Bytes()))
	}

	if len(stderr.Bytes()) > 0 {
		fmt.Fprint(os.Stderr, string(output.Bytes()))
	}

	return nil
}


func exeCmd(args []string) error {
	// Check for built-in commands
	// switch args[0] {
	// case "cd":
	// 	return cd(args)
	// case "echo", "ECHO":
	// 	return echo(args, false)
	// }

	var cmd *exec.Cmd
	var err error

	if args[0] == "cd" {
		_, err := cd(args)
		return err
	} else if args[0] == "echo" {
		args, err = echo(args)
		if err != nil {
			return err
		}
	}

	//Exec executes binary files
	// Pass the program and the arguments separately
	cmd = exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	err = cmd.Run()
	return err
}


func echo(args []string) ([]string, error) {
	quotesNum := 0
	quotesNum2 := 0
	// count quotes
	for _, arg := range args {
		quotesNum += strings.Count(arg, "'")
		quotesNum2 += strings.Count(arg, "\"")
	}
	if quotesNum % 2 != 0 || quotesNum2 % 2 != 0 {
		return nil, errors.New("unclosed quote")
	}

	// ENV
	for i, arg := range args {
		if len(arg) > 1 && (arg[0] == '$' || arg[:2] == "\"$") {
			arg = strings.Trim(arg, "\"")
			arg = strings.TrimPrefix(arg, "$")

			args[i] = os.Getenv(arg)
			if args[i] == "" {
				args = append(args[:i], args[i+1:]...)
			}
		}
	}

	// remove quotes
	for i := range args {
		args[i] = strings.Replace(args[i], "'", "", -1)
		args[i] = strings.Replace(args[i], "\"", "", -1)
	}
	// cmd := exec.Command(args[0], args[1:]...)
	// cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout
	// return cmd.Run()
	return args, nil
}

// return pwd and error
func cd(args []string) (string, error) {
	prevPed, err := os.Getwd()
	if err != nil {
		return prevPed, errors.New("pwd error")
	}
	
	var moveTo string
	if len(args) < 2 {
		// 'cd' to home dir
		moveTo = "~"
	} else {
		moveTo = args[1]
	}

	switch moveTo {
	// to old path
	case "-":
		moveTo = os.Getenv("OLDPWD")
		if moveTo == "" {
			return prevPed, errors.New("cd: OLDPWD not set") // check in bash
		}
	// to home directory
	case "~":
		moveTo = os.Getenv("HOME")
		if moveTo == "" {
			return prevPed, errors.New("cd: HOME not set")
		}
	}

	// remove last char if it's slash
	if moveTo[len(moveTo) - 1] == '/' {
		moveTo = moveTo[:len(moveTo) - 1]
	}

	// don't move .. if current dir is /
	if os.Getenv("PWD") == "/" && strings.Contains(moveTo, "..") {
		return prevPed, nil
	}

	// Change the directory and return the error
	err = os.Chdir(moveTo)
	if err != nil {
		return prevPed, err
	}

	// change pwd and oldpwd
	pwd, err := os.Getwd()
	if err != nil {
		return prevPed, errors.New("pwd error")
	}
	os.Setenv("PWD", pwd)
	os.Setenv("OLDPWD", prevPed)

	return pwd, nil
}

func changePWDInENV(dir string) {
	// os.Chdir(args[1]) не меняет в env pwd на новую
	pwd := os.Getenv("PWD")

	if strings.Contains(dir, "..") {
		slashNum := strings.Count(dir, "/")

		for ; slashNum > -1; slashNum-- {
			i := strings.LastIndex(pwd, "/")
			
			pwd = pwd[:i]
		}

	} else if dir != "." {
		pwd += "/" + dir
	}

	// fmt.Println(pwd)
	if pwd == "" {
		pwd = "/"
	}
	os.Setenv("PWD", pwd)
}

//В Linux запущенный экземпляр программы называется процессом.
// ps для получения списка запущенных в данный момент процессов
//ps выводит четыре столбца информации как минимум для двух процессов, запущенных
//в текущей оболочке, самой оболочки и процессов, запущенных в оболочке при вызове команды.
// PID - Идентификатор процесса
// TTY - Имя управляющего терминала для процесса.
// TIME - Совокупное время ЦП процесса, показанное в минутах и ​​секундах.
// CMD - Имя команды, которая использовалась для запуска процесса.
func ps(args []string) error {
	// https://stackoverflow.com/questions/9030680/list-of-currently-running-process-in-go
	return nil
}
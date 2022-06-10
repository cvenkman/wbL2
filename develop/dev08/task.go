package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
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


func main() {
	colorPurple := "\033[35m"
	colorReset := "\033[0m"
	user := os.Getenv("USER")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(colorPurple, user, "-> ", colorReset)

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if input == "\n" {
			continue
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// execute input command
func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)

	// Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		return cd(args)
	case "echo", "ECHO":
		return echo(args)
	}

	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()
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

func echo(args []string) error {
	quotesNum := 0
	quotesNum2 := 0
	// count quotes
	for _, arg := range args {
		quotesNum += strings.Count(arg, "'")
		quotesNum2 += strings.Count(arg, "\"")
	}
	if quotesNum % 2 != 0 || quotesNum2 % 2 != 0 {
		return errors.New("незакрытая кавычка")
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
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func cd(args []string) error {
	prevPed, err := os.Getwd()
	if err != nil {
		return errors.New("pwd error")
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
			return errors.New("cd: OLDPWD not set") // check in bash
		}
	// to home directory
	case "~":
		moveTo = os.Getenv("HOME")
		if moveTo == "" {
			return errors.New("cd: HOME not set")
		}
	}

	// remove last char if it's slash
	if moveTo[len(moveTo) - 1] == '/' {
		moveTo = moveTo[:len(moveTo) - 1]
	}

	// don't move .. if current dir is /
	if os.Getenv("PWD") == "/" && strings.Contains(moveTo, "..") {
		return nil
	}

	// Change the directory and return the error
	err = os.Chdir(moveTo)
	if err == nil {
		// change pwd and oldpwd
		pwd, err := os.Getwd()
		if err != nil {
			return errors.New("pwd error")
		}
		os.Setenv("PWD", pwd)
		os.Setenv("OLDPWD", prevPed)
	}
	return err
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
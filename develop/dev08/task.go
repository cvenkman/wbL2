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

	// Split the input to separate the command and the arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		return cd(args)
	case "echo":
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
	var moveTo string

	// 'cd' to home dir
	if len(args) < 2 {
		moveTo = os.Getenv("HOME")
		// empty "HOME" in env
		if moveTo == "" {
			return errors.New("cd: HOME not set")
		}
	} else {
		moveTo = args[1]
	}

	// remove last char if it's slash
	if moveTo[len(moveTo) - 1] == '/' {
		moveTo = moveTo[:len(moveTo) - 1]
	}

	// don't move if current dir is /
	if os.Getenv("PWD") == "/" && strings.Contains(moveTo, "..") {
		return nil
	}

	// Change the directory and return the error
	err := os.Chdir(moveTo)
	if err == nil {
		pwd, err := os.Getwd()
		if err != nil {
			return errors.New("pwd error")
		}
		// changePWDInENV(moveTo)
		os.Setenv("PWD", pwd)
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
package builtins

import (
	"errors"
	"os"
	"strings"
)

func Echo(args []string) ([]string, error) {
	quotesNum := 0
	quotesNum2 := 0
	// count quotes
	for _, arg := range args {
		quotesNum += strings.Count(arg, "'")
		quotesNum2 += strings.Count(arg, "\"")
	}
	if quotesNum%2 != 0 || quotesNum2%2 != 0 {
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
	return args, nil
}

// return pwd and error
func CD(args []string) (string, error) {
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
	if moveTo[len(moveTo)-1] == '/' {
		moveTo = moveTo[:len(moveTo)-1]
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

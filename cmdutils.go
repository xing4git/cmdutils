package cmdutils

import (
	"fmt"
	"os"
	"os/exec"
	// "strconv"
	"time"
)

// Execute some commands with bash. 
// The `cmdstr` param could be a single command such as `ls ~`, 
// or a pipeline command such as `ps aux | grep 'dropbox' | grep -v 'grep'`, 
// a complete script file content is also can be work.
// The `cmdstr` will be written to a temp file, and then execute that file with bash.
// The output of your command will be returned by this function.
func BashExecute(cmdstr string) (ret string, err error) {
	dir := os.TempDir() + pathSeperator() + "goexec"
	dirfile, err := os.Open(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0700)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	defer dirfile.Close()

	filename := fmt.Sprintf("%s%s%d", dir, pathSeperator(), time.Now().Unix())
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write([]byte(cmdstr))
	if err != nil {
		return "", err
	}

	cmd := exec.Command("bash", filename)
	retbytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(retbytes), nil
}

func pathSeperator() string {
	runs := make([]rune, 0, 1)
	runs = append(runs, os.PathSeparator)
	return string(runs)
}

func Run(cmdname string, params ...string) {
	cmd := exec.Command(cmdname, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Errorf("error: %s\n", err.Error())
		os.Exit(2)
	}
}

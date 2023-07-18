package test_tool

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func dateCmd() {
	icmd := exec.Command("date")

	out, err := icmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> date")
	fmt.Println(string(out))

	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing : ", err)
		case *exec.ExitError:
			fmt.Println("command exit rc = ", e.ExitCode())
		default:
			panic(err)
		}
	}
}
func grepCmd() {
	icmd := exec.Command("grep", "hello")

	in, _ := icmd.StdinPipe()
	out, _ := icmd.StdoutPipe()
	icmd.Start()
	in.Write([]byte("hello grep\ngoodbye grep"))
	in.Close()
	gb, _ := io.ReadAll(out)
	icmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(gb))

}
func lsCmd() {
	icmd := exec.Command("bash", "-c", "ls -a -l -h")
	out, err := icmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(out))
}

func Spawningproc() {
	fmt.Println("spawning process test start")
	dateCmd()
	grepCmd()
	lsCmd()
	fmt.Println("spawning process test end")
}

func Execproc() {
	fmt.Println("exec'ing process test start")

	binary, lookerr := exec.LookPath("ls")
	if lookerr != nil {
		panic(lookerr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execerr := syscall.Exec(binary, args, env)
	if execerr != nil {
		panic(execerr)
	}

	fmt.Println("exec'ing process test end")
}

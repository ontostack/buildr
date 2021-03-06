package buildr

import (
	"io"
	"log"
	"os"

	"github.com/codeskyblue/go-sh"
)

func Mkdir(d string) bool {
	if err := os.Mkdir(d, os.ModePerm); err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func Exists(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

func openFile(fname string, flag int, perm os.FileMode) (*os.File, bool) {
	f, err := os.OpenFile(fname, flag, perm)
	if err != nil {
		log.Fatalln(err)
		return nil, false
	}
	return f, true
}

func Cmd(cmd *sh.Session) bool {
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func Check(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return false
	}
	return true
}

func InDir(dir string, f func() bool) bool {
	prev, _ := os.Getwd()
	return Check(os.Chdir(dir)) && f() && Check(os.Chdir(prev))
}

func FillFile(fname string, fill func(io.Writer) bool) bool {
	if f, ok := openFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666); !ok {
		return false
	} else {
		defer f.Close()
		return fill(f)
	}
}

func AppendFile(fname string, fill func(io.Writer) bool) bool {
	if f, ok := openFile(fname, os.O_RDWR|os.O_APPEND, 0666); !ok {
		return false
	} else {
		defer f.Close()
		return fill(f)
	}
}

func GoBuild() bool {
	return Cmd(sh.Command("go", "build"))
}

func GoFmt() bool {
	return Cmd(sh.Command("go", "fmt"))
}

func GoGenerate() bool {
	return Cmd(sh.Command("go", "generate"))
}

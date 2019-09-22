package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/yoki123/ncmdump"
)

type filePointer struct {
	fp    *os.File
	path  string
	after []byte
}

func main() {
	arg := getAllUsefulArg()
	if len(arg) == 0 {
		return
	}
	_, err := os.Getwd()
	fp := make([]filePointer, 0)
	if err != nil {
		return
	}
	for k, v := range arg {
		f, err := os.Open(v)
		fmt.Printf("[%d|%d]...%s...%s\n", k+1, len(arg), ec(err), f.Name())
		if err != nil {
			continue
		} else {
			fp = append(fp, filePointer{
				fp:   f,
				path: v,
			})
		}
	}
	defer func() {
		for _, v := range fp {
			v.fp.Close()
		}
	}()
	//for k, v := range fp {
	//	fmt.Printf("[%d|%d]...%s\n", k+1, len(fp), v.fp.Name())
	//}
	fmt.Println("/////////////Changing dir///////////////")
	for k, v := range fp {
		fp[k].path = changeDir(v.path)
		fmt.Printf("[%d|%d]...%s...OK\n", k+1, len(fp), changeDir(v.path))
	}
	fmt.Println("//////Changing suffix&&Dumping...///////")
	for k, v := range fp {
		fmt.Printf("[%d|%d]...", k+1, len(fp))
		f, b, e := getFormat(v.fp)
		fmt.Print(f, "...", ec(e), "...")
		if err == nil {
			n := rename(v.path, f)
			fp[k].path = n
		} else {
			continue
		}
		fmt.Print(fp[k].path, "...OK\n")
		fp[k].after = b
	}
	fmt.Println("//////////////Write file////////////////")
	for k, v := range fp {
		if v.after == nil {
			fmt.Printf("[%d|%d]...%s...no data,skip!\n", k+1, len(fp), v.fp.Name())
			continue
		}
		fmt.Printf("[%d|%d]...%s...", k+1, len(fp), v.fp.Name())
		err = ioutil.WriteFile(v.path, v.after, 0666)
		if err != nil {
			fmt.Println(err.Error(), "fail")
		}
		fmt.Println("OK")
	}
	return
}

func getAllUsefulArg() []string {
	return os.Args[1:len(os.Args)]
}

func ec(err error) string {
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func getFormat(fp *os.File) (format string, b []byte, err error) {
	f, err := ncmdump.DumpMeta(fp)
	if err != nil {
		return "", nil, err
	}
	b, err = ncmdump.Dump(fp)
	if err != nil {
		return "", nil, err
	}
	return f.Format, b, err
}

func changeDir(path string) string {
	if string(path[0]) == "/" {
		return path
	} else if string(path[1]) == ":" {
		return path
	}
	c, _ := os.Getwd()
	return c + osDirCut() + path
}

func osDirCut() string {
	os := func() string {
		l := strings.Split(runtime.GOOS, "/")
		return l[0]
	}()
	if os == "windows" {
		return "\\"
	}
	return "/"
}

func rename(name string, rp string) string {
	l := strings.Split(name, ".")
	l[len(l)-1] = rp
	return strings.Join(l, ".")
}

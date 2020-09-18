package main

import (
	glg "fmt"
	//"github.com/kpango/glg"
	"github.com/schollz/progressbar/v2"
	"github.com/yoki123/ncmdump"
	"os"
	"runtime"
	"strings"
)

// Progress ...
var Progress int

var TaskCount int

var IfDone chan bool

var ProgressBarChan chan int

var FailedList []string

var TotalResult *ConvertResult

// ConvertStatus ...
type ConvertStatus struct {
	Code     int
	Err      error
	FilePath string
}

// ConvertResult ...
type ConvertResult struct {
	Total   int
	Success int
	Fail    int
}

func init() {
	FailedList = make([]string, 0)
	TotalResult = &ConvertResult{}
	ProgressBarChan = make(chan int)
	Progress = 0
	IfDone = make(chan bool)

}

func main() {
	Args := os.Args[1:]
	if len(Args) == 0 {
		os.Exit(0)
	}
	TaskCount = len(Args)
	pool, _ := MakePool(Args)
	ConvertManager(pool)

	glg.Print("Failed list\n")
	PrintListTree()

	glg.Printf("      ┏Total:   %d\n", TotalResult.Total)
	glg.Printf("Result┣Success: %d\n", TotalResult.Success)
	glg.Printf("      ┗Fail:    %d\n", TotalResult.Fail)

	var tmp string

	_, _ = glg.Scanln(&tmp)

}

// CheckIsNCMType ...
func CheckIsNCMType(filePath string) (isNCMType bool, fp *os.File, err error) {
	fp, err = os.Open(filePath)
	if err != nil {
		return false, nil, err
	}
	isNCMType, err = ncmdump.NCMFile(fp)
	return
}

// MakePool ...
func MakePool(args []string) (pool [][]string, goroutineCount int) {
	cpuCount := runtime.NumCPU()
	if len(args) < cpuCount {
		cpuCount = len(args)
	}
	pool = make([][]string, cpuCount)

	for k := range pool {
		pool[k] = make([]string, 0)
	}

	for k, v := range args {
		pool[k%cpuCount] = append(pool[k%cpuCount], v)
	}
	return pool, cpuCount
}

// ConvertManager ...
func ConvertManager(List [][]string) (result ConvertResult) {
	result = ConvertResult{}
	goroutineCount := len(List)

	resultChanList := make([]chan *ConvertResult, goroutineCount)

	for k := range resultChanList {
		go ConvertProgress(List[k])
	}
	go ProgressBarManager()

	<-IfDone

	return
}

// ConvertProgress ...
func ConvertProgress(convertList []string) {
	TotalResult.Total += len(convertList)
	for _, v := range convertList {
		if convert(v) == true {
			ProgressBarChan <- 1
			TotalResult.Success += 1
		} else {
			ProgressBarChan <- 1
			TotalResult.Fail += 1
		}
	}
}

//goland:noinspection GoNilness
func convert(filePath string) (stat bool) {
	stat = true
	if isNCMType, fp, err := CheckIsNCMType(filePath); err == nil || isNCMType == true {
		defer fp.Close()
		if result, err := ncmdump.Dump(fp); err == nil {
			if writeFile(func() string {
				fileMeta, _ := ncmdump.DumpMeta(fp)
				tmpSlice := strings.Split(fp.Name(), ".")
				if len(tmpSlice) < 2 {
					stat = false
					return ""
				}
				tmp := append(tmpSlice[:len(tmpSlice)-1], "."+fileMeta.Format)
				return strings.Join(tmp, "")
			}(), result) != nil {
				stat = false
			}
		}
	} else {
		stat = false
	}
	if stat == false {
		AddToFailedList(filePath)
	}
	return stat
}

func writeFile(filePath string, content []byte) (err error) {
	_, err = os.Open(filePath)
	if os.IsExist(err) {
		_ = os.Remove(filePath)
	} else {
		if fp, err := os.Create(filePath); err == nil {
			_, err = fp.Write(content)
			return err
		}
	}
	return
}

func ProgressBarManager() {
	Progress = 1
	bar := progressbar.NewOptions(TaskCount, progressbar.OptionSetRenderBlankState(true))
	for {
		if i, notClosed := <-ProgressBarChan; notClosed {
			_ = bar.Add(i)
			if Progress >= TaskCount {
				close(ProgressBarChan)
				glg.Println("")
				break
			}
			Progress += i
		} else {
			break
		}
	}
	IfDone <- true
}

func checkOS() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

func AddToFailedList(name string) {
	tmp := strings.Split(name, checkOS())[len(strings.Split(name, checkOS()))-1]
	FailedList = append(FailedList, tmp)
}

func PrintListTree() {
	length := len(FailedList)
	for k, v := range FailedList {
		copyString := make([]string, len(v))
		if k == length-1 {
			glg.Printf("┗[%d]:%s\n", k+1, v)
			if v[len(v)-4:] == "flac" {
				PrintFailedListWhenFlac(copyString, v)
				copyString[len(v)-5] = ">"
				copyString = append(copyString, "<")
				glg.Printf("  ┃  %s\n", strings.Join(copyString, ""))
				glg.Printf("  ┗ reason:This file is flac format!\n")
			} else if v[len(v)-3:] == "mp3" {
				PrintFailedListWhenMP3(copyString, v)
				copyString[len(v)-4] = ">"
				copyString = append(copyString, "<")
				glg.Printf("  ┃  %s\n", strings.Join(copyString, ""))
				glg.Printf("  ┗ reason:This file is mp3 format!\n")
			}
		} else {
			glg.Printf("┠[%d]:%s\n", k+1, v)
			if v[len(v)-4:] == "flac" {
				PrintFailedListWhenFlac(copyString, v)
				copyString[len(v)-5] = ">"
				copyString = append(copyString, "<")
				glg.Printf("┃ ┃  %s\n", strings.Join(copyString, ""))
				glg.Printf("┃ ┗ reason:This file is flac format!\n")
				glg.Printf("┃\n")
			} else if v[len(v)-3:] == "mp3" {
				PrintFailedListWhenMP3(copyString, v)
				copyString[len(v)-4] = ">"
				copyString = append(copyString, "<")
				glg.Printf("┃ ┃  %s\n", strings.Join(copyString, ""))
				glg.Printf("┃ ┗ reason:This file is mp3 format!\n")
				glg.Printf("┃\n")
			}
		}
	}
}

func PrintFailedListWhenMP3(copyString []string, v string) {
	for k := range v {
		if k >= len(v)-3 {
			copyString[k] = "~"
			continue
		}
		copyString[k] = " "
	}
}

func PrintFailedListWhenFlac(copyString []string, v string) {
	for k := range v {
		if k >= len(v)-4 {
			copyString[k] = "~"
			continue
		}
		copyString[k] = " "
	}
}

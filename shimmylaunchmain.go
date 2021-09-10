package yaenvmanshimmy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func Main() {

	cfg := ReadShimParams()

	args := []string{}
	for _, arg := range SplitLines(cfg) {
		arg = EscapeLine(arg)
		args = append(args, arg)
		fmt.Printf(`"%s" `, strings.ReplaceAll(arg, `"`, `\"`))
	}

	dir, err := os.Getwd()
	if err != nil {
		panic("This is genuinly problematic")
	}

	procAttr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   dir,
	}

	proc, err := os.StartProcess(args[0], args, &procAttr)
	if err != nil {
		fmt.Println("PROC START ERR: ", err)
		panic("")
	}
	proc.Release()
}

func ReadShimParams() string {
	execFileName := path.Base(os.Args[0])
	exeLoc := path.Dir(os.Args[0])
	cfgFname := execFileName + ".yaemshimmy"
	shimcfgFile := path.Join(exeLoc, cfgFname)

	fh, err := os.Open(shimcfgFile)
	if err != nil {
		panic(`File Error "` + cfgFname + `" in "` + exeLoc + "\"\n" + err.Error() + "\n")
	}
	defer fh.Close()
	// info, err := fh.Stat()
	// if err != nil {
	// 	panic(`File Error "` + cfgFname + `" in "` + exeLoc + "\"\n" + err.Error() + "\n")
	// }
	// fmt.Printf("ShimmyFileInfo: %v\n", info)

	fbuf, err := ioutil.ReadAll(fh)
	if err != nil {
		panic(`ReadError "` + cfgFname + `" in "` + exeLoc + "\"\n" + err.Error() + "\n")
	}

	// fmt.Println(fbuf)
	return string(fbuf)
}

func SplitLines(s string) []string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}
	return lines
}

func EscapeLine(line string) (rv string) {
	fmap := map[string](func(string) string){
		"ENV": func(a string) string {
			// log.Println("fmap ENV got param", a)
			return os.Getenv(a)
		},
		"PWD": func(a string) string {
			// log.Println("fmap PWD got param", a)
			if dir, err := os.Getwd(); err == nil {
				return dir
			}
			return ""
		},
		"@": func(a string) string {
			return "@" + a
		},
	}

	rv = line
	if strings.HasPrefix(line, "@") && len(line) > 3 {
		for key, vfunc := range fmap {
			arg := line[len(key)+1:]
			// log.Println("Line:", line)
			// log.Println("ARG:", arg, "key:", key)
			if strings.HasPrefix(line, "@"+key) {
				rv = vfunc(arg)
				// log.Println("Setto:", rv)
				break
			}
		}

	}
	return rv
}

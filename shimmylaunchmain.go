package yaenvmanshimmy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Main() {
	dir, err := os.Getwd()
	if err != nil {
		panic("This is genuinly problematic")
	}

	procAttr := os.ProcAttr{
		Dir: dir,
	}

	//TODO
	args := []string{}

	proc, err := os.StartProcess(args[0], args, &procAttr)
	proc.Release()
}

func ReadShimParams() {
	execFileName := path.Base(os.Args[0])
	exeLoc := path.Dir(os.Args[0])
	cfgFname := execFileName + ".yaemshimmy"
	shimcfgFile := path.Join(exeLoc)
	fh, err := os.Open(shimcfgFile)
	if err != nil {
		panic(`No matching Shimmy File -> "` + cfgFname + `" in "` + exeLoc + "\"\n")
	}
	fbuf, err := ioutil.ReadAll(fh)
	if err != nil {
		panic(`Failed to read -> "` + cfgFname + `" in "` + exeLoc + "\"\n")
	}
	fmt.Println(fbuf)

}

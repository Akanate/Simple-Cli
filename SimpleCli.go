package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/mkideal/cli"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type FileArgs struct {
	CreateFile bool   `cli:"c,createfile"`
	DeleteFile bool   `cli:"d,deleteFile"`
	ReadFile   bool   `cli:"r,ReadFile"`
	WriteFile  bool   `cli:"w,writeFile"`
	FileName   string `cli:"f,Filename"`
}

func main() {
	os.Exit(cli.Run(new(FileArgs), main_func))
}

func main_func(ctx *cli.Context) error {
	args := ctx.Argv().(*FileArgs)
	if args.CreateFile {
		return create(args.FileName)
	} else if args.DeleteFile {
		return delete(args.FileName)
	} else if args.ReadFile {
		return ReadFile(args.FileName)
	}
return nil
}

func create(path string) error {
	var err error
	var file *os.File
	file, err = os.Create(path)
	file.Close()
return err
}

func delete(path string) error {
	var err error
	e := os.Remove(path)
	if e != nil {
		fmt.Println(e)
	}
return err
}

func ReadFile(path string) error {
	var err error
	data, errs := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file")
		fmt.Println(errs)
	}
	fmt.Print(string(data))

return err
}

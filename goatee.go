package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"text/template"
)

func main() {
	tpl := template.New(flag.Arg(0))
	tpl.Funcs(template.FuncMap(map[string]interface{}{
		"string": func(filename string) (string, error) {
			buf, err := ioutil.ReadFile(filename)
			if err != nil {
				return "", err
			}

			return strconv.Quote(string(buf)), nil
		},
		"bytes": func(filename string) (string, error) {
			buf, err := ioutil.ReadFile(filename)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("%#v", buf), nil
		},
	}))
	_, err := tpl.ParseFiles(flag.Arg(0))
	catch(err)

	file, err := os.Create(regexp.MustCompile("\\.got$").ReplaceAllString(flag.Arg(0), "") + ".go")
	catch(err)
	defer file.Close()

	err = tpl.Execute(file, nil)
	catch(err)
}

func init() {
	flag.Parse()
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

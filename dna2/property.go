package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"text/template"
	"bytes"
	"fmt"
)

type Property struct {
	Prop string
	Value string
}

type Config []Property

func (prop Property) ParseConfig(filename string) (Config, error) {

	var config Config

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		// check if the line has = sign
		// and process the line. Ignore the rest.
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the config map
				property := Property{key,value}
				config = append(config, property)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

func (cf Config) FuncMaps() template.FuncMap{
	funcMap := template.FuncMap{}
	for _,v := range cf{
		funcMap[v.Prop]=v.Value
	}
	return funcMap
}


func (cf Config) ValMaps() map[string]string{
	funcMap := map[string]string{}
	for _,v := range cf{
		funcMap[v.Prop]=v.Value
	}
	return funcMap
}

func (cf Config) ParseApp(statement string) string{
	var buff bytes.Buffer
	//t := template.Must(template.New("").Parse(statement))
	t,err1 := template.New("titleTest").Parse(statement)
	t.Option("missingkey=error")
	if err1 != nil {
		fmt.Println(err1)
	}
	err := t.Execute(&buff, cf.ValMaps())

	if err != nil {
		return statement
	}
	rtnValue := buff.String()

	return rtnValue
}
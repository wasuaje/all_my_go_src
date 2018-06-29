package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
	"strings"
	"bytes"
	"fmt"
	"text/template"
)

type ServersJson struct {
	Servers []struct {
		ID             string   `json:"id"`
		Address        string   `json:"address"`
		FunctionalUser string   `json:"functional-user"`
		Location       string   `json:"location"`
		AppGroups      []string `json:"app-groups"`
	} `json:"servers"`
}

type Server struct {
	ID             string
	Address        string
	FunctionalUser string
	Location       string
	AppGroups      []string
}

type ServerList []Server

func (s ServersJson) Parse(filepath string) ServerList {

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &s)
	if err != nil {
		log.Fatal(err)
	}
	servers := make(ServerList,len(s.Servers))
	i := 0

	for _, v := range s.Servers {
		servers[i] = Server{
			Address: v.Address,
			AppGroups:v.AppGroups,
			FunctionalUser:v.FunctionalUser,
			Location:v.Location,
			ID:v.ID,
		}
		i++
	}
	return servers
}

func (s ServerList) GetCurrentHostData (hostname string) ServerList{
	newSrvLst := ServerList{}
	for _,v := range s {
		if v.ID == hostname{
			newSrvLst = append(newSrvLst,v)
		}
	}
	return newSrvLst
}

func (s ServerList) GetEnvVars() map[string]string{
	env_vars  := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		//fmt.Println(pair[0], "---", pair[1])
		env_vars[pair[0]]=pair[1]
	}
	return  env_vars
}

func (s ServerList) ParseValue(statement string) string{
	var buff bytes.Buffer
	t,err1 := template.New("titleTest").Parse(statement)
	t.Option("missingkey=error")
	if err1 != nil {
		fmt.Println("Template parsing error: ", err1)
	}
	err := t.Execute(&buff, s.GetEnvVars())
	if err != nil {
		return statement
	}
	rtnValue := buff.String()
	return rtnValue
}
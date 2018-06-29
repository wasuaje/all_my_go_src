package json_parser

import (
	"io/ioutil"
	"encoding/json"
	"log"
//	"fmt"
//	"reflect"
	"fmt"
)

type App struct {
	Applications []struct {
		Name string `json:"name"`
		StartOrder string `json:"start-order"`
		StopOrder string `json:"stop-order"`
		ChildProcesses []string `json:"child-processes"`
		Start string `json:"start"`
		Stop  string `json:"stop"`
		Check string `json:"check,omitempty"`
		AppGroup string `json:"app-group"`
	}`json:"applications"` }

type Env struct {
	Servers []struct {
		ID             string   `json:"id"`
		Address        string   `json:"address"`
		FunctionalUser string   `json:"functional-user"`
		Location       string   `json:"location"`
		AppGroups      []string `json:"app-groups"`
	} `json:"servers"`
}

type Prop struct {
		value            string
}

type Apps map[string]App

type Envs map[string]Env


func ParseApplications(filepath string) map[string]map[string]interface{} {
	var data App
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	set := make(map[string]map[string]interface{})
	for _, v := range data.Applications {
		val := make(map[string]interface{})
		val["StartOrder"]=v.StartOrder
		val["StopOrder"]=v.StopOrder
		val["ChildProcesses"]=v.ChildProcesses
		val["Start"]=v.Start
		val["Stop"]=v.Stop
		val["Check"]=v.Check
		val["AppGroup"]=v.AppGroup
		set[v.Name]=val
		//fmt.Println(set[v.Name])
	}
	fmt.Println(set)
	return set

}

func ParseEnvironments(filepath string) map[string]map[string]interface{} {
	var data Env
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	set := make(map[string]map[string]interface{})
	for _, v := range data.Servers {
		val := make(map[string]interface{})
		val["Address"]=v.Address
		val["AppGroups"]=v.AppGroups
		val["FunctionalUser"]=v.FunctionalUser
		val["Location"]=v.Location
		set[v.ID]=val
		//fmt.Println(set[v.Name])
	}
	return set

}

func ParseProperties(filepath string) map[string]string {
	var data map[string]string
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	//set := make(map[string]string)
	fmt.Println(data)
	//for _, v := range data {
	//	//val := make(map[string]interface{})
	//	//val["Address"]=v.Address
	//	//val["AppGroups"]=v.AppGroups
	//	//val["FunctionalUser"]=v.FunctionalUser
	//	//val["Location"]=v.Location
	//	//set[v.ID]=val
	//	fmt.Println(set[v])
	//}
	return data

}


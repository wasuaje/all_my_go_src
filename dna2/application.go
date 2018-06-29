package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"strconv"
)

type ApplicationsJson struct {
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

type Application struct {
	Name string
	StartOrder int
	StopOrder int
	ChildProcesses []string
	Start string
	Stop  string
	Check string
	AppGroup string
}

type AppList []Application

type GroupsWithApps struct{
	Appgroupname Application
	Applist      AppList
}

type GroupList []GroupsWithApps

func (app ApplicationsJson) Parse(filepath string) AppList {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &app)
	if err != nil {
		log.Fatal(err)
	}
	apps := make(AppList,len(app.Applications))
	i := 0
	for _, v := range app.Applications {
		//fmt.Println(k,v)
		startorder,_ := strconv.Atoi(v.StartOrder)
		stoporder,_ := strconv.Atoi(v.StopOrder)
		apps[i] = Application{Name:v.Name,
					StartOrder:startorder,
					StopOrder:stoporder,
					Start:v.Start,
					Stop:v.Stop,
					Check:v.Check,
					ChildProcesses:v.ChildProcesses,
					AppGroup:v.AppGroup,
		}
		i++
	}
	return apps
}


//TODO: Application groups field, filter by not name
//Recursive function !!!
func (app AppList) GetAppData(list []string) AppList {
	var filteredList AppList
	for _,k := range list{
		for _,v := range app{
			if k == v.AppGroup  {
				childList := app.GetAppData(v.ChildProcesses)
				filteredList = append(filteredList, childList...)
				if v.Start!="" && v.Stop !=""{
					filteredList = append(filteredList, v)
				}
				break
			}
			//group working as an app
			if k == v.Name && len(v.ChildProcesses) == 0 {
				filteredList = append(filteredList, v)
			}
		}
	}
	return filteredList
}

//To be able to order Appliction object by StartOrder or StopOrder
type ByStartOrder []Application

func (a ByStartOrder ) Len() int           { return len(a) }
func (a ByStartOrder ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStartOrder ) Less(i, j int) bool { return a[i].StartOrder < a[j].StartOrder}

type ByStopOrder []Application

func (a ByStopOrder ) Len() int           { return len(a) }
func (a ByStopOrder ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStopOrder ) Less(i, j int) bool { return a[i].StopOrder < a[j].StopOrder}
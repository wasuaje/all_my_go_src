package main

import (
	"fmt"
	"flag"
	"sort"
)

var currentHostData []Server

func main(){
	// TODO: Accept config location -c filename
	// TODO: Search for default locations and complains when nothing is found
	//interactive := flag.Bool("interactive", false, "Run DNA interactively")
	hostname := flag.String("host", "", "Host to run on")
	available_apps := flag.Bool("a", false, "List all available apps")
	//log_flag := flag.Bool("log_flag", false, "a bool")
	//include_grips := flag.Bool("include_grips", false, "a bool")
	environment_arg := flag.String("env", "", "The environment name to work with")
	app_status := flag.Bool("status", false, "Check the status of the app(s) list")
	//remote_server := flag.String("remote_server", "foo", "a string")
	//release_ems := flag.String("release_ems", "foo", "a string")
	stop := flag.Bool("stop", false, "Stop the app(s) list")
	start := flag.Bool("start", false, "Start the app(s) list")
	//restart := flag.Bool("restart", false, "a bool")
	//release_opts := flag.Bool("stop", false, "a bool")
	//ticket := flag.String("ticket", "", "a bool")
	//jstack := flag.Bool("jstack", false, "a bool")

	flag.Parse()

	//fmt.Println("host:", *environment_arg)
	//fmt.Println("int:", *interactive)
	//fmt.Println("start:", *start)
	//fmt.Println("stop:", *stop)
	//fmt.Println("jstak:", *jstack)
	//fmt.Println("args:", flag.Args())

	//fmt.Println("******:", replaces_env_vars_properties_placeholders("{GOROOT}/stop-application efp-datastore-query-service"))
	//fmt.Println("******:", re
	// places_env_vars_properties_placeholders("{GOROOT}"))

	//TODO: Config locations by flag and default location
	// Set application file to work on
	filename := "application.json"
	Apps := ApplicationsJson{}
	all_apps := Apps.Parse(filename)

	//Set environmet to work on
	//TODO: Get locations from env qa1-mcc ->location=mcc
	env := *environment_arg
	SRV := ServersJson{}
	srvrs := SRV.Parse(env+".json")

	//Properties parsing
	prop := Property{}
	props,_ := prop.ParseConfig("efp.properties")
	//current_apps := map[string]AppList{}

	if *hostname != "" {
		currentHostData = srvrs.GetCurrentHostData(*hostname)
	} else {
		currentHostData = srvrs
	}


	fmt.Println("Working with Enviroment: ", env, "Host:", currentHostData)



	//if *available_apps {
	//	//available_apps_list(*available_apps)
	//	for k,v := range myapps{
	//		fmt.Println(k,v["StartOrder"])
	//	}
	//}

	//Need to prepare data before any actio to save time and imporve performance



	if *available_apps  {

		fmt.Println("*********** Available Aplications ***********")
		sort.Sort(ByStartOrder(all_apps))
		for _,v := range all_apps{
			start_1 := props.ParseApp(v.Start)
			start_2 := srvrs.ParseValue(start_1)
			//fmt.Println("****",start_1,start_2)
			fmt.Println(v.Name,v.StartOrder,start_2)
		}
	}

	if *start || *stop || *app_status {
	//	action_app("start", env, flag.Args())
		HostAndCmds := ServerAppList{}

		var arg_apps []string

		for _,host := range currentHostData{
			if len(flag.Args()) > 0{
				arg_apps=flag.Args()
			}else{
				arg_apps=host.AppGroups
			}
			hst := ServerApps{Srv:host, Appl:all_apps.GetAppData(arg_apps)}
			HostAndCmds = append( HostAndCmds, hst)
		}


		for _,hst1 := range HostAndCmds{
			fmt.Println("*** Working with host:",hst1.Srv.Address)
			sort.Sort(ByStartOrder(hst1.Appl))
			for _,app := range hst1.Appl{
				fmt.Println(app.Name, app.StartOrder)
			}
		}


	}



}

package main

import (
	"flag"
	"fmt"
	"dna/json_parser"
	"strconv"
	"sort"
	"os"
	"strings"
	"dna/config_reader"
	"regexp"
	"time"
	"dna/remote_runner"
	"log"
)

// Candidate to be in a utilities functions package
func run_command(command string)bool{
	return true

}

// Candidate to be in a utilities functions package
func get_env_vars()map[string]string{
	env_vars  := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		//fmt.Println(pair[0], "---", pair[1])
		env_vars[pair[0]]=pair[1]
	}
	return  env_vars
}

// Candidate to be in a utilities functions package
func replaces_env_vars_properties_placeholders(value string)string{
	// TODO: take in account remote env vars
	env_vars := get_env_vars()
	properties, _ := config_reader.ReadConfig("test/efp.properties")
	re := regexp.MustCompile(`\{(.*)\}`)
	var rtn string
	// TODO: can't be both values existing in envs and properties
	if re.MatchString(value){
		clean_value := re.FindStringSubmatch(value)[1]
		if _, ok := properties[clean_value]; ok {
			rtn = re.ReplaceAllString(value, properties[clean_value])
		}
		if _, ok := env_vars[clean_value]; ok {
			rtn = re.ReplaceAllString(value, env_vars[clean_value])
		}
		//fmt.Println(re.FindString(value))
		//fmt.Printf("%q->",clean_value[1])
		//match_prop := properties[clean_value]
		//match_env := env_vars[clean_value]
		//fmt.Println("******",match_env, match_prop)
	}else{
		rtn = value
	}
	if rtn == ""{
		rtn = value
	}

	return rtn
}

// Candidate to be in a utilities functions package
func get_action_order(action string, order interface{}) int {
	//fmt.Println(app_data[the_app]["StartOrder"])
	ord, _ := order.(string)
	ord1, _  := strconv.Atoi(ord)
	rtn := ord1
	return rtn
}

// Candidate to be in a utilities functions package
func get_order_app_data_per_app(action string, app_name string, inter string, app_data  map[string]map[string]interface{}) (int, map[string]string){
	order := get_action_order(action, app_data[app_name][inter])
	start, _ := app_data[app_name]["Start"].(string)
	stop, _ := app_data[app_name]["Stop"].(string)
	check, _ := app_data[app_name]["Check"].(string)
	data := map[string]string{"Start" : start,
		"Stop" : stop,
		"Check" : check,
		"Name" : app_name,}
	return order, data
}

// Candidate to be in a utilities functions package
func get_servers_ordered_apps_by_env(env_data map[string]map[string]interface{}, app_data map[string]map[string]interface{}, app_list []string, action string) map[int]map[string]string{
	app_order := []int{}
	app_final := make(map[int]map[string]string)
	var inter string
	switch act := action; act {
	case "start":
		inter = "StartOrder"
	case "stop":
		inter = "StopOrder"
	default:
		inter = "StartOrder"
	}

	for _, the_app := range app_list {
		// The application exists in app data
		_, ok := app_data[the_app]
		if ok{
			// Type assertion to check if childprocess is a slice of strings
			children, _ := app_data[the_app]["ChildProcesses"].([]string)
			if len(children) > 0 {
				for _, child_app := range children{
					order, data := get_order_app_data_per_app(action, child_app, inter, app_data)
					app_order = append(app_order, order)
					app_final[order] = data
				}
			} else {
				order, data := get_order_app_data_per_app(action, the_app, inter, app_data)
				app_order = append(app_order, order)
				app_final[order] = data
			}
		}
	}
	//fmt.Println(app_order)
	//fmt.Println(app_final)
	return app_final
}

func action_app(action string, env string, app_list []string) {

	// TODO: get the env and build the .json depending on it
	// TODO: take in account env-locations to split and work with
	hosts := []string{"psin5p096.svr.us.jpmchase.net", "psin5p095.svr.us.jpmchase.net", "psin5p094.svr.us.jpmchase.net"}
	//call_ssh("ls -la","d569906", hosts)
	var msg string
	app_order := []int{}
	env_data := json_parser.ParseEnvironments("test/qa1.json")
	app_data := json_parser.ParseApplications("test/application.json")

	if len(app_data) == 0{
		fmt.Println("Couldn't get app data")
		os.Exit(1)
	}
	if len(env_data) == 0{
		fmt.Println("Couldn't get env data")
		os.Exit(1)
	}

	data_to_work_with := get_servers_ordered_apps_by_env(env_data, app_data, app_list, action)
	//fmt.Println(data_to_work_with)
	//fmt.Println(env_data)
	//fmt.Println(app_data)
	// Adding int keys to a slice to order
	for k, _ := range data_to_work_with{
		app_order = append(app_order, k)
	}

	// Sorting the slice
	sort.Ints(app_order)
	//fmt.Println( app_order)

	// Looping the slice to reference the ordered key into the actual data
	for _,v := range app_order{
		//fmt.Println(v)
		switch act := action; act {
		case "start":
			fmt.Println("Starting app: ", data_to_work_with[v]["Name"], v ,"\n")
			dat := data_to_work_with[v]["Start"]
			cmd := replaces_env_vars_properties_placeholders(dat)
			fmt.Println(cmd)
		case "stop":
			fmt.Println("Stopping app: ", data_to_work_with[v]["Name"], v ,"\n")
			dat := data_to_work_with[v]["Stop"]
			cmd := replaces_env_vars_properties_placeholders(dat)
			fmt.Println(cmd)
		case "status":
			log.Println("Checking status for app: ", data_to_work_with[v]["Name"], v )
			dat := data_to_work_with[v]["Check"]
			cmd := replaces_env_vars_properties_placeholders(dat)
			//log.fatal("Running: ", cmd)
			log.Printf("Running: %s", cmd)
			result, err := call_ssh(cmd,"d569906", hosts)
			if err {
				msg = "There was an error: "+ result
			}else{
				if result == ""{
					msg = "Process not found or not running"
				}else{
					msg = "Process Runnin with PID: "+ result
				}
			}
			log.Printf("Result: %v \n", msg)
		}
	}
}

// Candidate to be in a utilities functions package
func available_apps_list(av_app bool){
	app_data := json_parser.ParseApplications("test/application.json")
	for k, _ := range app_data {
		fmt.Println(k)
	}
}

// Candidate to be in a utilities functions package
func call_ssh(cmd string, user string, hosts []string) (result string, error bool) {
	results := make(chan string, 100)
	timeout := time.After(5 * time.Second)

	for _, hostname := range hosts {
		go func(hostname string) {
			results <- remote_runner.ExecuteCmd(cmd, user, hostname)
		}(hostname)
	}

	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			//fmt.Print(res)
			result = res
			error = false
		case <-timeout:
			//fmt.Println("Timed out!")
			result = "Time out!"
			error = true
		}
	}
	return
}

func main() {
	//interactive := flag.Bool("interactive", false, "Run DNA interactively")
	available_apps := flag.Bool("available_apps", false, "List all available apps")
	//log_flag := flag.Bool("log_flag", false, "a bool")
	//include_grips := flag.Bool("include_grips", false, "a bool")
	environment_arg := flag.String("environment", "", "The environment name to work with")
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

	//fmt.Println("env:", *environment_arg)
	//fmt.Println("int:", *interactive)
	//fmt.Println("start:", *start)
	//fmt.Println("stop:", *stop)
	//fmt.Println("jstak:", *jstack)
	//fmt.Println("args:", flag.Args())

	//fmt.Println("******:", replaces_env_vars_properties_placeholders("{GOROOT}/stop-application efp-datastore-query-service"))
	//fmt.Println("******:", replaces_env_vars_properties_placeholders("{GOROOT}"))
	env := *environment_arg
	//LOTS OF VALIDATION COMBINATION
	// TODO. remember handle when in TEST_ENV maybe with an existing en var called like this
	//

	//hosts := []string{"psin5p096.svr.us.jpmchase.net", "psin5p095.svr.us.jpmchase.net", "psin5p094.svr.us.jpmchase.net"}
	//call_ssh("ls -la","d569906", hosts)


	if *start && len(flag.Args()) > 0 {
		action_app("start", env, flag.Args())
	}
	if *stop && len(flag.Args()) > 0 {
		action_app("stop", env, flag.Args())
	}
	if *app_status && len(flag.Args()) > 0 {
		action_app("status", env, flag.Args())
	}
	if *available_apps {
		available_apps_list(*available_apps)
	}
}


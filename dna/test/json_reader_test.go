package json_readertest_test

import ("testing"
	"fmt"
	"dna/json_parser"
	"dna/config_reader"
)

func TestApps(t *testing.T){
	t.Log("Testing Apps json file parsing... (expecting: 'json data')")
	data := json_parser.ParseApplications("application.json")
	fmt.Println("data",data)
	//for k, v := range data {
		//fmt.Printf("Name: %v\n start: %v\n stop: %v\n children: %v\n app-group: %v\n\n",
		//	v.Name,v.Start, v.Stop, v.ChildProcesses, v.AppGroup)
		//fmt.Println(v["Name"],k)
	//}
	//if len(data.Applications) == 0{
	//	t.Error("Couldn't get app data")
	//}
	//fmt.Println(data.Applications[5].Name)

}

func TestEnvs(t *testing.T){
	t.Log("Testing Envs json file parsing... (expecting: 'json data')")
	//data := json_parser.ParseEnvironments("qa1")

	//for _, v := range data.Servers {
	//	fmt.Printf("ID: %v\n Address: %v\n F.User: %v\n Location: %v\n App-groups: %v\n\n",
	//		v.ID,v.Address, v.FunctionalUser, v.Location, v.AppGroups)
	//}
	//if len(data.Servers) == 0{
	//	t.Error("Couldn't get env data")
	//}
	//fmt.Println(data.Servers[5].ID)

}


func TestProps(t *testing.T){
	t.Log("Testing props json file parsing... (expecting: 'json data')")
	data,ok := config_reader.ReadConfig("efp.properties")

	fmt.Println(ok)

	fmt.Println("props",data)
	for k, v := range data {
		//fmt.Printf("Name: %v\n start: %v\n stop: %v\n children: %v\n app-group: %v\n\n",
		//	v.Name,v.Start, v.Stop, v.ChildProcesses, v.AppGroup)
		fmt.Println("Props",k,v )
	}
	//if len(data.Applications) == 0{
	//	t.Error("Couldn't get app data")
	//}
	//fmt.Println(data.Applications[5].Name)

}


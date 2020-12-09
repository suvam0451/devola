package collection

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

type AuthType int
const (
	passwd AuthType = iota
	publickey
	)

type CloudProvider int
const (
	aws CloudProvider = iota
	gcp
	digitalOcean
	other
)

type VPS struct {
	IP string
	AuthenticationType AuthType
	ID string
	Annotation string
	Provider CloudProvider
}

type VpsConfigData []VPS

func getConfigPath() (r string) {
	dir, _ := homedir.Dir()
	r = dir + "/devola_cfg.json"
	return r
}

func GetConfig() (retval VpsConfigData) {
	if jsonFile, err := os.Open(getConfigPath()); err != nil {
		fmt.Println(err)
	} else {
		defer jsonFile.Close()
		byteVal, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteVal, &retval)
	}
	return
}

func SetConfig(setval VpsConfigData) {
	outpath := getConfigPath()

	// Ensure filepath
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		fmt.Print("Generating configuration file: ~/devola_cfg.json")
		os.Create(outpath)
	} else {
		// os.Create(userdir  + "/devola_cfg.json")
	}

	file, _ := json.MarshalIndent(setval, "", " ")
	ioutil.WriteFile(outpath, file, 0644)
}

func Hello(id string) {
	var data VpsConfigData
	fmt.Print("First terminal")
	//cmd := &exec.Cmd{Path: "echo", Args: []string{"Bandoleros"}, Stdout: os.Stdout, Stdin: os.Stdin}
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}
	//out, err := exec.Command("ssh", "root@206.189.128.12").Output()
	//// 206.189.128.12
	//if err != nil {
	//	panic(err)
	//} else {
	//	output := string(out[:])
	//	fmt.Println("Command executed successfully", output)
	//}

	SetConfig(data)
}

// Returns true, if a key with same ID was detected
func EditConfig(data VpsConfigData, key string) int {
	for idx, val := range data {
		if val.ID == key {
			fmt.Println("The specified ID already exists. Override ?")
			return idx
		}
	}
	return -1
}

func AddServer(ip string, id string){
	data := GetConfig()
	extra := VPS{
		IP: ip,
		AuthenticationType: publickey,
		ID: id,
		Annotation: "ok",
		Provider: digitalOcean,
	}
	data = append(data, extra)
	EditConfig(data, id)
	SetConfig(data)
}

// func register(ip string, )
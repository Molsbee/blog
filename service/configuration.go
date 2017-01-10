package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Name     string                `json:"name"`
	Database databaseConfiguration `json:"database"`
}

type databaseConfiguration struct {
	Username   string             `json:"username"`
	Password   string             `json:"password"`
	Connection databaseConnection `json:"connection"`
}

type databaseConnection struct {
	HostName string `json:"hostName"`
	Port     int    `json:"port"`
}

func NewConfiguration(name, username, password, hostName string, port int) *Configuration {
	return &Configuration{
		Name: name,
		Database: databaseConfiguration{
			Username: username,
			Password: password,
			Connection: databaseConnection{
				HostName: hostName,
				Port:     port,
			},
		},
	}
}

// TODO: Encrypt data to hide credentials
// TODO: Handle Name conversion for filename convention
func (cs *Configuration) Save() error {
	json, err := json.MarshalIndent(cs, "", "\t")
	if err != nil {
		return err
	}

	log.Println(cs.Name)
	fileName := fmt.Sprintf(".blog-%s-config", cs.Name)
	log.Printf("creating configuration file - %s", fileName)
	return ioutil.WriteFile(fileName, json, 0644)
}

func GetConfiguration(blogName string) *Configuration {
	fileData, err := ioutil.ReadFile(fmt.Sprintf(".blog-%s-config", blogName))
	if err != nil {
		// TODO: How should I behave
	}

	configuration := Configuration{}
	if err := json.Unmarshal(fileData, &configuration); err != nil {
		/// TODO: HOw should I behave
	}

	return &configuration
}

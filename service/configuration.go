package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// TODO: Encrypt data to hide credentials
// TODO: Handle Name conversion for filename convention
func (cs *Configuration) Save() error {
	configJSON, err := json.MarshalIndent(cs, "", "\t")
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf(".blog-%s-config", cs.Name)
	return ioutil.WriteFile(fileName, configJSON, 0644)
}

func (cs *Configuration) ToString() string {
	return fmt.Sprintf(
		"name: %s\n"+
			"database:\n"+
			"\tusername: %s\n"+
			"\tpassword: %s\n"+
			"\tconnection:\n"+
			"\t\thostName: %s\n"+
			"\t\tport: %d\n", cs.Name, cs.Database.Username, cs.Database.Password, cs.Database.Connection.HostName, cs.Database.Connection.Port)
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

func GetConfiguration(blogName string) (*Configuration, error) {
	fileData, err := ioutil.ReadFile(fmt.Sprintf(".blog-%s-config", blogName))
	if err != nil {
		return nil, err
	}

	configuration := Configuration{}
	if err := json.Unmarshal(fileData, &configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}

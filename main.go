package main

import (
	//"io/ioutil"
	"fmt"
	"io/ioutil"
)

func main() {
	//var configPath = "nonexisting-file.yml"
	//var configPath = "my-file.yml"
	var configPath = "symlink-file.yml"

	configFile, err := ioutil.ReadFile(configPath)

	if err != nil {
		fmt.Errorf("failed to read task config: %s", err)
		return
	}

	str := string(configFile)
	fmt.Println(str)
}

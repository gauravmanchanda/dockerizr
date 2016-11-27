package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func generateFilename(file string) string {
	newFileName := strings.TrimPrefix(file, "templates/")
	newFileName = strings.TrimSuffix(newFileName, ".template")
	curDir, err := os.Getwd()
	if err != nil {
		panic("Unable to get Current Working Dir!!")
	}
	fmt.Println(curDir + "/" + newFileName)
	return curDir + "/" + newFileName
}

func copyAndMoveTemplates(config Config) {
	for _, file := range AssetNames() {
		b, err := Asset(file)
		if err != nil {
			panic(err)
		}
		fileDataStr := string(b[:])
		replaced := replaceWithConfig(fileDataStr, config)

		// write the whole body at once
		err = ioutil.WriteFile(generateFilename(file), []byte(replaced), 0644)
		if err != nil {
			panic(err)
		}
	}
	// read the whole file at once
}

func replaceWithConfig(fileDataStr string, config Config) string {
	replaceConfigs := []string{"AppPort", "ContainerMemory", "Lang", "LangVersion", "QuayOrganization", "ServiceName", "ServiceVisibility", "SlackChannel", "Template"}
	for i := 0; i < len(replaceConfigs); i++ {
		conf := replaceConfigs[i]
		fileDataStr = strings.Replace(fileDataStr, "<"+conf+">", config.getField(conf), -1)
	}
	return fileDataStr
}

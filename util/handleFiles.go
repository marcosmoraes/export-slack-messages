package util

import (
	"encoding/json"
	"export-slack-messages/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readDirectory() []string {
	var fileNames []string
	files, err := ioutil.ReadDir("json-files/")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func writingResultOnFile(filename, displayName string, name string, text string) {
	fileNameWithoutJSONExtension := strings.Split(filename, ".")
	file, err := os.OpenFile("csv-files/"+fileNameWithoutJSONExtension[0]+".csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao executar o arquivo ", err)
	}
	file.WriteString(displayName + ";" + name + ";" + text + "\n")
	file.Close()
}

/*ParsingToCSVFile create a csv files*/
func ParsingToCSVFile() {
	var slackMessage []model.TextMessage
	for _, fileName := range readDirectory() {
		slackMessage := parseToJSONObject(fileName, slackMessage)
		for _, textMessage := range slackMessage {
			if textMessage.UserProfile != nil {
				writingResultOnFile(fileName, textMessage.UserProfile.DisplayName, textMessage.UserProfile.Name, textMessage.Text)
			} else {
				writingResultOnFile(fileName, "", "", textMessage.Text)
			}
		}
	}
}

func parseToJSONObject(fileName string, slackMessage []model.TextMessage) []model.TextMessage {
	file, _ := os.Open("json-files/" + fileName)
	jsonObj, _ := ioutil.ReadAll(file)
	err := json.Unmarshal(jsonObj, &slackMessage)
	if err != nil {
		fmt.Println("erro on unmarshal file: ", err)
	}
	return slackMessage
}

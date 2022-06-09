package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/manifoldco/promptui"
)

type Data struct {
	EncodeBytes  []byte `json:"encode-bytes"`
	DecodeString string `json:"decode-string"`
}

const JSON_FILE = "data.json"

func main() {
	if _, err := os.Stat(JSON_FILE); err != nil {
		emptyCollection := Data{}
		file, _ := json.MarshalIndent(emptyCollection, "", " ")
		_ = ioutil.WriteFile(JSON_FILE, file, 0644)
	}

	for {
		var choices = []string{"Encode (byte to string)", "Decode (string to byte)"}

		prompt := promptui.Select{
			Label: "Menu",
			Items: choices,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		data := GetDataFromJson()

		switch result {
		case choices[0]:
			EncodeToString(data.EncodeBytes)
		case choices[1]:
			DecodeString(data.DecodeString)
		}
	}
}

func EncodeToString(byteArray []byte) {
	sEnc := b64.StdEncoding.EncodeToString(byteArray)
	fmt.Print("\n")
	fmt.Println(sEnc)
	fmt.Print("\n")
}

func DecodeString(text string) {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	fmt.Print("\n")
	fmt.Println(sDec)
	fmt.Print("\n")
}

func GetDataFromJson() Data {
	jsonFile, err := os.Open(JSON_FILE)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data

	json.Unmarshal([]byte(byteValue), &data)

	return data
}

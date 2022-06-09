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

func main() {
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
}

func DecodeString(text string) {
	sDec, _ := b64.StdEncoding.DecodeString(text)
	fmt.Print("\n")
	fmt.Println(sDec)
}

func GetDataFromJson() Data {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data

	json.Unmarshal([]byte(byteValue), &data)

	return data
}

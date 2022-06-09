package main

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/manifoldco/promptui"
)

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

		switch result {
		case choices[0]:
			EncodeToString()
		case choices[1]:
			DecodeString()
		}
	}
}

func EncodeToString() {
	byteArray := []byte{100, 110, 120, 130, 140, 150}
	sEnc := b64.StdEncoding.EncodeToString(byteArray)
	fmt.Print("\n")
	fmt.Println(sEnc)
}

func DecodeString() {
	text := "text"
	sDec, _ := b64.StdEncoding.DecodeString(text)
	fmt.Print("\n")
	fmt.Println(sDec)
}

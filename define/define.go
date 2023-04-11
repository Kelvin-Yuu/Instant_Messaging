package define

import (
	"io/ioutil"
	"log"
)

func ReadLocalPassword() string {
	data, err := ioutil.ReadFile("MyPassword.txt")
	if err != nil {
		log.Println("File reading error", err)
		return ""
	}
	return string(data)
}

// var MailPassword = os.Getenv("MailPassword")
var MailPassword = ReadLocalPassword()

type MessageStruct struct {
	Message      string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}

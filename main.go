package main

import (
	"fmt"

	"github.com/mmzou/geektime-dl/config"
	"github.com/mmzou/geektime-dl/login"
)

func init() {

	err := config.Config.Init()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	login := login.NewLoginClient()
	phone := "13240929571"
	password := "wys434390451"
	result := login.Login(phone, password)
	if !result.IsLoginSuccess() {
		fmt.Println(result.Error.Msg)
		return
	}

	geektime := &config.Geektime{
		User: config.User{
			ID:     result.Data.UID,
			Name:   result.Data.Name,
			Avatar: result.Data.Avatar,
		},
		GCID:         result.Data.GCID,
		GCESS:        result.Data.GCESS,
		Ticket:       result.Data.Ticket,
		ServerID:     result.Data.ServerID,
		CookieString: result.Data.CookieString,
	}

	fmt.Println("geektime", geektime)

	config.Config.Geektimes = append(config.Config.Geektimes, geektime)
	config.Config.AcitveUID = geektime.ID

	err := config.Config.Save()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(geektime)
}

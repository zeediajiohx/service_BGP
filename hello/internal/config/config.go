package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf `json:"rest_._rest_conf"`
	//Name      string
	//Host      string
	//Port      int
	DataBase  DataBase
	NoConfStr string `json:"noconfstr,optional"`
	//Address   string `json:"address,default=127.0.0.1:8080"`
	//Prop      string `json:"prop,optional"`
}
type DataBase struct {
	Url string `json:"url"`
}

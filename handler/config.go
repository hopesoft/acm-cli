package handler

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var AcmVersion string = "v0.0.1(beta)";

type AcmConf struct{
	Endpoint	string
	AccessKey	string
	SecretKey	string
	NamespaceId	string
}

type AcmNamespaceItem struct {
	DataId		string
	Group 		string
	Filename	string
}

// 配置文件结构
type AcmConfigTemplate struct {
	Config 	struct{
		Endpoint	string
		AccessKey	string
		SecretKey	string
	}
	Namespace 	map[string]struct{
		Id		string
		List 	[]AcmNamespaceItem
	}
}
// 选中的配置数据
type AcmActiveConfig struct {
	AcmConf
	List		[]AcmNamespaceItem
}

var AcmCfg = &AcmActiveConfig{}

func SetCfg(cfg string) {
	if _, err := os.Stat(cfg); err != nil {
		log.Fatalln("文件不存在")
	}
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	var c = &AcmConfigTemplate{}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalln(err)
	}
	ns, ok := c.Namespace[AcmEnv["namespace"]]
	if !ok {
		log.Fatalln("Namespace Not Found")
	}
	AcmCfg.Endpoint = c.Config.Endpoint
	AcmCfg.AccessKey = c.Config.AccessKey
	AcmCfg.SecretKey = c.Config.SecretKey
	AcmCfg.NamespaceId = ns.Id
	AcmCfg.List = ns.List
}

var AcmEnv = make(map[string]string)

func SetEnv(str string) {
	strs := strings.Split(str, ",")
	for _, str := range strs {
		val:= strings.Split(str, "=")
		AcmEnv[val[0]] = val[1]
	}
}
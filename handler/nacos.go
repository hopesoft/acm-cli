package handler

import (
	"acm-cli/utils"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type NacosConf struct {
	cli  config_client.IConfigClient
	Conf AcmConf
}

func NewNacosConf() *NacosConf {
	return &NacosConf{
		Conf: AcmCfg.AcmConf,
	}
}

func (nc *NacosConf) CreateClient() (config_client.IConfigClient, error) {
	config := constant.ClientConfig{
		TimeoutMs:            5 * 1000,
		ListenInterval:       30 * 1000,
		NamespaceId:          nc.Conf.NamespaceId,
		Endpoint:             nc.Conf.Endpoint,
		AccessKey:            nc.Conf.AccessKey,
		SecretKey:            nc.Conf.SecretKey,
	}
	return clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": config,
	})
}
// 监听配置
func (nc *NacosConf) ListenConfig(list []AcmNamespaceItem,fun func(data string, index int)) {
	for index, item := range list {
		i := index
		cli,err := nc.CreateClient()
		if err != nil {
			log.Fatalln(err)
		}
		err = cli.ListenConfig(vo.ConfigParam{
			DataId: item.DataId,
			Group:  item.Group,
			OnChange: func(namespace, group, dataId, data string) {
				fun(data,i)
			},
		})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
// 添加配置
func (nc *NacosConf) PublishConfig(dataId, group, content string) (bool, error) {
	if nc.cli == nil {
		cli,err := nc.CreateClient()
		if err != nil {
			log.Fatalln(err)
		}
		nc.cli = cli
	}
	return nc.cli.PublishConfig(vo.ConfigParam{
		DataId:  dataId,
		Group:   group,
		Content: content,
	})
}
// 读取配置
func (nc *NacosConf) GetConfig(dataId, group string) (content string,token string,err error) {
	content, err = nc.cli.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	return content, utils.Md5(content), err
}
// 删除配置
func (nc *NacosConf) DelConfig(dataId, group string) (bool, error) {
	return nc.cli.DeleteConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
}
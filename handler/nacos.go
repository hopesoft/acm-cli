package handler

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

type NacosConf struct {
	cli 		config_client.IConfigClient
	configItem 	AcmConfigItem
}

func NewNacosConf(item AcmConfigItem) (*NacosConf, error) {
	config := constant.ClientConfig{
		TimeoutMs:            5 * 1000,
		ListenInterval:       30 * 1000,
		NamespaceId:          item.NamespaceId,
		Endpoint:             item.Endpoint,
		AccessKey:            item.AccessKey,
		SecretKey:            item.SecretKey,
	}
	cli, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": config,
	})

	return &NacosConf{
		cli:cli,
		configItem:item,
	}, err
}
// 监听配置
func (nc *NacosConf) ListenConfig(list []AcmNamespaceItem,fun func(data string, index int)) {
	for index, item := range list {
		i := index
		err := nc.cli.ListenConfig(vo.ConfigParam{
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
func (nc *NacosConf) PublishConfig(content string) (bool, error) {
	return nc.cli.PublishConfig(vo.ConfigParam{
		DataId:  nc.configItem.DataId,
		Group:   nc.configItem.Group,
		Content: content,
	})
}
// 读取配置
func (nc *NacosConf) GetConfig() (string, error) {
	return nc.cli.GetConfig(vo.ConfigParam{
		DataId: nc.configItem.DataId,
		Group:  nc.configItem.Group,
	})
}
// 删除配置
func (nc *NacosConf) DelConfig() (bool, error) {
	return nc.cli.DeleteConfig(vo.ConfigParam{
		DataId:   nc.configItem.DataId,
		Group:    nc.configItem.Group,
	})
}
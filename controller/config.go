package controller

import (
	"acm-cli/handler"
)

func (c *Controller) Set() {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err)
	}
	_, err = conf.PublishConfig(c.Request.FormValue("content"))
	if err != nil {
		c.Failed(err)
	}
	c.Response("SUCCESS")
}

func (c *Controller) Get() {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err)
	}
	content, err := conf.GetConfig()
	if err != nil {
		c.Failed(err)
	}
	c.Response(content)
}

func (c *Controller) Del()  {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err)
	}
	_, err = conf.DelConfig()
	if err != nil {
		c.Failed(err)
	}
	c.Response("SUCCESS")
}

func (c *Controller) parseConfigItem() handler.AcmConfigItem {
	item := handler.AcmConfigItem{
		DataId:      c.Request.FormValue("data_id"),
		Group:       c.Request.FormValue("group"),
	}
	return handler.StuffConfigItem(item)
}
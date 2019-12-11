package controller

import (
	"acm-cli/handler"
	"errors"
)

func (c *Controller) Set() {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err, 400)
		return
	}
	_, ok := c.Request.Form["content"]
	if !ok {
		c.Failed(errors.New("parameter content cannot be empty"), 400)
		return
	}
	_, token, err := conf.GetConfig()
	if token != c.Request.FormValue("token") {
		c.Failed(errors.New("token expired"), 412)
		return
	}
	content := c.Request.FormValue("content")
	if len(content) <= 0 {
		_, err = conf.DelConfig()
	} else {
		_, err = conf.PublishConfig(content)
	}
	if err != nil {
		c.Failed(err, 400)
		return
	}
	c.Response("SUCCESS")
}

func (c *Controller) Get() {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err, 400)
		return
	}
	content, token, _ := conf.GetConfig()
	c.Writer.Header().Set("Content-Token", token);
	c.Response(content)
}

func (c *Controller) Del()  {
	conf, err := handler.NewNacosConf(c.parseConfigItem())
	if err != nil {
		c.Failed(err, 400)
		return
	}
	_, err = conf.DelConfig()
	if err != nil {
		c.Failed(err, 400)
		return
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
package controller

import (
	"acm-cli/handler"
	"errors"
)

func (c *Controller) Set() {
	_, ok := c.Request.Form["content"]
	if !ok {
		c.Failed(errors.New("parameter content cannot be empty"), 400)
		return
	}
	dataId := c.Request.FormValue("data_id")
	group := c.Request.FormValue("group")
	conf := handler.NewNacosConf()
	_, token, err := conf.GetConfig(dataId, group)
	if token != c.Request.FormValue("token") {
		c.Failed(errors.New("token expired"), 412)
		return
	}
	content := c.Request.FormValue("content")
	if len(content) <= 0 {
		_, err = conf.DelConfig(dataId, group)
	} else {
		_, err = conf.PublishConfig(dataId, group,content)
	}
	if err != nil {
		c.Failed(err, 400)
		return
	}
	c.Response("SUCCESS")
}

func (c *Controller) Get() {
	dataId := c.Request.FormValue("data_id")
	group := c.Request.FormValue("group")
	conf := handler.NewNacosConf()
	content, token, _ := conf.GetConfig(dataId, group)
	c.Writer.Header().Set("Content-Token", token);
	c.Response(content)
}

func (c *Controller) Del()  {
	dataId := c.Request.FormValue("data_id")
	group := c.Request.FormValue("group")
	conf := handler.NewNacosConf()
	_, err := conf.DelConfig(dataId, group)
	if err != nil {
		c.Failed(err, 400)
		return
	}
	c.Response("SUCCESS")
}
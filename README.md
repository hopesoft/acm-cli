# acm-cli
Aliyun Application Configuration Management Client

> ACM 是 Aliyun 提供的面向分布式系统的配置中心。
>
> 凭借配置变更、配置推送、历史版本管理、灰度发布、配置变更审计等配置管理工具，
>
> 更多详情请参照：[应用配置管理 ACM](https://help.aliyun.com/product/59604.html)

## 快速入门
* 下载 [acm-cli](https://github.com/xiaoliang0009/acm-cli/releases)
* 配置 [acm-config.yaml](https://github.com/xiaoliang0009/acm-cli/blob/master/acm-config.yaml)
```yaml
Config:
  Endpoint: acm.aliyun.com:8080
  AccessKey: your access key
  SecretKey: your secret key

Namespace:

  # custom develop namespace
  develop:
    Id: your namespace id
    List:
      - DataId: data id
        Group: group
        Filename: /example/develop/your-sync-config-file-name

  # custom beta namespace
  beta:
    Id: your namespace id
    List:
      - DataId: data id
        Group: group
        Filename: /example/beta/your-sync-config-file-name

  # custom product namespace
  product:
    Id: your namespace id
    List:
      - DataId: data id
        Group: group
        Filename: /example/product/your-sync-config-file-name
```
`Namespace` 对应你在 Aliyun 应用配置管理 所创建的

![namespace](https://flighter-img.oss-cn-hangzhou.aliyuncs.com/%E6%B7%B1%E5%BA%A6%E6%88%AA%E5%9B%BE_%E9%80%89%E6%8B%A9%E5%8C%BA%E5%9F%9F_20191211204404.png)

`DataId` 和 `Group` 则对应命名空间下的配置项

## 命令行参数

```shell
➜ acm acm-cli -h
Aliyun Application Configuration Management Client.

Usage:
  acm-cli [flags]

Flags:
  -c, --config string   设置配置文件 (default "./acm-config.yaml")
  -e, --env string      设置环境变量 (default "namespace=develop,port=10010")
  -h, --help            help for acm-cli
      --restful         开启restful
      --sync            开启自动同步配置 (default true)
```

`--restful=true` 监听端口`10010`, 支持动态设置配置, 实现动态服务注册功能

默认请求参数:

| 参数名  | 类型   | 说明                                            |
| ------- | ------ | ----------------------------------------------- |
| data_id | string | [配置集ID](https://help.aliyun.com/document_detail/59968.html?spm=5176.acm.0.0.56204a9bNqwmrE#h2-url-9) |
| group   | string | [配置分组](https://help.aliyun.com/document_detail/59968.html?spm=5176.acm.0.0.56204a9bNqwmrE#h2-url-10) |

* 读取配置
> GET  `/config`
> Response Header 中返回 `Content-Token`

* 写入配置
> POST  `/config`

| 参数名  | 类型   | 说明                                            |
| ------- | ------ | ----------------------------------------------- |
| content | string | 配置内容                                        |
| token   | string | GET /config Response Header 返回的Content-Token |

* 删除配置

> DELETE `/config`

  
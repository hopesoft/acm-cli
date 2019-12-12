package cmd

import (
  "acm-cli/controller"
  "acm-cli/handler"
  "fmt"
  "github.com/spf13/cobra"
  "io/ioutil"
  "log"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "acm-cli",
  Short: "aliyun acm cli",
  Long: `Aliyun Application Configuration Management Client.

关于ACM: 
    ACM 是面向分布式系统的配置中心。
    凭借配置变更、配置推送、历史版本管理、灰度发布、配置变更审计等配置管理工具，
    ACM 能帮助您集中管理所有应用环境中的配置，降低分布式系统中管理配置的成本.


详情链接: https://help.aliyun.com/product/59604.html`,
  Run: runCmd,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
type Flags struct{
  cfg       string
  env       string
  restful   bool
  sync      bool
}
var flags = &Flags{}

func init() {
  cobra.OnInitialize(func() {
    handler.SetEnv(flags.env)
    handler.SetCfg(flags.cfg)
  })

  rootCmd.PersistentFlags().StringVarP(&flags.cfg, "config", "c", "./acm-config.yaml", "设置配置文件")
  rootCmd.PersistentFlags().StringVarP(&flags.env, "env", "e", "namespace=develop,port=10010", "设置环境变量")
  rootCmd.PersistentFlags().BoolVar(&flags.restful, "restful", false, "开启restful")
  rootCmd.PersistentFlags().BoolVar(&flags.sync, "sync", true, "开启自动同步配置")
}

func runCmd(cmd *cobra.Command, args []string) {
  if flags.sync {
    log.Println("已开启自动同步配置")
    listenConfig()
  }
  if flags.restful {
    log.Println("已开启restful")
    controller.ListenServer()
  }
  log.Println("运行已停止，请检查运行参数")
}
// 监听 acm 配置
func listenConfig() {
  conf := handler.NewNacosConf()
  conf.ListenConfig(handler.AcmCfg.List,
    func(data string, index int) {
      err := ioutil.WriteFile(handler.AcmCfg.List[index].Filename, []byte(data), 0666)
      if err != nil {
        log.Fatal(err)
      }
  })
  // 持续监听配置
  if !flags.restful {
    ch := make(<-chan bool)
    <- ch
  }
}
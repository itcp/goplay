# goplay

#### 介绍
一个信息交流平台，可用于CMS（内容管理平台）、BBS（论坛）、wiki(知识文档管理)使用。

特点：所用框架组件都是极其轻量，前端加载快，后端运行性能好。简单也意味着方便扩展二次开发

#### 软件架构
后端基于golang语言开发，主框架由gin + gorm

前端是 frozenUI + axios

- frozenUI(http://frozenui.github.io/)

- axios(http://www.axios-js.com)


#### 安装教程

1. 请确认你的go环境是配置好的

   

2. 克隆项目代码和下载依赖包

   ```
   git clone ...
   cd goplay
   # 如果是在中国大陆，先指定加速源
   go env -w GO111MODULE=on
   go env -w GOPROXY=https://goproxy.cn,direct
   # 下载安装依赖
   go mod vendor
   ```

   

3. 数据库要准备好，项目配置文件修改

   ```
   cp config.toml.example conf/dev/config.toml
   vi conf/dev/config.toml
   ```

4.  启动运行

    `go run main.go`

#### 功能说明

1. 用户模块: 

   - 注册
   - 登录

2. 主题信息发布

   - 信息发布-图片上传
   - 信息列表页
   - 信息详情页

   


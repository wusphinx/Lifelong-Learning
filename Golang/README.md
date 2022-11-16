# Golang

## 多版本管理 
### [GVM](https://github.com/moovweb/gvm)
使用gvm进行golang的多版本管理的时候，go的源码是从github下载的，配置在`~/.gvm/scripts/install`
```
GO_SOURCE_URL=https://github.com/golang/go
```
国内用户为了加快安装，可将其改为
```
GO_SOURCE_URL=https://gitee.com/mirror_golang/go
```
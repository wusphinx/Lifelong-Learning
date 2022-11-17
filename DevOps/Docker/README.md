# Docker

## 不常用
### 从镜像中获取文件
用镜像方式打包是为了标准化构建环境，有时候只需要Linux环境下的可执行文件，一条命令就能解决：
```
$ docker run -v $(PWD):/somedir --rm youimage:$(IMG_TAG) sh -c "cp /your/bin/path /somedir/"
```
例如从ubuntu中提取`tty`，如下所示：
```
$ docker run -v $(PWD):/opt --rm ubuntu bash -c "cp /usr/bin/tty /opt/"
```

# Linux

## 不常使用
### 命令
容器化大行其道，许多部署在K8s环境中的容器基础镜像都是alpine这类极简镜像，有时想要去容器中查域名，却苦于没有可用的命令，这时候不妨试试[`getent`](https://man7.org/linux/man-pages/man1/getent.1.html)这个命令，你值得拥有
```
# getent hosts google.com
142.251.43.14     google.com  google.com
```

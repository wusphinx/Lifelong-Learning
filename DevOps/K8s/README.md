# K8s

## 命令

### cp
想要从pod容器中拷贝文件，通常需要使用cp这个命令，盲猜跟scp的使用习惯一致，尝试一下，发现报错了如下：
```
kubectl cp mypod:/etc/debconf.conf debconf.conf                                                   
tar: Removing leading `/' from member names
```
这可与文档描述有点不符
``
kubectl cp --help
 ……
  # Copy /tmp/foo from a remote pod to /tmp/bar locally
  kubectl cp <some-namespace>/<some-pod>:/tmp/foo /tmp/bar
```
进行容器中发现当前工作目录是根目录"/"
```
kubectl exec sidecar-demo /bin/pwd
/
```
所以从现象看是因为cp操作默认是从容器的当前工作目录开始的，所以将命令改为如下即可生效
```
kubectl cp mypod:etc/debconf.conf debconf.conf                                                   
```
且必须指定拷贝后文件名，不然会报错
```
kubectl cp mypod:etc/debconf.conf .                                               
error: open ..: is a directory
```
有些想当然了，算是一个小坑吧
# JMX
>JMX是Java平台上为应用程序、设备、系统等植入管理功能的框架。JMX可以跨越一系列异构操作系统平台、系统体系结构和网络传输协议，灵活的开发无缝集成的系统、网络和服务管理应用。

Kafka的监控其实社区也有解决方案，通过jmx agent的方式虽然还是有侵入性，但却是社区所推荐，笔者根据文档写了一个demo，其它直接JMX协议的也可用些方式接入Prometheus监控生态体系

## sidecar embed JMX 
公司的kafka是部署在K8s中的(是的，你没看错)，若我想将jmx接入kafka，势必要将jmx agent打入kafka镜像中，若jmx agent有更新，也需要更新镜像，这么看起来还是有点耦合的，如果用jmx作为sidecar，和原kafka容器在同一个Pod，共享jmx_prometheus_javaagent.jar所在目录，这个方法还是有点意思的。


参考：
- [JMX wiki](https://zh.wikipedia.org/zh-cn/JMX)
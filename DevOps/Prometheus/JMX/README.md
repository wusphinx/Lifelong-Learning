# JMX
>JMX是Java平台上为应用程序、设备、系统等植入管理功能的框架。JMX可以跨越一系列异构操作系统平台、系统体系结构和网络传输协议，灵活的开发无缝集成的系统、网络和服务管理应用。

Kafka的监控其实社区也有解决方案，通过jmx agent的方式虽然还是有侵入性，但却是社区所推荐，笔者根据文档写了一个demo，其它直接JMX协议的也可用些方式接入Prometheus监控生态体系

参考：
- [JMX wiki](https://zh.wikipedia.org/zh-cn/JMX)
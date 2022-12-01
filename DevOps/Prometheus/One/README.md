# Prometheus
公司的测试开发环境太乱了，时不时掉线、磁盘满了不可用，经常性地一天有小半天都不可用，
这对于号称注重效率的我来说实在是无法忍受，毕竟曾经也设计过监控告警平台，说干就干，不得不说，手有点生，
花费的时间超过预期。反思之后发现还是因为之前的工作内容没有积极下来，没有对这次的形成根本上的助益，多纪录是个好习惯。

## centos上安装docker-compose
```
yum install python3-pip -y 
python3 -m pip install --upgrade pip setuptools wheel -i https://pypi.tuna.tsinghua.edu.cn/simple some-package
pip3 install docker-compose -i https://pypi.tuna.tsinghua.edu.cn/simple some-package
```
## 接入Prometheus生态
- prometheus: 收集监控数据
- alertmanager: 收到告警后webhook到钉钉
- dingding: 钉钉机器人

参考：
- https://awesome-prometheus-alerts.grep.to/（这个网站绝对值得收藏）
- [将钉钉接入 Prometheus AlertManager WebHook
](https://theo.im/blog/2017/10/16/release-prometheus-alertmanager-webhook-for-dingtalk/)
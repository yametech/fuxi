# 主机探针
期望结合CRD/webhook实现有状态资源单副本集高可用

1. 监控主机的存活,以daemonSet方式运行在每个加入集群的主机/多集群主机
2. 主动上报主机 kubelet进程状态 , probe cAdvisor , 文件系统占用空间 CRI/OCI 实例状态
3. 支持被动请求上报数据
# 日志收集CRD资源
基于 io event 实现,收集当前目录/文件的操作事件,并通过事件上报数据日志

1. 日志收集器 接收kubernetes (web Admission webhook)/CRD 资源请求参数的方式运行
2. 针对每个pod收集，由用户定义收集参数，包括目录，收集存储（kafka/elk...)

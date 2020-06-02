# base

1. 组织架构人员等创建
    * 1.1 增删改查 

2. 运维角色定义资源
    * 2.1 存储资源定义(需要了解Storage class,PV,PVC)
    * 2.2 CIDR 规划定义(生成configMap)
    * 2.3 部门机器划分    
    
3. 一个部门对应一个kubernetes namespace,由独立的服务监听informer实现;
    * 3.1 部门创建需要初始化配额,存储,网络
    * 3.2 其他资源

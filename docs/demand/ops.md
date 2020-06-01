# ops 需求说明

1. 基于kubernetes operator(Tekton) 拓展
2. 服务命名ops (20191024)
    2.1 服务需记录用户提交模板,~~实现第个应用至少3个版本的模板~~(待斟酌)),方便用户回滚编排

3. 依赖gitea client,最好实现多几个client,服务[github.com/yametech/fuxi/srv/giter](20191024)   
    3.1 git服务器client 可以实现用户配置,用户仓库组织list

4. 依赖镜像仓库服务,需要安排审核harbor仓库,服务[github.com/yametech/fuxi/srv/registry](20191024)   
    4.1 harbor/registry仓库实现及想着接口规范勘查

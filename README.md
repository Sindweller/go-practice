## 简介
### http-server
位于/http-server目录下的是作业
- 模块八作业2:
  - service.yaml
  - ingress.yaml
  - secret.yaml tls证书相关信息

部署过程示例:
```shell
# 创建deploy
k create -f deployment.yaml
# 创建svc
k create -f service.yaml
# 创建证书issuer
k create -f issuer.yaml
# 创建ingress
k create -f ingress.yaml
```
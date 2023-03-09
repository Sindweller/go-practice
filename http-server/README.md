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
# 安装ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
# 创建ingress
k create -f ingress.yaml
```

### 监控

修改代码后推送至docker 仓库

```shell
docker build -t sindweller/httpserver:v3.0.0 .
docker push sindweller/httpserver:v3.0.0
```
部署：

```shell
k apply -f deployment.yaml
k apply -f service.yaml
k apply -f ingress.yaml 
```

准备prom

```shell
helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false
sindweller@xindeweiladeMacBook-Pro .kube % k get po
NAME                                            READY   STATUS    RESTARTS   AGE
httpserver-7755f5bd8d-gkqvj                     1/1     Running   0          6m59s
loki-0                                          1/1     Running   0          5m39s
loki-grafana-75b567d8f6-xrmqd                   2/2     Running   0          5m39s
loki-kube-state-metrics-5c6b9ddd4f-7v2wp        1/1     Running   0          5m39s
loki-prometheus-alertmanager-7d5bdfcb7b-nlqxx   2/2     Running   0          5m39s
loki-prometheus-node-exporter-g4wpz             1/1     Running   0          5m39s
loki-prometheus-pushgateway-7cdf755958-xpwtb    1/1     Running   0          5m39s
loki-prometheus-server-6764f67456-8dfm6         1/2     Running   0          5m39s
loki-promtail-rh7lz                             1/1     Running   0          5m39s
```
这里借鉴同学的安装步骤修改grafana service类型为nodeport

```shell
kubectl patch svc loki-grafana --type='json' -p '[{"op":"replace","path":"/spec/type","value":"NodePort"},{"op":"replace","path":"/spec/ports/0/nodePort","value":30066}]'
service/loki-grafana patched
```

可以看到自定义的指标：

```text
# HELP default_execution_latency_seconds Time spent.
# TYPE default_execution_latency_seconds histogram
default_execution_latency_seconds_bucket{step="total",le="0.001"} 0
default_execution_latency_seconds_bucket{step="total",le="0.002"} 0
default_execution_latency_seconds_bucket{step="total",le="0.004"} 0
default_execution_latency_seconds_bucket{step="total",le="0.008"} 0
default_execution_latency_seconds_bucket{step="total",le="0.016"} 0
default_execution_latency_seconds_bucket{step="total",le="0.032"} 0
default_execution_latency_seconds_bucket{step="total",le="0.064"} 0
default_execution_latency_seconds_bucket{step="total",le="0.128"} 0
default_execution_latency_seconds_bucket{step="total",le="0.256"} 0
default_execution_latency_seconds_bucket{step="total",le="0.512"} 0
default_execution_latency_seconds_bucket{step="total",le="1.024"} 0
default_execution_latency_seconds_bucket{step="total",le="2.048"} 0
default_execution_latency_seconds_bucket{step="total",le="4.096"} 0
default_execution_latency_seconds_bucket{step="total",le="8.192"} 0
default_execution_latency_seconds_bucket{step="total",le="16.384"} 0
default_execution_latency_seconds_bucket{step="total",le="+Inf"} 1
default_execution_latency_seconds_sum{step="total"} 18.000442819
default_execution_latency_seconds_count{step="total"} 1
```


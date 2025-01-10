## 下载依赖包
https://github.com/apache/skywalking-python/blob/master/docs/en/setup/Configuration.md

配置参数：
https://github.com/apache/skywalking-python/blob/320ef706ddeb1cbf6230bb5adf153ed71ad211a6/docs/en/setup/Configuration.md

对接参数：
1、直接代码config.init设置
2、环境变量设置，参考配置参数

Dockerfile所在目录

docker build -f Dockerfile -t 10.64.17.85:30085/demo/python-demo:latest .

kubectl create ns demo

helm install skywalking-demo -ndemo ../deploy/ --set skywalking.service=tracing.istio-system.svc.cluster.local:11800 --set image=10.64.17.85:30085/demo/python-demo:latest
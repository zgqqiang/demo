## 下载依赖包，编译的过程中修改对应的插件监听入口调用调用链
https://www.apache.org/dyn/closer.cgi/skywalking/go/0.5.0/apache-skywalking-go-0.5.0-bin.tgz

配置参数https://github.com/apache/skywalking-go/blob/72414bcb8cda68ad1a40d9ac0b3d6487cea4f7c6/tools/go-agent/config/agent.default.yaml

对接参数：
1、环境变量设置

Dockerfile所在目录

docker build -f Dockerfile -t 10.64.17.85:30085/demo/go-demo:latest .

kubectl create ns demo

helm install skywalking-demo -ndemo ../deploy/ --set skywalking.service=tracing.istio-system.svc.cluster.local:11800 --set image=10.64.17.85:30085/demo/go-demo:latest
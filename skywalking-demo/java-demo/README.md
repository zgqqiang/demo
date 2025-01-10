## 下载依赖包
https://dlcdn.apache.org/skywalking/java-agent/9.3.0/apache-skywalking-java-agent-9.3.0.tgz

配置参数：
https://github.com/apache/skywalking-java/blob/b358267905be8748d105014cc3da066aacd33b46/docs/en/setup/service-agent/java-agent/configurations.md

对接参数
1、直接修改skywalking-agent/config/agent.config
2、环境变量设置参考skywalking-agent/config/agent.config
3、通过系统属性配置agent.service_name
java -javaagent:/path/to/skywalking-agent.jar -Dskywalking.agent.service_name=<your_service_name> -jar your-project.jar
4、在JVM参数中的代理路径之后添加属性
# 模板
-javaagent:/path/to/skywalking-agent.jar=[key1]=[value1],[key2]=[value2]
# 举例
java -javaagent:/path/to/skywalking-agent.jar=agent.service_name=<your-service-name>,agent.authentication=<your-token> -jar your-project.jar

Dockerfile所在目录

docker build -f Dockerfile -t 10.64.17.85:30085/demo/java-demo:latest .

kubectl create ns demo

helm install skywalking-demo -ndemo ../deploy/ --set skywalking.service=tracing.istio-system.svc.cluster.local:11800 --set image=10.64.17.85:30085/demo/java-demo:latest
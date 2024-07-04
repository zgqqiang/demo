部署删除demo
kubectl apply -f <(helm template deploy)
kubectl delete -f <(helm template deploy)

打镜像
docker build -f Dockerfile-curl -t 10.64.17.48:30085/library/curl:1.0 .
docker build -f Dockerfile-curl -t 10.64.17.48:30085/demo/k8s-service:latest .
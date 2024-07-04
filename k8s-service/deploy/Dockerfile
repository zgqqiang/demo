FROM 10.64.17.48:30085/library/curl:1.0
workdir /test
COPY k8s-service /test
RUN chmod +x k8s-service
ENTRYPOINT ["./k8s-service"]

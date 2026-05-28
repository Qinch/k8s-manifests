FROM alpine:3.20

ARG PROJECT_NAME
ENV PROJECT_NAME=${PROJECT_NAME}
WORKDIR /k8s-manifests/part2/app-gateway
COPY target/bin ./bin
RUN chmod +x ./bin/${PROJECT_NAME}
CMD ["sh", "-c", "exec ./bin/${PROJECT_NAME}"]


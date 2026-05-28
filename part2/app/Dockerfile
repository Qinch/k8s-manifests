FROM alpine:3.20
ARG PROJECT_NAME
ENV PROJECT_NAME=${PROJECT_NAME}
WORKDIR /part2/app
COPY target/* ./
RUN chmod +x ./bin/${PROJECT_NAME}
CMD ["sh", "-c", "exec ./bin/${PROJECT_NAME}"]


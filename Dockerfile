FROM alpine
WORKDIR /app
COPY ${WERCKER_OUTPUT_DIR}/k8s /data/jzp/k8s
EXPOSE 1211
ENTRYPOINT /data/jzp/k8s
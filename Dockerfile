FROM alpine
ADD group_svc-service /group_svc-service
ENTRYPOINT [ "/group_svc-service" ]

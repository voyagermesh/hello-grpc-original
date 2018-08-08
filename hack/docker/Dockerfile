FROM alpine

RUN set -x \
  && apk add --update --no-cache ca-certificates curl

COPY hello-grpc /usr/bin/hello-grpc

USER 1971
ENTRYPOINT ["hello-grpc"]

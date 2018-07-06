#!/usr/bin/env bash

rm -rf dist

cat > hack/docker/Dockerfile_${ARCH} <<EOF
FROM $BASE

RUN set -x \
  && apk add --update --no-cache ca-certificates

COPY hello-grpc /usr/bin/hello-grpc

USER 1971
ENTRYPOINT ["hello-grpc"]
EOF

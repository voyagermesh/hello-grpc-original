FROM appscodeci/dind

RUN apt-get update && apt-get install -y sudo gnupg2
RUN curl -LO https://github.com/goreleaser/goreleaser/releases/download/v0.79.1/goreleaser_amd64.deb && dpkg -i goreleaser_amd64.deb

FROM golang:1.9 AS build-env
ENV GOPATH /go
ADD . /go
RUN cd /go \
    && go get github.com/google/go-github/github \
    && go get golang.org/x/oauth2 \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gh_scrap src/main.go

FROM alpine
ENV GITHUB_TOKEN ""
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /go/gh_scrap /go/input.txt /app/
ENTRYPOINT ./gh_scrap input.txt

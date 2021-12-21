FROM golang:1.15 as build
COPY . /go/gitops
RUN cd /go/gitops && go build -o bin/gitops -ldflags "-linkmode external -extldflags -static" cmd/gitops/main.go

FROM alpine:3
RUN apk add bash
COPY  --from=build /go/gitops/bin/gitops /usr/bin/gitops
ENTRYPOINT ["/bin/bash", "-c"]
CMD ["env"]

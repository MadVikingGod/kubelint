FROM golang:1.14.1-alpine3.11 AS BUILD
WORKDIR kubelint
COPY go.mod go.sum ./
COPY cmd cmd/
COPY pkg pkg/
RUN pwd; ls -l
RUN GOOS=linux GOARCH=amd64 go build -o bin/kubelint ./cmd

FROM alpine:3.11.5
COPY --from=BUILD /go/kubelint/bin/kubelint /bin/kubelint
CMD kubelint
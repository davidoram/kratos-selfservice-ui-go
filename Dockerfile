FROM golang:1.16 AS builder

WORKDIR /go/src/github.com/davidoram/kratos-selfservice-ui-go
ADD go.mod go.mod
ADD go.sum go.sum

ENV GO111MODULE on

RUN go mod download

ADD . .

RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o /usr/bin/kratos-selfservice-ui-go

FROM scratch
COPY --from=builder /usr/bin/kratos-selfservice-ui-go /

# Expose the default port that we will be listening to
EXPOSE 4455

ENTRYPOINT ["/kratos-selfservice-ui-go"]
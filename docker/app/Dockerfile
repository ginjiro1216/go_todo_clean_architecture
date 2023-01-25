FROM golang:latest

ENV ROOT=/go/src
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

COPY src/ /go/src
RUN go install github.com/cosmtrek/air@latest
RUN go mod download
EXPOSE 7050

CMD ["air", "-c", ".air.toml"]
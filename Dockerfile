FROM golang:1.18-bullseye

RUN apt-get update -y
RUN apt-get install iputils-ping -y

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CURZOR_HOME /go/src/curzor

RUN mkdir -p "$CURZOR_HOME"

COPY . "$CURZOR_HOME"
WORKDIR "$CURZOR_HOME"

RUN go mod vendor

EXPOSE 8080

CMD ["go", "run", "."]
FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY ./src /go/src/app

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 80 443

CMD ["go-wrapper", "run"] # ["app"]

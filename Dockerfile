FROM golang:1.8 AS build
WORKDIR /go/src/github.com/glassechidna/sshello/
RUN curl https://glide.sh/get | sh
RUN apt-get update
RUN apt-get install -y upx-ucl

COPY glide.yaml glide.lock ./
RUN glide install

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build
RUN upx sshello

FROM scratch
COPY --from=build /go/src/github.com/glassechidna/sshello/sshello .
CMD ["/sshello"]

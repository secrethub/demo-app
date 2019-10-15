FROM golang:1.13-alpine as build
WORKDIR /build
RUN apk add --update git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o demo ./cmd/main.go

FROM alpine
COPY --from=build /build/demo /usr/bin/demo

RUN apk add --repository https://alpine.secrethub.io/alpine/edge/main --allow-untrusted secrethub-cli
ADD secrethub.env .

EXPOSE 80
ENTRYPOINT ["secrethub", "run", "--"]
CMD ["demo", "serve", "--host", "0.0.0.0", "--port", "80"]

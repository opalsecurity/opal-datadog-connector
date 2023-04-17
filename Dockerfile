FROM golang:1.10 AS build
WORKDIR /go/src
COPY datadogconnector ./datadogconnector
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o datadogconnector .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/datadogconnector ./
EXPOSE 8080/tcp
ENTRYPOINT ["./datadogconnector"]

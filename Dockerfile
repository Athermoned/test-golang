FROM golang:latest AS build-http
ENV GO111MODULE=on
WORKDIR /usr/src/http-anagrams
ADD main.go ./
ADD go.mod ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest 
WORKDIR /root/
COPY --from=build-http /usr/src/http-anagrams/main ./
CMD ["./main"]

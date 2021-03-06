

# Build yemeksepetiCase in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /yemeksepetiCase
RUN cd /yemeksepetiCase && go mod download && go build

# Pull yemeksepetiCase into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /yemeksepetiCase/yemeksepetiCase /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["yemeksepetiCase"]

##
## Build
##

FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod .
# COPY go.sum .
RUN go mod download

COPY . . 

RUN go build -o /httpserv ./cmd/httpserv/*.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /httpserv /httpserv

# EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/httpserv"]

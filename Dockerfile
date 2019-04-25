# Maintain Golang Version
ARG GO_VERSION=1.11
################# MULTI STAGE BUILD ##############################
# Stage-1 Build
FROM golang:${GO_VERSION}-alpine AS build
LABEL MAINTAINER=https://github.com/ChaturvediSulabh
RUN apk add --update --no-cache ca-certificates git

WORKDIR /app
# Manage Dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy Source Code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app

# Stage-2 Build
FROM scratch
COPY --from=build /go/bin/app /go/bin/app
ENTRYPOINT [ "/go/bin/app" ] 
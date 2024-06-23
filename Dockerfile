# ------------------------------------------------------------------------------
# Base API image
# ------------------------------------------------------------------------------
FROM golang:alpine as api-base
WORKDIR $GOPATH/src/github.com/Sigma-Ratings/sigma-code-challenges/
RUN apk add --update \
    git \
    openssh

# Download go modules into local cache
COPY go.* ./
RUN go mod download

# ------------------------------------------------------------------------------
# Final Lib build
# ------------------------------------------------------------------------------
FROM api-base as api-build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o libsanctions  .

# ------------------------------------------------------------------------------
# Final image
# ------------------------------------------------------------------------------
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /api/
COPY --from=api-build /go/src/github.com/Sigma-Ratings/sigma-code-challenges/libsanctions ./libsanctions

ENTRYPOINT ["./libsanctions"]
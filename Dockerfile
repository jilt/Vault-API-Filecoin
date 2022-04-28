# syntax=docker/dockerfile:1

##
## build
##

FROM  golang:1.17.3-alpine AS base
WORKDIR /src
ENV CGO_ENABLED=0
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
COPY go.* ./

# Add the keys
ARG github_user
ARG github_personal_token
RUN apk add --no-cache git
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download


FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build go build -o /out/filecoin .

FROM base as unit-test
RUN --mount=target=. \
    --mount=type=cache,target=go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir /out && go test -v -coverprofile=/out/cover.out ./...


FROM golangci/golangci-lint:v1.31.0-alpine AS lint-base

FROM base as lint
RUN --mount=target=. \
    --mount=from=lint-base,src=/usr/bin/golangci-lint,target=/usr/bin/golangci-lint \
    --mount=type=cache,target=go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run --timeout 10m0s ./...

FROM scratch AS unit-test-coverage
COPY --from=unit-test /out/cover.out /cover.out

FROM scratch AS bin-unix
COPY --from=build /out/filecoin /

FROM bin-unix as bin
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./filecoin" ]
EXPOSE 8080

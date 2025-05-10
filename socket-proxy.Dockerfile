# Stage 1: deps
FROM golang:1.24.3-alpine AS deps
HEALTHCHECK NONE

# package version does not matter
# trunk-ignore(hadolint/DL3018)
RUN apk add --no-cache tzdata make libcap-setcap

ENV GOPATH=/root/go

WORKDIR /src

COPY socket-proxy/go.mod socket-proxy/go.sum ./

RUN go mod download -x

# Stage 2: builder
FROM deps AS builder

WORKDIR /src

COPY Makefile ./
COPY socket-proxy ./socket-proxy

ARG VERSION
ENV VERSION=${VERSION}

ARG MAKE_ARGS
ENV MAKE_ARGS=${MAKE_ARGS}

ENV GOCACHE=/root/.cache/go-build
ENV GOPATH=/root/go

RUN make ${MAKE_ARGS} docker=1 build

# Stage 3: Final image
FROM scratch

LABEL maintainer="yusing@6uo.me"
LABEL proxy.exclude=1

# copy timezone data
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# copy binary
COPY --from=builder /app/run /app/run

WORKDIR /app

CMD ["/app/run"]
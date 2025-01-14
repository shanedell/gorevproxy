FROM gcr.io/distroless/static-debian11:debug AS build

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /tmp

COPY gorevproxy /

ARG BUILD_DATE
ARG BUILD_VERSION
ARG VCS_REF
ARG VCS_URL

LABEL org.opencontainers.image.created=$BUILD_DATE
LABEL org.opencontainers.image.title="gorevproxy"
LABEL org.opencontainers.image.description="CLI to spin up Golang Reverse proxy. Configured via JSON or YAML file"
LABEL org.opencontainers.image.source=$VCS_URL
LABEL org.opencontainers.image.revision=$VCS_REF
LABEL org.opencontainers.image.vendor="Shane Dell"
LABEL org.opencontainers.image.version=$BUILD_VERSION

ENTRYPOINT ["/gorevproxy"]

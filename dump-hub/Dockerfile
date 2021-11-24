FROM golang:1.17.3 AS dh-builder
WORKDIR /opt/dump-hub
COPY . .
RUN make

FROM busybox:stable AS dump-hub-be
WORKDIR /usr/bin
COPY --from=dh-builder /opt/dump-hub/bin/dump-hub .
ENTRYPOINT ["/usr/bin/dump-hub"]
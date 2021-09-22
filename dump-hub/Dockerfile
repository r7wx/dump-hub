FROM golang AS dh-builder
WORKDIR /opt/dump-hub
COPY . .
RUN make

FROM busybox AS dump-hub-be
WORKDIR /usr/bin
COPY --from=dh-builder /opt/dump-hub/bin/dump-hub .
ENTRYPOINT ["/usr/bin/dump-hub"]
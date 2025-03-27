FROM golang:1.24 AS build

RUN useradd -u 10001 bento

WORKDIR /build/
COPY . /build/

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor ./cmd/geo-bento

FROM busybox AS package

WORKDIR /

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /build/geo-bento .
COPY ./testdata/s2.yaml /s2.yaml

USER bento

EXPOSE 4195

ENTRYPOINT ["/geo-bento"]

CMD ["-c", "/s2.yaml"]

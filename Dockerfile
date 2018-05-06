FROM ckeyer/go AS building

COPY . /go/src/github.com/funxdata/landlady
WORKDIR /go/src/github.com/funxdata/landlady

RUN make build

FROM alpine:edge

COPY --from=building /go/src/github.com/funxdata/landlady/bundles/landlady /bin/landlady

ENTRYPOINT ["/bin/landlady"]

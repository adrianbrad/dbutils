#BUILD
FROM golang:buster as builder

WORKDIR dbutils
COPY . .

RUN go test -c . -o /tmp/unit.test -ldflags="-w -s" -race

#RUN
FROM busybox:glibc as runner

COPY --from=builder /tmp/unit.test .
CMD ./unit.test -test.v -test.parallel 50 -test.failfast
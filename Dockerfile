FROM alpine:3.1
MAINTAINER Go Incremental Limited <info@goincremental.com>
COPY bin/enforcer enforcer
EXPOSE 9090
# CMD /enforcer

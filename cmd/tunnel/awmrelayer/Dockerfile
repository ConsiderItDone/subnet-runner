FROM golang:1.22.7 AS builder
WORKDIR /app
RUN git clone https://github.com/ava-labs/awm-relayer.git --branch v1.4.0
RUN cd /app/awm-relayer && git submodule update --init --recursive
RUN cd /app/awm-relayer && /bin/bash ./scripts/build_relayer.sh

FROM ubuntu:latest
COPY --from=builder /app/awm-relayer/build/awm-relayer /usr/bin/awmrelayer
EXPOSE 8080
USER 1001
CMD ["start"]
ENTRYPOINT [ "/usr/bin/awmrelayer" ]
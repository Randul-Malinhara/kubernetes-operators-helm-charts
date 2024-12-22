FROM golang:1.20 AS builder
WORKDIR /workspace
COPY . .
RUN make build

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532
ENTRYPOINT ["/manager"]
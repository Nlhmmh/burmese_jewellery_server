FROM golang:1.20.5 AS builder

# Copy all files and download dependencies
WORKDIR /build
COPY . .
RUN set -x && go mod download

# Go Build
WORKDIR /build
RUN set -x && \
	CGO_ENABLED=0 \
	GOOS=linux \
	go build -a -tags netgo -installsuffix netgo \
	--ldflags "-extldflags '-static'" \
	-o app .

# Set distroless image as runner and copy app binary to runtime
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static-debian11:nonroot AS runner
WORKDIR /runtime/
COPY --from=builder /build/app .
COPY config.yaml .
COPY certs/cert.pem /runtime/certs/
COPY certs/key.pem /runtime/certs/

# Run the app
ENTRYPOINT ["/runtime/app"]
CMD ["-debug"]

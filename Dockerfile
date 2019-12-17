# First stage, for building 
FROM golang

# Assuming the source code is collocated to this Dockerfile
COPY . .

# Build the Go app with CGO_ENABLED=0 so we use the pure-Go implementations for
# things like DNS resolution (so we don't build a binary that depends on system
# libraries)
# RUN CGO_ENABLED=0 go build 
RUN go version
RUN GOPATH=/tmp go test -v .
RUN GOPATH=/tmp CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /stockapp scripts/main.go

# Create a "nobody" non-root user for the next image by crafting an /etc/passwd
# file that the next image can copy in. This is necessary since the next image
# is based on scratch, which doesn't have adduser, cat, echo, or even sh.
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd

# The second and final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=0 /stockapp /stockapp

# Copy the /etc_passwd file we created in the builder stage into /etc/passwd in
# the target stage. This creates a new non-root user as a security best
# practice.
COPY --from=0 /etc_passwd /etc/passwd

# Run as the new non-root by default
USER nobody

EXPOSE 8080

CMD ["/stockapp"]
FROM ubuntu

# Install Dependencies
RUN apt-get update
RUN apt-get install -y curl
RUN apt-get upgrade -y
RUN curl -OL https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go1.24.4.linux-amd64.tar.gz

# Set Go PATH
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Set working directory
WORKDIR /app

# Copy Files
COPY go.mod go.mod
COPY README.md README.md
COPY main main

# Build the Go binary
RUN go build -o godex ./main/GoDex.go

# Run the binary
ENTRYPOINT ["./godex"]
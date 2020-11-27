# Start by building the application.
FROM centos:7 as builder

# Setup dependencies
RUN yum -y install wget gcc

# Download Go
RUN wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz && \
    tar zxvf go1.15.linux-amd64.tar.gz --no-overwrite-dir -C /usr

# Set the current working directory inside the container.
WORKDIR /go/src/app

# copy go.mod <optimization>
COPY go.mod .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN /usr/go/bin/go mod download
# </optizimation>

# Copy the source from the current directory to the Working Directory inside the container.
COPY . .

# Build the Go app.
RUN /usr/go/bin/go build cmd/main.go

# Command to run the executable.
CMD ["./main"]

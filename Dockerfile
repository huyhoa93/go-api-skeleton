FROM golang:1.14.0-alpine

# Add Maintainer Info
LABEL maintainer="HoaNH <hoanh160993@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
RUN mkdir -p /home/go/app
WORKDIR /home/go/app

# Copy go mod and sum files
#COPY ./go.mod ./go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

#RUN go get -d -v
RUN go get -d -v ./...
# Build the Go app
#RUN go build -o main .

# Expose port 8080 to the outside world
#EXPOSE 8084

#ENV CGO_ENABLED=0

#RUN go get github.com/githubnemo/CompileDaemon
#CMD [ "CompileDaemon", "--build=go build main.go", "--command=./main" ]
#ENTRYPOINT CompileDaemon -include=*.* -build="go build" -command=./main

RUN go get github.com/codegangsta/gin
CMD [ "gin","-i","-a","8084","run" ]


# Command to run the executable
#CMD [ "go", "run", "main.go" ]

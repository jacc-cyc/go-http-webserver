# Pull base image
FROM golang:latest

# Create working directory
RUN mkdir C:\\sample

# Set a working directory
WORKDIR c:\\sample

# Copy package to the image
ADD . .

# Build the go program
RUN go build

# Run the application
CMD ["go", "run", "httpwebserver.go"]

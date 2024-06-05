FROM golang:1.22-alpine



# Set the current working directory inside the container
WORKDIR /app

# install air
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go mod download golang.org/x/text

RUN go build -o ./tracking-inventory

# Expose port 8080 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["air"]
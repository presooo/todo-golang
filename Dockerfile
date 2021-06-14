# Use latest golang version
FROM golang

# Create the app directory in the docker container
RUN mkdir -p $GOPATH/src/todo

# Set the default directory
WORKDIR $GOPATH/src/todo

# Install dependencies
RUN go get -u github.com/gorilla/mux && go get -u github.com/go-martini/martini && go get -u go.mongodb.org/mongo-driver/mongo && go get -u github.com/spf13/viper

# Copy all files
COPY . .

# Expose port 8080
EXPOSE 8080

# cmd to start service
CMD [ "go", "run", "main.go" ]

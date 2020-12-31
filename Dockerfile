
# using golang:latest image
FROM golang:latest

MAINTAINER Prangan Majumder <pranganmajumder@appscode.com>

#it will create app directory in inside container
WORKDIR /app

# copy all file from project directory to app directory
COPY . .

# run this command to download all the required go module mentioned in go.mod
RUN go mod tidy


#build the file named apiserver or give what you want
RUN go build -o apiserver .

#expose the port to bind with the localport
EXPOSE 8080

# it will start like ./apiserver start
CMD ["./apiserver", "start"]
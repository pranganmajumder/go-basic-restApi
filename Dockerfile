
#-----------------------------------------------------------------------------------------------
# At firt build the binary file into your local directory run the follwoing command
# go build -o apiserver .
# it will create a binaryfile "apiserver" inside your project root directory
#just copy the previously build binaryfile "apiserver" into your workdirectory
#that's why don't neet to build the binary file so skipped it in RUN section
#
FROM busybox:glibc
WORKDIR /go/src/github.com/pranganmajumder/go-basic-restApi
COPY apiserver /go/src/github.com/pranganmajumder/go-basic-restApi
EXPOSE 8080
CMD ["start"]
ENTRYPOINT ["./apiserver"]





#-------------------------------------------------------------------------------------------------

# At firt build the binary file into your local directory run the follwoing command
# go build -o apiserver .
# it will create a binaryfile "apiserver" inside your project root directory
#just copy the previously build binaryfile "apiserver" into your workdirectory
#that's why don't neet to build the binary file so skipped it in RUN section

#FROM golang:latest
#WORKDIR /go/src/github.com/pranganmajumder/go-basic-restApi
#COPY apiserver  /go/src/github.com/pranganmajumder/go-basic-restApi
#
#EXPOSE 8080
#CMD ["start"]
#ENTRYPOINT ["./apiserver"]






#--------------------------------------------------------------------------------------------
# Copy all source file from project directory into your workdirectory
# build the binary file inside your container directory  " /go/src/github.com/pranganmajumder/go-basic-restApi "
# as you build the binary file inside the container directory
# like this  /go/src/github.com/pranganmajumder/go-basic-restApi/apiserver


## using golang:latest image
#FROM golang:latest
#WORKDIR /go/src/github.com/pranganmajumder/go-basic-restApi
#
#COPY .  /go/src/github.com/pranganmajumder/go-basic-restApi
#
#RUN go mod tidy && go build -o apiserver .
#EXPOSE 8080
#
## it will start like ./apiserver start
#CMD ["start"]
#ENTRYPOINT ["./apiserver"]












#----------------------------------- using golang latest image ---------------------------
# using golang:latest image
#FROM golang:latest
#MAINTAINER Prangan Majumder <pranganmajumder@appscode.com>

#it will create app directory in inside container
#WORKDIR /app

# copy all file from project directory to app directory
#COPY . .

# run this command to download all the required go module mentioned in go.mod
#RUN go mod tidy

#build the file named apiserver or give what you want
#RUN go build -o apiserver .

#expose the port to bind with the localport
#EXPOSE 8080

# it will start like ./apiserver start
#CMD ["./apiserver", "start"]

[![Go Report Card](https://goreportcard.com/badge/github.com/pranganmajumder/go-basic-restApi)](https://goreportcard.com/report/github.com/pranganmajumder/go-basic-restApi)
# Run api server in terminal
```bash
Build and run the app locally
$ go build -o <binary name you want>
$ ./binaryName start

to add flag
$ ./binaryName start -p=6000 
```

```$ go run main.go```

## To run unit testing

   * single function unit testing
        * generate main_test.go file then design the function you want to test
    
   * run main_test.go from drag and drop
   * if it shows the pass result then the programm is ok




## Add CLI command to run server

### To get cobra module , it'll add cobra in your go.mod file

```go get -u github.com/spf13/cobra/cobra```
### Init cobra in your project , it'll create a directory named cmd and there is a file named root.go and<br>
create a main.go file automatically in your project directory</br>
```cobra init --pkg-name <Project_name>```
```ex : cobra init --pkg-name go-basic-restApi```

### then to add a subcommand for root.go
```$ cobra add <command_name>```  ``` ex : $ cobra add start``` </br>
it will create a file command_name.go inside the cmd directory, modify it according to your needs


### To start api using cli
___
```$ go install go-basic-restApi``` </br>
It'll create a binary file named go-basic-restApi</br>

Run the binary file <br>
``` $ go-basic-restApi start```

#####another way</br>
 * build the file in your project directory</br>
```$ go build -o <binaryfilename_you_want>```</br>
   ex : ```$ go build -o apiserver``` </br>
It will create a binary file inside your project directory named "apiserver"
 * Run the binary file</br>
  ``` $ ./apiserver start``` or </br>
   ```$ ./apiserver start -p <port_name_you_want>```
 * Check the port from curl or postman

##### CLI command
___
* ```go-restApi_basic start -h``` or ```go-restApi_basic start -help```
* to change the default port  ```go-restApi_basic start --port=8000```
* ```$ go run main.go --port=8000``` it'll start the port at 8000





### Check basic authentication for /user/{id} endpoint
```
start api with the following command
$ go install go-restApi_basic
$ go-restApi_basic start
Now add username & password in postman, send the GET request & see the status code if for User: prangan Pass: 1234
```



#### Run using dockerfile
```bigquery
docker build -t <give a image name> .
docker run -d -p localport:containerport --name <give container name> <given_image_name> 
``` 

######Other Necessary command</br>
|Command | Use |
|--------|-----|
|`gofmt -s -w <filename.go>` |Format your code with gofmt|
|`golint <filename.go>`|Format your code with go lint |
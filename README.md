# Run api server in terminal
```$ go run main.go```
## To run unit testing
___
   * sinle function unit testing
        * generate main_test.go file then design the function you want to test
    
   * run main_test.go from drag and drop
   * if it shows the pass result then the programm is ok




## Implement CLI
___
### To get cobra module , it'll add cobra in your go.mod file

```go get -u github.com/spf13/cobra/cobra```
### Init cobra in your project , it'll create a directory named cmd and there is a file named root.go and<br>
create a main.go file automatically in your project directory</br>
```cobra init --pkg-name <Project_name>```

### then to add a subcommand for root.go
```cobra add <command_name>```</br>
it will create a file command_name.go inside the cmd directory, modify it according to your needs


### To start api using cli
___
```$ go install go-restApi_basic```
it'll create a binary file named go-restApi_basic</br>

Run the binary file ``` $ go-restApi_basic start```
##### CLI command
___
* ```go-restApi_basic start -h``` or ```go-restApi_basic start -help```
* to change the default port  ```go-restApi_basic start --port=8000```

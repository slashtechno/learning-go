### Creating a module  
Most Go code should be initialized as a module. Modules allow for easy management of dependencies and packages.  
To initalize the current directory as a module: `go mod init name`  
This creates the `go.mod` file which records dependencies  
### Imporing (local) packages from a Go program  
There are two ways to import packages in Go:  
`import package`  
```
import(
    "package"
    "anotherpackage"
)
```  
These packages are automatically added to `go.mod`  

### Compiling programs  
Programs can be compiled with `go build file`  
When compiling, dependencies are also compiled.  

### Importing remote packages  
Remote packages, such as from Github, can be imported with the same syntax as local packages. However, the full URL must be provided.  
For example, `import github.com/muxinc/mux-go`  

### Using go get  
Go get can be used to manually add packages. For example, if I wanted to add `mux` as a dependency, I could run the following  
`go get github.com/muxinc/mux-go`  

### Dependency Location  
When dependencies are downloaded, they are installed to `$GOPATH/pkg/`  

### Importing a local package  
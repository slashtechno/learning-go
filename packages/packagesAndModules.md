t### Creating a module  
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
These packages are automatically added to `go.mod` when `go mod tidy` is ran    

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
A local package can be imported by importing the path which contains the module.  
Example:  
`import /path/to/package/folder/`  
If the local package is in a subdirectory of the current module, for example, `import "example/strutil"` if "example" is the module name and "strutil" is a subdirectory.  

### Creating a local package  
* A local package can be created by making a file inside a folder.  
    * The folder should have the name of the package you are creating.  
    * The file name doesn't matter.  
The file should contain code that is similar  to the following:   
```  
package <packagename>
funcion <functioname> {
    <function code>
}
```  
#### Utilizing the local package  
Once the local package is imported, it can be used with `packagename.function`  

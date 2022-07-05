# Setup a Go Environment  
Go depends on two environment variables to be set  
* GOPATH  
* GOROOT  

### GOPATH  
GOPATH is where binaries and other packages are located. This can be set to any directory, but the default is `~/go`  
GOPATH can be set it two ways. Through system environment variable configuration, or through `go env`  
The command to set the GOPATH environment variable through `go env` is `go env -w GOPATH=<path>`  

### GOROOT  
GOROOT is where Go, and other Go files, are located. This should be set to the directory where Go is installed.  
GOROOT can be set it two ways. Through system environment variable configuration, or through `go env`  
The command to set the GOROOT environment variable through `go env` is `go env -w GOROOT=<path>`  

### PATH  
Your system's PATH contains executables that can be ran from the command line without needed an absolute path  
`$GOPATH/bin` and `$GOROOT/bin` both should be on your PATH  
#### PATH on Windows  
On Windows, this should already be the case. If not, search "Edit system environment variables" in Windows Start. Then, click Environment Variables and then select Path.  
#### PATH on Linux  
If you're on Linux, you should edit `~/.bashrc` or your shell's equivalent and add the following line:  
`export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"`  

### Example configuration  
```go
export GOPATH="$HOME/go"
export GOROOT="/usr/local/go"
export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"
```

### Conclusion  
Go relies on environment variables. There are ways to set these up, but `go env` is the simplest way. Also, GOPATH/bin and GOROOT/bin should be on your PATH.  

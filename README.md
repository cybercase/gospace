# gospace
Generates a script that set environment variables for Go workspace


```
$ go run gospace.go /path/to/your/workspace  # Generates "activate" script
$ source /path/to/your/workspace/activate

(workspace)$ echo $GOPATH
/path/to/your/workspace

(workspace)$ echo $PATH
/path/to/your/workspace/bin:...

(workspace)$ deactivate
$ echo $GOPATH

```

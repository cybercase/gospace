# gospace
Switch back and forth between a generic and a project-specific workspace.

## Install
```
go get github.com/cybercase/gospace
```

## Use
```
$ echo $GOPATH
/home/user/yourgopath

$ cd /path/to/custom/workspace
$ gospace
You're ready to go.

- Type 'source activate' jump into your new workspace
- Type 'deactivate' to step out
$ source activate

(workspace)$ echo $GOPATH
/path/to/custom/workspace

(workspace)$ echo $PATH
/path/to/custom/workspace/bin:...

(workspace)$ deactivate
$ echo $GOPATH
/home/user/yourgopath

```

## Note
- The `activate` script should be placed in your workspace root (aka the folder containig `src/ bin/ pkg/`)
- You can commit the `activate` script since it's not bound to any specific path

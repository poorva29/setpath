setpath is used to set GOPATH and handle go commands easily when using multiple workspaces.

Say, you have multiple workspaces which look like -

```
$HOME/workspace/client1
$HOME/workspace/client2
```

Then, for setting go path for these workspaces, everytime you `cd` to the path and this can be little cumbersome.
In this case you don't have a common GOPATH and you want to use unique GOPATH.


Usage:
To build a project:
```
poorva:~/workspace/client1$ setpath go build .
Setting GOPATH: /Users/poorva/workspace/client1/
```

To install dependencies
```
poorva:~/workspace/client1$ setpath go install github.com/hashicorp/terraform
Setting GOPATH: /Users/poorva/workspace/client1/
```


All the binaries by default will be stored in $GOPATH/bin
In this case /Users/poorva/workspace/client1/bin
```
poorva:~/workspace/client1$ setpath terraform plan
Setting GOPATH: /Users/poorva/workspace/client1/
```

All the commands are directly run in shell

Though it is not a mandate to have `src` folder before you build go files, this assumes that you have your files to build/install/run or for that matter any command in the `src` folder
In case you want to just run the commands *without* `src` folder then use `-no-src` flag
```
poorva:~/workspace/client1$ setpath -no-src go build . <other flags>
```

You can download mac and linux (64 bit) binaries from -
[mac](https://s3.ap-south-1.amazonaws.com/gosetpath/setpath_darwin)
[Linux](https://s3.ap-south-1.amazonaws.com/gosetpath/setpath_linux_64)

Fell free to compile for your operating system or file bugs




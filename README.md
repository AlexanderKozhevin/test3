# media-server-go
WebRTC media server for go



|         | x86 | x64 |
|:------- |:--- |:--- |
| Linux   | -   | ✔︎   | 
| macOS   | -   | ✔︎   | 

## How to generate swig code  

swig -go -cgo -c++ -intgosize 64 mediaserver.i


## How to build

1,  clone the code, include the submodule

2,  make all

3,  go build 

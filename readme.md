## 导出c静态库

### arm系统在导出前要进行操作

export GOOS="linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  

go build -buildmode=c-archive -o number.a

## c调用

gcc -o a.out _test_main.c number.a  
./a.out

## 在linux中操作

### 生成动态库

在linux中进行动态链接库文件生成，在main目录下执行  
go build -buildmode=c-shared -o ntisdk.so  
注意go调用要生成plugin  
go build -buildmode=plugin -o ntisdk-plugin.so main.go  

### 运行  

生成可执行文件  
gcc -o a.out main.c ntisdk.a  

将.so .h文件拷贝至指定目录 在main.c中调用相关方法之后设置下查找动态库路径的环境变量
用于指定查找共享库（动态链接库）时除了默认路径（./lib和./usr/lib）之外的其他路径。  
export LD_LIBRARY_PATH="/root/goProject/src/NTI-SDK/c"  

### 指定  

./a.out  

## build plugin

go build -buildmode=plugin -o  ntisdk1.so ntisdk-plugin.go

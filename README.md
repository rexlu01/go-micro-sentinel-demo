基于go micro框架简单实现sentinel模块的流控和熔断实例

## 依赖
* go-micro v1.18
* protoc、protoc-gen-go、protoc-gen-micro
* sentinel
* consul

#### 生成接口文件
普通：
protoc --proto_path=. --go_out=. --micro_out=. */*.proto
带引用：
protoc --proto_path=${GOPATH}/src:. --go_out=. --micro_out=. api/api.proto 


### micro api 命令
micro api --handler=api
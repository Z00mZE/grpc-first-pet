# grpc-first-pet

---
### Дока по `ProtoBuff`
https://developers.google.com/protocol-buffers/docs/gotutorial

---

### Дока по `gRPC`
https://grpc.io/docs/languages/go/quickstart/

---

Установка компонентов кодогенерации `protobuff`
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Кодогенерация по IDL proto
```shell
$ protoc --go_out=. --go-grpc_out=. api/proto/adder.proto
```

---

### Evans (gRPC-клиент)
https://github.com/ktr0731/evans

```shell
$ evans api/proto/adder.proto -p 8080

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___ 
 |  __|   ' ' / /  / _. | | '_ '  / __|
 | |____   ' V /  | (_| | | | | | '__ ,
 |______|   '_/    '__,_| |_| |_| |___/

 more expressive universal gRPC client


api.Adder@127.0.0.1:8080> call Add
x (TYPE_INT32) => 12
y (TYPE_INT32) => 354
{
  "result": 366
}
```
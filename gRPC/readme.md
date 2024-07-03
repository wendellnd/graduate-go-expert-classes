#### Generate pb files

```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

#### Test gRPC server with evans client

```
evans -r repl
```

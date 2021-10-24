# Learning Practical GO GRPC

Trying to follow (without downloading code) https://subscription.packtpub.com/book/web_development/9781839211744/1

## Some Findings

### Command Generator

In my local machine, this generator command did works

```
protoc --go_out=. --go-grpc_out=. ./proto/sfapi.proto
```

Refs: https://github.com/golang/protobuf/issues/1070#issuecomment-607465055


### Unimplented mustEmbedUnimplementedStarfriendsServer()

Refs: https://stackoverflow.com/a/69302744

After service code generation, we need to adjust this private method implementation, `mustEmbedUnimplementedStarfriendsServer()`.

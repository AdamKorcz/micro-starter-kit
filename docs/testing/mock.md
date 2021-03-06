# Mock

Running mock gRPC server from proto files for testing.

### Prerequisites

```bash
# gRPC mock server for testing
yarn global add bloomrpc-mock
# bloomrpc is a UI client for gRPC
# install `bloomrpc` via `brew` into ~/Applications)
brew cask install --appdir=~/Applications bloomrpc
```

> use certs generated from [mtls](mtls.md)

### Run

```bash
bloomrpc-mock service/greeter/proto/greeter/greeter.proto
# Or
bloomrpc-mock e2e/account.bloomrpc.proto \
-r config/certs/ca.crt \
-k config/certs/micro.key,config/certs/micro.crt \
-i ~/go/src  -i /usr/local/Cellar/protobuf/3.11.2/include \
-i ~/Developer/Work/go/micro-starter-kit
```

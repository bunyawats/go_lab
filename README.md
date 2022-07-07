

```shell
$ go get -u github.com/golang/mock/gomock
$ go install github.com/golang/mock/gomock
๓ mockgen
```


```shell
$ mockgen -destination=./lab/gomock_test/mocks/mock_doer.go -package=mocks go_lab/lab/gomock_test/doer Doer
```


```shell
$ cd lab/webassembly 
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js"
```

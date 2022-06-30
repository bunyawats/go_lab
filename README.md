

```shell
$ go get -u github.com/golang/mock/gomock
$ go install github.com/golang/mock/gomock
๓ mockgen
```


```shell
$ mockgen -destination=./lab/gomock_test/mocks/mock_doer.go -package=mocks go_lab/lab/gomock_test/doer Doer
```

# Golang Sample

## Directory structure

```
jsonserver
├── cmd
│   ├── api.go
│   ├── root.go
│   ├── sample.go
│   └── ...
├── internal // for apps
│   ├── api
│   │   ├── errors
│   │   ├── schemas
│   │   ├── storages
│   │   ├── transport
│   │   └── ...
│   └── another_packages
│       └── ....
├── pkg
│   ├── config
│   ├── databases
│   ├── errors
│   ├── healthcheck
│   ├── models
│   └── utils
│       ├── string
│       │   ├── hash.go
│       │   └── ...
│       └── ...
├── scripts
│   ├── golangci-lint.sh
│   └── ....
├── vendor
├── main.go
├── README.MD
└── ...
```

## Getting started

### Install Dependencies

From the project root, run:

```shell
go build ./...
go test ./...
go mod tidy
```


### Run dev

```shell
go run main.go api
```

Format swag comments on Go code
```shell
go fmt
```

### TODO
- [ ] Add more tests
- [ ] Add ci lint && reviewdog
- [ ] Add more examples

#!/bin/bash

(
  cd client
  npm install --loglevel=error
  npm run build
)

go get github.com/go-delve/delve/cmd/dlv

go generate -tags=dev ./pkg/...

go build -o goservd -gcflags="all=-N -l" cmd/goservd/main.go
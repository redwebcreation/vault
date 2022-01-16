# Vault

[![Tests status badge](https://github.com/redwebcreation/vault/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/redwebcreation/vault/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/redwebcreation/vault)](https://goreportcard.com/report/github.com/redwebcreation/vault)
[![Codebeat report](https://codebeat.co/badges/f80cdfad-f751-483f-9ff3-f9642e65ed50)](https://codebeat.co/projects/github-com-redwebcreation-vault-main)
[![Codecov badge](https://codecov.io/gh/redwebcreation/vault/branch/main/graph/badge.svg?token=BV9ZbA0vdg)](https://codecov.io/gh/redwebcreation/vault)

Vault provides a simple file-based key-value store for secrets.

## Installation

```
go get github.com/redwebcreation/vault
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/redwebcreation/vault"
)

func main() {
	v, _ := vault.NewVault("/path/to/vault")

	_ = v.Set("key", []byte("the-secret"), []byte("super-secret-password"))
	keyContents, _ := v.Get("key", []byte("super-secret-password"))

	fmt.Printf("%s", keyContents) // prints "the-secret"

	_ = v.Delete("key")

	v.Has("key") // returns false
}
```

## Testing

```bash
go test -v ./...
```

**vault** was created by [Félix Dorn](https://twitter.com/afelixdorn) under the [MIT license](LICENSE)

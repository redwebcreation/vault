# Vault

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

	// always returns nil
	_ = v.Delete("key")

	v.Has("key") // returns false
}
```

## Testing

```bash
go test -v ./...
```

**vault** was created by [Félix Dorn](https://twitter.com/afelixdorn) under the [MIT license](LICENSE)
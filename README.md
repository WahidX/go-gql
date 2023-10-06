1. Create a file `tools.go` with the below content.

```go
//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
)
```

2. Update mods

```sh
go mod tidy
```

3. Generate gql config.

```sh
go run github.com/99designs/gqlgen init
```

4. Now add your own `.graphqls` file in your own folder and modify the `gqlgen.yml` file like this.

```

```

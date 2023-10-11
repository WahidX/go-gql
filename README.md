# go-gql

> [WIP]

A sample graphql server using gqlgen.

1. Create a file `tools.go` with the below content.

   ```go
   //go:build tools
   // +build tools

   package tools

   import (
       "github.com/99designs/gqlgen"
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

4. Now add your own `.graphqls` file in your own folder and modify the `gqlgen.yml` accordingly.
   For this example added `schema.graphqls` file inside `todox` folder and adjusted the `gqlgen.yml` accordingly.

## Notes

- To introduce custom scaler datatypes(primitive datatypes) we have to mention them in gqlgen.yml like this.

  ```yml
  TimeStamp:
    model: github.com/wahidx/go-gql/todox/model.Timestamp
  ```

  And we have to implement the datatype in this file.\
  `github.com/wahidx/go-gql/todox/model.Timestamp`

- To Add persistent store.
  To add a persistent store which can be any DB or a volatile data store. We can add it in the `Resolver{} struct` inside `resolver.go`.

  ```go
  type Resolver struct {
    Store store.Store
  }
  ```

  It's called dependency injection.

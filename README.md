# go-gql

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

## Folder Structure

```txt
   ├── model            [contains types/structs]
   │   ├── Timestamp.go
   │   └── models_gen.go
   ├── resolvers        [contains resolvers - have access to service]
   │   ├── resolver.go
   │   └── schema.resolvers.go
   ├── schema.graphqls  [graphql schema]
   └── store            [persistent data layer]
       ├── store.go
       ├── todo.go
       └── user.go
```

So the app has 2 layers.

1. Resolvers (takes inputs and contains biz logic)
2. Store (data layer which accesses the database or other persistent data store)

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

- GraphQL endpoint
  - So a graphql endpoint acts like a normal rest endpoint with POST method.\
    `"/graphql" -> graphql handler`
  - So we can still send normal request headers with it.
  - And we can return response with normal http status codes.
-

## Concepts

- Schema Definition Language(SDL): A simple syntax for describing your data types, queries, mutations, subscriptions.
- 3 Main graphql types
  - Query
  - Mutation
  - Subscription
- Resolvers: functions that responds to graphql requests
- Queries and Mutations: Queries to fetch data, Mutations to do some write operation on data.
- Fragments are a piece of logic that can be shared between multiple queries and mutations.
  Example:
  While Querying users we can use NameParts fragments like this:

  ```graphql
  fragment NameParts on User {
  	firstName
  	lastName
  }

  query GetUsers {
  	users {
  		age
  		id
  		...NameParts
  	}
  }
  ```

## Things to learn

- Fragments: To reuse parts of your queries and keep them DRY (Don't Repeat Yourself).
- Variables: Know how to use variables to pass dynamic values to your GraphQL queries and mutations
- Alias: Learn how to use aliases to request multiple fields with the same name from a single query
- Directives: Understand GraphQL directives like @include and @skip to conditionally include or skip fields in a query
-

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

- ## Schema Definition Language(SDL)

  A simple syntax for describing your data types, queries, mutations, subscriptions.

- ## 3 Main graphql types

  - Query
  - Mutation
  - Subscription

- ## Resolvers

  Resolver Functions that responds to graphql requests

- ## Queries and Mutations

  Queries to fetch data, Mutations to do some write operation on data.

- ## Fragments

  These are a piece of logic that can be shared between multiple queries and mutations.
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

- ## Variables

  In graphql variables are used in parameterized queries and mutations making them more dynamic and reusable.

  ```graphql
  query GetUsers($limit: Int = 10) {
  	users(limit: $limit) {
  		id
  		name
  	}
  }
  ```

  In this case if no value is passed for limit, 10 is the default value.

- ## Directives:

  Directives in GraphQL provide a way to describe alternate runtime execution and type validation behavior in a GraphQL document. They are prefixed with an @ symbol and can appear after almost any form of syntax within the GraphQL query or schema languages.

  ### Common Use Cases

  - #### Conditional Field Inclusion

    Using @include and @skip directives to conditionally include or skip fields.

    ```graphql
    query getUser($includeEmail: Boolean!) {
    	user {
    		id
    		name
    		email @include(if: $includeEmail)
    	}
    }
    ```

  - #### Deprecation

    Marking fields as deprecated with @deprecated.

    ```graphql
    type User {
    	id: ID!
    	oldField: String @deprecated(reason: "Use newField.")
    	newField: String
    }
    ```

  - #### Server-side Directives:

    Custom directives defined on the server can perform various tasks like authorization, formatting, or other domain-specific logic.

    ```graphql
    type Query {
    	posts: [Post] @auth(requires: ADMIN)
    }
    ```

  - #### Unions

    We can declare types with union of multiple types. In result also we can use the new union type. If a query returns a union field then using `__typename` we should mention the schema of the result. [see docs](https://www.apollographql.com/docs/apollo-server/schema/unions-interfaces)

  - #### interfaces

    We can declare interfaces with children types inside. If a type imeplements an interface type then it must include all the children types mentioned in the interface.

    ```graphql
    interface Book {
    	title: String!
    	author: Author!
    }

    type Textbook implements Book {
    	title: String!
    	author: Author!
    	courses: [Course!]!
    }

    type ColoringBook implements Book {
    	title: String!
    	author: Author!
    	colors: [String!]!
    }

    type Query {
    	books: [Book!]!
    }
    ```

    While resolving we can again use the `__typename` to identify the proper object.

    ```graphql
    query GetBooks {
    	books {
    		# Querying for __typename is almost always recommended,
    		# but it's even more important when querying a field that
    		# might return one of multiple types.
    		__typename
    		title
    		... on Textbook {
    			courses {
    				# Only present in Textbook
    				name
    			}
    		}
    		... on ColoringBook {
    			colors # Only present in ColoringBook
    		}
    	}
    }
    ```

    [see docs](https://www.apollographql.com/docs/apollo-server/schema/unions-interfaces)

## Things to learn

- Fragments: To reuse parts of your queries and keep them DRY (Don't Repeat Yourself).
- Variables: Know how to use variables to pass dynamic values to your GraphQL queries and mutations
- Alias: Learn how to use aliases to request multiple fields with the same name from a single query
- Directives: Understand GraphQL directives like @include and @skip to conditionally include or skip fields in a query

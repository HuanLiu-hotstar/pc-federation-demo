# federation example

## build the server

```sh
create the project use create-project.sh file 
```

## define federation schema

```graphql

type Content @key(fields: "ID"){
  ID: ID!
  name: String!
}

type Query {
  GetContent(id: ID!):Content
}
```

## regenerate project to support federation

- configure the federation code generate file
- regenerate the project using `regenerate.sh`

```yml
# in gqlgen.yml file do the following action
# Uncomment to enable federation
federation:
  filename: graph/generated/federation.go
  package: generated

```

## resolver to find the resource

```go
// we should implement the method FindContentByID in entity.resolvers.go 
func (r *entityResolver) FindContentByID(ctx context.Context, id string) (*model.Content, error)

```

## start server

```go
go run server.go
```

# federation example

## build the server

```sh
create the project use create-project.sh file 
```

## define federation schema

```graphql

type Review {
  ID: ID!
  rating: Int!
  content: [Content] 
}

## Content defined in other subgraph
extend type Content @key(fields: "ID") {
  ID: ID! @external
}
input ReviewInput {
    id: ID!
}
type Query {
  GetReview(input: ReviewInput):Review
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
// we should implement the method GetReview in schema.resolvers.go
// model.Review should return with Content in schema definition
func (r *queryResolver) GetReview(ctx context.Context, input *model.ReviewInput) (*model.Review, error) 
```
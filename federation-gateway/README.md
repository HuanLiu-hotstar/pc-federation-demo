# federation example

- configure the endpoint for graphql server

```js
// configure the serviceList when new Apollo Gateway
serviceList: [
    { name: 'federation-one', url: 'http://localhost:8081/query' },
    { name: 'deferation-two', url: 'http://localhost:8082/query' },
  ]
```

## start the server

```
node index.js
```

## open playground to send query 

```graphql
# Write your query or mutation here
{
 	GetReview(input:{id:"100"}) {
    ID
    rating
    content{
      ID
      name
    }
  }
}
```

## repsonse for query 

```js
{
  "data": {
    "GetReview": {
      "ID": "100",
      "rating": 699,
      "content": [
        {
          "ID": "917",
          "name": "917_name"
        },
        {
          "ID": "252",
          "name": "252_name"
        },
        {
          "ID": "556",
          "name": "556_name"
        }
      ]
    }
  }
}

```
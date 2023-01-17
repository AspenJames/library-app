# Library app server

Directory containing the backend code for the library app.

## Go Fiber API

There is a [Go Fiber][gofiber] API app defined in `main.go`. Database connections, migrations, and queries are managed through [GORM][gorm].

## API At-A-Glance

* Create a book
```sh
$ curl -X POST \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d '{ "title": "Example Book", "authors": [{ "name": "Example Author" }] }' \
    localhost:8080/books | jq .
{
  "book": {
    "id": "8cc91fcf-6829-4a9d-8ff6-524ba3773422",
    "created_at": "2023-01-17T05:54:28.062585415Z",
    "updated_at": "2023-01-17T05:54:28.062585415Z",
    "title": "Example Book",
    "authors": [
      {
        "id": "81d40a20-9e16-44e0-8e11-d118109623ef",
        "created_at": "2023-01-17T05:54:27.997470655Z",
        "updated_at": "2023-01-17T05:54:27.997470655Z",
        "name": "Example Author"
      }
    ]
  }
}
```

* Get all books
```sh
$ curl localhost:8080/books | jq .
{
  "books": [
    {
      "id": "8cc91fcf-6829-4a9d-8ff6-524ba3773422",
      "created_at": "2023-01-17T05:54:28.062585Z",
      "updated_at": "2023-01-17T05:54:28.062585Z",
      "title": "Example Book",
      "authors": [
        {
          "id": "81d40a20-9e16-44e0-8e11-d118109623ef",
          "created_at": "2023-01-17T05:54:27.99747Z",
          "updated_at": "2023-01-17T05:54:27.99747Z",
          "name": "Example Author"
        }
      ]
    }
  ]
}
```
* Get a book by ID
```sh
$ curl localhost:8080/books/8cc91fcf-6829-4a9d-8ff6-524ba3773422 | jq .
{
  "book": {
    "id": "8cc91fcf-6829-4a9d-8ff6-524ba3773422",
    "created_at": "2023-01-17T05:54:28.062585Z",
    "updated_at": "2023-01-17T05:54:28.062585Z",
    "title": "Example Book",
    "authors": [
      {
        "id": "81d40a20-9e16-44e0-8e11-d118109623ef",
        "created_at": "2023-01-17T05:54:27.99747Z",
        "updated_at": "2023-01-17T05:54:27.99747Z",
        "name": "Example Author"
      }
    ]
  }
}
```

* Update a book by ID
```sh
$ curl -X POST \
    -H "Accept: application/json" \
    -H "Content-Type: application/json" \
    -d '{ "title": "Updated Title", "authors": [{ "name": "Updated Author" }] }' \
    localhost:8080/books/8cc91fcf-6829-4a9d-8ff6-524ba3773422 | jq .
{
  "book": {
    "id": "8cc91fcf-6829-4a9d-8ff6-524ba3773422",
    "created_at": "2023-01-17T05:54:28.062585Z",
    "updated_at": "2023-01-17T06:01:47.880325362Z",
    "title": "Updated Title",
    "authors": [
      {
        "id": "91ab4f72-c0f4-4f90-ab7b-1fbc6f27291e",
        "created_at": "2023-01-17T06:01:47.873754939Z",
        "updated_at": "2023-01-17T06:01:47.873754939Z",
        "name": "Updated Author"
      }
    ]
  }
}
```

* Delete a book by ID
```sh
$ curl -X DELETE localhost:8080/books/8cc91fcf-6829-4a9d-8ff6-524ba3773422 | jq .
{}
```

* Get all authors
```sh
$ curl localhost:8080/authors | jq .
{
  "authors": [
    {
      "id": "81d40a20-9e16-44e0-8e11-d118109623ef",
      "created_at": "2023-01-17T05:54:27.99747Z",
      "updated_at": "2023-01-17T05:54:27.99747Z",
      "name": "Example Author"
    },
    {
      "id": "91ab4f72-c0f4-4f90-ab7b-1fbc6f27291e",
      "created_at": "2023-01-17T06:01:47.873754Z",
      "updated_at": "2023-01-17T06:01:47.873754Z",
      "name": "Updated Author"
    }
  ]
}
```

* Get an author by ID
```sh
$ curl localhost:8080/authors/81d40a20-9e16-44e0-8e11-d118109623ef | jq .
{
  "author": {
    "id": "81d40a20-9e16-44e0-8e11-d118109623ef",
    "created_at": "2023-01-17T05:54:27.99747Z",
    "updated_at": "2023-01-17T05:54:27.99747Z",
    "name": "Example Author"
  }
}

```

### Database models

GORM model files are defined in the `models/` directory.

### Database migrations

Models listed in the [models/models.go][models_file] file will be auto-migrated
upon app initialization.

## Local development

There is a docker compose config and Dockerfile for local development. The env
variables defined in [.env][.env] are used by compose, where the necessary
values are piped to the appropriate services. Start a local cluster with
`docker compose up --build dev`.

[.env]: ./.env.sample
[gofiber]: https://gofiber.io/
[gorm]: https://gorm.io/
[models_file]: ./models/models.go

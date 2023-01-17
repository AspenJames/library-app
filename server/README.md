# Library app server

Directory containing the backend code for the library app.

## Go Fiber API

There is a [Go Fiber][gofiber] API app defined in `main.go`. Database connections, migrations, and queries are managed through [GORM][gorm].

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

# Gusher backend

Contains protocols service, will hopefully evolve to replace legacy backend api

## Running application

* Build with `go build -o gusher-backend` inside `/cmd/protocol` directory
* run with `./gusher-backend your.config.yml`
* swagger docs are available on following url: `/swagger/index.html`

## Updating swagger docs

At this moment swagger doc tags are not processed automatically during deployment, in order to update them follow these steps
* Install swaggo from https://github.com/swaggo/swag
* go to `/cmd/backend` and run `swag init --pd`

## Generating gorm models with gentool

````
gentool -dsn "root:root@tcp(localhost:3306)/gusher" -tables "protocol_task_startup" -fieldNullable -onlyModel -outPath "./protocol/protocol"  -modelPkgName "progress"
````

## Configuration



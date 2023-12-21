# Overview

Architecture explained 

## Routes, requests and responses

All HTTP related code is in `/cmd/backend` directory. It is convenient to keep routes, handlers, 
requests and responses together and group them by module. for example `/cmd/backend/protocol.go`, `/cmd/backend/startup.go`

All request handlers, should have appropriate swagger godoc comments

Handlers are responsible for processing user input, generating responses and wiring application blocks together. 
Avoid putting complex logic other than simple binding and struct initialization inside handlers. Also avoid accessing database
directly from handlers

## Business logic  

Application logic should exist inside modules (`/protocol/module.go`) and inside entities (`/protocol/protocol/proto.go`)

Design structs for entities with encapsulation in mind. Structs should define rules on how they can be manipulated. For example, 
it is usually bad idea to modify protocol ID, so this field is made readonly by making it private inside struct entity and 
exposing Id() method outside. For the same reason such entities are kept inside submodules.

Keep modules loose coupled. it is usually a bad idea to address one module from inside another. There are some exceptions to
this rule - `/access`, `/shared` access module has to work with multiple modules and call their methods. But you should not do
the opposite and call access from inside your modules. Shared module used to store structs and functions that may be needed in multiple
modules.

## Database access

Modules should not directly access database, instead they use repositories that are defined in `/repo` directory. Repositories
Are responsible for loading and saving entities. 

We are using GORM (https://gorm.io) as ORM, but we don't bind data directly to entities from DB. 
Instead, each entity has `ToDto()` and `FromDto()` functions that are responsible for passing data between entities and ORM 
structures. This adds additional layer of complexity to the code and forces us to repeat ourselves when creating structs, but 
its is needed for encapsulation. For example, we should prevent startups without signed agreement to change status to published, and
by making status field directly inaccessible we can achieve that

## Migrations 

Not implemented yet

## Unit testing

Not implemented, but we need it 


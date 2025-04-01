# Gomancer

⚠️ **Note:** Building in process. I hope to be able to create many things to deliver v1.0.0 ⚠️

Code generator to be BLAZINGLY! fast at coding API

Create your new project with `gomancer new <project_name>`
Create a new resource with `gomancer gen <model_name> <attribute:type[:optional]>`

Like:

`gomancer new github.com/user/great_api`

and

`gomancer gen User name:string age:int8 dob:time:optional --pk-uuid`

Pum! You're ready to go!

## Install

`go install github.com/manicar2093/gomancer@latest`

## Why?

I started with the following question: 

**Do we really need a fullstack framework for Golang?** 

I think Go has many powerful packages to handle many project's sizes, but there is no way to create API in a fast way.

Once I tried to lear Elixir and Phoenix and I was shocked by phx.gen tool. You can get a usable API in milliseconds, even you can create CRUDs in a faster way just by typing a CLI command. **Why not take this to Go?**

We got everything: ORM, HTTP Servers, UI Libraries for front end, Validation packages..._EVERYTHING!_. We just need to gather them all in one place ready to be used.

That's why Gomancer borned...

## What is our goal?

- [X] Start a project
- [X] Create controller
- [ ] Create controller testing
- [X] Create models
- [X] Create migrations with Prisma
- [X] Create crud repository
- [ ] Create crud testing
- [X] Create all at once: controller, model, migration and repository
- [ ] Create API documentation
- [ ] Create auth implementation
- [ ] Add crud HTML templates
- [ ] Add API testing

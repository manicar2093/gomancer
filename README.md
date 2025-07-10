# Gomancer

<div align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/manicar2093/gomancer" alt="Go Version">
  <img src="https://img.shields.io/github/v/release/manicar2093/gomancer" alt="Latest Release">
  <img src="https://img.shields.io/github/license/manicar2093/gomancer" alt="License">
</div>

⚠️ **Note:** Building in process. I hope to be able to create many things to deliver v1.0.0 ⚠️

## Overview

Gomancer is a powerful code generator designed to make API development in Go **BLAZINGLY** fast! It automates the creation of models, controllers, repositories, and database migrations, allowing you to focus on your business logic rather than boilerplate code.

Inspired by tools like Phoenix's generators in Elixir, Gomancer brings rapid API development to the Go ecosystem by leveraging existing powerful packages and organizing them in a cohesive way.

## Quick Start

Create a new project:
```bash
gomancer new github.com/user/great_api
cd great_api
go mod tidy
npm i  # for Prisma dependencies
```

Generate a complete API resource:
```bash
gomancer gen User name:string age:int8 dob:time:optional --pk-uuid
```

And just like that, you're ready to go with a fully functional API!

## Installation

```bash
go install github.com/manicar2093/gomancer@latest
```

## Features

### Project Initialization
Gomancer creates a well-structured project with:
- Echo web framework configuration
- Prisma for database migrations
- GORM for database access
- Environment configuration
- GitHub workflow (bump_version) for automatic semantic versioning and changelog generation
- Air for hot reloading
- And more!

### Code Generation
Generate complete API resources with a single command:
- Models with validation
- Controllers with CRUD endpoints
- Repositories with database operations
- Database migrations

### Supported Data Types
Gomancer supports a wide range of data types:
- Basic types: int, int8, int16, int32, int64, float32, float64, string, bool
- Complex types: time, decimal, uuid
- Enums (with custom values)
- Optional fields

### Primary Key Options
- Auto-increment integer (default)
- UUID (with `--pk-uuid` flag)

## Usage Examples

### Creating a New Project
```bash
gomancer new github.com/username/awesome-api
```

### Generating Resources

Basic resource with auto-increment ID:
```bash
gomancer gen Client name:string email:string age:int
```

Resource with UUID primary key:
```bash
gomancer gen Product name:string price:decimal stock:int --pk-uuid
```

Resource with optional fields:
```bash
gomancer gen User name:string bio:string:optional avatar:string:optional
```

Resource with enum field:
```bash
gomancer gen Order status:enum/pending/processing/shipped/delivered total:decimal
```

## Project Structure

When you create a new project with Gomancer, it generates the following structure:

```
<project_name>/
├── .air.toml                # Configuration for hot reloading
├── .cz.toml                 # Commitizen configuration
├── .env                     # Environment variables
├── README.md                # Project documentation
├── Taskfile.yml             # Task runner configuration
├── go.mod                   # Go module definition
├── package.json             # Node.js dependencies (for Prisma)
├── .github/
│   └── workflows/
│       └── bump_version.yml  # GitHub Actions workflow for automatic version bumping
├── cmd/
│   └── api/
│       ├── main.go          # Application entry point
│       └── controllers/
│           └── init.go      # Controllers initialization
├── internal/
│   └── domain/
│       └── models/
│           └── init.go      # Domain models initialization
├── pkg/
│   ├── config/
│   │   └── config.go        # Configuration settings
│   ├── generators/
│   │   └── generators.go    # Code generation utilities
│   └── versioning/
│       └── version.go       # Version management
└── prisma/
    └── schema/
        └── schema.prisma    # Database schema definitions
```

## Why Gomancer?

I started with the following question: 

**Do we really need a fullstack framework for Golang?** 

I think Go has many powerful packages to handle projects of all sizes, but there was no way to create APIs quickly and efficiently.

Once I tried to learn Elixir and Phoenix, I was shocked by the `phx.gen` tool. You can get a usable API in milliseconds, even create CRUDs with just a CLI command. **Why not bring this to Go?**

We have everything: ORMs, HTTP Servers, UI Libraries, Validation packages..._EVERYTHING!_ We just need to gather them all in one place, ready to be used.

That's why Gomancer was born.

## About types

Now the only thing to notice is all dates are handled as UTC and show as Local when data is shown in a HTML input :/. There is an issue [#27](https://github.com/manicar2093/gomancer/issues/27) to review this.

## Roadmap

- [X] Start a project
- [X] Create controller
- [ ] Create controller testing
- [X] Create models
- [X] Create migrations with Prisma
- [X] Create CRUD repository
- [ ] Create CRUD testing
- [X] Create all at once: controller, model, migration and repository
- [ ] Create API documentation
- [ ] Create auth implementation
- [ ] Add CRUD HTML templates
- [ ] Add API testing

## License

This project is licensed under the MIT License - see the LICENSE file for details.

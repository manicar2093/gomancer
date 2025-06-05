## v0.17.0 (2025-06-05)

### Feat

- add enum to cli docs

## v0.16.0 (2025-06-04)

### Feat

- change enum separator to /
- change enum parsing separator
- add enum in migration generation
- add enum in repository generation
- add enum in model generation and greater testfixture
- add enum data type to parser and change parsing implementation

### Fix

- add model package name to plural
- add controller and model files to  pluralized names
- add controller endpoint pluralized
- add conditional to generate model id and make difference on uuid and int

## v0.15.0 (2025-05-24)

### Feat

- add golang version 1.24 to generated code on boostrap package

## v0.14.0 (2025-05-24)

### Feat

- add deps implementation gorm repository package
- add deps implementation in models generator
- add new godeps to controllers package
- add new godeps to bootstrap package
- add first deps handling approach

### Refactor

- add better way to build dependencies without fmt.Sprintf

## v0.13.0 (2025-05-23)

### Feat

- add version 1.24 and ginkgo as go tool

## v0.12.0 (2025-05-23)

### Feat

- add package entity name to be shared on generation
- add package pluralization on repo generation

### Fix

- add passing isPkUuid argumento to parseArgs

## v0.11.0 (2025-05-15)

### Feat

- remove all winter references and point them to core package
- add validation on SelectiveUpdate generation to avoid empty returns
- remove code for prisma migrations due schema folder requirements
- change logger implementation delivering a Config method

## v0.10.0 (2025-04-17)

### Feat

- add core packages to init project

## v0.9.0 (2025-04-02)

### Feat

- add pluralized to migration table name
- add prisma migration creation on gen cmd
- change name for package creation

### Fix

- change package name into controller generation
- change download to tidy on bootstrap cmd success log

## v0.8.0 (2025-04-01)

### Feat

- add new package name for repository creation
- add LowerNoSpaceCase to name packages
- add time type handling

### Fix

- add changes to testingI for generator on bootstrap command
- add id type generator for controller
- add constructor generation for controller
- add clause import in repository
- change controller imports and method calls

## v0.7.0 (2025-03-28)

### Feat

- add check to readme task
- add log to bootstrap and gen cmd
- add info logs
- add parser args
- move cmd config to its package
- WIP add valid types map
- WIP add gen cmd
- change bootstrap cmd to new
- move version to its own package

### Fix

- add package directory creation

## v0.6.0 (2025-03-27)

### Feat

- add controller generation
- add log for repository creation

## v0.5.0 (2025-03-25)

### Feat

- add repo missing functions
- add repo struct, constructor and save method

### Refactor

- add model qualifier to generator data struct

## v0.4.0 (2025-03-24)

### Feat

- add prisma migration generation
- add supported types to domain

## v0.3.0 (2025-03-23)

### Feat

- add a new way to handle paths generation

## v0.2.0 (2025-03-22)

### Feat

- add model and generator code creation

## v0.1.0 (2025-03-21)

### Feat

- add readme for project
- add readme file for bootstrap
- add integration test for bootstrap result
- add install command
- add cobra with root and bootstrap commands
- add air file for development
- add taskfile to init project
- change .env.dev to just .env
- add taskfile to project
- add go.mod creation
- add GoVersion to templates execution
- add .github/workflows to bump version with commitizen
- move all templates to its own file to improve readability
- add versioning to init project
- add some file to start a project

### Fix

- add changes to templates

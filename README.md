# go-echo-graphql

## Adding new dependencies to wire
1. Create the factory method (e.g NewHealthHandler)
2. Import on `wire_{env}.go` and setup in the right call order on `wire.Build(...)`
3. At the terminal, execute `wire` command in the root folder.
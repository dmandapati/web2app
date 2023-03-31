# web2app

## VSCODE steps before we start coding

### Created the web2app repo on github and clone it locally and open it will VSCODE

### Initialise the project

```bash
go mod init web2app
```

### add dependence for CompileDaemon for Go

```bash
go get github.com/githubnemo/CompileDaemon
```

### also install Compile Daemon so we can run via command line

```bash
go install github.com/githubnemo/CompileDaemon
```

### add dependence for environment

```bash
go get github.com/joho/godotenv
```

### add dependency for gin web framework

```bash
go get github.com/gin-gonic/gin
```

### add dependency for database connection  

```bash
go get -u gorm.io/gorm
```

### add dependency for Postgres database in gorm

```bash
go get -u gorm.io/driver/postgres
```

### Before we start coding execute CompileDaemon so that it will build the package in backgroud

```bash
CompileDaemon -command="./web2app"
```

## We are using elephantsql.com for free postgres instance

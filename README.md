# web2app

## VSCODE steps before we start coding

### Created the web2app repo on github and clone it locally and open it will VSCODE

### Initialise the project

```bash
go mod init web2app
```

### add dependence for Compile Daemon for Go

```bash
go get github.com/githubnemo/CompileDaemon
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

## We are using elephantsql.com for free postgres instance

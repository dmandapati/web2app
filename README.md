# web2app

VSCODE steps before we start coding
# Initialise the project
go mod init example/web2app

# add dependence for Compile Daemon for Go
go get github.com/githubnemo/CompileDaemon

# add dependence for environment 
go get github.com/joho/godotenv

# add dependency for gin web framework 
go get github.com/gin-gonic/gin

# add dependency for database connection  
go get -u gorm.io/gorm

# add dependency for Postgres database in gorm
go get -u gorm.io/driver/postgres

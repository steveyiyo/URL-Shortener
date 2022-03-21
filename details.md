# Technical details

## Usage 3rd Library

- Web Engineering
  - Use gin for web framework (github.com/gin-gonic/gin)
  - Use cors for cross-domain requests. It will let you make requests from any domain. (github.com/gin-contrib/cors)
- Config file
  - Use viper to parse the config file (github.com/spf13/viper), let the code be more flexible.
- Database
  - Use sqlite3 for database (github.com/mattn/go-sqlite3)
- Cache Storage [Still Working on it]
  - Use Redis as cache storage. (github.com/go-redis/redis/v8)

For the full list, please check on [go.mod](go.mod).
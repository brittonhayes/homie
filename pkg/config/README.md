<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# config

```go
import "github.com/brittonhayes/homie/pkg/config"
```

## Index

- [type Configuration](<#type-configuration>)
  - [func LoadConfig(path string) (config Configuration, err error)](<#func-loadconfig>)


## type Configuration

```go
type Configuration struct {
    File   string
    Google struct {
        Secrets string
        Sheet   struct {
            ID        string
            HeaderRow int
            Title     string
        }
    }
    Telegram struct {
        Token   string
        Allowed []string
    }
}
```

### func LoadConfig

```go
func LoadConfig(path string) (config Configuration, err error)
```

LoadConfig reads configuration from file or environment variables\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

module cooky-go

go 1.12

require (
	github.com/Unknwon/com v0.0.0-20190321035513-0fed4efef755
	github.com/casbin/casbin v1.8.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-ini/ini v1.42.0
	github.com/jinzhu/gorm v1.9.9
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20190617133340-57b3e21c3d56
	gopkg.in/ini.v1 v1.42.0 // indirect
)

replace golang.org/x/crypto v0.0.0-20190617133340-57b3e21c3d56 => github.com/golang/crypto v0.0.0-20190617133340-57b3e21c3d56

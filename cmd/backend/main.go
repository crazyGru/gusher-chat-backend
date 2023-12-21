package main

import (
	"backend/access"
	"backend/chat"
	"backend/cmd/backend/docs"
	"backend/protocol"
	"backend/repo"
	"backend/shared"
	"backend/startup"
	"backend/user"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//	@title			Gusher backend API
//	@version		3.0
//	@description	Gusher backend api

//	@BasePath	/v3

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used
//	@tag.name	admin
//	@tag.description.markdown Admin
//	@tag.name	protocol
//	@tag.description.markdown Protocol
//	@tag.name	startup
//	@tag.description.markdown Startup

func main() {
	var err error

	log.SetReportCaller(true)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <config-path>", os.Args[0])
		os.Exit(0)
	}
	conf := &Config{}
	err = conf.ReadFromFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(0)
	}

	var db *sqlx.DB
	var gormDb *gorm.DB

	dbConf := mysql.Config{
		User:                 conf.DbUser,
		Passwd:               conf.DbPass,
		Net:                  "tcp",
		Addr:                 conf.DbAddr,
		AllowNativePasswords: true,
		DBName:               conf.DbName,
		ParseTime:            true,
	}
	db, err = sqlx.Open("mysql", dbConf.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	gormDb, err = gorm.Open(gormMySQL.Open(dbConf.FormatDSN()), &gorm.Config{
		//		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	protocolModule := protocol.NewModule(
		repo.NewProtocolRepository(gormDb),
		repo.NewProtocolTaskStartupRepository(gormDb),
	)
	startupModule := startup.NewModule(repo.NewStartupRepository(gormDb),
		repo.NewRoleRepository(gormDb), repo.NewApplicantRepository(gormDb))
	userModule := user.NewModule(repo.NewUserRepository(db))
	chatModule := chat.NewModule(repo.NewRoomRepository(gormDb), repo.NewRoomMemberRepository(gormDb), repo.NewMessageRepository(gormDb))
	var idProvider access.IdProvider
	if conf.DevelopmentMode == "yes" {
		idMap := map[string]access.Identity{}
		for k := range conf.MockIdentity {
			idMap[k] = access.Identity{UserId: conf.MockIdentity[k]}
		}
		idProvider = access.NewMockApiProvider(idMap)
	} else {
		idProvider = access.NewLegacyApiProvider(conf.AuthProviderUrl)
	}
	accessModule := access.NewModule(idProvider, startupModule, protocolModule, userModule, chatModule)
	srv := webserver{
		protocol: protocolModule,
		startup:  startupModule,
		access:   accessModule,
		user:     userModule,
		chat:     chatModule,
		conf:     conf,
	}
	srv.run()
}

type webserver struct {
	chat     *chat.Module
	protocol *protocol.Module
	startup  *startup.Module
	access   *access.Module
	identity *access.Identity
	user     *user.Module
	conf     *Config
}

type Response struct {
	Success bool `json:"success"` // true for successful response, false otherwise
	Data    any  `json:"data"`    // response data. when success=false, contains error message (string)
}

type ListResponse[T any] struct {
	Items []T `json:"items"`
	Total int `json:"total"`
}

// handle webserver request and return error on fail
func ginHandler(handler func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			c.Error(err)
			c.AbortWithStatusJSON(500, err.Error())
		}
	}
}

func (w *webserver) run() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200", "https://gushdev.com", "https://gusher.co"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowCredentials = true

	r.Use(cors.New(config)) // todo add secure CORS
	r.Use(ginHandler(w.getIdentity))

	// bind route handlers for modules here
	w.routesProtocol(r)
	w.routesStartup(r)
	w.routesChat(r)

	if w.conf.SwaggerGui == "yes" {
		docs.SwaggerInfo.BasePath = "/v3"
		if w.conf.DevelopmentMode == "yes" {
			docs.SwaggerInfo.BasePath = "/"
		}
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))
	}
	r.Run()
}

func (w *webserver) getIdentity(c *gin.Context) error {
	authValue := c.GetHeader("Authorization")
	if !strings.HasPrefix(authValue, "Bearer") {
		return nil
	}
	token := strings.TrimPrefix(authValue, "Bearer ")
	identity, err := w.access.GetIdentity(token)
	if err == nil {
		w.identity = identity
	} else {
		shared.ErrorLogger.Printf("getIdentity: %s", err.Error())
	}
	return nil
}

func (w *webserver) requireUser(c *gin.Context) error {
	if w.identity == nil {
		c.Abort()
		return errors.New("login required")
	}
	return nil
}

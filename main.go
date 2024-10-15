package main

import (
	"errors"
	"fmt"
	"github.com/Octek/resource-profile-management-backend.git/api/bookings"
	"github.com/Octek/resource-profile-management-backend.git/api/experience"
	"github.com/Octek/resource-profile-management-backend.git/api/projects"
	"github.com/Octek/resource-profile-management-backend.git/api/questions"
	"github.com/Octek/resource-profile-management-backend.git/api/seed"
	"github.com/Octek/resource-profile-management-backend.git/api/skills"
	user "github.com/Octek/resource-profile-management-backend.git/api/users"
	"github.com/Octek/resource-profile-management-backend.git/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginlogrus "github.com/toorop/gin-logrus"
	"gopkg.in/matryer/try.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	log.SetReportCaller(true)
	// TODO: toggle json on when deploying
	// log.SetFormatter(&log.JSONFormatter{})
	logger = log.New()
	logger.SetReportCaller(true)
	// logger.SetFormatter(&log.JSONFormatter{})
}

func SetupDatabase(connString string) (*gorm.DB, error) {
	var db *gorm.DB
	const attempts = 5
	err := try.Do(func(attempt int) (bool, error) {
		fmt.Printf("Connecting to db, attempt %v\n", attempt)
		var err error
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
		if err == nil {
			return true, nil
		}
		fmt.Printf("failed to connect database, attempt # %v", attempt)
		fmt.Println(fmt.Errorf("error: %w", err))
		sleepTime := attempt * attempts * 2
		time.Sleep(time.Second * time.Duration(sleepTime))
		return attempt < attempts, err
	})
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	// ping the DB to ensure that it is connected
	postgresDB, _ := db.DB()
	err = postgresDB.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	// setup database connection
	db, err := SetupDatabase(utils.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("setup database connection successful")

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(ginlogrus.Logger(logger), gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// programmatically set swagger info
	//docs.SwaggerInfo.Title = "Profile Management"
	//docs.SwaggerInfo.Description = "Profile Management API"
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = utils.GetSwaggerHostUrl()
	//docs.SwaggerInfo.BasePath = "/"
	//docs.SwaggerInfo.Schemes = []string{"https", "http"}
	//docs.SwaggerInfo.InfoInstanceName = "swagger"

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK", "statusCode": http.StatusOK})
	})

	// Skill
	var skillRepo = skills.NewSkillRepositoryPostgres(db)
	skillService := skills.NewService(skillRepo)
	skills.Routes(router, skillService)

	// Experience
	var experienceRepo = experience.NewExperienceRepositoryPostgres(db)
	experienceService := experience.NewService(experienceRepo)
	experience.Routes(router, experienceService)

	// Question
	var questionRepo = questions.NewQuestionRepositoryPostgres(db)
	questionService := questions.NewService(questionRepo)
	questions.Routes(router, questionService)

	// Project
	var projectRepo = projects.NewProjectRepositoryPostgres(db)
	projectService := projects.NewService(projectRepo)
	projects.Routes(router, projectService)

	// Booking
	var bookingRepo = bookings.NewBookingRepositoryPostgres(db)
	bookingService := bookings.NewService(bookingRepo)
	bookings.Routes(router, bookingService)

	// User
	var userRepo = user.NewUserRepositoryPostgres(db)
	userService := user.NewService(userRepo)
	user.Routes(router, userService)

	seed.SeedData(userService)

	API_SERVER_PORT := os.Getenv("SERVER_PORT")
	if len(API_SERVER_PORT) == 0 {
		API_SERVER_PORT = "4001"
	}
	router.Run(fmt.Sprintf(":%v", API_SERVER_PORT))
}

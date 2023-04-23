// @Author nono.he 2023/4/23 17:49:00
package database

import (
	"fmt"
	"os"
	"testing"

	"xorm.io/xorm"

	"nonoDemo/internal/infrastructure/client"
	"nonoDemo/internal/infrastructure/database/entity"
	"nonoDemo/pkg/config"
	"nonoDemo/pkg/framework"
	"nonoDemo/pkg/utils/observability/log"
)

func TestOption(t *testing.T) {
	dbEngine, logger := initTest()
	repo := initRepoImpl(dbEngine, logger)
	//var cam BlogRepository
	cams, tal, err := repo.FindBlogByOptions(
		entity.TableName("blog_manager"),
		entity.WithID("123"),
		entity.WithName("123"),
	)
	logger.Error("err", err)
	fmt.Println("tal：", tal)
	fmt.Println(cams)
}

func initTest() (*xorm.Engine, framework.Logger) {
	/*	user := os.Getenv("MYSQL_USERNAME")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_ADDRESS")
		name := "ads-admin"
		port := 3306
		idleConns := 100
		openConns := 100
		showSql := true*/

	var configFile = "D:\\Go_WorkSpace\\src\\learning\\github_repository\\nonoDemo/configs/config.yaml"

	cfg := config.Config{}
	viperOption := &framework.ViperOption{CfPath: configFile, EnvPrefix: ""}
	_, _ = framework.LoadConfiguration(viperOption, &cfg)
	cfg.Database.Host = os.Getenv("MYSQL_HOST")
	cfg.Database.User = os.Getenv("MYSQL_USER")
	cfg.Database.Port = os.Getenv("MYSQL_PORT")
	cfg.Database.Password = os.Getenv("MYSQL_PASSWORD")
	dbEngine := client.NewXorm(cfg)

	logger := log.NewLogger(cfg.Log)
	return dbEngine, logger
}

func initRepoImpl(dbE *xorm.Engine, logger framework.Logger) *BlogRepository {
	repoImpl := NewBlogRepository(dbE, logger)
	return repoImpl
}

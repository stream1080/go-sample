package global

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/golang-migrate/migrate/v4"
// 	_ "github.com/golang-migrate/migrate/v4/database/postgres"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	"go.uber.org/zap"
// )

// func InitMigration() {

// 	if !Cfg.Migration.Enabled {
// 		return
// 	}

// 	// 构建数据库连接字符串 postgres://mattes:secret@localhost:5432/database?sslmode=disable&search_path=
// 	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
// 		Cfg.PostgreSQL.UserName,
// 		Cfg.PostgreSQL.PassWord,
// 		Cfg.PostgreSQL.Host,
// 		Cfg.PostgreSQL.Port,
// 		Cfg.PostgreSQL.DbName,
// 		Cfg.PostgreSQL.SslMode,
// 	)

// 	m, err := migrate.New(fmt.Sprintf("file://%s", Cfg.Migration.Path), dsn)
// 	if err != nil {
// 		zap.S().Panic("failed to init migration with %s", err)
// 	}

// 	if err = m.Up(); err != nil {
// 		if !errors.Is(err, migrate.ErrNoChange) {
// 			zap.S().Panic(err)
// 		}
// 		zap.L().Debug("migrations no change to apply")
// 	}
// }

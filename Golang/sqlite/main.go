package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// refer: https://www.cnblogs.com/failymao/p/17197166.html
type User struct {
	gorm.Model
	Name string
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("latency: %.2fs", time.Now().Sub(start).Seconds())
	}()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	// 创建 SQLite3 数据库连接
	db, err := gorm.Open(sqlite.Open("test.db?_journal_mode=WAL"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	// 设置连接池大小
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to set database pool size")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// 自动迁移 User 模型对应的表
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate table")
	}

	var wg sync.WaitGroup
	wg.Add(2000)
	// 并发写入 1000 条数据
	for i := 0; i < 1000; i++ {
		go func(i int) {
			defer wg.Done()
			err := db.Transaction(func(tx *gorm.DB) error {
				user := User{Name: fmt.Sprintf("user_%d", i)}
				result := tx.Create(&user)
				return result.Error
			})
			if err != nil {
				fmt.Printf("failed to write data: %v\n", err)
			}
		}(i)
	}

	// 并发读取数据
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			var users []User
			err := db.Transaction(func(tx *gorm.DB) error {
				result := tx.Find(&users)
				return result.Error
			})
			if err != nil {
				fmt.Printf("failed to read data: %v\n", err)
			}
		}()
	}

	wg.Wait()
}

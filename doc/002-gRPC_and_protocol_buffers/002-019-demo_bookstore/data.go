package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// 使用GORM
func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Shelf{}, Book{})
	return db, nil
}

// 定义Model

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

package db

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/mayuka-c/book-management-system/internal/pkg/config"
	"github.com/mayuka-c/book-management-system/pkg/log"
)

const (
	tablename = "books"
)

type DBClient struct {
	dbClient *gorm.DB
}

func Connect(dbConfig config.DBConfig) *DBClient {
	user := dbConfig.DB_USERNAME
	pass := dbConfig.DB_PASSWORD
	url := dbConfig.DB_URL

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := user + ":" + pass + "@tcp(" + url + ")/books?charset=utf8&parseTime=True&loc=Local"
	dbClient, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Creation of table based on go struct
	// It creates table with name `books`
	dbClient.AutoMigrate(&Book{})

	return &DBClient{
		dbClient: dbClient,
	}
}

func (d *DBClient) InsertBook(ctx context.Context, input Book) error {

	result := d.dbClient.Table(tablename).Create(&input)
	if err := result.Error; err != nil {
		log.Errorf(ctx, err.Error())
		return err
	}

	log.Infof(ctx, "Data added succeessfully to DB")
	return nil
}

func (d *DBClient) GetAllBooks(ctx context.Context) ([]Book, error) {

	var books []Book
	result := d.dbClient.Table(tablename).Find(&books)
	if err := result.Error; err != nil {
		log.Errorf(ctx, err.Error())
		panic(err)
	}

	fmt.Printf("Books: %#v\n", books)

	log.Infof(ctx, "Returned all books from DB succeessfully")
	return books, nil
}

func (d *DBClient) GetBookByID(ctx context.Context, id uint) (Book, error) {

	var book Book
	result := d.dbClient.Table(tablename).First(&book, id)
	if err := result.Error; err != nil {
		log.Errorf(ctx, err.Error())
		return book, err
	}

	log.Infof(ctx, "Returned book with id:%v from DB succeessfully", id)
	return book, nil
}

func (d *DBClient) UpdateBook(ctx context.Context, input UpdateBook) error {

	result := d.dbClient.Table(tablename).Save(&input)
	if err := result.Error; err != nil {
		log.Errorf(ctx, err.Error())
		return err
	}

	log.Infof(ctx, "Book with id: %v updated succeessfully in DB", input.ID)
	return nil
}

func (d *DBClient) DeleteBook(ctx context.Context, id uint) error {

	var book Book
	result := d.dbClient.Table(tablename).Delete(&book, id)
	if err := result.Error; err != nil {
		log.Errorf(ctx, err.Error())
		return err
	}

	log.Infof(ctx, "Book with id: %v deleted succeessfully from DB", id)
	return nil
}

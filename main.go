package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type Item struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 初始化数据库表格
	createTable(db)

	router := gin.Default()

	// 读取所有条目
	router.GET("/items", func(c *gin.Context) {
		items := []Item{}
		rows, err := db.Query("SELECT id, title, body FROM items")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var item Item
			err = rows.Scan(&item.ID, &item.Title, &item.Body)
			if err != nil {
				log.Fatal(err)
			}
			items = append(items, item)
		}
		c.JSON(http.StatusOK, gin.H{"items": items})
	})

	// 读取一个条目
	router.GET("/items/:id", func(c *gin.Context) {
		item := Item{}
		id := c.Param("id")
		row := db.QueryRow("SELECT id, title, body FROM items WHERE id = ?", id)
		err = row.Scan(&item.ID, &item.Title, &item.Body)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"item": item})
	})

	// 新增一个条目
	router.POST("/items", func(c *gin.Context) {
		var item Item
		if err := c.BindJSON(&item); err != nil {
			log.Fatal(err)
		}
		result, err := db.Exec("INSERT INTO items (title, body) VALUES (?, ?)", item.Title, item.Body)
		if err != nil {
			log.Fatal(err)
		}
		item.ID, err = result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"item": item})
	})

	// 更新一个条目
	router.PUT("/items/:id", func(c *gin.Context) {
		var item Item
		if err := c.BindJSON(&item); err != nil {
			log.Fatal(err)
		}
		id := c.Param("id")
		_, err := db.Exec("UPDATE items SET title=?, body=? WHERE id=?", item.Title, item.Body, id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"item": item})
	})

	// 删除一个条目
	router.DELETE("/items/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM items WHERE id=?", id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	})

	router.Run(":8080")
}

func createTable(db *sql.DB) {
	sqlTable := `
    CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        body TEXT
    )
    `
	_, err := db.Exec(sqlTable)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created")
}

package controllers

import (
	"log"
	"net/http"
	"pratice/config"
	"pratice/structs"

	"github.com/gin-gonic/gin"
)

func GetNote(c *gin.Context) {
	var notes structs.Note
	var arr_note []structs.Note

	db := config.Dbconn()
	defer db.Close()

	rows, err := db.Query("SELECT id_note, title, note FROM tb_note")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&notes.Id, &notes.Title, &notes.Note); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_note = append(arr_note, notes)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   arr_note,
	})
}

func AddNote(c *gin.Context) {

	db := config.Dbconn()
	defer db.Close()

	var input structs.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	title := input.Title
	note := input.Note
	db.Exec("INSERT INTO tb_note (title, note) VALUES (?,?)",
		title,
		note,
	)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "Success add to database",
	})

}

func UpdateNote(c *gin.Context) {
		id := c.Param("id")
		db := config.Dbconn()
		defer db.Close()
		var input structs.Note
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		title := input.Title
		note := input.Note
		db.Exec("UPDATE tb_note SET title = ? , note = ? WHERE id_note = ?",
			title,
			note,
			id,
		)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Success update database",

		})

	}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")
		db := config.Dbconn()
		defer db.Close()
		db.Exec("DELETE FROM tb_note WHERE id_note = ?",
			id,
		)
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Success delete database",

		})
}

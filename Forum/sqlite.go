package Forum

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID         int
	Pseudo     string
	Password   string
	Mail       string
	Number     int
	ProfilePic string
}

type Post struct {
	ID           int
	Content      string
	IsTopic      int
	Title        sql.NullString
	Category     sql.NullString
	ParentPostId sql.NullInt64
	UserId       int
}

func GetUserRows(rows *sql.Rows) []User {
	final := make([]User, 0)
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.ProfilePic, &u.Pseudo, &u.Mail, &u.Number, &u.Password)
		if err != nil {
			log.Fatal(err)
		}
		final = append(final, u)
		// fmt.Println(u)
	}
	fmt.Println(final)
	return final
}

func GetPostRows(rows *sql.Rows) []Post {
	final := make([]Post, 0)
	for rows.Next() {
		var u Post
		err := rows.Scan(&u.ID, &u.Content, &u.IsTopic, &u.Title, &u.Category, &u.ParentPostId, &u.UserId)
		if err != nil {
			log.Fatal(err)
		}
		final = append(final, u)
		// fmt.Println(u)
	}
	return final
}

func InitDatabase(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)

	if err != nil {
		log.Fatal(err)
	}

	sqlStmnt := `
				PRAGMA foreign_keys = ON;
				CREATE TABLE IF NOT EXISTS user (
					userId INTEGER PRIMARY KEY AUTOINCREMENT,
					profilePic STRING SECONDARY KEY,
					pseudo STRING NOT NULL,
					mail STRING,
					number INT,
					password STRING NOT NULL
				);
				CREATE TABLE IF NOT EXISTS post (
					postId INTEGER PRIMARY KEY AUTOINCREMENT,
					content STRING NOT NULL,
					isTopic INTEGER NOT NULL,
					title STRING,
					category STRING,
					parentPostId INT,
					userId INT NOT NULL,
					FOREIGN KEY (userId) REFERENCES user(userId),
					FOREIGN KEY (parentPostId) REFERENCES post(postId)
				)
				`
	_, err = db.Exec(sqlStmnt)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateUser(db *sql.DB, pseudo string, password string, mail string, number int, profilePic string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO user (pseudo, password, mail, number, profilePic) VALUES (?,?,?,?,?)`, pseudo, password, mail, number, profilePic)
	return result.LastInsertId()
}

func CreateTopic(db *sql.DB, content string, userId int, isTopic int, titre string, categorie string) (int64, error) {
	result, _ := db.Exec(`INSERT INTO post (content, userId, isTopic, title, category) VALUES (?,?,?,?,?)`, content, userId, isTopic, titre, categorie)
	return result.LastInsertId()
}

func CreatePost(db *sql.DB, content string, userId int, isTopic int, parentPostId int) (int64, error) {
	result, _ := db.Exec(`INSERT INTO post (content, userId, isTopic, parentPostId) VALUES (?,?,?,?)`, content, userId, isTopic, parentPostId)
	return result.LastInsertId()
}

func GetTopic(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table + " WHERE isTopic = 1"
	result, _ := db.Query(query)
	return result
}

// à retirer
func GetTable(db *sql.DB, table string) *sql.Rows {
	query := "SELECT * FROM " + table
	result, _ := db.Query(query)
	return result
}

func DeletePostById(db *sql.DB, id string) (int64, error) {
	result, _ := db.Exec(`DELETE FROM post WHERE postId = ?`, id)
	return result.LastInsertId()
}

// reriter à

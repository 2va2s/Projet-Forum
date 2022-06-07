package Forum

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

type Category struct {
	ID    int
	Name  string
	Color string
}

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
	Category     sql.NullInt64
	ParentPostId sql.NullInt64
	UserId       int
	Date         string
	UpVote       int
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
	}
	return final
}

func GetPostRows(rows *sql.Rows) []Post {
	final := make([]Post, 0)
	for rows.Next() {
		var u Post
		err := rows.Scan(&u.ID, &u.Content, &u.IsTopic, &u.Title, &u.Category, &u.ParentPostId, &u.UserId, &u.Date, &u.UpVote)
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
					ID INTEGER PRIMARY KEY AUTOINCREMENT,
					ProfilePic STRING SECONDARY KEY,
					Pseudo STRING NOT NULL,
					Mail STRING,
					Number STRING,
					Password STRING NOT NULL
				);
				CREATE TABLE IF NOT EXISTS category (
					ID INTEGER PRIMARY KEY AUTOINCREMENT,
					Name STRING NOT NULL,
					Color STRING NOT NULL
				);
				CREATE TABLE IF NOT EXISTS post (
					ID INTEGER PRIMARY KEY AUTOINCREMENT,
					Content STRING NOT NULL,
					IsTopic INTEGER NOT NULL,
					Title STRING,
					Category INTEGER NOT NULL,
					ParentPostId INTEGER,
					UserId INTEGER NOT NULL ,
					Date STRING NOT NULL,
					UpVote STRING NOT NULL,
					FOREIGN KEY (UserId) REFERENCES user(ID) ,
					FOREIGN KEY (ParentPostId) REFERENCES post(ID),
					FOREIGN KEY (Category) REFERENCES category(ID)
				);
				`
	_, err = db.Exec(sqlStmnt)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ParseTable(model interface{}, table string) (a string) {

	result := ""
	e := reflect.ValueOf(model)
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		if varName != "ID" {

			result += string(varName) + ", "
		}
	}
	result = result[:len(result)-2]
	return result
}

func Create(db *sql.DB, table string, model interface{}, t ...interface{}) (int64, error) {

	result := "INSERT INTO " + table + " (" + ParseTable(model, table) + ")" + " VALUES " + "("
	for i := 0; i < len(t); i++ {
		result += "?, "
	}
	result = result[:len(result)-2]
	result += ")"

	request, err := db.Exec(result, t...)
	fmt.Println(result)

	if err != nil {
		fmt.Println(err)
		return -1, nil
	}

	return request.LastInsertId()
}

func Get(db *sql.DB, table string, mode string) *sql.Rows {
	query := "SELECT * FROM " + table
	if mode == "topic" {
		query += " WHERE isTopic = 1"
	}
	result, _ := db.Query(query)
	return result
}

func DeletePostById(db *sql.DB, id string) (int64, error) {
	result, _ := db.Exec(`DELETE FROM post WHERE postId = ?`, id)
	return result.LastInsertId()
}

// reriter à

// func CreateUser(db *sql.DB, pseudo string, password string, mail string, number int, profilePic string) (int64, error) {
// 	result, _ := db.Exec(`INSERT INTO user (pseudo, password, mail, number, profilePic) VALUES (?,?,?,?,?)`, pseudo, password, mail, number, profilePic)
// 	return result.LastInsertId()
// }

// func CreateTopic(db *sql.DB, content string, userId int, isTopic int, titre string, categorie string) (int64, error) {
// 	result, _ := db.Exec(`INSERT INTO post (content, userId, isTopic, title, category) VALUES (?,?,?,?,?)`, content, userId, isTopic, titre, categorie)
// 	return result.LastInsertId()
// }

// func CreatePost(db *sql.DB, content string, userId int, isTopic int, parentPostId int) (int64, error) {
// 	result, _ := db.Exec(`INSERT INTO post (content, userId, isTopic, parentPostId) VALUES (?,?,?,?)`, content, userId, isTopic, parentPostId)
// 	return result.LastInsertId()
// }

// func GetTopic(db *sql.DB, table string) *sql.Rows {
// 	query := "SELECT * FROM " + table + " WHERE isTopic = 1"
// 	result, _ := db.Query(query)
// 	return result
// }

// func GetTable(db *sql.DB, table string) *sql.Rows {
// 	query := "SELECT * FROM " + table
// 	result, _ := db.Query(query)
// 	return result
// }

package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"pratice/config"
	"pratice/controllers"
	"pratice/structs"
	"strings"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(config))
	// User Route
	router.POST("/login", loginHandler)
	router.POST("/register", registHandler)
	router.GET("/me", auth, me)
	// Note Route
	router.GET("/note", auth, controllers.GetNote)
	router.Run(":3000")
}
func me(c *gin.Context) {
	db := config.Dbconn()
	defer db.Close()
	mytoken := c.Request.Header.Get("Authorization")
	token := strings.Replace(mytoken, "Bearer ", "", -1)
	u := structs.User{}
	db.QueryRow("SELECT username FROM tb_user WHERE token = ?;", token).Scan(&u.Username)
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"username": &u.Username,
	})
}

func registHandler(c *gin.Context) {
	db := config.Dbconn()
	defer db.Close()
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "" && password != "" {
		err := c.Request.ParseForm()
		if err != nil {
			panic(err)
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
		_, err = db.Exec("INSERT INTO tb_user (username, password) VALUES (?,?)",
			username,
			hash,
		)

		if err != nil {
			log.Print(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Success Regist :)",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Failed, Check again :(",
		})
	}

}
func loginHandler(c *gin.Context) {
	db := config.Dbconn()
	defer db.Close()
	var input structs.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := input.Username
	password := input.Password
	log.Println(username)
	u := structs.User{}
	db.QueryRow("SELECT username, password FROM tb_user WHERE username= ?;", username).Scan(&u.Username, &u.Password)
	hash := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if username != "" && password != "" {
		if u.Username == username && hash == nil {
			sign := jwt.New(jwt.GetSigningMethod("HS256"))
			claims := sign.Claims.(jwt.MapClaims)
			claims["user"] = u.Username
			token, err := sign.SignedString([]byte("Note2019_"))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				c.Abort()
			}
			db.Exec("UPDATE tb_user SET token =  ? WHERE username = ? ", token, username)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}

}

func auth(c *gin.Context) {
	mytoken := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(mytoken, "Bearer ", "", -1)
	log.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("Note2019_"), nil
	})
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"status":  http.StatusOK,
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

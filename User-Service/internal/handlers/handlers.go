package handlers

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/Sapiyulla/Orders-Manager/User-Service/internal/database"
	uniqid "github.com/Sapiyulla/Orders-Manager/User-Service/internal/tools/uuid"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type user struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var pool = database.NewPool(os.Getenv("DSN"))

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableQuote:     true,
		DisableTimestamp: true,
	})
	if os.Getenv("DSN") == "" {
		logrus.Fatalln("empty DSN string with environment variable!")
		return
	}
	pool.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS users (
            uuid VARCHAR(36),
            login VARCHAR(12),
            password VARCHAR(60),
            reg_data DATE
        )`,
	)
	logrus.Infoln("Table `users` initialized")
}

func Register(c *gin.Context) {
	conn := database.NewConnect(pool)
	defer conn.Release()

	var user user
	c.BindJSON(&user)

	if user.Login == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "`login` or `password` is empty",
		})
		return
	}

	row := conn.QueryRow(context.Background(),
		`SELECT uuid FROM users WHERE login=$1`,
		user.Login,
	)

	var newUser string
	err := row.Scan(&newUser)

	if err != nil && newUser == "" {
		if errors.Is(err, pgx.ErrNoRows) {
			// логика создания нового пользователя и присвоение уникального id
			NSuid, err := uuid.NewUUID()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			// проверяем на отсутствие uuid в бд
			uniqid.SetFreeUUID(conn, &NSuid)
			// добавление в бд
			_, err = conn.Exec(context.Background(),
				`INSERT INTO users VALUES ($1, $2, $3, $4)`,
				NSuid.String(), user.Login, user.Password, c.GetHeader("date"))
			if err != nil {
				logrus.Warnln(err.Error())
			}
			// отправка uuid в ответ на запрос пользователя
			c.Header("uuid", NSuid.String())
			c.Header("Status", "autorized")
			c.JSON(http.StatusCreated, gin.H{
				"status": "user created",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logrus.Warn(err.Error())
		return
	} else if newUser != "" {
		c.JSON(409, gin.H{
			"error": "user already exists",
		})
		return
	}
}

func Login(c *gin.Context) {
	conn := database.NewConnect(pool)
	defer conn.Release()

	var user user
	c.BindJSON(&user)

	if user.Login == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "`login` or `password` is empty",
		})
		return
	}

	row := conn.QueryRow(context.Background(),
		`SELECT uuid, password FROM users WHERE login=$1`,
		user.Login,
	)

	var UserUUID, pw string
	err := row.Scan(&UserUUID, &pw)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		logrus.Warn(err.Error())
		return
	}

	// check correct password
	if pw != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect login or password",
		})
		return
	}

	if c.GetHeader("Status") == "autorized" && c.GetHeader("uuid") == UserUUID {
		c.JSON(http.StatusOK, gin.H{
			"status": "already autorized",
		})
		return
	}
	c.Header("Status", "autorized")
	c.Header("uuid", UserUUID)
	c.JSON(http.StatusOK, gin.H{
		"status": "loged in account",
	})
}

func GetUserData(c *gin.Context) {
	uuid := c.Param("id")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty uuid field"})
		return
	}

	conn := database.NewConnect(pool)
	defer conn.Release()

	var row = conn.QueryRow(
		context.Background(),
		`SELECT uuid, login, password FROM users WHERE uuid=$1`,
		uuid,
	)

	type UserInfo = struct {
		UUID     string `json:"uuid"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	var user_info UserInfo
	err := row.Scan(&user_info.UUID, &user_info.Login, &user_info.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user with this id not found",
			})
			return
		}
		c.Status(http.StatusInternalServerError)
		logrus.Warnln(err.Error())
		return
	}

	if c.GetHeader("Role-X2CE") == "Admin-x-1" {
		c.JSON(http.StatusOK, user_info)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"login": user_info.Login,
		"uuid":  user_info.UUID,
	})
}

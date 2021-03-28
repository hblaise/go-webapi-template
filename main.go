package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoginDTO struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type LoginResponse struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
	Token    string `json:"token" form:"token" query:"token"`
}

type MessageBag struct {
	Message string `json:"message" form:"message" query:"message"`
	Type    string `json:"type" form:"type" query:"type"`
}

type BaseResponse struct {
	Success    bool         `json:"success" form:"success" query:"success"`
	Message    string       `json:"message" form:"message" query:"message"`
	MessageBag []MessageBag `json:"messageBag" form:"message_bag" query:"message_bag"`
	Data       interface{}  `json:"data" form:"data" query:"data"`
}

func login(c echo.Context) error {
	//username := c.FormValue("username")
	//password := c.FormValue("password")
	loginDTO := new(LoginDTO)

	er := c.Bind(loginDTO) // bind the structure with the context body

	if er != nil {
		panic(er)
	}

	//Get the user data
	username := loginDTO.Username
	password := loginDTO.Password

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Jon Snow"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	/*return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})*/

	var messageBag = []MessageBag{}

	/*messageBag = append(messageBag, MessageBag{
		Message: "egy üzi",
		Type:    "error",
	})*/

	/*messageBag = append(messageBag,MessageBag{
		Message: "második üzi",
		Type:    "error",
	})*/

	return c.JSON(http.StatusOK, BaseResponse{
		Success:    true,
		Message:    "Sikeres művelet",
		MessageBag: messageBag,
		Data: LoginResponse{
			Username: loginDTO.Username,
			Password: loginDTO.Password,
			Token:    t,
		},
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}

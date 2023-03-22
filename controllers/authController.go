package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/constants"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/database"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/interfaces"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/models"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/validations"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

func Register(c *fiber.Ctx) error {

	data := new(interfaces.User)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	resp, errors := validations.ValidateStruct(*data)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MaxCost)

	user := models.User{
		Name:         data.Name,
		Email:        data.Email,
		PasswordHash: passwordHash,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        strconv.Itoa(int(user.ID)),
	})

	tokenString, err := claims.SignedString(constants.SigningKey)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func Me(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	accessToken := strings.Split(headers["Authorization"], " ")[1]

	token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SigningKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	if err := database.DB.Where("ID = ?", claims.ID).First(&user).Error; err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.JSON(user)
}

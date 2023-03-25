package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/constants"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/database"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/handlers"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/initializers"
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

	resp, errorValidation := validations.ValidateStruct(*data)
	if errorValidation != nil {
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:         data.Name,
		Email:        data.Email,
		PasswordHash: passwordHash,
	}

	var exists bool

	// Check exist email
	database.DB.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", data.Email).Find(&exists)
	if exists {
		response, _ := handlers.ValidatorResponse([]interfaces.ErrorResponse{
			interfaces.ErrorResponse{
				Field:   "User.Email",
				Message: "email existed",
				Tag:     "",
			},
		})
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	// Check exist username
	database.DB.Model(&models.User{}).Select("count(*) > 0").Where("name = ?", data.Name).Find(&exists)
	if exists {
		response, _ := handlers.ValidatorResponse([]interfaces.ErrorResponse{
			interfaces.ErrorResponse{
				Field:   "User.Name",
				Message: "name existed",
				Tag:     "",
			},
		})
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err := database.DB.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var exists bool

	// Check exist email
	database.DB.Model(&models.User{}).Select("count(*) > 0").Where("email = ?", data["email"]).Find(&exists)

	if !exists {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var user models.User

	database.DB.Select([]string{"id", "password_hash", "email"}).Where("email = ?", data["email"]).First(&user)

	err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(data["password"]))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	s, er := strconv.Atoi(initializers.Config.TokenExpiresIn)
	if er == nil {
		s = 60
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s) * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        user.ID.String(),
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

package public_user_controller

import (
	publicuserdto "flower-backend/dto/public"
	"flower-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetUserByID godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve public user information by ID
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int						true	"User ID"
//	@Success		200	{object}	map[string]interface{}	"User fetched successfully"
//	@Failure		400	{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		404	{object}	map[string]interface{}	"User not found"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/{id} [get]
func (uc *userController) GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uc.svc.GetUserByID(uint(userIdUint))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("user_id", userId))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

// GetUserByEmail godoc
//
//	@Summary		Get user by email
//	@Description	Retrieve public user information by email
//	@Tags			users
//	@Produce		json
//	@Param			email	path		string					true	"User email"
//	@Success		200		{object}	map[string]interface{}	"User fetched successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		404		{object}	map[string]interface{}	"User not found"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/email/{email} [get]
func (uc *userController) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	if !utils.ValidateEmail(email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}
	user, err := uc.svc.GetUserByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("email", email))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.String("email", email))
}

// GetUserByUsername godoc
//
//	@Summary		Get user by username
//	@Description	Retrieve public user information by username
//	@Tags			users
//	@Produce		json
//	@Param			username	path		string					true	"User username"
//	@Success		200			{object}	map[string]interface{}	"User fetched successfully"
//	@Failure		400			{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		404			{object}	map[string]interface{}	"User not found"
//	@Failure		500			{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/username/{username} [get]
func (uc *userController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if !utils.ValidateUsername(username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username"})
		return
	}
	user, err := uc.svc.GetUserByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.String("username", username))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.String("username", username))
}

// GetUserAll godoc
//
//	@Summary		Get all users
//	@Description	Retrieve all public users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	map[string]interface{}	"All users fetched successfully"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/all [get]
func (uc *userController) GetUserAll(c *gin.Context) {
	users, err := uc.svc.GetUserAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": publicuserdto.ToPublicUsers(users)})
	uc.logger.Info("all users fetched successfully", zap.Int("users_count", len(users)))
}

// GetUserByIDWithSelect godoc
//
//	@Summary		Get user by ID with select
//	@Description	Retrieve public user information by ID with select fields
//	@Tags			users
//	@Produce		json
//	@Param			id		path		int						true	"User ID"
//	@Param			select	query		string					true	"Fields to update (comma-separated)"
//	@Success		200		{object}	map[string]interface{}	"User fetched successfully"
//	@Failure		400		{object}	map[string]interface{}	"Bad request - invalid input"
//	@Failure		404		{object}	map[string]interface{}	"User not found"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Securuty		BearerAuth
//	@Router			/user/id/{id}/select [get]
func (uc *userController) GetUserByIDWithSelect(c *gin.Context) {
	userId := c.Param("id")
	userIdUint, err := utils.ParseUint(userId, uc.logger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	selectFieldsString := c.Query("select")
	if selectFieldsString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Select fields are required"})
		return
	}
	selectFields := publicuserdto.EnsurePublicUserSelectFields(strings.Split(selectFieldsString, ","))
	user, err := uc.svc.GetUserByIDWithSelect(uint(userIdUint), selectFields)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uc.logger.Error("user not found", zap.Uint("user_id", uint(userIdUint)))
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": publicuserdto.ToPublicUser(user)})
	uc.logger.Info("user fetched successfully", zap.Uint("user_id", uint(userIdUint)))
}

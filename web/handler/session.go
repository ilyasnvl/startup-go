package handler

import (
	"fmt"
	"net/http"
	"startup/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "session_new.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		fmt.Println("error:", err)
		return
	}

	user, err := h.userService.Login(input)
	if err != nil || user.Role != "admin" {
		c.Redirect(http.StatusFound, "/login")
		fmt.Println("error2:", err)
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")

}

func (h *sessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}

// func (h *sessionHandler) Navbar(c *gin.Context) {
// 	session := sessions.Default(c)
// 	userName := session.Get("userName")

// 	// Jika userName tidak ditemukan, redirect ke login
// 	if userName == nil {
// 		c.Redirect(http.StatusFound, "/login")
// 		return
// 	}

// 	// Render template dengan mengirimkan userName ke template
// 	c.HTML(http.StatusOK, "base.html", gin.H{
// 		"userName": userName,
// 	})
// }

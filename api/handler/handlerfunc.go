package handler

import (
	comp "MusicMesh/api_geteway-MusicMesh/generate/composition"
	"MusicMesh/api_geteway-MusicMesh/generate/invitation"
	u "MusicMesh/api_geteway-MusicMesh/generate/user"
	user "MusicMesh/api_geteway-MusicMesh/generate/user"
	"MusicMesh/api_geteway-MusicMesh/generate/user_profile"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func (h *Handler) Post(c *gin.Context) {
	url := c.Request.URL.String()

	switch {
	case strings.HasPrefix(url, "/user"):
		fmt.Println("Nma Gap EEEEe")
		user := u.User{}
		c.ShouldBindJSON(&user)
		ctx := context.Background()
		_, err := h.User.Create(ctx, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"success": " user created"})
		return

	case strings.HasPrefix(url, "/composition"):
		composition := comp.Composition{}
		c.ShouldBindJSON(&composition)
		ctx := context.Background()
		_, err := h.Composition.Create(ctx, &composition)
		if err != nil {
			fmt.Println("hello world")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"success": " composition created"})
		return

	case strings.HasPrefix(url, "/collaboration"):
		invite := invitation.Invitation{}
		err := c.ShouldBindJSON(&invite)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println("Error in shouldBindJson", err)
		}
		ctx := context.Background()
		_, err = h.Invitation.Create(ctx, &invite)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"success": " invitation created"})
	}
}

func (h *Handler) PostUProfile(c *gin.Context) {
	fmt.Println("Nma Gap EEEEe")
	userProfile := user_profile.UserProfile{}
	err := c.ShouldBindJSON(&userProfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userProfile.Bio, userProfile.Role)
	_, err = h.UserProfile.Create(context.Background(), &userProfile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": " user_profile created"})

}

func (h *Handler) Get(c *gin.Context) {
	url := c.Request.URL.String()
	switch {
	case strings.HasPrefix(url, "/user/profile"):
		usPr, err := h.UserProfile.Get(context.Background(),
			&user_profile.FilterRequest{
				Query: "select user_id, full_name, bio, role, location, avatar_url, website from user_profiles",
				Arr:   []string{},
			})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, usPr)
		return

	case strings.HasPrefix(url, "/user"):
		usr, err := h.User.Get(context.Background(),
			&user.FilterRequest{
				Query: "select user_id, user_name, email, password, created_at, updated_at from  users where deleted_at = 0",
				Arr:   []string{}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusCreated, usr)
	case strings.HasPrefix(url, "/composition"):
		id := c.Param("id")
		cID := &comp.UserId{UserID: id}
		comps, err := h.Composition.GetByUserID(context.Background(), cID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"user_id": id, "composition": comps})
	}
}

func (h *Handler) GetById(c *gin.Context) {
	url := c.Request.URL.String()
	switch {
	case strings.HasSuffix(url, "/profile"):
		id := c.Param("id")
		ufID := &user_profile.UserProfileId{Id: id}
		uf, err := h.UserProfile.GetByID(context.Background(), ufID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, uf)
		return
	case strings.HasPrefix(url, "/user"):
		id := c.Param("id")
		uID := &user.UserId{Id: id}
		us, err := h.User.GetByID(context.Background(), uID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, us)
	case strings.HasPrefix(url, "/composition"):
		id := c.Param("id")
		cID := &comp.CompositionId{CompositionID: id}
		cm, err := h.Composition.GetByID(context.Background(), cID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println(err)
			return
		}
		c.JSON(http.StatusOK, cm)
	}
}

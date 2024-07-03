package handler

import (
	c "MusicMesh/api_geteway-MusicMesh/generate/composition-service/composition"
	t "MusicMesh/api_geteway-MusicMesh/generate/composition-service/tracks"
	u "MusicMesh/api_geteway-MusicMesh/generate/user-service/user"
	up "MusicMesh/api_geteway-MusicMesh/generate/user-service/user_profile"
)

type Handler struct {
	User        u.UserServiceClient
	UserProfile up.UserServiceClient
	Composition c.CompositionServiceClient
	Track       t.TrackServiceClient
}

func NewHandler(user u.UserServiceClient,
	userProfile up.UserServiceClient,
	composition c.CompositionServiceClient,
	track t.TrackServiceClient) *Handler {

	return &Handler{user, userProfile, composition, track}
}

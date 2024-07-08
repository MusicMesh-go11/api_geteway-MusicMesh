package handler

import (
	"MusicMesh/api_geteway-MusicMesh/generate/collaboration"
	"MusicMesh/api_geteway-MusicMesh/generate/comment"
	c "MusicMesh/api_geteway-MusicMesh/generate/composition"
	mdata "MusicMesh/api_geteway-MusicMesh/generate/composition_metadata"
	"MusicMesh/api_geteway-MusicMesh/generate/invitation"
	t "MusicMesh/api_geteway-MusicMesh/generate/tracks"
	u "MusicMesh/api_geteway-MusicMesh/generate/user"
	uinterac "MusicMesh/api_geteway-MusicMesh/generate/user_interactions"
	up "MusicMesh/api_geteway-MusicMesh/generate/user_profile"
)

type Handler struct {
	User                u.UserServiceClient
	UserProfile         up.UserServiceClient
	Composition         c.CompositionServiceClient
	Track               t.TrackServiceClient
	Collaboration       collaboration.CollaborationServiceClient
	Comment             comment.CommentServiceClient
	Invitation          invitation.InvitationServiceClient
	CompositionMetadata mdata.CompositionMetadataServiceClient
	UserInteraction     uinterac.UserInteractionsServiceClient
}

package api

import (
	"MusicMesh/api_geteway-MusicMesh/api/handler"
	"MusicMesh/api_geteway-MusicMesh/generate/collaboration"
	"MusicMesh/api_geteway-MusicMesh/generate/comment"
	comp "MusicMesh/api_geteway-MusicMesh/generate/composition"
	"MusicMesh/api_geteway-MusicMesh/generate/composition_metadata"
	"MusicMesh/api_geteway-MusicMesh/generate/invitation"
	"MusicMesh/api_geteway-MusicMesh/generate/tracks"
	use "MusicMesh/api_geteway-MusicMesh/generate/user"
	"MusicMesh/api_geteway-MusicMesh/generate/user_interactions"
	userProf "MusicMesh/api_geteway-MusicMesh/generate/user_profile"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func API(us, com, col, dis *grpc.ClientConn) *gin.Engine {

	UserClient := use.NewUserServiceClient(us)
	UserProfileClient := userProf.NewUserServiceClient(us)
	CompositionClient := comp.NewCompositionServiceClient(com)
	TrackClient := tracks.NewTrackServiceClient(com)
	CollaborationClient := collaboration.NewCollaborationServiceClient(col)
	CommentClient := comment.NewCommentServiceClient(col)
	InvitationClient := invitation.NewInvitationServiceClient(col)
	CompositionMetadataClient := composition_metadata.NewCompositionMetadataServiceClient(dis)
	UserInteractionCLient := user_interactions.NewUserInteractionsServiceClient(dis)

	handler := &handler.Handler{
		User:                UserClient,
		UserProfile:         UserProfileClient,
		Composition:         CompositionClient,
		Track:               TrackClient,
		Collaboration:       CollaborationClient,
		Comment:             CommentClient,
		Invitation:          InvitationClient,
		CompositionMetadata: CompositionMetadataClient,
		UserInteraction:     UserInteractionCLient,
	}

	router := gin.Default()

	user := router.Group("/user")
	composition := router.Group("/composition")
	collaboration := router.Group("/collaboration")
	discovery := router.Group("/discovery")

	user.POST("", handler.Post)       //go
	user.GET("", handler.Get)         // go
	user.GET("/:id", handler.GetById) // go
	user.PUT("/:id")
	user.DELETE("/:id")

	user.POST("/:id/profile", handler.PostUProfile) // go
	user.GET("/profile", handler.Get)               // go
	user.GET("/:id/profile", handler.GetById)       // go
	user.PUT("/:id/profile")
	user.DELETE("/:id/profile")

	composition.POST("", handler.Post)        // go
	composition.GET("/:id", handler.GetById)  // go
	composition.GET("/:id/user", handler.Get) // go
	composition.PUT("/:id")
	composition.DELETE("/:id")
	user.GET("/:id/compositions")

	composition.POST("/:id/track")
	composition.GET("/track")
	composition.GET("/:id/track")
	composition.PUT("/:id/track")
	composition.DELETE("/:id/track")

	collaboration.POST("/invite", handler.Post)
	collaboration.PUT("/invite/:id/respond")
	collaboration.PUT("/invite")
	collaboration.DELETE("/invite/:id")
	composition.GET("/:id/collaborators")
	composition.PUT("/:id/collaborators/:userId")
	composition.DELETE("/:id/collaborators/:userId")
	composition.POST("/:id/comment")
	composition.GET("/:id/comments")

	discovery.GET("/trending")
	discovery.GET("/recommended")
	discovery.GET("/genres/:genre")
	discovery.GET("/search")
	composition.POST("/:id/like")
	composition.DELETE("/:id/like")
	composition.POST("/:id/listen")
	return router
}

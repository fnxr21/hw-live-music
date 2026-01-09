package router

import (
	"github.com/labstack/echo/v4"
)

func RouterInit(r *echo.Group) {
	User(r)
	Song(r)
	Table(r)
	LivePlaylist(r)
	SongRequest(r)
	WebSocketRoute(r)
}



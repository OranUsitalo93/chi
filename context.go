package chi

import (
	"context"
	"net/http"
)

// ... (existing code)

func (c *Context) Reset() {
	c.RoutePattern = ""
	c.RoutePath = ""
	c.RouteParams.Keys = c.RouteParams.Keys[:0]
	c.RouteParams.Values = c.RouteParams.Values[:0]
	c.Params = c.Params[:0]
}

// ... (rest of the file)
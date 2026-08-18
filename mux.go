// ... (existing code)

func (mx *Mux) routeHTTP(w http.ResponseWriter, r *http.Request) bool {
	ctx := RouteContext(r.Context())
	prevLen := len(ctx.RouteParams.Keys)

	// ... (routing logic)

	// If match fails, restore the params length
	if !matched {
		ctx.RouteParams.Keys = ctx.RouteParams.Keys[:prevLen]
		ctx.RouteParams.Values = ctx.RouteParams.Values[:prevLen]
		return false
	}

	return true
}

// ... (rest of the file)
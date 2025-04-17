package httphealthcheck

type (
	HealthStatusData struct {
		IsAvailable bool  `json:"is_available"`
		Error       error `json:"error,omitempty"`
	}

	healthResponse struct {
		Errors  map[string]error            `json:"errors,omitempty"`
		Healths map[string]HealthStatusData `json:"healths"`
		Msg     string                      `json:"msg,omitempty"`
	}
)

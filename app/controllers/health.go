package controllers

// HealthController regroup all endpoints for health check
type HealthController struct {
	Controller
}

type status struct {
	Status string `json:"status"`
}

// Get return status of the API
func (c *HealthController) Get(code string) {
	s := status{Status: "OK"}

	c.Reply().Ok().JSON(s)
}

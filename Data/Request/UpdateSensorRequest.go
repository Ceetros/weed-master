package Request

type UpdateSensorRequest struct {
	WeedId         string `json:"sensor" binding:"required"`
	UmidityPercent int    `json:"value" binding:"required"`
}

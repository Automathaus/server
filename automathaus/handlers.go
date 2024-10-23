package automathaus

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func registerNode(app *pocketbase.PocketBase, c echo.Context) error {
	type RegisterNodeRequest struct {
		Name       string `json:"name" validate:"required"`
		Ip         string `json:"ip" validate:"required,ip"`
		MacAddress string `json:"macAddress" validate:"required,mac"`
		Type       string `json:"type" validate:"required"`
	}

	var request RegisterNodeRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format: " + err.Error()})
	}

	// Add validation check
	if request.Name == "" || request.Ip == "" || request.MacAddress == "" || request.Type == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	collection, err := app.Dao().FindCollectionByNameOrId("nodes")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	record.Set("name", request.Name)
	record.Set("ip", request.Ip)
	record.Set("macAddress", request.MacAddress)
	record.Set("nodeType", request.Type)

	if err := app.Dao().SaveRecord(record); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Node registered successfully"})
}

func registerTempHumidity(app *pocketbase.PocketBase, c echo.Context) error {
	type RegisterTempHumidityRequest struct {
		Temperature float64 `json:"temperature" validate:"required"`
		Humidity    float64 `json:"humidity" validate:"required"`
		Room        string  `json:"room" validate:"required"`
		Sensor      string  `json:"sensor" validate:"required"`
	}

	var request RegisterTempHumidityRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format: " + err.Error()})
	}

	// Add validation check
	if request.Room == "" || request.Sensor == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	// Add range validation for temperature and humidity
	if request.Temperature == 0 || request.Humidity == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Temperature and humidity must be provided"})
	}

	collection, err := app.Dao().FindCollectionByNameOrId("tempData")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	record.Set("temperature", request.Temperature)
	record.Set("humidity", request.Humidity)
	record.Set("room", request.Room)
	record.Set("sensor", request.Sensor)

	if err := app.Dao().SaveRecord(record); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "data registered successfully"})
}

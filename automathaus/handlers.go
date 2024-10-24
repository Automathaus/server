package automathaus

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func registerNode(app *pocketbase.PocketBase, c echo.Context) error {
	type RegisterNodeRequest struct {
		Name       string `json:"nodeName" validate:"required"`
		Ip         string `json:"ip" validate:"required,ip"`
		MacAddress string `json:"macAddress" validate:"required,mac"`
		Type       string `json:"nodeType" validate:"required"`
	}

	var request RegisterNodeRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format: " + err.Error()})
	}

	collection, err := app.Dao().FindCollectionByNameOrId("nodes")
	if err != nil {
		return err
	}

	//check if the node already exists
	existingNode, err := app.Dao().FindFirstRecordByData("nodes", "macAddress", request.MacAddress)
	if err == nil && existingNode != nil {
		form := forms.NewRecordUpsert(app, existingNode)
		form.LoadData(map[string]any{
			"name":   request.Name,
			"ip":     request.Ip,
			"type":   request.Type,
			"online": true,
		})
		if err := form.Submit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"id": existingNode.Id})
	}
	if err != nil && err != sql.ErrNoRows {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error checking for existing node: " + err.Error()})
	}
	// Node doesn't exist, continue with creation

	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)

	form.LoadData(map[string]any{
		"name":       request.Name,
		"ip":         request.Ip,
		"macAddress": request.MacAddress,
		"type":       request.Type,
		"online":     true,
	})

	if err := form.Submit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	//respond with the id of the new record
	return c.JSON(http.StatusOK, map[string]string{"id": record.Id})
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

	collection, err := app.Dao().FindCollectionByNameOrId("tempData")
	if err != nil {
		return err
	}

	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)

	form.LoadData(map[string]any{
		"temperature": request.Temperature,
		"humidity":    request.Humidity,
		"room":        request.Room,
		"sensor":      request.Sensor,
	})

	if err := form.Submit(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "data registered successfully"})
}

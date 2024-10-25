package automathaus

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

func sendStateChangeToNode(ip string, gpio float64, state bool) error {
	fmt.Printf("Sending state change to node - IP: %s, Pin: %f, State: %t\n", ip, gpio, state)
	url := fmt.Sprintf("http://%s/bindings/AutomathausRelayControl/relayControl", ip)
	payload := map[string]interface{}{
		"pin":   gpio,
		"state": state,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

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

func lightsControl(app *pocketbase.PocketBase, c echo.Context) error {
	// Extract path parameters
	roomName := c.PathParam("roomName")
	lightName := c.PathParam("lightName")

	// First find the room
	roomRecord, err := app.Dao().FindFirstRecordByData("rooms", "name", roomName)
	if err != nil {
		fmt.Printf("Error finding room: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Find light with matching name and position (room ID)
	lightRecord, err := app.Dao().FindFirstRecordByFilter(
		"lights",
		"name = {:name} && position = {:position}",
		map[string]any{
			"name":     lightName,
			"position": roomRecord.Id,
		},
	)
	if err != nil {
		fmt.Printf("Error finding light: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	actuatorRecord, err := app.Dao().FindRecordById("actuators", lightRecord.Get("actuator").(string))
	if err != nil {
		fmt.Printf("Error finding actuator: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	controllerRecord, err := app.Dao().FindRecordById("nodes", actuatorRecord.Get("controller").(string))
	if err != nil {
		fmt.Printf("Error finding controller: %v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Get query parameters
	state := c.QueryParam("state")
	duration := c.QueryParam("duration")

	// Check if both parameters are provided
	if state != "" && duration != "" {
		fmt.Printf("Error: Cannot use both state and duration parameters simultaneously\n")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Cannot use both state and duration parameters simultaneously",
		})
	}

	// Log the request details
	fmt.Printf("Light control request - Room: %s, Light: %s, State: %s, Duration: %s, Controller IP: %s\n",
		roomName, lightName, state, duration, controllerRecord.Get("ip").(string))

	switch {
	case state != "":
		fmt.Printf("Setting light state to: %s\n", state)
		// Parse the state parameter
		stateToSet := state == "on" || state == "true" || state == "1"
		stateSetErr := sendStateChangeToNode(controllerRecord.Get("ip").(string), actuatorRecord.Get("gpio").(float64), stateToSet)
		if stateSetErr != nil {
			fmt.Printf("Error setting light state: %v\n", stateSetErr)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to set light state: " + stateSetErr.Error()})
		}
		// Update the light record with the new state
		actuatorRecord.Set("state", stateToSet)
		if err := app.Dao().SaveRecord(actuatorRecord); err != nil {
			fmt.Printf("Error updating light state in database: %v\n", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update light state: " + err.Error()})
		}
	case duration != "":
		d, err := time.ParseDuration(duration + "ms")
		if err != nil {
			fmt.Printf("Error parsing duration: %v\n", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid duration format"})
		}

		go func() {
			fmt.Printf("Setting light state to on\n")
			stateSetErr := sendStateChangeToNode(controllerRecord.Get("ip").(string), actuatorRecord.Get("gpio").(float64), true)
			if stateSetErr != nil {
				fmt.Printf("Error setting light state: %v\n", stateSetErr)
				return
			}

			actuatorRecord.Set("state", true)
			if err := app.Dao().SaveRecord(actuatorRecord); err != nil {
				fmt.Printf("Failed to update light state to on: %v\n", err)
				return
			}

			time.Sleep(d)
			fmt.Printf("Setting light state to off\n")
			stateSetErr = sendStateChangeToNode(controllerRecord.Get("ip").(string), actuatorRecord.Get("gpio").(float64), false)
			if stateSetErr != nil {
				fmt.Printf("Error setting light state: %v\n", stateSetErr)
				return
			}
			actuatorRecord.Set("state", false)
			if err := app.Dao().SaveRecord(actuatorRecord); err != nil {
				fmt.Printf("Failed to update light state to off: %v\n", err)
				return
			}
		}()
	default:
		fmt.Printf("Error: Either state or duration parameter is required\n")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Either state or duration parameter is required",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Light control executed successfully"})
}

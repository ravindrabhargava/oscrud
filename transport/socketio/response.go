package socketio

import (
	"encoding/json"
	"oscrud"
)

func parseResponse(response oscrud.TransportResponse) string {

	if response.Error() != nil {
		errJSON, err := json.Marshal(
			map[string]interface{}{
				"status":  response.Status(),
				"error":   response.ErrorMap(),
				"headers": response.Headers(),
			},
		)
		if err != nil {
			return err.Error()
		}
		return string(errJSON)
	}
	resJSON := map[string]interface{}{
		"status":  response.Status(),
		"headers": response.Headers(),
	}
	switch response.ContentType() {
	case oscrud.ContentTypeJSON:
		resJSON["result"] = response.Result()
		res, err := json.Marshal(resJSON)
		if err != nil {
			return err.Error()
		}
		return string(res)
	}
	return oscrud.ErrResponseFailed.Error()
}

package socketio

import (
	"encoding/json"
	"oscrud"
)

func parseError(headers map[string]string, exception *oscrud.ErrorResponse) string {
	errJSON, err := json.Marshal(
		map[string]interface{}{
			"status":  exception.Status(),
			"error":   exception.ErrorMap(),
			"headers": headers,
		},
	)
	if err != nil {
		return err.Error()
	}
	return string(errJSON)
}

func parseResult(headers map[string]string, result *oscrud.ResultResponse) string {
	resJSON := map[string]interface{}{
		"status":  result.Status(),
		"headers": headers,
	}
	switch result.ContentType() {
	case oscrud.ContentTypeJSON:
		resJSON["result"] = result.Result()
		res, err := json.Marshal(resJSON)
		if err != nil {
			return err.Error()
		}
		return string(res)
	}
	return oscrud.ErrResponseFailed.Error()
}

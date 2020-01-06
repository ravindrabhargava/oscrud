package socketio

import (
	"encoding/json"
	"oscrud"
)

func parseError(exception *oscrud.ErrorResponse) string {
	errJson, err := json.Marshal(exception.ErrorMap())
	if err != nil {
		return err.Error()
	}
	return string(errJson)
}

func parseResult(result *oscrud.ResultResponse) string {
	switch result.ContentType() {
	case oscrud.ContentTypeJSON:
		resJson, err := json.Marshal(result.Result())
		if err != nil {
			return err.Error()
		}
		return string(resJson)
	}
	return oscrud.ErrResponseFailed.Error()
}

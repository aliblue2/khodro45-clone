package logging

func mapToZapParam(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}

func mapToZeroLogParam(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{}, 0)
	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}

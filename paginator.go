package plexgo

import (
	"net/http"
	"reflect"
)

func PaginateWithDefaults[T any](f interface{}) ([]T, *http.Response, error) {
	return Paginate[T](f, 0, 250, 10000)
}

func Paginate[T any](f interface{}, initialOffset int32, increment int32, limit int32) ([]T, *http.Response, error) {
	var offset int32 = initialOffset
	var returnObject []T
	var latestResponse *http.Response
	incrementResp := Invoke(f, "Limit", increment)
	for offset < limit {
		// first invoke the Offset command to set the new offset
		offsetResp := Invoke(incrementResp[0].Interface(), "Offset", offset)
		// invoke the Execute function to get the response
		resp := Invoke(offsetResp[0].Interface(), "Execute")

		// convert the expected return values to their respective types
		actualValue := resp[0].Interface().([]T)
		latestResponse = resp[1].Interface().(*http.Response)
		err := resp[2].Interface()

		if err != nil {
			return returnObject, latestResponse, err.(error)
		}

		// append the results to the main return object
		returnObject = append(returnObject, actualValue...)

		// check if this is the last set in the response. This could be enhanced by inspecting the header for the max results
		if int32(len(actualValue)) < increment {
			break
		}

		offset += increment
	}
	return returnObject, latestResponse, nil
}

func Invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, v := range args {
		inputs[i] = reflect.ValueOf(v)
	}
	return reflect.ValueOf(any).MethodByName(name).Call(inputs)
}

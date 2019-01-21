package main

import(
	"net/http"
	"strconv"
)


func getParamFromQueryUrl(r *http.Request, param string) string {

    values, ok := r.URL.Query()[param]
    if !ok || len(values[0]) < 1 {
        panic("Url Param" + param + "is missing")
        return ""
    }
    return values[0]
}

func getNumericParamFromQueryUrl(r *http.Request, param string) int {
	value := getParamFromQueryUrl(r, param)
    numericValue, _ := strconv.Atoi(value)
    return numericValue
}
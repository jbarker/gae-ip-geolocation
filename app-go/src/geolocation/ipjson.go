package geolocation

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type GeoResponse struct {
    Country string
    Region string
    City string
    CityLatLong string
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    h := r.Header
    response := GeoResponse{h.Get("X-AppEngine-Country"),
        h.Get("X-AppEngine-Region"),
        h.Get("X-AppEngine-City"),
        h.Get("X-AppEngine-CityLatLong")}

    b, err := json.MarshalIndent(response, "", "  ")

    if (nil == err) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, string(b))
    } else {
        fmt.Fprint(w, err)
    }
}

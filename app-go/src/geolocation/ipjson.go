package geolocation

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type GeoResponse struct {
    Country string
    Region string
    City string
    CityLatLong string
}

const (
    enableCORS bool = false
)

func init() {
    http.HandleFunc("/api/ip.json", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    rh := r.Header
    rwh := w.Header()
    
    // CORS: Cross-Origin Resource Sharing
    if (enableCORS) {
        if origin := rh.Get("Origin"); "" != origin {
            rwh.Set("Access-Control-Allow-Origin", origin)
            rwh.Set("Access-Control-Allow-Methods", "GET")
        	rwh.Set("Vary", "Origin")
        }
        if "OPTIONS" == r.Method  {
            return
        }
    }
    
    response := GeoResponse{rh.Get("X-AppEngine-Country"),
        rh.Get("X-AppEngine-Region"),
        rh.Get("X-AppEngine-City"),
        rh.Get("X-AppEngine-CityLatLong")}

    b, err := json.MarshalIndent(response, "", "  ")

    if (nil == err) {
        rwh.Set("Content-Type", "application/json;charset=utf-8")
        
        // cache control, expires in 60 seconds
        rwh.Set("Cache-Control", "private")
        expires := time.Now().Add(time.Duration(60 * time.Second)).UTC()
        rwh.Set("Expires", expires.Format("Mon, Jan 02 2006 15:04:05 GMT"))
        
        fmt.Fprint(w, string(b))
    } else {
        fmt.Fprint(w, err)
    }
}

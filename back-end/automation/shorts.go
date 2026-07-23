package automation

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Short struct {
    Title string `json:"title"`
    URL   string `json:"url"`
}

func ShortsHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println("Request masuk")

    var short Short

    err := json.NewDecoder(r.Body).Decode(&short)
    if err != nil {
        fmt.Println(err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println("Title :", short.Title)
    fmt.Println("URL   :", short.URL)

    w.WriteHeader(http.StatusOK)

    // json.NewEncoder(w).Encode(map[string]interface{}{
    //     "status": "success",
    //     "title": short.Title,
    //     "url": short.URL,
    // })
}

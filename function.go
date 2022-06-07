package artifacthub_gchat_updates

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("NotificationHandler", notificationHandler)
}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	var p ah_payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		_ = fmt.Errorf("error decoding request body: %s", err)
		return
	}
	fmt.Println(p.Data.Package.Name)
	fmt.Println("C'est bon !")
	return
}

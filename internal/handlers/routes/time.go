package routes

import (
	"gortth/web/templates"
	"net/http"
	"time"
)

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	templates.Time(templates.TimeProps{
		Time:          now.Format("15:04:05 MST"),
		Date:          now.Format("Monday, January 2, 2006"),
		Timezone:      now.Location().String(),
		UnixTimestamp: now.Unix(),
	}).Render(r.Context(), w)

}

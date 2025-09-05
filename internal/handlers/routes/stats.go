package routes

import (
	"gortth/web/templates"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now()

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	uptime := time.Since(startTime)

	// Convert bytes to MB for better readability
	allocMB := float64(m.Alloc) / 1024 / 1024
	totalAllocMB := float64(m.TotalAlloc) / 1024 / 1024
	sysMB := float64(m.Sys) / 1024 / 1024

	templates.Stats(templates.StatsProps{
		GoVersion:    runtime.Version(),
		CpuCores:     runtime.NumCPU(),
		Goroutines:   runtime.NumGoroutine(),
		Uptime:       uptime.Round(time.Second).String(),
		CurrentAlloc: allocMB,
		TotalAlloc:   totalAllocMB,
		SystemMemory: sysMB,
		GcCycles:     m.NumGC,
	}).Render(r.Context(), w)
}

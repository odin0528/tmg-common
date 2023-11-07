package utils

import (
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/charts"
)

func DrawLineChart(filename, chartName, yUnitString string, xPoints, yPoints []float64) string {
	if filename == "" {
		filename = "example.html"
	}

	if filepath.Ext(filename) == "" {
		filename += ".html"
	}

	if _, err := os.Stat("./reports"); os.IsNotExist(err) {
		os.Mkdir("reports", os.ModePerm)
	}
	f, _ := os.Create(filepath.Join("reports", filename))
	page := charts.NewPage()
	page.Add(createLineChart(xPoints, yPoints, chartName, yUnitString))
	page.Render(f)

	return filename
}

func createLineChart(xPoints, yPoints []float64, chartName, yUnitString string) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: chartName}, charts.ToolboxOpts{Show: true})

	for i := 0; i < len(xPoints); i++ {
		value := int(xPoints[i] * 100)
		xPoints[i] = float64(value) / 100
	}

	line.AddXAxis(xPoints).
		AddYAxis(yUnitString, yPoints)

	return line
}

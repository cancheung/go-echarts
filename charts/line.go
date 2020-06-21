package charts

import "github.com/go-echarts/go-echarts/types"

// Line represents a line chart.
type Line struct {
	RectChart
}

// Type returns the chart type.
func (Line) Type() string { return types.ChartLine }

// NewLine creates a new line chart.
func NewLine() *Line {
	chart := &Line{}
	chart.initBaseConfiguration()
	chart.HasXYAxis = true
	return chart
}

// SetXAxis adds the X axis.
func (c *Line) SetXAxis(xAxis interface{}) *Line {
	c.xAxisData = xAxis
	return c
}

// AddSeries adds the Y axis.
func (c *Line) AddSeries(name string, yAxis interface{}, opts ...SeriesOpts) *Line {
	series := SingleSeries{Name: name, Type: types.ChartLine, Data: yAxis}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}
package logger

var (
	colorIndex = 0
	colorDict  = map[string]string{}
	white      = "\033[37m"
	reset      = "\033[0m"
	bold       = "\033[1m"
	red        = "\033[31m"
	cyan       = "\033[36m"
	colors     = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}
)

func nextColor() string {
	colorIndex = colorIndex + 1
	return colors[colorIndex%len(colors)]
}

func colorFor(key string) string {
	if color, ok := colorDict[key]; ok {
		return color
	}

	colorDict[key] = nextColor()
	return colorDict[key]
}

package logger

import (
	"fmt"
	"sync"
)

var (
	colors  sync.Map
	white   = "\033[37m"
	reset   = "\033[0m"
	bold    = "\033[1m"
	red     = "\033[31m"
	blue    = "\033[34m"
	green   = "\033[32m"
	cyan    = "\033[36m"
	yellow  = "\033[33m"
	magenta = "\033[35m"

	//colorIndex      = 0
	//listOfColors    sync.Map
	//colorLoggerDict sync.Map
	/*colors          = []string{
		"\033[34m", // blue
		"\033[32m", // green
		"\033[36m", // cyan
		"\033[33m", // yellow
		"\033[35m", // magenta
	}*/
)

func init() {
	colors.Store("index", 0)
	colors.Store("index:0", blue)
	colors.Store("index:1", green)
	colors.Store("index:2", cyan)
	colors.Store("index:3", yellow)
	colors.Store("index:4", magenta)
	colors.Store("len", 5)
}

func nextColor() string {
	//colorIndex = colorIndex + 1
	//return colors[colorIndex%len(colors)]
	currentIndex, _ := colors.Load("index")
	len, _ := colors.Load("len")
	color, _ := colors.Load(fmt.Sprintf("index:%d", currentIndex.(int)%len.(int)))

	colors.Store("index", currentIndex.(int)+1)

	return color.(string)
}

func colorFor(key string) string {
	if color, ok := colors.Load(fmt.Sprintf("module:%s", key)); ok {
		return color.(string)
	}

	color := nextColor()
	colors.Store(fmt.Sprintf("module:%s", key), color)
	return color
}

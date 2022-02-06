package pipes

type Logmessage struct {
	severity int
	info     string
}

var DebugMessages chan string

func Init() {

}

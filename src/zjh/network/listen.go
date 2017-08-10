package network

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:19850", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func request(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	routeObj := &Route{}
	go routeObj.Run(c)
	//	defer c.Close()
	//	for {
	//		mt, message, err := c.ReadMessage()
	//		if err != nil {
	//			log.Println("read:", err)
	//			break
	//		}
	//		log.Printf("recv: %s", message)
	//		err = c.WriteMessage(mt, message)
	//		if err != nil {
	//			log.Println("write:", err)
	//			break
	//		}
	//	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func Listen() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", request)
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
</head>
<body>
</body>
</html>
`))

package ws

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/datalayer/kuber/cluster"
	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	//	writeWait = 10 * time.Second
	writeWait = 100 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// EnableCompression: true, # Experimental
}

var mutex = &sync.Mutex{}

type runner func(m WsMessage)

func Ws(w http.ResponseWriter, r *http.Request) {

	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	defer con.Close()

	for {

		m := WsMessage{}
		err = con.ReadJSON(&m)
		if err != nil {
			log.Println("Error reading json.", err)
			internalError(con, "stdout 1:", err)
			//			return
		} else {

			if m.Op == "PING" {
				doEcho(m, con, w, r)
			}
			if m.Op == "CREATE_CLUSTER_DEF" {
				doCommand(m, con, w, r, createClusterDefCommand)
			}
			if m.Op == "CREATE_CLUSTER" {
				doCommand(m, con, w, r, createClusterCommand)
			}
			if m.Op == "DELETE_CLUSTER" {
				doCommand(m, con, w, r, deleteClusterCommand)
			}
		}

	}

}

func doEcho(m WsMessage, con *websocket.Conn, w http.ResponseWriter, r *http.Request) {
	writeJsonToConn(m, con, w, r)
}

func doCommand(m WsMessage, con *websocket.Conn, w http.ResponseWriter, r *http.Request, run runner) {

	stdoutDone := make(chan struct{})

	old := os.Stdout
	outr, outw, err := os.Pipe()
	if err != nil {
		internalError(con, "stdout:", err)
		return
	}
	defer outr.Close()
	defer outw.Close()
	os.Stdout = outw

	go pumpStdoutWs(m, con, outr, stdoutDone, w, r)
	//	go pingWs(con, stdoutDone)

	run(m)

	outw.Close()
	os.Stdout = old

}

func createClusterDefCommand(m WsMessage) {
	cluster.CreateClusterDef(cluster.Options(m.Cluster.ClusterName, m.Cluster.AwsProfile))
}

func createClusterCommand(m WsMessage) {
	cluster.CreateCluster(cluster.Options(m.Cluster.ClusterName, m.Cluster.AwsProfile))
}

func deleteClusterCommand(m WsMessage) {
	cluster.DeleteCluster(cluster.Options(m.Cluster.ClusterName, m.Cluster.AwsProfile))
}

func pumpStdoutWs(m WsMessage, con *websocket.Conn, r io.Reader, done chan struct{}, hw http.ResponseWriter, hr *http.Request) {

	defer func() {}()

	s := bufio.NewScanner(r)

	for s.Scan() {
		//		ws.SetWriteDeadline(time.Now().Add(writeWait))
		b := s.Bytes()
		mm := WsMessage{}
		mm.Op = m.Op
		mm.Message = string(b)
		err := writeJsonToConn(mm, con, hw, hr)
		if err != nil {
			log.Println("error", err)
			break
		}
	}
	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}

	close(done)

	//	ws.SetWriteDeadline(time.Now().Add(writeWait))
	//	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//	time.Sleep(closeGracePeriod)
	//	ws.Close()
}

func pingWs(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				log.Println("ping:", err)
			}
		case <-done:
			return
		}
	}
}

func writeJsonToConn(m WsMessage, con *websocket.Conn, w http.ResponseWriter, r *http.Request) error {
	mutex.Lock()
	err := con.WriteJSON(m)
	if err != nil {
		fmt.Println(err)
		//			con.Close()
	}
	mutex.Unlock()
	return err
}

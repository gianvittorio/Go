package hydraproto

import (
	"errors"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

type ProtoHandler struct{}


func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

func (pSender *ProtoHandler) EncodeAndSend(obj interface{}, destination string) error {
	v, ok := obj.(*Ship)
	if !ok {
		return errors.New("Proto: Unknown message type")
	}
	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	return sendmessage(data, destination)
}

func (pSender *ProtoHandler) EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	return nil, errors.New("Proto: Unknown message type")
}

func (pSender *ProtoHandler) DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	return pb, proto.Unmarshal(buffer, pb)
}

func (pSender *ProtoHandler) ListenAndDecode(listenadress string) (chan interface{}, error) {
	outChan := make(chan interface{})
	l, err := net.Listen("tcp", listenadress)
	if err != nil {
		return outChan, err
	}
	log.Println("Listening to ", listenadress)
	go func(){
		defer l.Close()
		for {
			c, err := l.Accept()
			if err != nil {
				break
			}
			log.Println("Accepted connection from ", c.RemoteAddr())
			go func(c net.Conn){
				defer c.Close()
				for {
					buffer, err := ioutil.ReadAll(c)
					if err != nil {
						break
					}
					if len(buffer) == 0 {
						continue
					}
					obj, err := pSender.DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case outChan <- obj:
					default:
					}
				}
			}(c)
		}
	}()
	return outChan, nil
}

func sendmessage(buffer []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d bytes to %s \n", len(buffer), destination)
	_, err = conn.Write(buffer)
	return err
}

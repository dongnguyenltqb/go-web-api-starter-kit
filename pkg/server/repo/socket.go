package repo

import "github.com/gorilla/websocket"

type SocketRepo interface {
	GetById(string) *Socket
}

type socketRepo struct {
}

type Socket struct {
	Conn   *websocket.Conn
	Rooms  []string
	UserID string
}

var Sk map[string]*Socket

func (s *Socket) AddRoomToSocket(name string) *Socket {
	s.Rooms = append(s.Rooms, name)
	return s
}

func (s *Socket) RemoveRoom(name string) *Socket {
	r := []string{}
	for _, v := range s.Rooms {
		if v != name {
			r = append(r, v)
		}
	}
	s.Rooms = r
	return s
}

func (srp *socketRepo) GetById(id string) *Socket {
	return Sk[id]
}

func (srp *socketRepo) GetByRoomName(name string) (s []*Socket) {
	for _, v := range Sk {
		ok := false
		for _, x := range v.Rooms {
			if x == name {
				ok = true
			}
		}
		if ok {
			s = append(s, v)
		}
	}
	return
}

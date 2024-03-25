package chat

import (
	"strconv"

	desc "github.com/Sysleec/chat-server/pkg/chat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	s.mxChannel.RLock()
	id := strconv.Itoa(int(req.GetChat().ChatId))

	chatChan, ok := s.channels[id]
	s.mxChannel.RUnlock()

	if !ok {
		return status.Errorf(codes.NotFound, "chat not found")
	}

	s.mxChat.Lock()
	if _, chatOk := s.chats[id]; !chatOk {
		s.chats[id] = &Chat{
			streams: map[string]desc.ChatV1_ConnectChatServer{},
		}
	}
	s.mxChat.Unlock()

	s.chats[id].m.Lock()
	s.chats[id].streams[req.GetUsername()] = stream
	s.chats[id].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChan:
			if !okCh {
				return nil
			}

			for _, st := range s.chats[id].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}
		case <-stream.Context().Done():
			s.chats[id].m.Lock()
			delete(s.chats[id].streams, req.GetUsername())
			s.chats[id].m.Unlock()
			return nil
		}

	}
}

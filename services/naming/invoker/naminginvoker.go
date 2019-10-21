package invoker

import (
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/clientproxy"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/marshaller"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/miop"
	"Plataformas-Exercicio03/RobertoMiddleware/infrastructure/srh"
	"Plataformas-Exercicio03/RobertoMiddleware/services/naming"
	"Plataformas-Exercicio03/util"
)

type NamingInvoker struct{}

func (NamingInvoker) Invoke() {

	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: util.NAMING_PORT}
	marshallerImpl := marshaller.Marshaller{}
	namingImpl := naming.NamingService{}
	miopPacketReply := miop.Packet{}
	replyParams := make([]interface{}, 1)

	// control loop
	for {
		// receive request packet
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall request packet
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)

		// extract operation name
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		// demux request
		switch operation {
		case "Register":
			_p1 := miopPacketRequest.Bd.ReqBody.Body[0].(string)
			_map := miopPacketRequest.Bd.ReqBody.Body[1].(map[string]interface{})
			_proxyTemp := _map["Proxy"].(map[string]interface{})
			_p2 := clientproxy.ConversorProxy{TypeName: _proxyTemp["TypeName"].(string), Host: _proxyTemp["Host"].(string), Port: int(_proxyTemp["Port"].(float64)), Id: int(_proxyTemp["Id"].(float64))}

			// dispatch request
			replyParams[0] = namingImpl.Register(_p1, _p2)
		case "Lookup":
			_p1 := miopPacketRequest.Bd.ReqBody.Body[0].(string)

			// dispatch request
			replyParams[0] = namingImpl.Lookup(_p1)
		}

		// assembly reply packet
		repHeader := miop.ReplyHeader{Context: "", RequestId: miopPacketRequest.Bd.ReqHeader.RequestId, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replyParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		// marshall reply packet
		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)

		// send reply packet
		srhImpl.Send(msgToClientBytes)
	}
}

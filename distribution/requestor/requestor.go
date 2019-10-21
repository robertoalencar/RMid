package requestor

import (
	"Plataformas-Exercicio03/RobertoMiddleware/aux"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/marshaller"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/miop"
	"Plataformas-Exercicio03/RobertoMiddleware/infrastructure/crh"
	"Plataformas-Exercicio03/util"
)

type Requestor struct{}

func (Requestor) Invoke(inv aux.Invocation) interface{} {

	marshallerInst := marshaller.Marshaller{}
	crhInst := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	// create request packet
	reqHeader := miop.RequestHeader{Context: "Context", RequestId: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: util.MIOP_REQUEST}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	// serialise request packet
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)

	// send request packet and receive reply packet
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes)
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	// extract result from reply packet
	r := miopPacketReply.Bd.RepBody.OperationResult

	return r
}

package protocol

import (
	"github.com/iotaledger/goshimmer/packages/daemon"
	"github.com/iotaledger/goshimmer/packages/node"
	"github.com/iotaledger/goshimmer/plugins/autopeering/server/tcp"
	"github.com/iotaledger/goshimmer/plugins/autopeering/server/udp"
)

func Configure(plugin *node.Plugin) {
	errorHandler := createErrorHandler(plugin)

	udp.Events.ReceiveDrop.Attach(createIncomingDropProcessor(plugin))
	udp.Events.ReceivePing.Attach(createIncomingPingProcessor(plugin))
	udp.Events.Error.Attach(errorHandler)

	tcp.Events.ReceiveRequest.Attach(createIncomingRequestProcessor(plugin))
	tcp.Events.ReceiveResponse.Attach(createIncomingResponseProcessor(plugin))
	tcp.Events.Error.Attach(errorHandler)
}

func Run(plugin *node.Plugin) {
	daemon.BackgroundWorker(createChosenNeighborDropper(plugin))
	daemon.BackgroundWorker(createAcceptedNeighborDropper(plugin))
	daemon.BackgroundWorker(createOutgoingRequestProcessor(plugin))
	daemon.BackgroundWorker(createOutgoingPingProcessor(plugin))
}

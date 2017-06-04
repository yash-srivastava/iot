package conf

import (
	"github.com/orcaman/concurrent-map"
)

var(
	SGU_TCP_CONNECTION cmap.ConcurrentMap
	SGU_SCU_LIST cmap.ConcurrentMap
	Retry_3000 cmap.ConcurrentMap
	PACKET_CONFIG Sgu_packet
	RESPONSE_PACKET_CONFIG Sgu_response_packet
	CUSTOM_PACKET_CONFIG Sgu_packet
	SERVER_PACKET_CONFIG Server_packet
)

func Init()  {
	SGU_TCP_CONNECTION = cmap.New()
	SGU_SCU_LIST = cmap.New()
	Retry_3000 = cmap.New()
	PACKET_CONFIG = GetSguPacket()
	RESPONSE_PACKET_CONFIG = GetSguResponsePacket()
	CUSTOM_PACKET_CONFIG = GetCustomPackets()
	SERVER_PACKET_CONFIG = GetServerPacket()
}
package sender

import (
	"time"
)

func AddCommonParameters(delim byte, sgu_id uint64, seq_no uint64, length int, packet_type int) []byte {

	var response []byte
	response = make([]byte, length)


	data := convertToByteArray(uint64(delim),1)
	response = add_byte_array_to_response(0,1,data,response)

	data = convertToByteArray(uint64(length-3),2)
	response = add_byte_array_to_response(1,2,data,response)

	data = convertToByteArray(sgu_id,6)
	response = add_byte_array_to_response(3,6,data,response)

	currentTime := time.Now().Local()
	timestamp := currentTime.Format("20060102150405")

	data = []byte(timestamp)
	response = add_byte_array_to_response(9,14,data,response)

	data = convertToByteArray(uint64(seq_no),4)
	response = add_byte_array_to_response(23,4,data,response)

	data = convertToByteArray(uint64(packet_type),2)
	response = add_byte_array_to_response(27,2,data,response)

	return response
}

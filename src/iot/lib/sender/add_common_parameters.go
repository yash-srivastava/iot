package sender

import "time"

func AddCommonParameters(delim byte, sgu_id uint64, seq_no uint64, length int, packet_type int,) []byte {

	var response []byte
	response = make([]byte, length)

	response = convertToByteArray(uint64(delim),1)
	response = append(response, convertToByteArray(uint64(length),2)...)
	response = append(response, convertToByteArray(sgu_id,6)...)

	currentTime := time.Now().Local()
	timestamp := currentTime.Format("20060102150405")

	response = append(response, []byte(timestamp)...)


	response = append(response, convertToByteArray(uint64(seq_no),4)...)
	response = append(response, convertToByteArray(uint64(packet_type),2)...)

	return response
}

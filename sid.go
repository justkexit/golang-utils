package golang-utils

import (
	"encoding/binary"
	"strings"
)

// StringSID converts bytes Windows SID to string representation
func StringSID(sid []byte) string {
	var par = []string{"S"}
	par = append(par, fmt.Sprintf("%x", uint(sid[0])))
	cnt := int(sid[1])
	par = append(par, fmt.Sprintf("%d", binary.BigEndian.Uint64(append([]byte{0, 0}, sid[2:8]...))))
	var pos = 8
	for i := 0; i < cnt; i++ {
		par = append(par, fmt.Sprintf("%d", binary.LittleEndian.Uint32(sid[pos:pos+4])))
		pos += 4
	}
	return strings.Join(par, "-")
}

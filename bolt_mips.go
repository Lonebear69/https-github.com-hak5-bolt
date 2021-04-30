// +build mips
package bolt

const maxMapSize = 0x40000000 // 1GB
const maxAllocSize = 0xFFFFFFF
var brokenUnaligned = false

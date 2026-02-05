package tipos

import (
	"fmt"
	"math"
)

func TiposBasicos() {
	var intnormal int = 42
	var int8normal int8 = 127
	var int16normal int16 = 32767
	var int32normal int32 = 2147483647
	var int64normal int64 = 9223372036854775807

	fmt.Printf("Valor de intnormal: %d Tipo -> %T\n", intnormal, intnormal)
	fmt.Printf("Valor de int8normal: %d Tipo -> %T\n", int8normal, int8normal)
	fmt.Printf("Valor de int16normal: %d Tipo -> %T\n", int16normal, int16normal)
	fmt.Printf("Valor de int32normal: %d Tipo -> %T\n", int32normal, int32normal)
	fmt.Printf("Valor de int64normal: %d Tipo -> %T\n", int64normal, int64normal)

	fmt.Println(math.SmallestNonzeroFloat32)
}

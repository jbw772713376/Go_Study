package main
// 没有明确初始值的变量声明会被赋予它们的 零值。

import "fmt"

func zero_study() {
	//int类型的默认值为0
	var i int
	//flatt类型的默认值为0.00
	var f float32
	//bool类型的默认值为false
	var b bool
	//string类型的默认值为null
	var s string

	//Printf的格式化输出：
	//int类型：
		//%b 二进制表示
		//%c 相应Unicode码点所表示的字符
		//%d 十进制表示
		//%o 八进制表示
		//%q 单引号围绕的字符字面值，由Go语法安全地转义
		//%x 十六进制表示，字母形式为小写 a-f
		//%X 十六进制表示，字母形式为大写 A-F
		//%U Unicode格式：U+1234，等同于 "U+%04X"
	//float类型：
		//%b 无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat中的 'b' 转换格式一致。例如 -123456p-78
		//%e 科学计数法，例如 -1234.456e+78
		//%E 科学计数法，例如 -1234.456E+78
		//%f 有小数点而无指数，例如 123.456
		//%g 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
		//%G 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	//string和bytes类型：
		//%s 字符串或切片的无解译字节
		//%q 双引号围绕的字符串，由Go语法安全地转义
		//%x 十六进制，小写字母，每字节两个字符
		//%X 十六进制，大写字母，每字节两个字符
	//bool类型：
		//%t true 或 false
	//对于%v而言：
		//bool:                    %t
		//int, int8 etc.:          %d
		//uint, uint8 etc.:        %d, %x if printed with %#v
		//float32, complex64, etc: %g
		//string:                  %s
		//chan:                    %p
		//pointer:                 %p
	fmt.Printf("%d, %f, %t, %q\n", i, f, b, s)
}
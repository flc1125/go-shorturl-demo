package url

var chars string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 将长网址转换为短网址
func Encode(num int) string {
	if num == 0 {
		return "0"
	}

	var bytes []byte

	for num > 0 {
		bytes = append(bytes, chars[num%62])
		num = num / 62
	}

	Reverse(bytes)

	return string(bytes)
}

func Reverse(a []byte) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

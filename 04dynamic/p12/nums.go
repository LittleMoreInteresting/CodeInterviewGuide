package p12

func num1(str string) int {
	if str == "" {
		return 0
	}
	return process(str, 0)

}

func process(str string, i int) int {
	if i == len(str) {
		return 1
	}
	if str[i] == '0' {
		return 0
	}
	res := process(str, i+1)
	if i+1 < len(str) && (str[i]-'0')*10+str[i+1]-'0' < 27 {
		res += process(str, i+2)
	}
	return res
}

func num2(str string) int {
	if str == "" {
		return 0
	}
	cur := 1
	if str[len(str)-1] == '0' {
		cur = 0
	}
	next := 1
	tmp := 0
	for i := len(str) - 2; i >= 0; i-- {
		if str[i] == '0' {
			next = cur
			cur = 0
		} else {
			tmp = cur
			if (str[i]-'0')*10+str[i+1]-'0' < 27 {
				cur += next
			}
			next = tmp
		}
	}
	return cur
}

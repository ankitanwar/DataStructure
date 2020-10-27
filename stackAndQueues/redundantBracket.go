package stackandqueues

//CheckRedundant : To check whether the given string conatain redundant bracket or not
func CheckRedundant(str string) bool {
	s := []string{}
	for i := 0; i < len(str); i++ {
		current := string(str[i])
		if current == ")" {
			if s[len(s)-1] == "(" {
				return true
			}
			flag := true
			for {
				current := s[len(s)-1]
				if current == "+" || current == "-" || current == "*" {
					flag = false
				}
				s = s[:len(s)-1]
				if len(s) == 0 || current == "(" {
					break
				}
			}
			if flag == true {
				return true
			}
		} else {
			s = append(s, current)
		}
	}
	return false
}

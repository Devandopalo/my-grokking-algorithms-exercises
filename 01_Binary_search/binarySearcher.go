package binarySercher

type intArr []int
type bSerch interface {
	Search()
}

func (a intArr) Search(x int) (exist bool, index int) {
	exist = false
	s := a
	middleIndex := len(s) / 2

	for !exist {
		if x < s[middleIndex] {
			s = s[:middleIndex]
			exist, index = s.Search(x)
		}

		if x > s[middleIndex] {
			s = s[middleIndex:]
			exist, index = s.Search(x)
		}
		/*
			 stop here need to sleep and continue
			 need change if  to switch like:
			 switch midleIndex := len(s) / 2;  midleIndex {
			 	case midleIndex > s[midleIndex]:
					s = s[:middleIndex]
					exist, index = s.Search(x)
			}*/

	}

	if x == s[middleIndex] {
		exist = true
		return exist, middleIndex
	}

	return exist, index
}

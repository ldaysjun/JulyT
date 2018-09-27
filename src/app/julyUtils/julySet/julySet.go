package julySet

var Exists = struct {}{}

type Set struct {
	m map[interface{}]struct{}
}

func New(items ...interface{}) *Set{
	s := &Set{}
	s.m = make(map[interface{}]struct{})

	return s
}

//添加
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = Exists
	}
	return nil
}
//判断师傅包含
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
//获取大小
func (s *Set) Size() int {
	return len(s.m)
}
//清除
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

//对比
func (s *Set) Equal(other *Set) bool {
	// 如果两者Size不相等，就不用比较了
	if s.Size() != other.Size() {
		return false
	}

	// 迭代查询遍历
	for key := range s.m {
		// 只要有一个不存在就返回false
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//判断子集
func (s *Set) IsSubset(other *Set) bool {
	// s的size长于other，不用说了
	if s.Size() > other.Size() {
		return false
	}
	// 迭代遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
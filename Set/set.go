package set

type Set struct {
	Content []interface{}
}

func NewSet(values ...interface{}) *Set {
	return &Set{
		Content: values,
	}
}

func (s *Set) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Set) Add(element interface{}) {
	newContent := []interface{}{element}
	if shouldAdd(element, s) {
		s.Content = append(newContent, s.Content...)
	}
}

func (s *Set) Size() int {
	return len(s.Content)
}

func (s *Set) Delete(element interface{}) {
	indexOf := s.indexOf(element)
	if indexOf != -1 {
		s.Content = append(s.Content[:indexOf], s.Content[indexOf + 1 :]...)
	}
}

func (s *Set) indexOf(element interface{}) int {
	indexOf := -1
	for i, content := range s.Content {
		if content == element {
			indexOf = i
			break
		}
	}
	return indexOf
}

func shouldAdd(element interface{}, s *Set) bool {
	return s.indexOf(element) == -1
}
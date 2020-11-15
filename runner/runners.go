package runner

type runners []*runner

func NewRunners(sets CommandSets) runners {
	result := runners{}
	for _, set := range sets {
		result = append(result, NewRunner(set))
	}
	return result
}

func (rs runners) Do() {
	for _, r := range rs {
		go r.Do()
	}
}
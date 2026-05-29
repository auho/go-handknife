package suites

func (se *Suite) pfIndexIn() {
	se.mutex.Lock()
	se.pfIndex++
	se.mutex.Unlock()
}

func (se *Suite) pfIndexOut() {
	se.mutex.Lock()
	se.pfIndex--
	se.mutex.Unlock()
}

func (se *Suite) pfIndexExec(fn func(func() string)) {
	for _, _fn := range se.pfFunc[se.pfIndex] {
		fn(_fn)
	}

	delete(se.pfFunc, se.pfIndex)
	se.pfIndexOut()
}

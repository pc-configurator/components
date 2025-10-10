//go:build integration

package test

func (s *Suite) Test_CreateComponent() {
	age := 33

	s.Equal(33, age)
}

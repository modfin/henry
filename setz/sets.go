package setz

type Set[A comparable] map[A]struct{}

func Of[A comparable](slice []A) Set[A] {
	set := make(Set[A])
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func Has[A comparable](set Set[A], v A) bool {
	_, ok := set[v]
	return ok
}

func Add[A comparable](set Set[A], v A) Set[A] {
	set[v] = struct{}{}
	return set
}

func Remove[A comparable](set Set[A], v A) Set[A] {
	delete(set, v)
	return set
}

func Union[A comparable](set1 Set[A], set2 Set[A]) Set[A] {
	set := make(Set[A])
	for k := range set1 {
		set[k] = struct{}{}
	}
	for k := range set2 {
		set[k] = struct{}{}
	}
	return set
}

func Intersection[A comparable](set1 Set[A], set2 Set[A]) Set[A] {
	set := make(Set[A])
	for k := range set1 {
		if _, ok := set2[k]; ok {
			set[k] = struct{}{}
		}
	}
	return set
}

func Difference[A comparable](set1 Set[A], set2 Set[A]) Set[A] {
	set := make(Set[A])
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			set[k] = struct{}{}
		}
	}
	return set
}

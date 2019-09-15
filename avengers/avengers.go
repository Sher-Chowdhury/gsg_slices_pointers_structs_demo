package avengers

type Avenger struct {
	ID            int
	SuperHeroName string
	RealName      string
}

// This will hold info about all the avengers
var avengers []*Avenger
var id = 1

func AddAvenger(aSuperHeroName string, aRealName string) {
	var hero Avenger
	hero = Avenger{
		ID:            id,
		SuperHeroName: aSuperHeroName,
		RealName:      aRealName,
	}
	id++

	/*
			We use "&" so that the address of the hero object is appended to the
		    avengers pointers slice.
	*/
	avengers = append(avengers, &hero)
}

func GetAvengers() []*Avenger {
	// fmt.Println(avengers)
	return avengers
}

func RemoveAvengers() {

}

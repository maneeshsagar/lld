package domain

import "slices"

type TheaterController struct {
	CityVsTheater map[string][]*Theater
	Theaters      []*Theater
}

func NewTheaterController() *TheaterController {
	return &TheaterController{
		CityVsTheater: make(map[string][]*Theater),
	}
}

func (t *TheaterController) AddTheater(theater *Theater) {
	t.CityVsTheater[theater.City] = append(t.CityVsTheater[theater.City], theater)
	t.Theaters = append(t.Theaters, theater)
}

func (t *TheaterController) RemoveTheater(theater *Theater) {

	if theater == nil {
		return
	}

	allTheaterIndex := -1
	for i := 0; i < len(t.Theaters); i++ {
		if t.Theaters[i].Id == theater.Id {
			allTheaterIndex = i
		}
	}

	if allTheaterIndex != -1 {
		slices.Delete(t.Theaters, allTheaterIndex, allTheaterIndex+1)
	}

	cityVsTheaterIndex := -1

	for i := 0; i < len(t.CityVsTheater[theater.City]); i++ {
		if t.CityVsTheater[theater.City][i].Id == theater.Id {
			cityVsTheaterIndex = i
		}
	}

	if allTheaterIndex != -1 {
		slices.Delete(t.CityVsTheater[theater.City], cityVsTheaterIndex, cityVsTheaterIndex+1)
	}

}

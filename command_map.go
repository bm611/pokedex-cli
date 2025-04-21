package main

import (
	"fmt"
)

func callbackMap(cfg *config) error {
	locationAreasResp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range locationAreasResp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		fmt.Println("No previous location area")
		return nil
	}
	locationAreasResp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range locationAreasResp.Results {
		fmt.Println(area.Name)
	}
	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous

	return nil
}

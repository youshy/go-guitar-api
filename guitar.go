package main

type Guitar struct {
	Name         string `json:"name:omitempty`
	Strings      int    `json:"strings:omitempty`
	BodyWood     string `json:"bodywood:omitempty`
	NeckWood     string `json:"neckwood:omitempty`
	FingerWood   string `json:"fingerwood:omitempty`
	PickupsNo    int    `json:"pickupsno:omitempty`
	NeckPickup   string `json:"neckpickup:omitempty`
	BridgePickup string `json:"bridgepickup:omitempty`
	Bridge       string `json:"bridge:omitempty`
}

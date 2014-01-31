package vindinium

type State struct {
	Game    Game   `json:"game"`
	Hero    Hero   `json:"hero"`
	Token   string `json:"token"`
	ViewURL string `json:"viewURL"`
	PlayURL string `json:"playURL"`
}

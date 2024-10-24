package config

// Config содержит строковое обозначение типов Generator, Solver, UI и Renderer.
type Config struct {
	GeneratorType string `json:"GeneratorType"`
	SolverType    string `json:"SolverType"`
	UIType        string `json:"UIType"`
	RendererType  string `json:"RendererType"`
}

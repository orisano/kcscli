package kcscli

// Subtask is component of kcs problem
type Subtask struct {
	Name  string   `json:"name"`
	Files []string `json:"files"`
	Score int      `json:"score"`
}

// Problem is problem infomation
type Problem struct {
	Title       string    `json:"title"`
	TimeLimit   float64   `json:"time_limit"`
	MemoryLimit int       `json:"memory_limit"`
	Author      string    `json:"author"`
	Subtasks    []Subtask `json:"subtasks"`
	Solvers     []string  `json:"solvers"`
}

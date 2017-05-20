package kcscli

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

const (
	problemFilename = "problem.json"
)

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
	MemoryLimit uint64    `json:"memory_limit"`
	Author      string    `json:"author"`
	Subtasks    []Subtask `json:"subtasks"`
	Solvers     []string  `json:"solvers"`
}

func LoadProblem() (Problem, error) {
	raw, err := ioutil.ReadFile(problemFilename)
	if err != nil {
		return Problem{}, err
	}

	var problem Problem
	json.Unmarshal(raw, problem)
	return problem, nil
}

func (p *Problem) Save(withStdout bool) error {
	file, err := os.Create(problemFilename)
	if err != nil {
		return err
	}

	var writer io.Writer
	if withStdout {
		writer = io.MultiWriter(file, os.Stdout)
	} else {
		writer = file
	}
	return p.SaveTo(writer)
}

func (p *Problem) SaveTo(writer io.Writer) error {
	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	_, err = writer.Write(b)
	return err
}

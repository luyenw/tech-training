package main

import "fmt"

type Dependency interface {
	toString() string
}

type DependencyImpl struct{}

func (d DependencyImpl) toString() string {
	return "Dependency implementation"
}

type Container struct {
	dependencies map[string]Dependency
}

func NewContainer() *Container {
	return &Container{
		dependencies: make(map[string]Dependency),
	}
}

func (c *Container) Provide(name string, dep Dependency) {
	c.dependencies[name] = dep
}

func (c *Container) Resolve(name string) Dependency {
	return c.dependencies[name]
}

func main() {
	container := NewContainer()
	container.Provide("myDependency", DependencyImpl{})
	dependency := container.Resolve("myDependency")
	fmt.Println(dependency.toString())
}

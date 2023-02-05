package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"strings"
)

func main() {
	data :=[][]string{
		[]string("Alfred", "The Good", "500")
		[]string("Bob", "The Bad", "200")
		[]string("Alice", "The Ugly", "300")
		[]string("Hortense", "21", "(1, 1, 1)")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.setHeader([]string{"Name", "Alias", "Power"})
	table.appendBulk(data)
	table.Render()
}
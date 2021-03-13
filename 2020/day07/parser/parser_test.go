package parser

import (
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {

	testcases := [][]Rule{
		[]Rule{
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
			Rule{"dark orange", map[string]int{"bright white": 2, "muted yellow": 1}},
		},
		[]Rule{
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
			Rule{"light red", map[string]int{"bright blue": 1, "muted yellow": 1}},
		},
		[]Rule{
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
			Rule{"light red", map[string]int{"muted yellow": 1}},
		},
		[]Rule{
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 2}},
		},
		[]Rule{
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
			Rule{"light red", map[string]int{"bright white": 2, "muted yellow": 1}},
		},
	}

	expected := []bool{false, false, false, false, true}

	for i, test := range testcases {
		got := test[0].Equal(test[1])
		if got != expected[i] {
			t.Errorf("expected %+v, got %+v", expected[i], got)
		}
	}

}

const s = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

func TestParse(t *testing.T) {
	expected := []Rule{
		Rule{"light red", map[string]int{"bright white": 1, "muted yellow": 2}},
		Rule{"dark orange", map[string]int{"bright white": 3, "muted yellow": 4}},
		Rule{"bright white", map[string]int{"shiny gold": 1}},
		Rule{"muted yellow", map[string]int{"shiny gold": 2, "faded blue": 9}},
		Rule{"shiny gold", map[string]int{"dark olive": 1, "vibrant plum": 2}},
		Rule{"dark olive", map[string]int{"faded blue": 3, "dotted black": 4}},
		Rule{"vibrant plum", map[string]int{"faded blue": 5, "dotted black": 6}},
		Rule{"faded blue", map[string]int{}},
		Rule{"dotted black", map[string]int{}},
	}

	gots := Parse(strings.NewReader(s))

	if len(expected) != len(gots) {
		t.Fatalf("expected len (%d) diffs from got len (%d)", len(expected), len(gots))
	}

	for i, e := range expected {
		if !e.Equal(gots[i]) {
			t.Errorf("expected %+v, got %+v", e, gots[i])
		}
	}
}

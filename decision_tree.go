package main

// Yes, No, Idk

// each level of a decision tree has a
type DecisionTree struct {
	Action   Action
	Decision map[YesNoIdk]*DecisionTree
}

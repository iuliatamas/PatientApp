package main

var DemoTree *DecisionTree

func NewDemoTree(p *Patient) *DecisionTree {
	root := NewDT(NewPA("Good Morning! Did you weight yourself yet?", p))

	n := NewDT(NewPA("Make sure you do that first thing every morning! You should do that soon.", p))
	root.Decision[No] = n
	y := NewDT(NewPA("Did you gain over 2 pounds since last time?", p))
	root.Decision[Yes] = y
	i := NewDT(NewPA("I am sorry I couldn't understand that... yet", p))
	root.Decision[Idk] = i

	yn := NewDT(NewPA("You are all good then :)", p))
	y.Decision[No] = yn
	yy := NewDT(NewPA("Did you happen to eat anything salty recently?", p))
	y.Decision[Yes] = yy
	yi := NewDT(NewPA("I am sorry I couldn't understand that... yet", p))
	y.Decision[Idk] = yi

	yyn := NewDT(NewPA("Did you take your Lasix?", p))
	yy.Decision[No] = yyn
	yyy := NewDT(NewPA("That is probably it :)", p))
	yy.Decision[Yes] = yyy
	yyi := NewDT(NewPA("I am sorry I couldn't understand that... yet", p))
	yy.Decision[Idk] = yyi

	yynn := NewDT(NewPA("Oh in case you forgot it is time to take Lasix. It's the circular white one with the number 40 and the word Lasix on it :) Try not to forget next time!", p))
	yyn.Decision[No] = yynn
	yyny := NewDT(NewPA("If this continues we should probably contact your doctor :/", p))
	yyn.Decision[Yes] = yyny
	yyni := NewDT(NewPA("I am sorry I couldn't understand that... yet", p))
	yyn.Decision[Idk] = yyni

	return root
}

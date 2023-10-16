package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultListWidth   = 28
	defaultListHeight  = 40
	defaultDetailWidth = 45
	defaultInputWidth  = 22
	defaultHelpHeight  = 4
	cError             = "#CF002E"
	cItemTitleDark     = "#F5EB6D"
	cItemTitleLight    = "#F3B512"
	cItemDescDark      = "#9E9742"
	cItemDescLight     = "#FFD975"
	cTitle             = "#232923"
	cDetailTitle       = "#D32389"
	cPromptBorder      = "#D32389"
	cDimmedTitleDark   = "#DDDDDD"
	cDimmedTitleLight  = "#222222"
	cDimmedDescDark    = "#999999"
	cDimmedDescLight   = "#555555"
	cTextLightGray     = "#FFFDF5"
)


var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240")).Height(5)

var FocusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("250")).Height(5)

var DocStyle = lipgloss.NewStyle().Margin(1, 2)

var TableStyle = table.DefaultStyles()

var AppStyle = lipgloss.NewStyle().Margin(0, 1)

var H1Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color(cTextLightGray)).
	Background(lipgloss.Color(cTitle)).
	Padding(0, 1).Bold(true).MarginBottom(1)

var H2Style = lipgloss.NewStyle().
	Foreground(lipgloss.Color(cTextLightGray)).
	Background(lipgloss.Color(cTitle)).
	Padding(0, 1).Bold(true)
var BrightTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedTitleLight, Dark: cDimmedTitleDark}).Render
var NormalTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedDescLight, Dark: cDimmedDescDark})
var HelpStyle = list.DefaultStyles().HelpStyle.Width(defaultListWidth).Height(5)

func init() {
	TableStyle.Header = TableStyle.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	TableStyle.Selected = TableStyle.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
}

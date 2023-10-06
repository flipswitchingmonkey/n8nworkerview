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
	BorderForeground(lipgloss.Color("240")).Height(20)

var FocusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("250")).Height(20)

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

var DetailTitleStyle = lipgloss.NewStyle().
	Width(defaultDetailWidth).
	Foreground(lipgloss.Color(cTextLightGray)).
	Background(lipgloss.Color(cDetailTitle)).
	Padding(0, 1).
	Align(lipgloss.Center)
var InputTitleStyle = lipgloss.NewStyle().
	Width(defaultInputWidth).
	Foreground(lipgloss.Color(cTextLightGray)).
	Background(lipgloss.Color(cDetailTitle)).
	Padding(0, 1).
	Align(lipgloss.Center)
var SelectedTitle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: cItemTitleLight, Dark: cItemTitleDark}).
	Foreground(lipgloss.AdaptiveColor{Light: cItemTitleLight, Dark: cItemTitleDark}).
	Padding(0, 0, 0, 1)
var SelectedDesc = SelectedTitle.Copy().
	Foreground(lipgloss.AdaptiveColor{Light: cItemDescLight, Dark: cItemDescDark})
var DimmedTitle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedTitleLight, Dark: cDimmedTitleDark}).
	Padding(0, 0, 0, 2)
var DimmedDesc = DimmedTitle.Copy().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedDescDark, Dark: cDimmedDescLight})
var InputStyle = lipgloss.NewStyle().
	Margin(1, 1).
	Padding(1, 2).
	Border(lipgloss.RoundedBorder(), true, true, true, true).
	BorderForeground(lipgloss.Color(cPromptBorder)).
	Render
var DetailStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.ThickBorder(), false, false, false, true).
	BorderForeground(lipgloss.AdaptiveColor{Light: cItemTitleLight, Dark: cItemTitleDark}).
	Render
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(cError)).Render
var NoStyle = lipgloss.NewStyle()
var FocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(cPromptBorder))
var BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
var BrightTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedTitleLight, Dark: cDimmedTitleDark}).Render
var NormalTextStyle = lipgloss.NewStyle().
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedDescLight, Dark: cDimmedDescDark})
var SpecialTextStyle = lipgloss.NewStyle().
	Width(defaultDetailWidth).
	Margin(0, 0, 1, 0).
	Foreground(lipgloss.AdaptiveColor{Light: cItemTitleLight, Dark: cItemTitleDark}).
	Align(lipgloss.Center).Render
var DetailsBlockLeft = lipgloss.NewStyle().
	Width(defaultDetailWidth / 2).
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedTitleLight, Dark: cDimmedTitleDark}).
	Align(lipgloss.Right).
	Render
var DetailsBlockRight = lipgloss.NewStyle().
	Width(defaultDetailWidth / 2).
	Foreground(lipgloss.AdaptiveColor{Light: cDimmedDescLight, Dark: cDimmedDescDark}).
	Align(lipgloss.Left).
	Render
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

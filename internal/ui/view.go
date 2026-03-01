package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	if m.err != nil {
		return tea.View{
			Content: fmt.Sprintf("\n  %s\n\n  %s\n\n  Press q to quit.\n",
				errorStyle.Render("Failed to load brew data:"),
				m.err.Error(),
			),
			AltScreen: true,
		}
	}

	title := titleStyle.Render("🍺 Brew Potato ---  Packages")

	if m.isShowingAbout {
		return tea.View{
			Content:   lipgloss.JoinVertical(lipgloss.Left, title, renderAboutView()),
			AltScreen: true,
		}
	}

	tableView := tableWrapStyle.Render(m.table.View())

	var footer string
	switch {
	case m.isInitialLoading:
		footer = footerStyle.Render(
			fmt.Sprintf("  %s  Loading packages...", m.spinner.View()),
		)
	case m.isConfirming:
		footer = m.renderConfirmDialog()
	case m.isLoading:
		footer = footerStyle.Render(
			fmt.Sprintf("  %s  Uninstalling %s...", m.spinner.View(), m.loadingPkg),
		)
	default:
		footer = footerStyle.Render("  ↑/↓  navigate    d  uninstall    ?  about    q  quit")
	}

	return tea.View{
		Content:   lipgloss.JoinVertical(lipgloss.Left, title, tableView, footer),
		AltScreen: true,
	}
}

func renderAboutView() string {
	author := lipgloss.NewStyle().Hyperlink("https://github.com/mrSamDev").Render("mrSamDev")
	bubble := lipgloss.NewStyle().Hyperlink("https://github.com/charmbracelet/bubbletea").Render("Bubble Tea")
	bubbles := lipgloss.NewStyle().Hyperlink("https://github.com/charmbracelet/bubbles").Render("Bubbles")
	lipGloss := lipgloss.NewStyle().Hyperlink("https://github.com/charmbracelet/lipgloss").Render("Lip Gloss")

	box := dialogStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			keyHintStyle.Render("Built with")+"  "+bubble+"  ·  "+bubbles+"  ·  "+lipGloss,
			"",
			keyHintStyle.Render("By")+"  "+author,
		),
	)

	back := footerStyle.Render("  esc  back    q  quit")

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().MarginTop(1).Render(box),
		back,
	)
}

func (m Model) renderConfirmDialog() string {
	pkg := m.packages[m.confirmIdx].Name

	prompt := fmt.Sprintf("Uninstall  %s ?", dialogWarningStyle.Render(pkg))

	yKey := keyHintStyle.Render("y")
	nKey := keyHintStyle.Render("n")
	keys := fmt.Sprintf("%s  confirm    %s  cancel", yKey, nKey)

	box := dialogStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left, prompt, "", keys),
	)

	return lipgloss.NewStyle().MarginTop(1).Render(
		lipgloss.PlaceHorizontal(m.width, lipgloss.Center, box),
	)
}

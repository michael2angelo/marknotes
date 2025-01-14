package pages

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muhwyndhamhp/marknotes/pkg/models"
	"github.com/muhwyndhamhp/marknotes/pkg/post/values"
	"github.com/muhwyndhamhp/marknotes/ssh/base"
	"github.com/muhwyndhamhp/marknotes/utils/errs"
	"github.com/muhwyndhamhp/marknotes/utils/scopes"
)

type Home struct {
	PostRepo models.PostRepository
	Posts    []models.Post
}

// MatchKeyAction implements base.Page.
func (h *Home) MatchKeyAction(m base.Model, key string, sc base.ScreenMetadata) (base.Model, bool, tea.Cmd) {
	for i := range h.Posts {
		if key == fmt.Sprintf("%d", i+1) {
			a := NewArticle(&h.Posts[i])
			m.Content = a.RenderPage(m.Style, sc)

			m.ActiveTab = -1
			m.Page = a
			return m, true, nil
		}
	}
	// Handle own's page navigation
	if key != h.GetAccessKey() && m.Content != "" {
		return m, false, nil
	}

	m.Content = h.RenderPage(m.Style, sc)

	return m, true, nil
}

// GetAccessKey implements base.Page.
func (h *Home) GetAccessKey() string {
	return "h"
}

// GetName implements base.Page.
func (h *Home) GetName() string {
	return "Home"
}

func NewHome(postRepo models.PostRepository) base.Page {
	return &Home{PostRepo: postRepo}
}

// RenderPage implements base.Page.
func (h *Home) RenderPage(style lipgloss.Style, sm base.ScreenMetadata) string {
	doc := strings.Builder{}

	doc.WriteString(base.DescStyle.AlignHorizontal(lipgloss.Center).Width(sm.Width-8).Render(intro) + "\n")

	scopes := []scopes.QueryScope{
		scopes.OrderBy("published_at", scopes.Descending),
		scopes.Paginate(1, 5),
		scopes.Where("status = ?", values.Published),
		scopes.PostIndexedOnly(),
	}

	posts, err := h.PostRepo.Get(context.Background(), scopes...)
	if err != nil {
		panic(errs.Wrap(err))
	}

	h.Posts = posts

	for i, post := range posts {
		st := base.SubduedDescStyle.PaddingTop(1).Width(sm.Width - 2)
		body := lipgloss.JoinVertical(
			lipgloss.Top,
			base.PostTitle.Render(fmt.Sprintf("%s %s", post.Title, base.DescStyle.Render(fmt.Sprintf("[%d]", i+1)))),
			st.Render(
				fmt.Sprintf(
					"Published: %s | Updated: %s",
					post.PublishedAt.Format("Jan, 02 2006"),
					post.UpdatedAt.Format("Jan, 02 2006"),
				),
			),
		)
		doc.WriteString(base.PostItem.Width(sm.Width - 8).Render(body))
		doc.WriteString("\n")
	}

	return lipgloss.NewStyle().Padding(0, 4).Render(doc.String())
}

const intro = `
Hello! 🖖 My name is Wyndham

And I'm a Software Engineer that really into Pragmatic, Practical, and Beautiful Code.

My stack mostly consist of Android Development (Kotlin, Java), Backend Development (Golang, Rails), and Frontend Development (Plain HTML, CSS, JS).

More on my resume [r].

Hit me up via Twitter [t], LinkedIn [l] or Email [e] whenever you need help or just want to have some chit-chat!


Latest Articles:`

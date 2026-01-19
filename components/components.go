package components

import (
	about "two-x-dot-org/components/about_page"
	home "two-x-dot-org/components/home_page"
	"two-x-dot-org/components/page"

	"github.com/a-h/templ"
)

func GenHomePage() templ.Component {
	return page.Page(home.HomePage())
}

func GenAboutPage() templ.Component {
	return page.Page(about.AboutPage())
}

package components

import (
	about "two-x-dot-org/components/about_page"
	gallery "two-x-dot-org/components/gallery_page"
	home "two-x-dot-org/components/home_page"
	music "two-x-dot-org/components/music_page"
	"two-x-dot-org/components/page"

	"github.com/a-h/templ"
)

func GenHomePage() templ.Component {
	return page.Page(home.HomePage())
}

func GenAboutPage() templ.Component {
	return page.Page(about.AboutPage())
}

func GenMusicPage() templ.Component {
	return page.Page(music.MusicPage())
}

func GenGalleryPage() templ.Component {
	return page.Page(gallery.GalleryPage())
}

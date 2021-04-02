package handlers

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/benbjohnson/hashfs"
)

// Each HTML file has an associated variable that is populated
// with its content at build time - see https://golang.org/pkg/embed/
//
// This means the binary contains everything it needs to serve the site
//
var (
	// Shared templates
	//
	//go:embed layout.html
	layoutTemplate string
	//go:embed navbar.html
	navbarTemplate string
	//go:embed common_stimulus.html
	commonStimulusTemplate string
	//go:embed page_heading.html
	pageHeadingTemplate string

	// Template per page
	//
	//go:embed home.html
	homeTemplate string
	//go:embed dashboard.html
	dashboardTemplate string
	//go:embed login.html
	loginTemplate string
	//go:embed logout.html
	logoutTemplate string
	//go:embed recovery.html
	recoveryTemplate string
	//go:embed registration.html
	registrationTemplate string
	//go:embed settings.html
	settingsTemplate string

	emptyFuncMap         = template.FuncMap{}
	emptyStmulusTemplate = `
	{{define "stimulus"}}
	<!-- Empty stimulus template -->
	{{end}}`
)

// TemplateName provides a typesafe way of referring to templates
type TemplateName string

const (
	homePage         = TemplateName("home")
	dashboardPage    = TemplateName("dashboard")
	loginPage        = TemplateName("login")
	logoutPage       = TemplateName("logout")
	recoveryPage     = TemplateName("recovery")
	registrationPage = TemplateName("registration")
	settingsPage     = TemplateName("settings")
)

// Register all the Templates during initialisation
func init() {
	type tmpl struct {
		name      TemplateName     // Template name that handler code will refer to - one for each 'page'
		fmap      template.FuncMap // List of functions used inside the template
		templates []string         // List of HTML templates, snippets etc that make up the page
		stimulus  string           // Optional stimulus controller code
	}
	// All pages get the commonTemplates
	commonTemplates := []string{layoutTemplate, navbarTemplate, commonStimulusTemplate, pageHeadingTemplate}

	// The templates and their associated functions to include etc
	templates := []tmpl{
		{name: homePage, fmap: emptyFuncMap, templates: []string{homeTemplate}},
		{name: dashboardPage, fmap: emptyFuncMap, templates: []string{dashboardTemplate}},
		{name: loginPage, fmap: loginFuncMap(), templates: []string{loginTemplate}},
		{name: logoutPage, fmap: emptyFuncMap, templates: []string{logoutTemplate}},
		{name: recoveryPage, fmap: recoveryFuncMap(), templates: []string{recoveryTemplate}},
		{name: registrationPage, fmap: registrationFuncMap(), templates: []string{registrationTemplate}},
		{name: settingsPage, fmap: settingsFuncMap(), templates: []string{settingsTemplate}},
		{name: homePage, fmap: emptyFuncMap, templates: []string{homeTemplate}},
	}
	for _, t := range templates {
		stimulusTemplate := emptyStmulusTemplate
		if t.stimulus != "" {
			stimulusTemplate = t.stimulus
		}
		tmpl := append(commonTemplates, t.templates...)
		tmpl = append(tmpl, stimulusTemplate)

		// Ammend the global functions to the funcMap
		for k, v := range globalFuncMap() {
			t.fmap[k] = v
		}

		if err := RegisterTemplate(t.name, t.fmap, tmpl...); err != nil {
			// If we have a problem with a template, abort the app
			log.Fatalf("%v template error: %v", t.name, err)
		}
	}
}

// Default template functions, added to all templates
func globalFuncMap() template.FuncMap {

	return template.FuncMap{
		"assetPath": func(fs hashfs.FS, name string) string {
			if strings.HasPrefix(name, "/") {
				log.Printf("Error assetPath called with name: '%s' should not start with '/'", name)
			}
			path := fs.HashName(name)
			if strings.HasPrefix(path, "/") {
				return path
			}
			return fmt.Sprintf("/%s", path)
		},
	}
}

// Functions used by the 'settingsPage' templates
func settingsFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"password":          "Password",
		"traits.email":      "Email",
		"traits.name.first": "First name",
		"traits.name.last":  "Last name",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			log.Printf("No labelFor name: %s", name)
			return ""
		},
	}
}

// Functions used by the 'loginPage' templates
func loginFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"password":   "Password",
		"identifier": "Email",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			log.Printf("No labelFor name: %s", name)
			return ""
		},
	}
}

// Functions used by the 'recoveryPage' templates
func recoveryFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"email": "Email",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			log.Printf("No labelFor name: %s", name)
			return ""
		},
	}
}

// Functions used by the 'registration' templates
func registrationFuncMap() template.FuncMap {

	fieldLabel := map[string]string{
		"password":          "Password",
		"traits.email":      "Email",
		"traits.name.first": "First name",
		"traits.name.last":  "Last name",
	}

	return template.FuncMap{
		"labelFor": func(name string) string {
			if lbl, ok := fieldLabel[name]; ok {
				return lbl
			}
			log.Printf("No labelFor name: %s", name)
			return ""
		},
	}
}

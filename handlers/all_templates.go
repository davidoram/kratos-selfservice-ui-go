package handlers

import (
	_ "embed"
	"html/template"
	"log"
)

var (
	// Shared templates
	//
	//go:embed layout.html
	layoutTemplate string

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

	emptyFuncMap = template.FuncMap{}
)

const (
	homePage         = "home"
	dashboardPage    = "dashboard"
	loginPage        = "login"
	logoutPage       = "logout"
	recoveryPage     = "recovery"
	registrationPage = "registration"
	settingsPage     = "settings"
)

// Register the templates used by this handler
func init() {
	if err := RegisterTemplate(homePage, emptyFuncMap, homeTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", homePage, err)
	}

	if err := RegisterTemplate(dashboardPage, emptyFuncMap, dashboardTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", dashboardPage, err)
	}

	if err := RegisterTemplate(loginPage, loginFuncMap(), loginTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", loginPage, err)
	}

	if err := RegisterTemplate(logoutPage, emptyFuncMap, logoutTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", logoutPage, err)
	}

	if err := RegisterTemplate(recoveryPage, recoveryFuncMap(), recoveryTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", recoveryPage, err)
	}

	if err := RegisterTemplate(registrationPage, registrationFuncMap(), registrationTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", registrationPage, err)
	}

	if err := RegisterTemplate(settingsPage, settingsFuncMap(), settingsTemplate, layoutTemplate); err != nil {
		log.Fatalf("%v template error: %v", settingsPage, err)
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

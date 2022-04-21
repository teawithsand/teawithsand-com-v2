package remtrans

import (
	"strings"

	"github.com/teawithsand/webpage/util/trans"
)

var Translations = trans.NewBuilder().
	AddMiddleware(func(lang, key, val string) (string, string, string) {
		val = trans.Dedent(val)
		return lang, key, val
	}).
	AddMiddleware(func(lang, key, val string) (string, string, string) {
		val = strings.Trim(val, "\n\t\v ")
		return lang, key, val
	}).
	Lang("en").
	// TODO(teawithsand): prefix these with page
	WithPrefix("common.home", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
		return lb.
			WithPrefix("header", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("title", "teawithsand.com").
					Put("paragraph", "Teawithsand's webpage for all the things")
			}).
			WithPrefix("features", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("langka", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("title", "Langka").
							Put("paragraph", `Simple app for learning languages`).
							Put("button", `Go to langka`)
					}).
					WithPrefix("portfolio", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("title", `My portfolio`).
							Put("paragraph", `Links to GitHub/LinkedIn and my past and current projects with a few notes from me`).
							Put("button", `Go to portfolio`)
					}).
					WithPrefix("cheatsheet", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("title", `Programming cheatsheet`).
							Put("paragraph", `Ordered notes about weird behavior of computers, with emphasis on programming languages like JS and Go(but mostly JS)`).
							Put("button", `Go to cheatsheet`)
					})
			}).
			WithPrefix("contact", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("title", "Contact information").
					Put("paragraph", "You can call contact me using phone number or via my email address")
			})
	}).
	WithPrefix("common", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
		return lb.
			WithPrefix("navbar", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("home", "Home").
					Put("about_me", "About me").
					Put("langka_home", "Langka").
					Put("logged_in", "Logged in: {publicName}").
					Put("logout", "Log out").
					Put("my_profile", "My profile").
					Put("register", "Register").
					Put("login", "Log in")
			}).
			WithPrefix("footer", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("text", "By teawithsand; 2022 - echo date(\"Y\");")
			}).
			WithPrefix("contact", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("phone", "+48 883 910 432").
					Put("phone.raw", "+48883910432").
					Put("email", "teawithsand@gmail.com").
					Put("email.raw", "teawithsand@gmail.com")
			}).
			WithPrefix("form", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("submiterror.title", "An error occurred")
			}).
			WithPrefix("error", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("front", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("api_error", "Unknown remote server error occurred").
							Put("unknown", "Unknown error")

					}).
					WithPrefix("http", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("forbidden", "Access denied").
							Put("internal_server_error", "Internal server error").
							Put("not_found", "Not found").
							Put("unauthorized", "Authorization required, please login first").
							Put("unknown", "Unknown error")
					})
			})
	}).
	WithPrefix("langka", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
		return lb.
			WithPrefix("browsewordsgame", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb
			}).
			WithPrefix("form", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.WithPrefix("wordset.create", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
					return lb.
						WithPrefix("name", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
							return lb.
								Put("label", "Name of word set").
								Put("placeholder", "Name of word set")
						}).
						WithPrefix("from_language", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
							return lb.
								Put("label", "Source words language").
								Put("placeholder", "In format like en-US or fr-FR(case sensitive)")
						}).
						WithPrefix("to_language", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
							return lb.
								Put("label", "Destination words language").
								Put("placeholder", "In format like en-US or fr-FR(case sensitive)")
						}).
						WithPrefix("description", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
							return lb.
								Put("label", "Description").
								Put("placeholder", "Description of words contained in this word set")
						}).
						WithPrefix("submit", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
							return lb.
								Put("label", "Create word set")
						})
				})
			}).
			WithPrefix("page", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("home", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("title", "Langka - learn foreign languages").
							Put("subtitle", "Simple app for learning foreign languages").
							WithPrefix("browse", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
								return lb.
									Put("title", "Browse public resources").
									Put("description", "Langka allows creating and publishing learning resources. Here you can browse public ones.").
									Put("button", "Browse")
							}).
							WithPrefix("join", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
								return lb.
									Put("title", "Create new langka resources").
									Put("description", "Only registered user can create new resources, so you have to register first.").
									Put("button", "Register")
							}).
							WithPrefix("browse_owned", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
								return lb.
									Put("title", "Browse your own resources").
									Put("description", "You can also create new ones").
									Put("button", "Browse your resources")
							}).
							WithPrefix("create", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
								return lb.
									Put("title", "Create new word set").
									Put("description", "None of word sets you found is good enough? You can create new one.").
									Put("button", "Create new word set")
							})
					}).
					WithPrefix("wordset", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							WithPrefix("create", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
								return lb.
									Put("title", "Create new word set").
									Put("subtitle", "You will be able to add words later, once word set is created.")
							})
					})
			})
	}).
	WithPrefix("user", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
		return lb.
			WithPrefix("page", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("register", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.Put("title", "Register")
					}).
					WithPrefix("login", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.Put("title", "Log in")
					}).
					WithPrefix("change_password", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.Put("title", "Change password")
					}).
					WithPrefix("change_email", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.Put("title", "Change email address")
					})
			}).
			WithPrefix("form", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("login", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("username.label", "Username").
							Put("password.label", "Password").
							Put("submit.label", "Log in")
					}).
					WithPrefix("register", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("username.label", "Username").
							Put("username.placeholder", "Username").
							Put("email.label", "Email").
							Put("email.placeholder", "Email").
							Put("captcha.label", "Captcha").
							Put("submit.label", "Register")
					}).
					WithPrefix("change_password", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("password.label", "New password").
							Put("password.placeholder", "New password").
							Put("repeat_password.label", "Repeat new password").
							Put("repeat_password.placeholder", "Repeat new password").
							Put("submit.label", "Change password")
					}).
					WithPrefix("change_email", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("email.label", "New email").
							Put("email.placeholder", "New email").
							Put("submit.label", "Change email")
					})
			}).
			WithPrefix("validator", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("username.invalid", "Given user name is not valid").
					Put("email.invalid", "Given email is not valid").
					Put("password.invalid", "Given password is not valid").
					Put("repeat_password.mismatch", "Password must match the other one")

			}).
			WithPrefix("error", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					Put("email_in_use", "Given email address is already in use").
					Put("login_in_use", "Given login is already in use").
					Put("auth_token_invalid", "Session provided by your client is not valid. Try loging out and in again.").
					Put("not_found_for_login", "User for given username and password was not found").
					Put("user_for_registration_already_registered", "User for email that this link was sent to was already registered.").
					Put("must_log_in", "user must be logged in before continuing.")
			}).
			WithPrefix("profile", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
				return lb.
					WithPrefix("secret", func(lb *trans.LanguageBuilder) *trans.LanguageBuilder {
						return lb.
							Put("title", "{publicName}'s profile").
							Put("registered_at", "Registered at: {registeredAt}").
							Put("email", "Email: {email}").
							Put("public_name", "Name: {publicName}").
							Put("email_confirmed_at", "Email confirmed at: {confirmedAt}").
							Put("email_not_confirmed", "Email not confirmed").
							Put("email_not_confirmed.description", "Check your mailbox. If you did not receive any email, change email to same one.")
					}).
					Put("delete_account", "Delete account").
					Put("change_email", "Change email").
					Put("change_password", "Change password")
			})

	}).
	Done().
	MustBuild()

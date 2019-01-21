package templates

type Templates = map[string]string

var (
	Packr = Templates{
		"_README.md":       "README.md",
		"_CHANGELOG.md":    "CHANGELOG.md",
		"_CONTRIBUTORS.md": "CONTRIBUTORS.md",
		"_editorconfig":    ".editorconfig",
		"_gitignore":       ".gitignore",
	}
)

load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix bazel-go
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=external_dependencies/go.mod",
        "-to_macro=external_dependencies/deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

# make sure that others can import this either using the
# import or go_default_library naming conventions
# gazelle:go_naming_convention import_alias

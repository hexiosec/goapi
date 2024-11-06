# Developing Templates

Templates are grouped under [client](./client/) and [server](./server/) depending on whether they handle creating requests, or creating server stubs for a specific language/framework.

Each template should consist of:

- manifest.yml - Template configuration manifest.
- *.tmpl - Go text/template compatible template files.

## Manifest Format

- `name`: Template Name
- `description`: Template Description
- `render`: List of render outputs
  - `path`: Path of the generated file, including `*` as a wildcard if needed
  - `for`: One of:
    - `root`: One copy of this file is generated for a given spec. No wildcard support.
    - `tags`: A file per tag is generated, wildcard will be replaced with the tag name.

## Template Functions

All functions from the text/template module and [sprig](https://masterminds.github.io/sprig/) are supported.

Additional supporting functions defined by this package are documented here:

### `include <template_name>`
### `json <data>`
### `comment <prefix> <content>`
### `toGoPascalCase <string>`
### `toGoCamelCase <string>`
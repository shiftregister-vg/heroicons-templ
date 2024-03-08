# heroicons-templ
Provides heroicons for Go projects that use templ

## Install

```bash
go get github.com/stevegood/heroicons-templ/v2
```

## Usave

Within your templ template you could do the following:

```templ
package templates

import "github.com/stevegood/heroicons-templ/v2/24/outline"

templ MyTemplate() {
  <div>
    @outline.AcademicCap()
    Contgrats, you're using heroicons!
  </div>
}
```

Or with attributes (this will override the default attribute values if any are set):

```templ
package templates

import "github.com/stevegood/heroicons-templ/v2/24/outline"

templ MyTemplate() {
  <div>
    @outline.AcademicCapWithAttrs(templ.Attributes{"fill": "currentColor", "aria-hidden": "true"})
    Contgrats, you're using heroicons!
  </div>
}
```

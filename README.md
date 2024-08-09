[![.github/workflows/build.yml](https://github.com/bdreece/echo-renderer/actions/workflows/build.yml/badge.svg)](https://github.com/bdreece/echo-renderer/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bdreece/echo-renderer.svg)](https://pkg.go.dev/github.com/bdreece/echo-renderer)

# echo-renderer

A basic implementation of [echo.Renderer] using [html/template].

## Usage

The [echo-renderer] package is designed to be used in applications where shared templates
are used in conjunction with page templates to compose complete HTML documents. Using [html/template],
this is achieved by overriding a shared template `block` nodes with the target template content.

As such, custom template names are not supported for the target page template (i.e., a page template
must be referred to by its filename).

In order to use [echo-renderer], simple create and assign a new `echorenderer.Renderer` to the router's
`Renderer` property. See the [examples](./examples/) directory for more information.

## License

MIT License

Copyright (c) 2024 Brian Reece

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

[echo-renderer]: https://pkg.go.dev/github.com/bdreece/echo-renderer
[echo.Renderer]: https://pkg.go.dev/github.com/labstack/echo/v4#Renderer
[html/template]: https://pkg.go.dev/html/template

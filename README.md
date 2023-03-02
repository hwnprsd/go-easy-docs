# Runtime API Doc generation for Go-Fiber

I created this project to provide a minimalistic, runtime API documentation generation tool for Go-Fiber. Unlike OpenAPI and Swagger, I wanted to create a tool that is simple and easy to use for small projects. This tool registers all routes on Fiber, generates documentation that includes the request body type, and creates a simple documentation page that can be shared with your team.

The tool currently supports GET and POST requests, and includes opinionated request handler wrappers that provide faster workflow. Some features of the tool include:

- Auto documentation generation: easily serve the documentation on any route, with any level of authentication.
- Function wrappers for GET and POST requests: handlers can be functions or methods that take one argument of any DTO you have declared for that route. This allows for more composable handlers and cleaner, typed code.
- Pre-validation of request body before feeding it to the handler.
- Post-request error handling and response structuring: all handlers return `(any, error)` which can be used to construct predictable response structures and handle errors with appropriate HTTP codes using the custom `RequestError` implementation of `error`.
- Access to ctx via closure, keeping handlers clean and typed at all times.

While this project currently only supports Go-Fiber, I may explore porting it to other frameworks in the future. Additionally, the Swagger-like UI may be improved in future versions of the tool.


UI Example
![image](https://user-images.githubusercontent.com/23007190/222368037-86af7d05-59fe-4479-9907-3531268ae0b9.png)



Open Source - Take it




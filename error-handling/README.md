### Custom Error Handling

If you implement `Error` interface with pointer the function return you this:
```
{Errors:[some thing went wrong some thing went wrong 2]} 
```
Its actually your `validationErr` struct. Why?

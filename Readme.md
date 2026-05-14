# Short Url Decoder
Decode that fishy link
## Usage
```
Usage of Short Url Decoder
  -output string
        specify the output type(final, trace, csv) (default "final")
  -url string
        the url to be decoded
```
example: `./sud --url https://bit.ly/3OIdoeL --output=trace`
# Writeup
## Objctives
1. Refresh my familiarity with go
1. Gain experience with how redirects are handled
## What I learned
This was my first time implementing a roundtripper(or even the non default transport). I implemented the roundtripper interface so that I could log each destination. While the technical work was not challenging. It provided a friendly reminder that the world is not going to end if I throw myself into something new.
## Resources
https://oneuptime.com/blog/post/2026-01-30-how-to-build-custom-http-transport-in-go/view
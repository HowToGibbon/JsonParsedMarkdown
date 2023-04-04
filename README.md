# JsonParsedMarkdown
JsonParsedMarkdown JPM to create markdown from Oreilly Notes and Highlights 

## Run
Fill the sample.json with the response you get from https://learning.oreilly.com/api/v1/annotations/all/?page_size=100

- if you have more than 100 get the rest by adding the page attribute https://learning.oreilly.com/api/v1/annotations/all/?page_size=100&page=2

In the root directory run: `go run .`
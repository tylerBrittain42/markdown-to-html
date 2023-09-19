# markdown-to-html
A project to familarize myself with Go

## About
This project takes in a markdown file and outputs an equivalent html file.

## Oddities
### Handling Paragraphs
- A paragraph tag will be created on the first occurence of a line that does not begin with a block element
- A paragraph will only end when the next line begins with a block tag or the end of file is reached
## Instructions

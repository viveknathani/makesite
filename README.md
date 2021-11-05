# makesite

makesite is an experimental, site generator tool. It will take markdown files as input and generate the corresponding HTML files, as per the directory structure.

## description of packages 

- converter: everything related to generating HTML from Markdown 
- fileio: everything related to doing file I/O in this application 
- meta: everything about parsing metadata from the `.md` files 
 
Here is an example of writing metadata that will be parsed by the program.
```
[meta]: # (CSS_URL=./theme.css)
[meta]: # (DOCUMENT_TITLE=title)

```
This should be present at the beginning of your document.

## license

[MIT](./LICENSE)

# makesite

makesite is an experimental, site generator tool. It will take markdown files as input and generate the corresponding HTML files, as per the directory structure.

## description of packages 

- converter: everything related to generating HTML from Markdown 
- fileio: everything related to doing file I/O in this application 
- meta: everything about parsing metadata from the `.md` files 
 
Over here, `metadata` includes information for `<meta>` tags, inner text for `<head>` and `<title>` tags, path to the CSS file to be used.

## license

[MIT](./LICENSE)

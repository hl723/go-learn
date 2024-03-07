# Custom Wiki Pages

This project creates custom wiki pages. 

- Users can create, edit, and view wiki pages. 
- The code uses a Page data structure to store each page with a title and body. 
- For persistent storage, we are writing pages to files. 
- Used net/http package to build a simple web app. 
- Used html/template package to have dynamic HTML pages that can be changed with code.
- Added error handling on all functions to ensure the server does not crash.
- Implemented template caching to as an optimization to avoid redundant computation.
- Performs input validation with a regexp to avoid user accessing invalid paths.
- Used function literals and closures to write a wrapper for the request handlers for avoiding code dupication.
- Restructured code to have dedicated template and data directories.
- Added a handler for a default Home page redirection.
- Added the ability to add interpage links by specifying `[PageName]` in the page text.

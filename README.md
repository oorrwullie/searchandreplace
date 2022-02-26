## searchandreplace
Recursively search files or filenames and replace text excluding the node_modules, .git, .idea, .vscode and .svn directories and bundle.js and DS_Store files.

### Usage
  --dir string
    	Directory to search
  --file
    	Search and replace file name
  --replace string
    	Text to replace with
  --search string
    	Text to search for
  --text
    	Search and replace text
  #### Example
  `searchAndReplace --text --search "foo" --replace "bar" --dir .`
  
  ## Installation
  ```
  git clone https://github.com/oorrwullie/searchandreplace
  cd searchandreplace
  go install searchAndReplace.go
  ```

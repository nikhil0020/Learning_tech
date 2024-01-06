* Open Sidebar
    * Leader by default is Space key
    * toggle side meny ( Leader + e)
    * Press m to move the file
    * Press d for delete file
    * Press r for rename of file

* File Navigation
    * Press double g to reach start of the page
    * Press capital G to reach end of the page
    * Press $ to read end of line
    * Press 0 to reach start of line
    * E or ( Shift + right arrow) to fast forward
    * (Shift + left arrow ) to fast backward

* Find a file (Leader + FF) 

* Search across codebase (Leader + SG)( S = search , G = grep)

* Find and Replace (Local)
    * ( :%s/ToReplace/Replace/s )
    * Ex -> ( :%s/app.get/app.post/s )

* Splitting windows
    * Split horizontally (Leader - )
    * Split vertically (Leader | )
    * Navigate windows (Ctrl + (H/J/K/L))

* Switching between tabs
    * (Shift + L) for moving forward between tabs
    * (Shift + H) for mmoving backward between tabs

* LSP
    * Go to Defination (Leader + gd)
    * Autocomplete
    * Hover Documentation (k)
    * Smart replace (Leader + cr) -> when we change name of a variable it will be replaced wherever it is used
    * format (Leader + cf)
    * Code Action (Leader + ca ) -> give suggestion of how to fix warning related to code , ex- unused variables
    * Remove unused imports (Leader + cR)
    * organise import (Leader + co)

* Switch buffers
    * "[ + b" -> go to left buffer (tab)
    * "] + b" -> go to right buffer (tab)

* Line numbers
    * To go to a certain line 
        * press number + J (to go down)
        * press number + k (to go up)
        * number is relative position from current line

* Open terminal
    * Leader + ft

* Explore (First go to explore mode)
    * press : when enter "Explore" to get into explore mode
    * Add a file (%)
    * Add a folder (D)

* Delete a line ( dd ) -> it is equivalent to cut not delete
* Undo (u)
* Redo (Ctrl + r)
* Paste (p)

* Visual Mode (v)
    * after going to visual mode we can select a segment which we want to copy , cut
    * copy (y) -> yanked means copy
    * select between symbols like string or bracket
        * Press vi + symbol
        * Ex-> vi + " , this will select everything inside string

* Find ( /wordToFind )
    * Press n to go to next occurance
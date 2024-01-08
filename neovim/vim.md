1. :q! -> quit override

2. w to jump to next word

3. b to jump to prev word

4. W -> Next WORD

5. B -> Prev WORD

6. r to replace letter

7. R -> Replace mode

8. cw -> change word (It will delete the word and now we can replace that)

9. 8w -> jump eight words

10. c7w -> change 7 word

11. 4j or 4 \<Arrow down> -> move four line down

12. C -> to delete rest of the line

13. dw -> delete the word

14. D -> Delete rest of word

15. d4w -> delete 4 words

16. db -> delete prev word

17. dd -> to delete line

    - 4dd -> delete 4 lines

18. cc -> change line

    - 4cc -> change 4 lines

19. u to undo

    - 3u -> undo last 3 changes

20. ctrl + r to redo

    - 5 ctrl + r to redo 5 last things

21. ciw -> change inner word

22. diw -> delete inner word

23. Change between particular symbols

    - ci) -> change inner panentheses
    - ci( -> change inner parentheses
    - ci] -> change inner brackets
    - ci[ -> change inner brackets
    - ci} -> change inner brackets
    - ci{ -> change inner brackets

24. % -> jump to brackets

    - If we are at start of bracket then by % we can go to end and vice versa

25. c% -> change until brackets

26. gg -> beginning of File
27. G -> End of File
28. 19G -> go to line 19
29. :17 -> go to line 17
30. $ -> End of line
31. 0 -> beginning of line
32. d$ -> delete till the end of the line

33. Copy and Paste (First we have to enter into visual mode)

    - d -> Delete
    - c -> change
    - y -> Yanking (copy)
    - p -> paste After
    - P -> paste Before
    - yy -> yank line
    - 5yy -> yank next 5 lines
    - 9p -> paste 9 times
    - y3w -> yank next 3 words , y3b -> yank prev 3 words
    - yi) -> yank inner bracket
    - yiw -> yank inner word
    - shift + v -> visual line
    - ctrl + v -> visual block
    - . -> press dot to repeat the last operation
    - zz -> to center selected line
    - \> shify right
    - < shift left
    - = indent
    - \>> ? << -> Shift line
    - == -> indent line
    - gg=G -> indent whole File
    - ggdG -> delete whole File

34. Search and replace
    - /word -> to search word
    - n -> next Occurance
    - N -> prev Occurance
    - \+ -> for the next token Occurance
    - \# -> for the prev token Occurance
    - :s/oldWord/newWord/g - Replace
    - :%s/oldWord/newWord/g - Replace everywhere
35. Set Properties
    - :set number -> Line numbers
    - :set relativenumber -> Rel. numbers
    - :colorscheme scheme -> Select theme (Press tab after :colorscheme to see multiple schemes)
    - :set tabstop=4 -> tab width to 4
    - :set autoindent -> auto indentation
    - :set mouse=a -> activate mouse
    - :set mouse="" -> deactivate mouse

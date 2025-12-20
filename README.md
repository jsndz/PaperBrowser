                                                                                                                               
,-.----.                                                                                                                       
\    /  \                                              ,---,.                                                                  
|   :    \            ,-.----.                       ,'  .'  \                                                                 
|   |  .\ :           \    /  \             __  ,-.,---.' .' |  __  ,-.   ,---.           .---.                        __  ,-. 
.   :  |: |           |   :    |          ,' ,'/ /||   |  |: |,' ,'/ /|  '   ,'\         /. ./|  .--.--.             ,' ,'/ /| 
|   |   \ : ,--.--.   |   | .\ :   ,---.  '  | |' |:   :  :  /'  | |' | /   /   |     .-'-. ' | /  /    '     ,---.  '  | |' | 
|   : .   //       \  .   : |: |  /     \ |  |   ,':   |    ; |  |   ,'.   ; ,. :    /___/ \: ||  :  /`./    /     \ |  |   ,' 
;   | |`-'.--.  .-. | |   |  \ : /    /  |'  :  /  |   :     \'  :  /  '   | |: : .-'.. '   ' .|  :  ;_     /    /  |'  :  /   
|   | ;    \__\/: . . |   : .  |.    ' / ||  | '   |   |   . ||  | '   '   | .; :/___/ \:     ' \  \    `. .    ' / ||  | '    
:   ' |    ," .--.; | :     |`-''   ;   /|;  : |   '   :  '; |;  : |   |   :    |.   \  ' .\     `----.   \'   ;   /|;  : |    
:   : :   /  /  ,.  | :   : :   '   |  / ||  , ;   |   |  | ; |  , ;    \   \  /  \   \   ' \ | /  /`--'  /'   |  / ||  , ;    
|   | :  ;  :   .'   \|   | :   |   :    | ---'    |   :   /   ---'      `----'    \   \  |--" '--'.     / |   :    | ---'     
`---'.|  |  ,     .-./`---'.|    \   \  /          |   | ,'                         \   \ |      `--'---'   \   \  /           
  `---`   `--`---'      `---`     `----'           `----'                            '---"                   `----'            
                                                                                                                               
# PaperBrowser


PaperBrowser is a tiny HTML-to-image rendering engine written in Go.
It supports only a minimal, strict subset of HTML and inline styles,
with the goal of learning how real browser engines work — by building
one from scratch.

No layout engines.  
No hidden CSS rules.  
No shortcuts.

This is **not a browser**.  
It’s a **paper renderer** — turning HTML into pixels as if printed on a page.


The idea is simple, convert the HTML into a image.

For the Version 1,

- The Input should take a HTML file like a simple one with no metadata. And convert it to a image with paragraph and text.
- No colors/CSS/JS for now.
- Every element should have a closing tag.



---
    Tokenizing HTML → creating a DOM

    Parsing inline styles

    Block vs inline layout

    Text measurement and line breaking

    Painting onto a pixel canvas


🛠 Tech Stack

    Go (core engine)

    gg + freetype for font shaping & drawing

    No existing HTML/CSS rendering libraries



 Vision

PaperBrowser starts very small.
Every new feature must teach one new rendering principle.

When it grows beyond that, you’ve already won —
because you’ll understand the foundations of real browsers.


📎 License

MIT — Learn. Hack. Break. Improve.


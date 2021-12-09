```
  ,----..              .---.  .--.--.    
 /   /   \            /. ./| /  /    '.  
|   :     :       .--'.  ' ;|  :  /`. /  
.   |  ;. /      /__./ \ : |;  |  |--`   
.   ; /--`   .--'.  '   \' .|  :  ;_     
;   | ;  __ /___/ \ |    ' ' \  \    `.  
|   : |.' .';   \  \;      :  `----.   \ 
.   | '_.' : \   ;  `      |  __ \  \  | 
'   ; : \  |  .   \    .\  ; /  /`--'  / 
'   | '/  .'   \   \   ' \ |'--'.     /  
|   :    /      :   '  |--"   `--'---'   
 \   \ .'        \   \ ;                 
  `---`           '---"                  
(g)o (w)eb (s)erver

Bare-bones GET-only web server in golang for personal static site hosting.

(c) 2021 Brian Chen ihasdapi <at> gmail.com
This is licensed under the 'unlicense' license. See LICENSE for details.
```

A simple web server in golang used for hosting [chenbrian.ca](https://chenbrian.ca)
As it is a static site, this server only has to serve `GET` requests and deal with simple routing/error handling.

Configuration is done in the `CONFIG.yml` file. 
By convention this will serve the `index.html` found at the path requested by the URL.




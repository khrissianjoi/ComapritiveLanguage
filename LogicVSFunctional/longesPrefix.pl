
longestPrefix(A,[B|C]) :- prefixMatcher(A,B,C).

prefixMatcher("","",_) :- !.
prefixMatcher(A,A,[C|[]]) :- string_chars(A,BX), string_chars(C,CX), charMatcher(BX, CX,A), !.
prefixMatcher(A,B,[C|D]) :- string_chars(B,BX), string_chars(C,CX), charMatcher(BX, CX, BC), prefixMatcher(A,BC,D).

charMatcher([A|B],[A|C],D) :- getHead(B,BX), getHead(C,CX), CX == BX, charMatcher(B,C,E), string_concat(A,E,D), !.
charMatcher([A|[]],[A|_],D) :- string_concat(A,"",D), !.
charMatcher([A|B],[A|C],"") :- getHead(B,BX), getHead(C,CX), CX \= BX, !.
charMatcher(_,_,"").

getHead([A|_],A).

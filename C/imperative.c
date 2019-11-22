#include <stdio.h>
#include <ctype.h>

char * lowerChar(char * myChar) {

    return "c";

}

int main()
{
    printf("Hi there!\n");
    printf("Welcome to the To-Do-List manager.\n");
    printf("Would you like to enter an Event or a Task?\n");
    char toDo[5];
    scanf("%s:", toDo);
    printf("You have chosen to add %s.\n", toDo);
    // char check = tolower('c');
    char check = lowerChar(toDo);
}

// cc -c first.c "compile"
// cc -o first first.c "create executable"
// ./first "execute"	

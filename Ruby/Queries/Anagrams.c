#include <stdio.h>
#include <string.h>
void bubble_sort(char str[]){
    for(int i = 0; i<strlen(str); i++){
        for(int j = 0; j < (strlen(str) -1) ; j++){
            if(str[j] > str[j+1]){
                char temp = str[j];
                str[j] = str[j+1];
                str[j+1] = temp;
            }
        }
    }
}
int main(){
    char str_1[50], str_2[50];
    printf("Enter 1st string \n");
    scanf("%s", str_1);
    bubble_sort(str_1);
    printf("Enter 2nd string \n");
    scanf("%s", str_2);
    bubble_sort(str_2);
    int lenght = strlen(str_1);
    int i = 0;
    while(i<lenght){
        if(strlen(str_1) != strlen(str_2)){
            printf("not anagrams");
            break;
        }
        else if(str_1[i] != str_2[i]){
            printf("not anagrams");
            break;
        }
        else if(i == lenght-1){
            printf("strings are anagrams");
        }
        i++;
    }
    return 0;
}
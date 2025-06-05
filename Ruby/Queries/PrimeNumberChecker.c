#include <stdio.h>
#include <stdbool.h>
int main(){
    int num;
    scanf("%d",&num);
    int temp = 2;
    bool prime = true;
    while(temp < num){
        if(num%temp == 0){
            printf("Not Prime");
            prime = false;
            break;
        }
        temp += 1;
    }
    if(prime){
        printf("Prime Num");
    }
    return 0;
}
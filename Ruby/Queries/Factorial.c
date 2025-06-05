#include <stdio.h>
int factorial(int num){
    if(num == 0 || num == 1){
        return 1;
    }
    return (num * factorial(num - 1));
}
int main(){
    int fact;
    printf("Enter the num of which factorial require \n");
    scanf("%d",& fact);
    int number = factorial(fact);
    printf("%d",number);
    return 0;
}
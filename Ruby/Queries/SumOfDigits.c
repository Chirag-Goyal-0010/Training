#include <stdio.h>
int sum(int num){
    int sum = 0, digit, temp = num;
    while(temp>0){
        digit = temp % 10;
        sum = sum + digit;
        temp = temp / 10;
    }
    return sum;
}
int main(){
    int sum_digits;
    printf("Enter the num of which sum of digits require \n");
    scanf("%d",& sum_digits);
    int number = sum(sum_digits);
    printf("%d",number);
    return 0;
}
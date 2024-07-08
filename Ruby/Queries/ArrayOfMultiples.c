#include <stdio.h>

int main()
{
    int num, length = 0;
    scanf("%d",& num);
    int arr[num];
    for(int i = 1; i <= num; i++){
        if(num % i == 0){
            arr[length] = i;
            printf("%d, ", arr[length]);
            length ++;
        }
    }
    return 0;
}
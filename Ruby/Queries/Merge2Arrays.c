#include <stdio.h>
int main(){
    printf("no of elements in array 1 \n");
    int element_1;
    scanf("%d", & element_1);
    int arr_1[element_1];
    for(int i = 0; i<element_1; i++){
        scanf("%d", & arr_1[i]);
    }
    printf("no of elements in array 2 \n");
    int element_2;
    scanf("%d", & element_2);
    int arr_2[element_2];
    for(int i = 0; i<element_2; i++){
        scanf("%d", & arr_2[i]);
    }
    int element = element_1 + element_2;
    printf("total no of elements %d \n", element);
    int arr[element];
    for(int i = 0; i<element; i++){
        if(i<element_1){
            arr[i] = arr_1[i];
        }
        else {
            arr[i] = arr_2[i - element_1];
        }
    }
    for (int i = 0; i < element; i++){
        printf(" %d element is %d \n", i+1, arr[i]);
    }
    return 0;
}
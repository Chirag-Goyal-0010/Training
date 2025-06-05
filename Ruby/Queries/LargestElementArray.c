#include <stdio.h>
int large_element(int num[50]){
    
}
void bubble_sort(int arr[], int num){
    for(int i = 0; i<num; i++){
        for(int j = 0; j<num - 1; j++){
            if(arr[j] > arr[j+1]){
                int temp = arr[j+1];
                arr[j+1] = arr[j];
                arr[j] = temp;
            }
        }
    }
}
int main(){
    printf("no of elements \n");
    int element;
    scanf("%d", & element);
    int arr[element];
    for(int i = 0; i<element; i++){
        scanf("%d", & arr[i]);
    }
    bubble_sort(arr, element);
    // for (int i = 0; i < element; i++){
    //     printf(" %d element is %d \n", i+1, arr[i]);
    // }
    printf("Largest element of array is %d",arr[element-1]);
    return 0;
}
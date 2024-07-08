#include <stdio.h>
void remove_duplicate(int arr[], int *elements){
    for(int i = 0; i < *elements; i++){
        for(int j = i+1; j < *elements; j++){
            if(arr[i] == arr[j]){
                for(int k = j; k < (*elements - 1); k++){
                    arr[k] = arr[k+1];
                }
                (*elements)--;
            }
        }
    }
}
int main(){
    int arr[50];
    printf("no of elements: ");
    int elements;
    scanf("%d",& elements);
    for(int i = 0; i < elements; i++){
        scanf("%d", & arr[i]);
    }
    remove_duplicate(arr , &elements);
    for(int i = 0; i < elements; i++){
        printf("%d ",arr[i]);
    }
    return 0;
}
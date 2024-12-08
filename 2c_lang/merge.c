#include <stdio.h>

// Function to merge two subarrays
void merge(int A[], int p, int q, int r) {
    int nL = q - p + 1; // Length of the left subarray
    int nR = r - q;     // Length of the right subarray

    int L[nL], R[nR]; // Temporary arrays to hold the subarrays

    // Copy data to the left subarray
    for (int i = 0; i < nL; i++) {
        L[i] = A[p + i];
    }
    // Copy data to the right subarray
    for (int j = 0; j < nR; j++) {
        R[j] = A[q + 1 + j];
    }

    int i = 0, j = 0, k = p;

    // Merge the two subarrays back into A
    while (i < nL && j < nR) {
        if (L[i] <= R[j]) {
            A[k] = L[i];
            i++;
        } else {
            A[k] = R[j];
            j++;
        }
        k++;
    }

    // Copy any remaining elements from the left subarray
    while (i < nL) {
        A[k] = L[i];
        i++;
        k++;
    }

    // Copy any remaining elements from the right subarray
    while (j < nR) {
        A[k] = R[j];
        j++;
        k++;
    }
}

// Recursive merge sort function
void mergeSort(int A[], int p, int r) {
    if (p < r) {
        int q = (p + r) / 2;

        // Recursively sort the left half
        mergeSort(A, p, q);

        // Recursively sort the right half
        mergeSort(A, q + 1, r);

        // Merge the two halves
        merge(A, p, q, r);
    }
}

// Function to print the elements of an array
void printArray(int A[], int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", A[i]);
    }
    printf("\n");
}

// Main function
int main() {
    int A[] = {12, 11, 13, 5, 6, 7};
    int size = sizeof(A) / sizeof(A[0]);

    printf("Array before sorting:\n");
    printArray(A, size);

    mergeSort(A, 0, size - 1);

    printf("Array after sorting:\n");
    printArray(A, size);

    return 0;
}

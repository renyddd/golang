#include <stdio.h>
#include <string.h>

#define N 100

// 剑指 offer 题目 4：空格替换

void ReplaceBlank (char string[], int length) {
    if (string == NULL && length <= 0) {
        return;
    }

    int originalLength = 0;
    int numberOfBlank = 0;
    int i = 0;

    while (string[i] != '\0') {
        ++originalLength;

        if (string[i] == ' ') {
            ++numberOfBlank;
        }
        ++i;
    }
    printf("originalLength: %d, numberOfBlank: %d\n", originalLength, numberOfBlank);

    // newLength 为替换空格后的长度
    int newLength = originalLength + numberOfBlank * 2;

    // 注意此处的逻辑判断
    // 既然这是在元数组地址上进行复制，因此待处理数字的本身长度要有足够的空间。
    // length: 函数的形参
    if (newLength > length) {
        return;
    }

    int indexOfOriginal = originalLength;
    int indexOfNew = newLength;
    printf("indexOfOriginal: %d, indexOfNew: %d\n", indexOfOriginal, indexOfNew);

    while (indexOfOriginal >= 0 && indexOfNew > indexOfOriginal) {
        if (string[indexOfOriginal] == ' ') {
            string[indexOfNew--] = '0';
            string[indexOfNew--] = '2';
            string[indexOfNew--] = '%';
        } else {
            string[indexOfNew--] = string[indexOfOriginal];
        }
        --indexOfOriginal;
    }
}


int main() {
    char string[N] = "ab c d.";
    int i = 0;

    for (; i < strlen(string); i++) {
        printf("%c", string[i]);
    }
    printf("\n");

    ReplaceBlank(string, N);

    i = 0;

    for (; i < strlen(string); i++) {
        printf("%c", string[i]);
    }
    printf("\n");

    return 0;
}

#include <iostream>

using namespace std;

int main(void)
{
    bool arr[30] = {
        false,
    };
    int num;
    const int MAX_SUM = 465;
    int sum = 0;

    for (int i = 0; i < 28; i++)
    {
        cin >> num;
        arr[num] = true;
        sum += num;
    }
    sum = MAX_SUM - sum;

    for (int j = 0; j < 30; j++)
    {
        if (arr[j + 1] == false)
        {
            cout << j + 1 << endl;
            cout << sum - (j + 1) << endl;
            return 0;
        }
    
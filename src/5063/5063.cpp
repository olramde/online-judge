#include <iostream>

using namespace std;

int adBenefit(const int &r, const int &e, const int &c)
{
    int clearProfit = e - c;

    if (clearProfit == r)
        return 0; // does not matter
    else if (clearProfit > r)
        return 1; // advertise
    else if (clearProfit < r)
        return 2;
}

void printResultString(int num)
{
    switch (num)
    {
    case 0:
        cout << "does not matter" << endl;
        break;
    case 1:
        cout << "advertise" << endl;
        break;
    case 2:
        cout << "do not advertise" << endl;
        break;

    default:
        break;
    }
}

int main(void)
{
    int num;
    cin >> num;
    int *results = new int[num];

    for (int i = 0; i < num; i++)
    {
        int r, e, c;
        cin >> r >> e >> c;

        results[i] = adBenefit(r, e, c);
    }

    for (int i = 0; i < num; i++)
    {
        printResultString(results[i]);
    }
}
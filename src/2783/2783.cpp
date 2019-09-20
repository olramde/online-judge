#include <iostream>
#include <float.h>

using namespace std;

void checkMinPrice(double price, double gram, double &minPricePerGram)
{
    if (minPricePerGram > price / gram)
        minPricePerGram = price / gram;
}

int main(void)
{
    double price, gram;
    int count = 0;
    double minPricePerGram = DBL_MAX;

    cin >> price >> gram;
    checkMinPrice(price, gram, minPricePerGram);

    cin >> count;
    for (int i; i < count; i++)
    {
        cin >> price >> gram;
        checkMinPrice(price, gram, minPricePerGram);
    }

    cout << minPricePerGram * 1000 << endl;
}
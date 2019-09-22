#include <iostream>
#include <string>

using namespace std;

int main(void)
{
    int numOfStudent;
    int timeOfShow;
    bool *timeTable;
    int countTime = 0;

    cin >> numOfStudent >> timeOfShow;
    timeTable = new bool[timeOfShow + 1];

    for (int i = 0; i < numOfStudent; i++)
    {
        int cycle;
        int currentCycle;

        cin >> cycle;
        currentCycle = cycle;
        while (true)
        {
            if (currentCycle > timeOfShow)
                break;
            else if (timeTable[currentCycle] == false)
            {
                timeTable[currentCycle] = true;
                countTime += 1;
            }
            currentCycle += cycle;
        }
    }

    cout << countTime << endl;
}
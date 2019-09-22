#include <iostream>
#include <utility>
#include <vector>
#include <algorithm>

using namespace std;

bool dsecVectorValue(const pair<int, int> &a, const pair<int, int> &b)
{
    return (a.second > b.second);
}

int main(void)
{
    vector<pair<int, int>> scoreDashboard;
    int topFiveScsoreUser[5];
    int amountScore = 0;

    for (int i = 1; i <= 8; i++)
    {
        int score;
        cin >> score;

        scoreDashboard.push_back(make_pair(i, score));
    }
    sort(scoreDashboard.begin(), scoreDashboard.end(), dsecVectorValue);

    for (int i = 0; i < 5; i++)
    {
        topFiveScsoreUser[i] = scoreDashboard[i].first;
        amountScore += scoreDashboard[i].second;
    }

    sort(topFiveScsoreUser, topFiveScsoreUser + 5);
    cout << amountScore << endl;

    for (int i = 0; i < 5; i++)
    {
        cout << topFiveScsoreUser[i] << " ";
    }
}
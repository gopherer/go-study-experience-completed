package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>

int argMax(int t[5]);
void HandleProcess(int t);
void UI();

struct Process
{
    char name;
    int cupTime;
    int needTime;
    int round;
    int count;
    char state[10];
};

const int processCount = 5;
struct Process pro[5];
char arrState[10];

int main()
{

    int needTimeCount;
    srand((unsigned)time(NULL));
    int temp;

    for (size_t i = 0; i < processCount; i++)
    {
    Flag:
        temp = rand() % 6 + 1;
        if (temp % 2 == 0)
        {
            pro[i].needTime = temp;
        }
        else
        {
            goto Flag;
        }

        pro[i].name = 'a' + i;
        pro[i].cupTime = 0;
        pro[i].round = 0;
        pro[i].count = 0;
        strcpy(pro[i].state, "waiting");
    }
    for (size_t i = 0; i < processCount; i++)
    {
        needTimeCount += pro[i].needTime;
    }
    for (size_t i = 0; i < needTimeCount / 2; i++)
    {
        int t = i % processCount;

        if (pro[t].needTime != 0)
        {
            HandleProcess(t);
        }
        else
        {
            for (size_t i = 0; i < processCount; i++)
            {
                t = (t + 1) % processCount;
                if (pro[t].needTime != 0)
                {
                    HandleProcess(t);
                    break;
                }
            }
        }
    }
    return 0;
}

void HandleProcess(int t)
{
    pro[t].needTime -= 2;
    pro[t].cupTime += 2;
    pro[t].round += 1;
    pro[t].count += 1;
    if (pro[t].needTime != 0)
    {
        strcpy(pro[t].state, "running");
        UI();
        strcpy(pro[t].state, "waiting");
    }
    else
    {
        strcpy(pro[t].state, "terminated");
        UI();
    }
    return;
}
void UI()
{
    printf("name  cupTime  needTime  round  count  state\n");
    for (size_t i = 0; i < processCount; i++)
    {
        printf("%c         %d     %d        %d       %d     %s\n", pro[i].name, pro[i].cupTime, pro[i].needTime, pro[i].round, pro[i].count, pro[i].state);
    }
    printf("-----------------------------------------------------\n");
}
*/

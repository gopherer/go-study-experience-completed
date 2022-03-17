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
int priority;
char state[10];
};

const int processCount = 5;
struct Process pro[5];
int arr[5];
int arrPriority[5];
char arrState[10];

int main()
{
int needTimeCount;
srand((unsigned)time(NULL));

for (size_t i = 0; i < processCount; i++)
{
pro[i].name = 'a' + i;
pro[i].cupTime = 0;
pro[i].needTime = rand() % 3 + 1;
pro[i].priority = rand() % 10 + 40;
arrPriority[i] = pro[i].priority;
strcpy(pro[i].state, "waiting");
}
for (size_t i = 0; i < processCount; i++)
{
arr[i] = pro[i].priority;
needTimeCount += pro[i].needTime;
}
for (size_t i = 0; i < needTimeCount; i++)
{
int t = argMax(arr);
if (t == 0)
{
HandleProcess(0);
}
else if (t == 1)
{
HandleProcess(1);
}
else if (t == 2)
{
HandleProcess(2);
}
else if (t == 3)
{
HandleProcess(3);
}
else if (t == 4)
{
HandleProcess(4);
}
for (size_t i = 0; i < processCount; i++)
{
arr[i] = pro[i].priority;
}
}
return 0;
}

int argMax(int arr[5])
{
int maxIndex = 0;
int maxVal = arr[0];

for (int i = 0; i < processCount; i++)
{
if (maxVal < arr[i])
{
maxVal = arr[i];
maxIndex = i;
}
}
return maxIndex;
}

void HandleProcess(int t)
{
pro[t].priority -= 3;
pro[t].needTime -= 1;
pro[t].cupTime += 1;
arrPriority[t] = pro[t].priority;
if (pro[t].needTime != 0)
{
strcpy(pro[t].state, "running");
UI();
strcpy(pro[t].state, "waiting");
}
else
{
arrPriority[t] = pro[t].priority;
pro[t].priority = -1;
strcpy(pro[t].state, "terminated");
UI();
}
return;
}
void UI()
{
printf("name  cupTime  needTime  priority  state\n");
for (size_t i = 0; i < processCount; i++)
{
printf("%c         %d     %d        %d        %s\n", pro[i].name, pro[i].cupTime, pro[i].needTime, arrPriority[i], pro[i].state);
}
printf("-----------------------------------------------------\n");
}
*/

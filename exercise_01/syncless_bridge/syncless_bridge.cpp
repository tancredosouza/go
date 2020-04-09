#include <iostream>
#include <string>
#include <thread>
#include <chrono>
#include <mutex>

using namespace std;

enum Direction
{
  LEFT,
  RIGHT
};

bool isGoingLeft, isGoingRight;
string bridge;

Direction getDirection()
{
  return isGoingLeft ? LEFT : RIGHT;
}

void cross()
{
  Direction direction = getDirection();

  cout << "=================================" << endl;
  cout << "🚗 Car on the bridge is going " << direction << endl;

  this_thread::sleep_for(chrono::seconds(5));

  cout << "✅ Car finished going to the " << direction << endl;
  cout << "---------------------------------" << endl;
}

void tryCrossingBridge(Direction direction)
{
  switch (direction)
  {
  case LEFT:
    isGoingLeft = true;
    cross();
    isGoingLeft = false;
    break;

  case RIGHT:
    isGoingRight = true;
    cross();
    isGoingRight = false;
    break;
  }
}

void keepTryingCrossingBrigeToThe(Direction direction, int numberSecondsSleeping)
{
  while (true)
  {
    cout << "🕑 Car is waiting to go to the " << direction << endl;
    this_thread::sleep_for(chrono::seconds(numberSecondsSleeping));

    tryCrossingBridge(direction);
  }
}

void bridgeWatcher()
{
  while (true)
  {
    if (isGoingLeft && isGoingRight)
    {
      cout << "💥 CRASH HAPPENED!" << endl;
      throw 255;
    }
  }
}

int main()
{
  thread first(keepTryingCrossingBrigeToThe, LEFT, 2);
  thread second(keepTryingCrossingBrigeToThe, RIGHT, 3);
  thread third(bridgeWatcher);

  while (true)
    ;

  return 0;
}
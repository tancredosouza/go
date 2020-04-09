#include <iostream>
#include <string>
#include <thread>
#include <chrono>
#include <mutex>

using namespace std;

mutex bridgeMutex;

bool isGoingLeft, isGoingRight;
string bridge;

string getDirection()
{
  return isGoingLeft ? "left" : "right";
}

void cross()
{
  string direction = getDirection();

  cout << "=================================" << endl;
  cout << "🚗 Car on the bridge is going " << direction << endl;

  this_thread::sleep_for(chrono::seconds(5));

  cout << "✅ Car finished going to the " << direction << endl;
  cout << "---------------------------------" << endl;
}

void tryCrossingBridge(string s)
{
  if (s == "left")
  {
    isGoingLeft = true;
    cross();
    isGoingLeft = false;
  }
  if (s == "right")
  {
    isGoingRight = true;
    cross();
    isGoingRight = false;
  }
}

void keepTryingCrossingBrigeToThe(string direction, int numberSecondsSleeping) {
  while(true) {
    cout << "🕑 Car is waiting to go to the " << direction << endl;
    this_thread::sleep_for(chrono::seconds(numberSecondsSleeping));

    bridgeMutex.lock();

    tryCrossingBridge(direction);

    bridgeMutex.unlock();
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
  thread first(keepTryingCrossingBrigeToThe, "left", 2);
  thread second(keepTryingCrossingBrigeToThe, "right", 3);
  thread third(bridgeWatcher);

  while (true)
    ;

  return 0;
}
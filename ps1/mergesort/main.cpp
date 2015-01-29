#include <string>
#include <iostream>
#include <fstream>
#include <cstdlib>
#include <vector>

using std::string;
using std::vector;
using std::ifstream;
using std::getline;
using std::stoi;
using std::cout;
using std::endl;
using std::distance;
using std::copy;

namespace {
vector<int> get_data(const string& filename) {
  ifstream file(filename.c_str());
  if(!file.good())
    return vector<int>();

  string line;
  getline(file, line); // get the header
  getline(file, line); // get the count
  const int len = stoi(line);
  vector<int> data(len);
  for(int i = 0; getline(file, line); ++i) {
    data[i] = stoi(line);
  }
  return data;
}

inline void print_data(const vector<int>& data) {
  cout << "[";
  for(vector<int>::const_iterator i = data.begin(), end = data.end(); i != end; ++i)
    cout << *i << " ";
  cout << "]";
}

inline void print_data(vector<int>::const_iterator first, vector<int>::const_iterator last) {
  cout << "[";
  for(; first != last; ++first)
    cout << *first << " ";
  cout << "]";
}

void mergesort(vector<int>::iterator first, vector<int>::iterator last) {
  auto swap = [](int* a, int* b) {
    const auto c = *a;
    *a = *b;
    *b = c;
  };

  auto merge = [](vector<int>::const_iterator abegin, vector<int>::const_iterator aend,
                  vector<int>::const_iterator bbegin, vector<int>::const_iterator bend) {
    vector<int> buffer(distance(abegin, aend) + distance(bbegin, bend));
    auto bufferi = buffer.begin();
    auto ai = abegin;
    auto bi = bbegin;
    for(; ai != aend && bi != bend; ++bufferi) {
      if(*ai < *bi) {
        *bufferi = *ai;
        ++ai;
      } else {
        *bufferi = *bi;
        ++bi;
      }
    }

    if(ai != aend)
      copy(ai, aend, bufferi);
    else if(bi != bend)
      copy(bi, bend, bufferi);

    return buffer;
  };

  const auto len = distance(first, last);

  if(len > 2) {
    const auto split = len / 2;
    mergesort(first, first + split);
    mergesort(first + split, last);
    const auto buffer = merge(first, first + split, first + split, last);
    copy(buffer.begin(), buffer.end(), first);
    return;
  }

  if(len == 2) {
    if(*first > *(first + 1))
      swap(&(*first), &(*(first + 1)));
    return;
  }

  return;
}

} // namespace

int main(int argc, char* argv[]) {
  if(argc < 2) {
    cout << "usage: " << argv[0] << " inputfile" << endl;
    return 0;
  }

  const string filename = argv[1];
  vector<int> data = get_data(filename);
  if(data.size() == 0) {
    cout << "failed to get data" << endl;
    return 0;
  }

  print_data(data);
  cout << endl;
  cout << "sorting..." << endl;
  mergesort(data.begin(), data.end());

  print_data(data);
  cout << endl;

  return 0;
}

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

/// This is a subset of a true radix sort, for input values from [0..99].
/// A true radix sort for arbitrary integers is significantly slower, having to check many more digits.
void selection_sort(vector<int>* data_) {
  vector<int>& data = *data_;

  auto swap = [](int* a, int* b) {
    const int c = *a;
    *a = *b;
    *b = c;
  };

  vector<int> buckets(100, 0);

	// i is the current position, before which all elements are sorted
  for(int i = 0, end = data.size(); i != end; ++i) {
		int nextLowest = i; ///< the index of the lowest value in the array, which needs to be 'selected' and swapped with data[i]
		// j is the position of the iterator, which finds the next lowest value in the array
    for(int j = i, jend = data.size(); j != jend; ++j) {
      if(data[j] < data[nextLowest])
        nextLowest = j;
    }
		swap(&data[i], &data[nextLowest]);
  }
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
  selection_sort(&data);

  print_data(data);
  cout << endl;

  return 0;
}

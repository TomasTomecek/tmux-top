// how do you display specific interval
struct interval_display {
    double threshold;
    char * bg;
    char * fg;
};

struct interval_display** init_interval_list(const char *args[], int count);
void print_interval_item(struct interval_display **entries, int count, double value, char * suffix);

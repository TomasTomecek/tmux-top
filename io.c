#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of...
#include "display.h"
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>
#include <string.h>  // strsep


const char* DEFAULT_INTERVAL[] = {
    "0.0,colour4,default",
    "20.0,green,default",
    "100.0,colour166,default",
    "500.0,colour1,default",
};


struct io_stat {
    unsigned long value;
    char* device;
};

struct io_stat_list {
    struct io_stat * data;
    struct io_stat_list * next;
};

struct io_stat_list *
get_io_stats(void) {
    FILE *fp;
    char *line = NULL;
    char *dev_buffer;
    size_t len = 0;
    ssize_t read;
    unsigned long value_l;
    struct io_stat_list * head = NULL, *current;
    struct io_stat* data;

    fp = fopen("/proc/diskstats", "r");
    if (fp == NULL) {
        printf("Can't open /proc/diskstats. Is /proc mounted?");
        exit(EXIT_FAILURE);
    }
    while ((read = getline(&line, &len, fp)) != -1) {
        //    1       0 ram0 0 0 0 0 0 0 0 0 0 0 0
        sscanf(line, "%*s %*s %*s %*s %*s %*s %*s %*s %*s %*s %*s %lu", &value_l);
        if (value_l > 0) {
            data = (struct io_stat*) malloc(sizeof(struct io_stat));
            data->value = value_l;
            sscanf(line, "%*s %*s %s", dev_buffer);
            data->device = strdup(dev_buffer);
            //printf(dev_name);
            current = (struct io_stat_list *) malloc(sizeof(struct io_stat_list));
            current->data = data;
            current->next = head;
            head = current;
        }
    }
    return head;
}

int main(int argc, char **argv) {
    struct io_stat_list * head = get_io_stats();
    struct io_stat_list * current;
    current = head;
    while(current) {
        printf("%s: %lu ", current->data->device, current->data->value);
        current = current->next;
    }
    return 0;
}

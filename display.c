#define _DEFAULT_SOURCE  // shut up gcc: implicit declaration of...
#include "display.h"
#include <stdlib.h>
#include <stdio.h>
#include <string.h>  // strcmp

int cmp_func(const void * a, const void * b) {
    struct interval_display *a_s = (struct interval_display *) a;
    struct interval_display *b_s = (struct interval_display *) b;
    return (a_s->threshold < b_s->threshold) - (a_s->threshold > b_s->threshold);
}

struct interval_display**
init_interval_list(const char *args[], int count) {
    struct interval_display **entries = calloc(count, sizeof(struct interval_display*));

    char *token, *string;
    char *triple[3];
    int loop_iteration;
    for (int i = 0; i < count; i++) {
        string = strdup(args[i]);
        //printf("%s\n", string);
        loop_iteration = 0;
        while ((token = strsep(&string, ",")) != NULL) {
            //printf("token[%d] = %s\n", loop_iteration, token);
            triple[loop_iteration] = strdup(token);
            ++loop_iteration;
        }
        entries[i] = (struct interval_display*) malloc(sizeof(struct interval_display));
        //printf("T[0] = %s, T[1] = %s, T[2] = %s", triple[0], triple[1], triple[2]);
        entries[i]->fg = strdup(triple[1]);
        //printf("fg = %s\n", entries[i]->fg);
        sscanf(triple[0], "%lf", &entries[i]->threshold);
        //printf("load = %.2f\n", entries[i]->threshold);
        if (loop_iteration == 3) {
            entries[i]->bg = strdup(triple[2]);
        } else {
            entries[i]->bg = "";
        }
        //printf("bg = %s\n", entries[i]->bg);
        //printf("fg=%s, bg=%s, %.2f\n", entries[i]->fg, entries[i]->bg, entries[i]->threshold);
    }

    // sort array
    qsort(entries, count, sizeof(struct interval_display*), cmp_func);
    /*for (int i = 0; i < *entries_count; i++) {
        printf("[%d] fg=%s, bg=%s, %.2f\n", i, entries[i]->fg, entries[i]->bg, entries[i]->threshold);
    }*/
    return entries;
}

void
print_interval_item(struct interval_display **entries, int count, double value, char * suffix) {
    for (int i = 0; i < count; i++) {
        // printf("\n[%d] load = %.2f, threshold = %.2f\n", i, value, entries[i]->threshold);
        if (value > entries[i]->threshold || i == count - 1) {
            if (strcmp(entries[i]->bg, "") == 0) {
                printf("#[fg=%s]%.2f%s#[fg=default]",
                       entries[i]->fg, value, suffix);
            } else {
                printf("#[bg=%s,fg=%s]%.2f%s#[fg=default,bg=default]",
                       entries[i]->bg, entries[i]->fg, value, suffix);
            }
            return;
        }
    }
}

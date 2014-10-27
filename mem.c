#define _DEFAULT_SOURCE
#include <stdlib.h>  // getloadavg
#include <stdio.h>   // printf
#include <argp.h>
#include <string.h>  // strsep
#include <sys/sysinfo.h>
#include "display.h"


#define MEM_TOTAL "MemTotal"
#define MEM_FREE "MemFree"  // total - (used + cache)
#define MEM_AVAILABLE "MemAvailable"  // total - used
#define MEM_CACHED "Cached"
#define MEM_BUFFERS "Buffers"
#define MEM_ACTIVE "Active"


const char* DEFAULT_INTERVAL[] = {
    "0.0,colour4,default",
    "2.0,green,default",
    "3.5,colour166,default",
    "5.0,colour1,default",
};


char* units[] = {"k", "M", "G", "T"};


struct mem_info {
    unsigned long total;
    unsigned long free;
    unsigned long available;
    unsigned long cached;
    unsigned long buffers;
    unsigned long active;
};


struct mem_info_f_item {
    double value;
    char * unit;
};


struct mem_info_f {
    struct mem_info_f_item total;
    struct mem_info_f_item free;
    struct mem_info_f_item free_free;  // amount according to program free
    struct mem_info_f_item available;
    struct mem_info_f_item used;
    struct mem_info_f_item cached;
    struct mem_info_f_item buffers;
    struct mem_info_f_item active;
};


struct mem_info*
get_mem_info(void) {
    FILE *fp;
    char *line = NULL;
    char *tmp, *name, *value = NULL;
    size_t len = 0;
    ssize_t read;
    unsigned long value_l;
    struct mem_info * info = (struct mem_info*) malloc(sizeof(struct mem_info));

    fp = fopen("/proc/meminfo", "r");
    if (fp == NULL)
        exit(EXIT_FAILURE);

    while ((read = getline(&line, &len, fp)) != -1) {
        tmp = strdup(line);
        name = strsep(&tmp, ":");

        tmp = strdup(line);
        value = strtok(tmp, " ");
        value = strtok(NULL, " ");
        sscanf(value, "%lu", &value_l);
        // TODO: refactor
        if (strcmp(name, MEM_TOTAL) == 0) {
            info->total = value_l;
        } else if (strcmp(name, MEM_FREE) == 0) {
            info->free = value_l;
        } else if (strcmp(name, MEM_AVAILABLE) == 0) {
            info->available = value_l;
        } else if (strcmp(name, MEM_CACHED) == 0) {
            info->cached = value_l;
        } else if (strcmp(name, MEM_BUFFERS) == 0) {
            info->buffers = value_l;
        } else if (strcmp(name, MEM_ACTIVE) == 0) {
            info->active = value_l;
        }
    }

    fclose(fp);

    return info;
}

void
format_value(unsigned long value, struct mem_info_f_item * item) {
    int i = 0;
    double response = (double) value;
    while (response > 1024.0) {
        response = response / 1024.0;
        ++i;
    }
    item->unit = units[i];
    item->value = response;
}

struct mem_info_f*
format_mem_info(struct mem_info* info) {
    struct mem_info_f * info_f = (struct mem_info_f*) malloc(sizeof(struct mem_info_f));

    format_value(info->total, &info_f->total);
    format_value(info->cached, &info_f->cached);
    format_value(info->buffers, &info_f->buffers);
    format_value(info->free, &info_f->free);
    format_value(info->active, &info_f->active);
    format_value(info->available, &info_f->available);
    format_value(info->total - (info->buffers + info->cached + info->free), &info_f->used);
    format_value(info->buffers + info->cached + info->free, &info_f->free_free);

    return info_f;
}

void
print_mem_info(struct mem_info_f * info, struct interval_display ** entries, int count) {
    print_interval_item(entries, count, info->used.value, info->used.unit);
    printf("#[fg=white]/");
    printf("#[fg=colour14]%.2f%s", info->total.value, info->total.unit);
    //printf("%.1f%s", info->free_free.value, info->free_free.unit);
}

/* ARGPARSE */

const char *argp_program_version =
  "load 0.1";

const char *argp_program_bug_address =
  "github.com/TomasTomecek/tmux-status-helpers/";

/* Program documentation. */
static char doc[] = "doc";

/* A description of the arguments we accept. */
static char args_doc[] = "[THRESHOLD_ENTRY [THRESHOLD_ENTRY...]]";

/* The options we understand. */
static struct argp_option options[] = {
    { 0 }
};

/* Used by main to communicate with parse_opt. */
struct arguments
{
  char *args[1024];
  int args_count;
};

static error_t
parse_opt (int key, char *arg, struct argp_state *state)
{
  /* Get the input argument from argp_parse, which we
     know is a pointer to our arguments structure. */
  struct arguments *arguments = state->input;

  switch (key)
    {
    case ARGP_KEY_ARG:
      arguments->args[state->arg_num] = arg;
      ++arguments->args_count;
      break;

    case ARGP_KEY_END:
      break;

    default:
      return ARGP_ERR_UNKNOWN;
    }
  return 0;
}

static struct argp argp = { options, parse_opt, args_doc, doc, 0, 0, 0 };

/* MAIN */

int
main(int argc, char **argv) {
    struct arguments arguments;
    arguments.args_count = 0;
    argp_parse(&argp, argc, argv, 0, 0, &arguments);

    int entries_count = arguments.args_count;
    struct interval_display **entries;
    if (entries_count == 0) {
        entries_count = (int)(sizeof(DEFAULT_INTERVAL) / sizeof(DEFAULT_INTERVAL[0]));
        entries = init_interval_list(DEFAULT_INTERVAL, entries_count);
    } else {
        entries = init_interval_list((const char**) arguments.args, entries_count);
    }
    struct mem_info * mem_info = get_mem_info();
    struct mem_info_f * mem_info_f = format_mem_info(mem_info);
    print_mem_info(mem_info_f, entries, entries_count);

    return 0;
}

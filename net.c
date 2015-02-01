#define _DEFAULT_SOURCE
#include <unistd.h>  // sleep
#include <stdio.h>   // printf
#include <string.h>  // strsep
#include <sys/sysinfo.h>
#include <sys/types.h>
#include <dirent.h>
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
const char * SYS_NET_DIR = "/sys/class/net/";
const char * RX_PATH = "/statistics/rx_bytes";
const char * TX_PATH = "/statistics/tx_bytes";


char* units[] = {"k", "M", "G", "T"};


struct net_info {
    struct net_info_f_item * received;
    struct net_info_f_item * sent;
    char * name;
};


struct net_info_f_item {
    double value;
    char * unit;
    double bytes;
};


struct net_info_iter {
    struct net_info * data;
    struct net_info_ter * next;
};


//struct net_info_f_item **
void
get_net_info(void) {
    FILE *fp;
    char *line = NULL;
    char rx_path[255], tx_path[255];
    char if_path[255];
    size_t len = 0;
    ssize_t read;
    DIR *dp;
    struct dirent *ep;
    struct net_info_iter * info = NULL;

    dp = opendir(SYS_NET_DIR);
    if (dp != NULL) {
        while ((ep = readdir (dp)) != NULL) {
            if (strcmp(ep->d_name, ".") == 0 || strcmp(ep->d_name, (const char *) "..") == 0) {
                continue;
            }
            struct net_info_f_item * rx = (struct net_info_f_item *) malloc(sizeof(struct net_info_f_item));
            strcpy(if_path, SYS_NET_DIR);
            strcat(if_path, ep->d_name);
            strcpy(rx_path, if_path);
            strcat(rx_path, RX_PATH);
            fp = fopen(rx_path, "r");
            if (fp == NULL) {
                printf("%s\n", rx_path);
                perror("Can't open file.");
            }
            read = getline(&line, &len, fp);
            sscanf(line, "%lf", &rx->bytes);
            printf("%s: %lf\n", ep->d_name, rx->bytes);
            fclose(fp);

        }
        closedir(dp);
    } else {
        perror("Couldn't open the directory");
    }
    //return info;
}

/* MAIN */

int
main(int argc, char **argv) {
    get_net_info();
    return 0;
}


# event-cron
A GO program which runs traditional way of cron job and also supports http based event triggers to run tasks

# Configuration
There are two Configuration involved one which configures the server and other which listing the Jobs
- config/config.json
- config/cronjob

#Server config
This is a JSON file containing information about server host and port, log level, and cronjob file path

#Cronjob config
The cron module is based on  github.com/robfig/cron  following is the usage.


### Cron spec format

- The "standard" cron format, described on [the Cron wikipedia page] and used by
  the cron Linux system utility.

[the Cron wikipedia page]: https://en.wikipedia.org/wiki/Cron

### Event spec format

- An Event is defined by double @@ symbol followed by command to execute for example

    ```
        #@@my-event ls -lrt
    ```
- The above event can be invoked by a http request as http://[host]:[port]/run-cron?event=my-event

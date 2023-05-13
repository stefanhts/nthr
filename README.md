# NTHR

A source of truth file syncing service which supports personal deployment and file storage

## The PRTL

The entrance to the NTHR.

## The NTHR

Mysterious realm where your files live.

## Things we want to be able to do and how to achieve them

- create the nthr
    - set up filesystem
        - create directory if it doesn't exist
        - if it exits, confirm with user whether they intended that
    - ensure we're allowed to run traffic through the network port (no firewall)
- start the nthr
    - listen on our port for api calls (as follows)
        - Create file
        - Read all files
        - Read specific files
        - Update file
        - Delete file
        - Check if update is necessary for one file
            - Check which version is more up to date
        - Check if update is necessary for all files
- stop the nthr
    - stop listening...
- check the status of the nthr
    - return whether we're running or not

- create a nthr prtl
    - check if a connection can be made to the specified nthr
    - set up the filesystem
        - create directory if it doesn't exist
        - figure out what to do if it does exist (fail, maybe?)
    - perform the initial pull from the nthr
        - read all files from nthr
        - write all files to local filesystem
- enable syncing with the nthr
    - check if we're up to date with nthr
        - if we are, good
        - if not, pull and replace out of date files
            - if our files are more recent... either write to nthr or decide what to do
    - start following loop of actions
        - check if we're synced with nthr
            - if we are, good
            - if not, see which files are more recent
                - if our file is more recent, write file to nthr
                - if nthr file is more recent, pull file from nthr
- disable syncing with the nthr
    - just stop doing things
- ping the nthr
    - test if we can reach the nthr
- check the status of the nthr prtl
    - see if we can reach nthr (ping)
    - see if we're currently syncing (enabled)
    - see if we're up to date


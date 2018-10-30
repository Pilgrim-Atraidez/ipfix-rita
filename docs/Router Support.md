# Thank You For Your Interest In Helping Us Support More Devices

IPFIX-RITA needs your help in order to support more IPFIX/ Netflow v9 enabled devices. In this guide, we will walk through gathering an error log, capturing relevant network traffic, and packaging up the results for analysis.

First off, run the installer as explained in the [Readme](../README.md).

After IPFIX-RITA is installed and running, ensure your IPFIX/ Netflow v9 device is actively sending records to your collector using tcpdump. Tcpdump can be installed using your system package manager. Once tcpdump is installed, run `tcpdump -i [IPFIX/ Netflow v9 Interface] 'udp port 2055'`. This will bring up a live stream of the IPFIX/ Netflow v9 data entering your system. Before continuing, ensure you see active IPFIX/ Netflow v9 traffic appear in your terminal. You can exit out of tcpdump by hitting `CTRL-C`.

Now we can begin gathering the data needed to support your device.

Note: You may need to open up two terminal sessions.

In the first terminal:
- Begin capturing an error log by running `sudo /opt/ipfix-rita/bin/ipfix-rita logs -f | grep "ERR" | tee ipfix-rita-errs.log`
- Wait a a few minutes and see if any error messages come up
- If no error messages come up, good news! Your device may be supported but undocumented.
  - On your RITA system, run `rita show-databases`. If you see a new database with today's date, everything is working as it should!
  - Let us know the type of device you're using and whether you are using IPFIX or Netflow v9 by sending us an email at support@activecountermeasures.com
- If you see some error messages, don't worry. Continue through the guide, and we will work with you to fix everything up.
- Go ahead, and leave the error log running

Open up another terminal session and follow these steps:
- Begin a packet capture using `sudo tcpdump -i -C 50 -w ipfix-rita-debug.pcap -s 0 'udp port 2055'`
- Leave the packet capture running

Once both the error log and packet capture are running, continue to let them run for ten minutes or so. After some time has elapsed, hit `CTRL-C` in both terminal sessions.

In either terminal:
- Stop IPFIX-RITA: `/opt/ipfix-rita/bin/ipfix-rita stop`
- Make a directory to hold the error log and packet capture: `mkdir ipfix-rita-debug`
- Move the error log into the folder: `mv /path/to/ipfix-rita-errs.log ./ipfix-rita-debug`
- Move the first packet capture into the folder: `mv /path/to/ipfix-rita-debug.pcap1 ./ipfix-rita-debug`
- Compress the folder: `tar cjf ipfix-rita-debug.tar.bz2 ipfix-rita-debug`

Finally, send us an email at support@activecountermeasures.com with the name of the device you are using, the type of traffic you are collecting (IPFIX vs Netflow v9), and the compressed data we just gathered.
### A tool for programming Nuvoton devices, particularly focusing on their modern 8051 family 

### \*\* Experimental MS51FB version \*\*
	
### Notes: 
Possible targets are Nuvoton N76E003 and MS51FB9AE. 
The programmer utility has been tested with Nu-Link-Me on a NT-MS51FB 8051 NuTiny dev board.

The Program memory is limited to 12KB for both the N76E003 and MS51FB9AE processors because the image split command does not 
parse the chip configuration to determine the split between the program flash and load flash memory, 
So it defaults to the worst case of 4KB of load flash. (The N76E003 might be able to be upped to 14KB, but it was set to 
12KB by original author).

### If you are going to rebuild this program:
The include paths are setup as relative, so this program source code should be copied to your system 
(git clone of zip file - https://github.com/mountaintom/nuvoprog.git - at this time the latest version may be in one of the branches) and 
compiled (go build) in-place on your computer. The nuvoprog command should be run from there or manually moved to where you want it. 

You can compile the nuvoprog utility as another name (such as nuvoprog-test) by changing the main directory name (such as nuvoprog to nuvoprog-test) then run go build.

### Examples:
```
	Download flash data from chip:
	./nuvoprog read ./flash-read.ihx --target MS51FB9AE 

	Split downloaded flash data into Program, Load ROM and chip configuration files:
	./nuvoprog image split -i ./flash-read.ihx --target MS51FB9AE  -a program-flash-data.ihx -l loader-flas-data.ihx -c chip-configuration.json 
	Note: This is how to get an example chip-config json file to work with.

	Program flash in chip:
	./nuvoprog program --target ms51fb9ae -a ./program-to-flash.ihx -c @chip-configuration.json
        Note: The files may be combined with "image merge" and the resulting ihx file programmed with the "-i" flag.
```
Usage:
```
  nuvoprog [command]

Available Commands:
  config      Configuration tools
  devices     List connected programmers
  help        Help about any command
  image       Image manipulation commands
  program     Program a target device
  read        Read device flash contents

Flags:
  -h, --help            help for nuvoprog
  -t, --target string   target device
  -v, --verbose         make verbose (enable debug logging)

Use "nuvoprog [command] --help" for more information about a command.
subcommand is required
```


\*\*\*\*\*\*


#nuvoprog - Nuvoton microcontroller programmer

`nuvoprog` is an open source tool for programming Nuvoton microcontollers;
previously, they could only be programmed under Windows using Nuvoton's
proprietary tools. This tool assumes a Nuvoton NuLink family programmer
or compatible; no support is provided (yet) for other programmers.

This tool should be reasonably robust but presently has very limited device
support; if you wish to add support for new devices, that would be much
appreciated. Information on how to do so is at the bottom of this readme

Additionally, a human-friendly (JSON) interface to the configuration bits
is provided.

The tool provides integrated help

Example usage:
```
$ nuvoprog read -t n76e003 dev.ihx
$ nuvoprog config decode -i dev.ihx
$ nuvoprog program -t n76e003 -c @config.json -a aprom.ihx -l ldrom.ihx

```

You may also be interested in [libn76](https://github.com/erincandescent/libn76),
a SDCC-supporting BSP for the Nuvoton N76 family.

*Cortex-M devices*: While I have no objections to someone adding support for
these, have you considered OpenOCD?

# Installing
This is a Go project; install a Go toolchain and install it
using `go get -u github.com/erincandescent/nuvoprog`. Ensure
that `$GOPATH/bin` is on your path (`GOPATH` defaults to `$HOME/go`);
alternatively, move the resulting binary to a location of your choice.

The `hidapi` and `libusb` packages are [vendored by our upstream](https://github.com/karalabe/hid)

# Supported Devices
## Programmers

 *  Nu-Link-Me (as found on Nu-Tiny devboards)

Coming soon:

 * Nu-Link

## Target devices

 * N76E003 (8051T1 family)

# Missing functionality

* Firmware upgrades
* Debugging?

# Adding support for new devices

To add support for new devices, you will need:

 * Windows
 * The Nuvoton ICP tool, and
 * Wireshark

A Wireshark dissector for the protocol can be found in the misc directory.

Nuvoton have [an OpenOCD patch](http://openocd.zylin.com/#/c/4739/1) which you may find useful as reference material

## Other NuLink Programmers
If this is a protocol v2 programmer, you'll need to add support for that (The leading length field
changes from 8 to 16 bits, but othewise things are unchanged).

Add the VID and PID to the table in `protocol/device.go` and see if `nuvoprog` connects successfully.
If it doesn't, compare protocol exchanges in Wireshark

## Other Microcontrollers
First step is to see if the microcontroller belongs to the same family and if the connection and
programming flow is the same (The flow should be the same for the 8051T1 family, may differ for
others).

If they are, you probably just need to define target devide details:

 * Configuration bit codec
 * Target definition (see `target/n76/n76e003.go`)

You may need to get details like LDROM offsets from Wireshark dumps

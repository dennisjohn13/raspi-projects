# raspi-projects

Home Projects

- led_blink

# Compile the project

`GOARM=7 GOARCH=arm GOOS=linux go build led_blink.go`

# RaspBerry Pi commands

## Log on to Raspberry Pi

`ssh pi@raspberrypi.local`

## To Copy the binaries from Local(MAC) to Raspberry Pi

`scp motion_sensor pi@raspberrypi.local:/home/pi/Desktop/projects`

## To Run the program in Raspberry Pi and tail the output to Local (MAC)

`ssh -t pi@raspberrypi.local "./Desktop/projects/motion_sensor"`

## Raspberry Pi Pin Diagram

![Screenshot](GPIO-Pinout-Diagram.png)
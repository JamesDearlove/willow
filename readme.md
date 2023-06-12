# Project Willow

A WIP prototype GUI for the Beepberry.

Still a very loose goal of creating a GUI for the device, more to follow

## Setup notes

I don't have a Beepberry yet, but I do have the screen it comes with.

0. Image Pi with Raspberry Pi OS Lite, connect screen to same pins as Beepberry.

    Pinout can be found in the [Beepberry docs](https://beepberry.sqfmi.com/docs/hardware/pinouts#display)

    Note: If you're using a Pi 4 for testing, force the OS to be in 32-bit mode by adding this to your config.txt
    ```
    arm_64bit=0
    ```

1. Install the Sharp LCD Driver

    This is a very quickly moving area, currently I'm running a patched version from: [TheMediocritist](https://github.com/TheMediocritist/Sharp-Memory-LCD-Kernel-Driver/tree/8bit)

    Additionally as DRM driver work is ongoing, I'm using [snag](https://github.com/TheMediocritist/snag) (based on raspi2fb) which copies HDMI buffer to the LCD buffer.

2. Download and compile Raylib

    Details can be found in the [Raylib Wiki](https://github.com/raysan5/raylib/wiki/Working-on-Raspberry-Pi)

3. Download and install Go

4. Build/Run with Go.

    This is patchy for now, tags are required as audio crashes out and it attempts to build for X otherwise.

    Note: First time running may take a moment to compile the Raylib-go library.

    ```bash
    go run -tags drm,noaudio .
    ```

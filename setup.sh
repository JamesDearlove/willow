#!/bin/bash
set -e

echo "Enabling I2C, SPI, and Console Auto login..."
sudo raspi-config nonint do_i2c 0 || { echo "Error: Failed to enable I2C."; exit 1; }
sudo raspi-config nonint do_spi 0 || { echo "Error: Failed to enable SPI."; exit 1; }
sudo raspi-config nonint do_boot_behaviour B2 || { echo "Error: Failed to enable Console Auto login."; exit 1; }

echo "Updating and installing dependencies..."
sudo apt-get -y install git raspberrypi-kernel-headers < "/dev/null" || { echo "Error: Failed to install dependencies."; exit 1; }

echo "Compiling and installing display driver..."
cd ~/
git clone https://github.com/TheMediocritist/Sharp-Memory-LCD-Kernel-Driver.git || { echo "Error: Failed to clone display driver repository."; exit 1; }
cd ~/Sharp-Memory-LCD-Kernel-Driver
git checkout 8bit
make || { echo "Error: Failed to compile display driver."; exit 1; }
sudo make modules_install || { echo "Error: Failed to install display driver."; exit 1; }
sudo depmod -A || { echo "Error: Failed to update module dependencies."; exit 1; }
dtc -@ -I dts -O dtb -o sharp.dtbo sharp.dts || { echo "Error: Failed to compile device tree."; exit 1; }
sudo cp sharp.dtbo /boot/overlays

# fbset -fb /dev/fb0 -i
# fbset -fb /dev/fb1 -i

# cd ~/
# sudo apt-get install cmake libbsd-dev -y

# git clone https://github.com/TheMediocritist/snag.git
# cd snag

# cd ~/
# sudo apt-get install libdrm-dev libegl1-mesa-dev libgles2-mesa-dev libgbm-dev -y
# git clone https://github.com/raysan5/raylib.git
# cd raylib/src
# make PLATFORM=PLATFORM_DRM


import spidev as SPI
import ST7789

from PIL import Image

# Raspberry Pi pin configuration:
RST = 27
DC = 25
BL = 24
bus = 0
device = 0

# 240x240 display with hardware SPI:
display = ST7789.ST7789(SPI.SpiDev(bus, device), RST, DC, BL)

# Initialize library.
display.Init()

# Clear display.
display.clear()

image = Image.open('/etc/raspichamber/display/status.jpg')
display.ShowImage(image, 0, 0)
